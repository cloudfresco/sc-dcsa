// https://github.com/dcsaorg/DCSA-EBL/blob/master/src/main/java/org/dcsa/ebl/controller/ShippingInstructionSummariesController.java

package eblcontrollers

import (
	"net/http"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
)

// ShippingInstructionSummaryController - Create ShippingInstructionSummary Controller
type ShippingInstructionSummaryController struct {
	log                                     *zap.Logger
	UserServiceClient                       partyproto.UserServiceClient
	ShippingInstructionSummaryServiceClient eblproto.ShippingInstructionSummaryServiceClient
	ServerOpt                               *config.ServerOptions
}

// NewShippingInstructionSummaryController - Create ShippingInstructionSummary Handler
func NewShippingInstructionSummaryController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, shippingInstructionSummaryServiceClient eblproto.ShippingInstructionSummaryServiceClient, serverOpt *config.ServerOptions) *ShippingInstructionSummaryController {
	return &ShippingInstructionSummaryController{
		log:                                     log,
		UserServiceClient:                       userServiceClient,
		ShippingInstructionSummaryServiceClient: shippingInstructionSummaryServiceClient,
		ServerOpt:                               serverOpt,
	}
}

// GetShippingInstructionSummaries - list ShippingInstruction Headers
func (sis *ShippingInstructionSummaryController) GetShippingInstructionSummaries(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"shippinginstr:read"}, sis.ServerOpt.Auth0Audience, sis.ServerOpt.Auth0Domain, sis.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")

	shippingInstructionSummaries, err := sis.ShippingInstructionSummaryServiceClient.GetShippingInstructionSummaries(ctx,
		&eblproto.GetShippingInstructionSummariesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		sis.log.Error("Error",
			zap.String("user", user.Email),
			zap.String("reqid", user.RequestId),

			zap.Error(err))
		common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, shippingInstructionSummaries)
}
