// https://github.com/dcsaorg/DCSA-EBL/blob/master/src/main/java/org/dcsa/ebl/controller/TransportDocumentController.java

package eblcontrollers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	"github.com/cloudfresco/sc-dcsa/internal/workflows/eblworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// TransportDocumentController - Create TransportDocument Controller
type TransportDocumentController struct {
	log                            *zap.Logger
	UserServiceClient              partyproto.UserServiceClient
	TransportDocumentServiceClient eblproto.TransportDocumentServiceClient
	wfHelper                       common.WfHelper
	workflowClient                 client.Client
}

// NewTransportDocumentController - Create TransportDocument Handler
func NewTransportDocumentController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, transportDocumentServiceClient eblproto.TransportDocumentServiceClient, wfHelper common.WfHelper, workflowClient client.Client) *TransportDocumentController {
	return &TransportDocumentController{
		log:                            log,
		UserServiceClient:              userServiceClient,
		TransportDocumentServiceClient: transportDocumentServiceClient,
		wfHelper:                       wfHelper,
		workflowClient:                 workflowClient,
	}
}

// ServeHTTP - parse url and call controller action
func (td *TransportDocumentController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := common.GetAuthData(r)
	cdata := partyproto.GetAuthUserDetailsRequest{}
	cdata.TokenString = data.TokenString
	cdata.Email = data.Email

	md := metadata.Pairs("authorization", "Bearer "+cdata.TokenString)
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	user, err := td.UserServiceClient.GetAuthUserDetails(ctx, &cdata)
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
		td.processGet(ctx, w, r, user, pathParts, queryString)
	case http.MethodPost:
		td.processPost(ctx, w, r, user, pathParts, data.TokenString)
	case http.MethodPut:
		td.processPut(ctx, w, r, user, pathParts)
	case http.MethodPatch:
		td.processPatch(ctx, w, r, user, pathParts, data.TokenString)
	default:
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processGet - Parse URL for all the GET paths and call the controller action
/*
	GET    "/v2/transport_documents/{transportDocumentRequestReference}"
*/

func (td *TransportDocumentController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
	if len(pathParts) == 3 {
		// GET    "/v2/transport_documents/:id"
		td.GetTransportDocumentByTransportDocumentReference(ctx, w, r, pathParts[2], user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPost - Parse URL for all the POST paths and call the controller action
/*
	POST "/v2/transport_documents"

*/
func (td *TransportDocumentController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
	if (len(pathParts) == 2) && (pathParts[1] == "transport_documents") {
		td.CreateTransportDocument(ctx, w, r, user, tokenString)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPut - Parse URL for all the put paths and call the controller action

func (td *TransportDocumentController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// processPatch - Parse URL for all the patch paths and call the controller action
/*
  PATCH "/v2/transport-documents/{transportDocumentReference}"
*/

func (td *TransportDocumentController) processPatch(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
	if len(pathParts) == 3 {
		td.ApproveTransportDocument(ctx, w, r, pathParts[2], user, tokenString)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// CreateTransportDocument - Create TransportDocument Header
func (td *TransportDocumentController) CreateTransportDocument(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, tokenString string) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        eblworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := eblproto.CreateTransportDocumentRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		td.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := td.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, eblworkflows.CreateTransportDocumentWorkflow, &form, tokenString, user, td.log)
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
func (td *TransportDocumentController) GetTransportDocumentByTransportDocumentReference(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
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
}

// ApproveTransportDocument - ApproveTransportDocument
func (td *TransportDocumentController) ApproveTransportDocument(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse, tokenString string) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        eblworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := eblproto.ApproveTransportDocumentRequest{TransportDocumentReference: id, UserEmail: user.Email, RequestId: user.RequestId}
	wHelper := td.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, eblworkflows.ApproveTransportDocumentWorkflow, &form, tokenString, user, td.log)
	workflowClient := td.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var response string
	err := workflowRun.Get(ctx, &response)
	if err != nil {
		td.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, response)
}
