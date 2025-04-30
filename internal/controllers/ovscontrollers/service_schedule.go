// https://app.swaggerhub.com/apis-docs/dcsaorg/DCSA_OVS/3.0.0-Beta-1#/

package ovscontrollers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-dcsa/internal/workflows/ovsworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)

// ServiceScheduleController - Create ServiceSchedule Controller
type ServiceScheduleController struct {
	log                          *zap.Logger
	UserServiceClient            partyproto.UserServiceClient
	ServiceScheduleServiceClient ovsproto.ServiceScheduleServiceClient
	wfHelper                     common.WfHelper
	workflowClient               client.Client
	ServerOpt                    *config.ServerOptions
}

// NewServiceScheduleController - Create ServiceSchedule Handler
func NewServiceScheduleController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, serviceScheduleServiceClient ovsproto.ServiceScheduleServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *ServiceScheduleController {
	return &ServiceScheduleController{
		log:                          log,
		UserServiceClient:            userServiceClient,
		ServiceScheduleServiceClient: serviceScheduleServiceClient,
		wfHelper:                     wfHelper,
		workflowClient:               workflowClient,
		ServerOpt:                    serverOpt,
	}
}

// CreateServiceSchedule - Create ServiceSchedule Header
func (sc *ServiceScheduleController) CreateServiceSchedule(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"srvsched:cud"}, sc.ServerOpt.Auth0Audience, sc.ServerOpt.Auth0Domain, sc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        ovsworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := ovsproto.CreateServiceScheduleRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		sc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := sc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, ovsworkflows.CreateServiceScheduleWorkflow, &form, token, user, sc.log)
	workflowClient := sc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var serviceSchedule ovsproto.CreateServiceScheduleResponse
	err = workflowRun.Get(ctx, &serviceSchedule)

	if err != nil {
		sc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, serviceSchedule)
}

// GetServiceSchedules - list ServiceSchedule Headers
func (sc *ServiceScheduleController) GetServiceSchedules(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"srvsched:read"}, sc.ServerOpt.Auth0Audience, sc.ServerOpt.Auth0Domain, sc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")

	serviceSchedules, err := sc.ServiceScheduleServiceClient.GetServiceSchedules(ctx,
		&ovsproto.GetServiceSchedulesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		sc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, serviceSchedules)
}

// GetServiceScheduleByUniversalServiceReference - Get ServiceSchedule By UniversalServiceReference
func (sc *ServiceScheduleController) GetServiceScheduleByUniversalServiceReference(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"srvsched:read"}, sc.ServerOpt.Auth0Audience, sc.ServerOpt.Auth0Domain, sc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("universalServiceReference")

	serviceSchedule, err := sc.ServiceScheduleServiceClient.GetServiceScheduleByUniversalServiceReference(ctx, &ovsproto.GetServiceScheduleByUniversalServiceReferenceRequest{UniversalServiceReference: id, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		sc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, serviceSchedule)
}

// UpdateServiceSchedule - Update ServiceSchedule
func (sc *ServiceScheduleController) UpdateServiceSchedule(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"srvsched:cud"}, sc.ServerOpt.Auth0Audience, sc.ServerOpt.Auth0Domain, sc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("universalServiceReference")

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        ovsworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := ovsproto.UpdateServiceScheduleByUniversalServiceReferenceRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
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
	result := wHelper.StartWorkflow(workflowOptions, ovsworkflows.UpdateServiceScheduleWorkflow, &form, token, user, sc.log)
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
