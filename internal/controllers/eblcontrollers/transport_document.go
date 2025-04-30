// https://github.com/dcsaorg/DCSA-EBL/blob/master/src/main/java/org/dcsa/ebl/controller/TransportDocumentController.java

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

// TransportDocumentController - Create TransportDocument Controller
type TransportDocumentController struct {
	log                            *zap.Logger
	UserServiceClient              partyproto.UserServiceClient
	TransportDocumentServiceClient eblproto.TransportDocumentServiceClient
	wfHelper                       common.WfHelper
	workflowClient                 client.Client
	ServerOpt                      *config.ServerOptions
}

// NewTransportDocumentController - Create TransportDocument Handler
func NewTransportDocumentController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, transportDocumentServiceClient eblproto.TransportDocumentServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *TransportDocumentController {
	return &TransportDocumentController{
		log:                            log,
		UserServiceClient:              userServiceClient,
		TransportDocumentServiceClient: transportDocumentServiceClient,
		wfHelper:                       wfHelper,
		workflowClient:                 workflowClient,
		ServerOpt:                      serverOpt,
	}
}

// CreateTransportDocument - Create TransportDocument Header
func (td *TransportDocumentController) CreateTransportDocument(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"transportdoc:cud"}, td.ServerOpt.Auth0Audience, td.ServerOpt.Auth0Domain, td.UserServiceClient)
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

	form := eblproto.CreateTransportDocumentRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		td.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := td.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, eblworkflows.CreateTransportDocumentWorkflow, &form, token, user, td.log)
	workflowClient := td.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var transportDocument eblproto.CreateTransportDocumentResponse
	err = workflowRun.Get(ctx, &transportDocument)
	if err != nil {
		td.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, transportDocument)
}

// GetTransportDocumentByTransportDocumentReference - GetTransportDocumentByTransportDocumentReference TransportDocument
func (td *TransportDocumentController) GetTransportDocumentByTransportDocumentReference(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"transportdoc:read"}, td.ServerOpt.Auth0Audience, td.ServerOpt.Auth0Domain, td.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("transportDocumentRequestReference")

	transportDocument, err := td.TransportDocumentServiceClient.FindByTransportDocumentReference(ctx, &eblproto.FindByTransportDocumentReferenceRequest{TransportDocumentReference: id, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		td.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, transportDocument)
}

// ApproveTransportDocument - ApproveTransportDocument
func (td *TransportDocumentController) ApproveTransportDocument(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"transportdoc:cud"}, td.ServerOpt.Auth0Audience, td.ServerOpt.Auth0Domain, td.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("transportDocumentRequestReference")

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        eblworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := eblproto.ApproveTransportDocumentRequest{TransportDocumentReference: id, UserEmail: user.Email, RequestId: user.RequestId}
	wHelper := td.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, eblworkflows.ApproveTransportDocumentWorkflow, &form, token, user, td.log)
	workflowClient := td.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var response string
	err = workflowRun.Get(ctx, &response)
	if err != nil {
		td.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, response)
}
