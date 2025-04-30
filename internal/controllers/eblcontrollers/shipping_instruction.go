// https://github.com/dcsaorg/DCSA-EBL/blob/master/src/main/java/org/dcsa/ebl/controller/ShippingInstructionController.java

package eblcontrollers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-dcsa/internal/workflows/eblworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)

// ShippingInstructionController - Create ShippingInstruction Controller
type ShippingInstructionController struct {
	log                              *zap.Logger
	UserServiceClient                partyproto.UserServiceClient
	ShippingInstructionServiceClient eblproto.ShippingInstructionServiceClient
	wfHelper                         common.WfHelper
	workflowClient                   client.Client
	ServerOpt                        *config.ServerOptions
}

// NewShippingInstructionController - Create ShippingInstruction Handler
func NewShippingInstructionController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, shippingInstructionServiceClient eblproto.ShippingInstructionServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *ShippingInstructionController {
	return &ShippingInstructionController{
		log:                              log,
		UserServiceClient:                userServiceClient,
		ShippingInstructionServiceClient: shippingInstructionServiceClient,
		wfHelper:                         wfHelper,
		workflowClient:                   workflowClient,
		ServerOpt:                        serverOpt,
	}
}

// CreateShippingInstruction - Create ShippingInstruction Header
func (sis *ShippingInstructionController) CreateShippingInstruction(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"shippinginstr:cud"}, sis.ServerOpt.Auth0Audience, sis.ServerOpt.Auth0Domain, sis.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        eblworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := eblproto.CreateShippingInstructionRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		sis.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := sis.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, eblworkflows.CreateShippingInstructionWorkflow, &form, token, user, sis.log)
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
func (sis *ShippingInstructionController) GetShippingInstructionByShippingInstructionReference(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"shippinginstr:read"}, sis.ServerOpt.Auth0Audience, sis.ServerOpt.Auth0Domain, sis.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("shippingInstructionRequestReference")

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

// UpdateShippingInstruction - Update ShippingInstruction
func (sis *ShippingInstructionController) UpdateShippingInstruction(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"shippinginstr:cud"}, sis.ServerOpt.Auth0Audience, sis.ServerOpt.Auth0Domain, sis.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("shippingInstructionRequestReference")

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        eblworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := eblproto.UpdateShippingInstructionByShippingInstructionReferenceRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
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
	result := wHelper.StartWorkflow(workflowOptions, eblworkflows.UpdateShippingInstructionWorkflow, &form, token, user, sis.log)
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
