// https://github.com/dcsaorg/DCSA-EBL/blob/master/src/main/java/org/dcsa/ebl/controller/TransportDocumentSummariesController.java
package eblcontrollers

import (
	"net/http"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
)

// TransportDocumentSummaryController - Create TransportDocumentSummary Controller
type TransportDocumentSummaryController struct {
	log                            *zap.Logger
	UserServiceClient              partyproto.UserServiceClient
	TransportDocumentServiceClient eblproto.TransportDocumentServiceClient
	ServerOpt                      *config.ServerOptions
}

// NewTransportDocumentSummaryController - Create TransportDocumentSummary Handler
func NewTransportDocumentSummaryController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, transportDocumentServiceClient eblproto.TransportDocumentServiceClient, serverOpt *config.ServerOptions) *TransportDocumentSummaryController {
	return &TransportDocumentSummaryController{
		log:                            log,
		UserServiceClient:              userServiceClient,
		TransportDocumentServiceClient: transportDocumentServiceClient,
		ServerOpt:                      serverOpt,
	}
}

// GetTransportDocumentSummaries - list TransportDocument Headers
func (td *TransportDocumentSummaryController) GetTransportDocumentSummaries(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"transportdoc:read"}, td.ServerOpt.Auth0Audience, td.ServerOpt.Auth0Domain, td.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")

	transportDocumentSummaries, err := td.TransportDocumentServiceClient.GetTransportDocumentSummaries(ctx,
		&eblproto.GetTransportDocumentSummariesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		td.log.Error("Error",
			zap.String("user", user.Email),
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, transportDocumentSummaries)
}
