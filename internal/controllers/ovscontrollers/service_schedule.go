// https://app.swaggerhub.com/apis-docs/dcsaorg/DCSA_OVS/3.0.0-Beta-1#/

package ovscontrollers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/ovs/v3"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	"github.com/cloudfresco/sc-dcsa/internal/workflows/ovsworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// ServiceScheduleController - Create ServiceSchedule Controller
type ServiceScheduleController struct {
	log                          *zap.Logger
	UserServiceClient            partyproto.UserServiceClient
	ServiceScheduleServiceClient ovsproto.ServiceScheduleServiceClient
	wfHelper                     common.WfHelper
	workflowClient               client.Client
}

// NewServiceScheduleController - Create ServiceSchedule Handler
func NewServiceScheduleController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, serviceScheduleServiceClient ovsproto.ServiceScheduleServiceClient, wfHelper common.WfHelper, workflowClient client.Client) *ServiceScheduleController {
	return &ServiceScheduleController{
		log:                          log,
		UserServiceClient:            userServiceClient,
		ServiceScheduleServiceClient: serviceScheduleServiceClient,
		wfHelper:                     wfHelper,
		workflowClient:               workflowClient,
	}
}

// ServeHTTP - parse url and call controller action
func (sc *ServiceScheduleController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := common.GetAuthData(r)
	cdata := partyproto.GetAuthUserDetailsRequest{}
	cdata.TokenString = data.TokenString
	cdata.Email = data.Email

	md := metadata.Pairs("authorization", "Bearer "+cdata.TokenString)
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	fmt.Println("cdata", cdata)
	fmt.Println("data.TokenString", data.TokenString)
	user, err := sc.UserServiceClient.GetAuthUserDetails(ctx, &cdata)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}
	pathParts, queryString, err := common.ParseURL(r.URL.String())
	if err != nil {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}

	switch r.Method {
	case http.MethodGet:
		sc.processGet(ctx, w, r, user, pathParts, queryString)
	case http.MethodPost:
		sc.processPost(ctx, w, r, user, pathParts, data.TokenString)
	case http.MethodPut:
		sc.processPut(ctx, w, r, user, pathParts, data.TokenString)
	case http.MethodPatch:
		sc.processPatch(ctx, w, r, user, pathParts)
	case http.MethodDelete:
		sc.processDelete(ctx, w, r, user, pathParts)
	default:
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processGet - Parse URL for all the GET paths and call the controller action
/*
	GET    "/v3/service-schedules"
  GET    "/v3/service-schedules/{universalServiceReference}"
*/

func (sc *ServiceScheduleController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
	if (len(pathParts) == 2) && (pathParts[1] == "service-schedules") {
		limit := queryString.Get("limit")
		cursor := queryString.Get("cursor")
		// GET    "/v3/service-schedules"
		sc.GetServiceSchedules(ctx, w, r, limit, cursor, user)
	} else if len(pathParts) == 3 {
		sc.GetServiceScheduleByUniversalServiceReference(ctx, w, r, pathParts[2], user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPost - Parse URL for all the POST paths and call the controller action
/*
	POST "/v3/service-schedules"

*/
func (sc *ServiceScheduleController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
	fmt.Println("processPost pathParts", pathParts)
	fmt.Println("processPost pathParts[1]", pathParts[1])
	if (len(pathParts) == 2) && (pathParts[1] == "service-schedules") {
		sc.CreateServiceSchedule(ctx, w, r, user, tokenString)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPut - Parse URL for all the put paths and call the controller action
/*
  PUT    "/v3/service-schedules/{universalServiceReference}"
*/

func (sc *ServiceScheduleController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
	if len(pathParts) == 3 {
		sc.UpdateServiceSchedule(ctx, w, r, pathParts[2], user, tokenString)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPatch - Parse URL for all the patch paths and call the controller action

func (sc *ServiceScheduleController) processPatch(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// processDelete - Parse URL for all the delete paths and call the controller action
/*
 DELETE   ""
*/

func (sc *ServiceScheduleController) processDelete(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// CreateServiceSchedule - Create ServiceSchedule Header
func (sc *ServiceScheduleController) CreateServiceSchedule(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, tokenString string) {
	fmt.Println("tokenstring", tokenString)
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        ovsworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := ovsproto.CreateServiceScheduleRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		sc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := sc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, ovsworkflows.CreateServiceScheduleWorkflow, &form, tokenString, user, sc.log)
	workflowClient := sc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var serviceSchedule ovsproto.CreateServiceScheduleResponse
	err = workflowRun.Get(ctx, &serviceSchedule)
	fmt.Println("serviceschedule controller err", err)
	if err != nil {
		sc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, serviceSchedule)
}

// GetServiceSchedules - list ServiceSchedule Headers
func (sc *ServiceScheduleController) GetServiceSchedules(ctx context.Context, w http.ResponseWriter, r *http.Request, limit string, cursor string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		serviceSchedules, err := sc.ServiceScheduleServiceClient.GetServiceSchedules(ctx,
			&ovsproto.GetServiceSchedulesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
		if err != nil {
			sc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
			return
		}

		common.RenderJSON(w, serviceSchedules)
	}
}

// GetServiceScheduleByUniversalServiceReference - Get ServiceSchedule By UniversalServiceReference
func (sc *ServiceScheduleController) GetServiceScheduleByUniversalServiceReference(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		serviceSchedule, err := sc.ServiceScheduleServiceClient.GetServiceScheduleByUniversalServiceReference(ctx, &ovsproto.GetServiceScheduleByUniversalServiceReferenceRequest{UniversalServiceReference: id, UserEmail: user.Email, RequestId: user.RequestId})
		if err != nil {
			sc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
			return
		}
		common.RenderJSON(w, serviceSchedule)
	}
}

// UpdateServiceSchedule - Update ServiceSchedule
func (sc *ServiceScheduleController) UpdateServiceSchedule(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse, tokenString string) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        ovsworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := ovsproto.UpdateServiceScheduleByUniversalServiceReferenceRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		sc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UniversalServiceReference = id
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := sc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, ovsworkflows.UpdateServiceScheduleWorkflow, &form, tokenString, user, sc.log)
	workflowClient := sc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var response string
	err = workflowRun.Get(ctx, &response)
	if err != nil {
		sc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, response)
}
