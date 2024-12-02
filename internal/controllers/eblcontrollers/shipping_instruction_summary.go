// https://github.com/dcsaorg/DCSA-EBL/blob/master/src/main/java/org/dcsa/ebl/controller/ShippingInstructionSummariesController.java

package eblcontrollers

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// ShippingInstructionSummaryController - Create ShippingInstructionSummary Controller
type ShippingInstructionSummaryController struct {
	log                                     *zap.Logger
	UserServiceClient                       partyproto.UserServiceClient
	ShippingInstructionSummaryServiceClient eblproto.ShippingInstructionSummaryServiceClient
}

// NewShippingInstructionSummaryController - Create ShippingInstructionSummary Handler
func NewShippingInstructionSummaryController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, shippingInstructionSummaryServiceClient eblproto.ShippingInstructionSummaryServiceClient) *ShippingInstructionSummaryController {
	return &ShippingInstructionSummaryController{
		log:                                     log,
		UserServiceClient:                       userServiceClient,
		ShippingInstructionSummaryServiceClient: shippingInstructionSummaryServiceClient,
	}
}

// ServeHTTP - parse url and call controller action
func (cc *ShippingInstructionSummaryController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	GET    "/v2/shipping-instructions-summaries"
*/

func (cc *ShippingInstructionSummaryController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
	if len(pathParts) == 2 && (pathParts[1] == "shipping-instructions-summaries") {
		limit := queryString.Get("limit")
		cursor := queryString.Get("cursor")
		// GET    "/v2/shipping-instructions-summaries"
		cc.GetShippingInstructionSummaries(ctx, w, r, limit, cursor, user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPost - Parse URL for all the POST paths and call the controller action

func (cc *ShippingInstructionSummaryController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// processPut - Parse URL for all the put paths and call the controller action

func (cc *ShippingInstructionSummaryController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// GetShippingInstructionSummaries - list ShippingInstruction Headers
func (cc *ShippingInstructionSummaryController) GetShippingInstructionSummaries(ctx context.Context, w http.ResponseWriter, r *http.Request, limit string, cursor string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		shippingInstructionSummaries, err := cc.ShippingInstructionSummaryServiceClient.GetShippingInstructionSummaries(ctx,
			&eblproto.GetShippingInstructionSummariesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
		if err != nil {
			cc.log.Error("Error",
				zap.String("user", user.Email),
				zap.String("reqid", user.RequestId),

				zap.Error(err))
			common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
			return
		}

		common.RenderJSON(w, shippingInstructionSummaries)
	}
}
