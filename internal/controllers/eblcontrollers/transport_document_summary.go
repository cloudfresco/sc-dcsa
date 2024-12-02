// https://github.com/dcsaorg/DCSA-EBL/blob/master/src/main/java/org/dcsa/ebl/controller/TransportDocumentSummariesController.java
package eblcontrollers

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// TransportDocumentSummaryController - Create TransportDocumentSummary Controller
type TransportDocumentSummaryController struct {
	log                            *zap.Logger
	UserServiceClient              partyproto.UserServiceClient
	TransportDocumentServiceClient eblproto.TransportDocumentServiceClient
}

// NewTransportDocumentSummaryController - Create TransportDocumentSummary Handler
func NewTransportDocumentSummaryController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, transportDocumentServiceClient eblproto.TransportDocumentServiceClient) *TransportDocumentSummaryController {
	return &TransportDocumentSummaryController{
		log:                            log,
		UserServiceClient:              userServiceClient,
		TransportDocumentServiceClient: transportDocumentServiceClient,
	}
}

// ServeHTTP - parse url and call controller action
func (cc *TransportDocumentSummaryController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := common.GetAuthData(r)
	cdata := partyproto.GetAuthUserDetailsRequest{}
	cdata.TokenString = data.TokenString
	cdata.Email = data.Email

	md := metadata.Pairs("authorization", "Bearer "+cdata.TokenString)
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	user, err := cc.UserServiceClient.GetAuthUserDetails(ctx, &cdata)
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
		cc.processGet(ctx, w, r, user, pathParts, queryString)
	case http.MethodPost:
		cc.processPost(ctx, w, r, user, pathParts)
	case http.MethodPut:
		cc.processPut(ctx, w, r, user, pathParts)
	default:
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processGet - Parse URL for all the GET paths and call the controller action
/*
	GET    "/v2/transport-document-summaries"
*/

func (cc *TransportDocumentSummaryController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
	if len(pathParts) == 2 && (pathParts[1] == "transport-document-summaries") {
		// GET    "/v2/transport-document-summaries"
		limit := queryString.Get("limit")
		cursor := queryString.Get("cursor")
		cc.GetTransportDocumentSummaries(ctx, w, r, limit, cursor, user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPost - Parse URL for all the POST paths and call the controller action

func (cc *TransportDocumentSummaryController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// processPut - Parse URL for all the put paths and call the controller action

func (cc *TransportDocumentSummaryController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// GetTransportDocumentSummaries - list TransportDocument Headers
func (cc *TransportDocumentSummaryController) GetTransportDocumentSummaries(ctx context.Context, w http.ResponseWriter, r *http.Request, limit string, cursor string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		c, err := cc.TransportDocumentServiceClient.GetTransportDocumentSummaries(ctx,
			&eblproto.GetTransportDocumentSummariesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
		if err != nil {
			cc.log.Error("Error",
				zap.String("user", user.Email),
				zap.String("reqid", user.RequestId),
				zap.Error(err))
			common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
			return
		}

		common.RenderJSON(w, c)
	}
}
