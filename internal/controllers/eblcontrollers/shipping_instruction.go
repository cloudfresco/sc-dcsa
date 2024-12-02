// https://github.com/dcsaorg/DCSA-EBL/blob/master/src/main/java/org/dcsa/ebl/controller/ShippingInstructionController.java

package eblcontrollers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-dcsa/internal/workflows/eblworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// ShippingInstructionController - Create ShippingInstruction Controller
type ShippingInstructionController struct {
	log                              *zap.Logger
	UserServiceClient                partyproto.UserServiceClient
	ShippingInstructionServiceClient eblproto.ShippingInstructionServiceClient
	wfHelper                         common.WfHelper
	workflowClient                   client.Client
}

// NewShippingInstructionController - Create ShippingInstruction Handler
func NewShippingInstructionController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, shippingInstructionServiceClient eblproto.ShippingInstructionServiceClient, wfHelper common.WfHelper, workflowClient client.Client) *ShippingInstructionController {
	return &ShippingInstructionController{
		log:                              log,
		UserServiceClient:                userServiceClient,
		ShippingInstructionServiceClient: shippingInstructionServiceClient,
		wfHelper:                         wfHelper,
		workflowClient:                   workflowClient,
	}
}

// ServeHTTP - parse url and call controller action
func (sis *ShippingInstructionController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := common.GetAuthData(r)
	cdata := partyproto.GetAuthUserDetailsRequest{}
	cdata.TokenString = data.TokenString
	cdata.Email = data.Email

	md := metadata.Pairs("authorization", "Bearer "+cdata.TokenString)
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	user, err := sis.UserServiceClient.GetAuthUserDetails(ctx, &cdata)
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
		sis.processGet(ctx, w, r, user, pathParts, queryString)
	case http.MethodPost:
		sis.processPost(ctx, w, r, user, pathParts, data.TokenString)
	case http.MethodPut:
		sis.processPut(ctx, w, r, user, pathParts, data.TokenString)
	case http.MethodPatch:
		sis.processPatch(ctx, w, r, user, pathParts)
	default:
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processGet - Parse URL for all the GET paths and call the controller action
/*
	GET    "/v2/shipping-instructions/{shippingInstructionRequestReference}"
*/

func (sis *ShippingInstructionController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
	if len(pathParts) == 3 {
		// GET    "/v2/shipping-instructions/:id"
		sis.GetShippingInstructionByShippingInstructionReference(ctx, w, r, pathParts[2], user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPost - Parse URL for all the POST paths and call the controller action
/*
	POST "/v2/shipping-instructions"

*/
func (sis *ShippingInstructionController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
	if (len(pathParts) == 2) && (pathParts[1] == "shipping-instructions") {
		sis.CreateShippingInstruction(ctx, w, r, user, tokenString)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPut - Parse URL for all the put paths and call the controller action
/*
PUT   "/v2/shipping-instructions/{shippingInstructionRequestReference}" 3
*/

func (sis *ShippingInstructionController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
	if len(pathParts) == 3 {
		sis.UpdateShippingInstruction(ctx, w, r, pathParts[2], user, tokenString)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPatch - Parse URL for all the patch paths and call the controller action

func (sis *ShippingInstructionController) processPatch(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// CreateShippingInstruction - Create ShippingInstruction Header
func (sis *ShippingInstructionController) CreateShippingInstruction(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, tokenString string) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        eblworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := eblproto.CreateShippingInstructionRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		sis.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := sis.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, eblworkflows.CreateShippingInstructionWorkflow, &form, tokenString, user, sis.log)
	workflowClient := sis.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var shippingInstruction eblproto.CreateShippingInstructionResponse
	err = workflowRun.Get(ctx, &shippingInstruction)
	if err != nil {
		sis.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, shippingInstruction)
}

// GetShippingInstructionByShippingInstructionReference - GetShippingInstructionByShippingInstructionReference ShippingInstruction
func (sis *ShippingInstructionController) GetShippingInstructionByShippingInstructionReference(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		shippingInstruction, err := sis.ShippingInstructionServiceClient.FindByReference(ctx, &eblproto.FindByReferenceRequest{ShippingInstructionReference: id, UserEmail: user.Email, RequestId: user.RequestId})
		if err != nil {
			sis.log.Error("Error",
				zap.String("reqid", user.RequestId),
				zap.Error(err))
			common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
			return
		}
		common.RenderJSON(w, shippingInstruction)
	}
}

// UpdateShippingInstruction - Update ShippingInstruction
func (sis *ShippingInstructionController) UpdateShippingInstruction(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse, tokenString string) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        eblworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := eblproto.UpdateShippingInstructionByShippingInstructionReferenceRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		sis.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	form.ShippingInstructionReference = id
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := sis.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, eblworkflows.UpdateShippingInstructionWorkflow, &form, tokenString, user, sis.log)
	workflowClient := sis.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var response string
	err = workflowRun.Get(ctx, &response)
	if err != nil {
		sis.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, response)
}
