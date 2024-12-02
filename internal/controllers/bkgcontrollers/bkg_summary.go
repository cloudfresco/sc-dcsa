// https://github.com/dcsaorg/DCSA-BKG/blob/master/src/main/java/org/dcsa/bkg/controller/BKGSummariesController.java

package bkgcontrollers

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/bkg/v2"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// BkgSummaryController - Create Booking Controller
type BkgSummaryController struct {
	log                     *zap.Logger
	UserServiceClient       partyproto.UserServiceClient
	BkgSummaryServiceClient bkgproto.BkgSummaryServiceClient
}

// NewBkgSummaryController - Create Booking Handler
func NewBkgSummaryController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, bkgSummaryServiceClient bkgproto.BkgSummaryServiceClient) *BkgSummaryController {
	return &BkgSummaryController{
		log:                     log,
		UserServiceClient:       userServiceClient,
		BkgSummaryServiceClient: bkgSummaryServiceClient,
	}
}

// ServeHTTP - parse url and call controller action
func (cc *BkgSummaryController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	case http.MethodDelete:
		cc.processDelete(ctx, w, r, user, pathParts)
	default:
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processGet - Parse URL for all the GET paths and call the controller action
/*
	GET    "/v2/booking-summaries"
*/

func (cc *BkgSummaryController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
	if len(pathParts) == 2 && (pathParts[1] == "booking-summaries") {
		limit := queryString.Get("limit")
		cursor := queryString.Get("cursor")
		// GET    "/v2/booking-summaries"
		cc.GetBookingSummaries(ctx, w, r, limit, cursor, user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPost - Parse URL for all the POST paths and call the controller action

func (cc *BkgSummaryController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// processPut - Parse URL for all the put paths and call the controller action

func (cc *BkgSummaryController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// processDelete - Parse URL for all the delete paths and call the controller action
/*
 DELETE   ""
*/

func (cc *BkgSummaryController) processDelete(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// GetBookingSummaries - list Booking Summaries
func (cc *BkgSummaryController) GetBookingSummaries(ctx context.Context, w http.ResponseWriter, r *http.Request, limit string, cursor string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		bookingSummaries, err := cc.BkgSummaryServiceClient.GetBookingSummaries(ctx,
			&bkgproto.GetBookingSummariesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
		if err != nil {
			cc.log.Error("Error",
				zap.String("user", user.Email),
				zap.String("reqid", user.RequestId),

				zap.Error(err))
			common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
			return
		}

		common.RenderJSON(w, bookingSummaries)
	}
}
