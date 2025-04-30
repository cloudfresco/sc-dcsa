// https://github.com/dcsaorg/DCSA-BKG/blob/master/src/main/java/org/dcsa/bkg/controller/BKGSummariesController.java

package bkgcontrollers

import (
	"net/http"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
)

// BkgSummaryController - Create Booking Controller
type BkgSummaryController struct {
	log                     *zap.Logger
	UserServiceClient       partyproto.UserServiceClient
	BkgSummaryServiceClient bkgproto.BkgSummaryServiceClient
	ServerOpt               *config.ServerOptions
}

// NewBkgSummaryController - Create Booking Handler
func NewBkgSummaryController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, bkgSummaryServiceClient bkgproto.BkgSummaryServiceClient, serverOpt *config.ServerOptions) *BkgSummaryController {
	return &BkgSummaryController{
		log:                     log,
		UserServiceClient:       userServiceClient,
		BkgSummaryServiceClient: bkgSummaryServiceClient,
		ServerOpt:               serverOpt,
	}
}

// GetBookingSummaries - list Booking Summaries
func (bs *BkgSummaryController) GetBookingSummaries(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"bkgs:read"}, bs.ServerOpt.Auth0Audience, bs.ServerOpt.Auth0Domain, bs.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")

	bookingSummaries, err := bs.BkgSummaryServiceClient.GetBookingSummaries(ctx,
		&bkgproto.GetBookingSummariesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		bs.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, bookingSummaries)
}
