// https://github.com/dcsaorg/DCSA-BKG/blob/master/src/main/java/org/dcsa/bkg/controller/BKGBkgShipmentSummariesController.java

package bkgcontrollers

import (
	"net/http"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
)

// BkgShipmentSummaryController - Create Booking Controller
type BkgShipmentSummaryController struct {
	log                             *zap.Logger
	UserServiceClient               partyproto.UserServiceClient
	BkgShipmentSummaryServiceClient bkgproto.BkgShipmentSummaryServiceClient
	ServerOpt                       *config.ServerOptions
}

// NewBkgShipmentSummaryController - Create Booking Handler
func NewBkgShipmentSummaryController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, bkgShipmentSummaryServiceClient bkgproto.BkgShipmentSummaryServiceClient, serverOpt *config.ServerOptions) *BkgShipmentSummaryController {
	return &BkgShipmentSummaryController{
		log:                             log,
		UserServiceClient:               userServiceClient,
		BkgShipmentSummaryServiceClient: bkgShipmentSummaryServiceClient,
		ServerOpt:                       serverOpt,
	}
}

// GetBkgShipmentSummaries - list Shipment Summaries
func (bss *BkgShipmentSummaryController) GetBkgShipmentSummaries(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"bkgss:read"}, bss.ServerOpt.Auth0Audience, bss.ServerOpt.Auth0Domain, bss.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")

	bkgShipmentSummaries, err := bss.BkgShipmentSummaryServiceClient.GetBkgShipmentSummaries(ctx,
		&bkgproto.GetBkgShipmentSummariesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		bss.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, bkgShipmentSummaries)
}
