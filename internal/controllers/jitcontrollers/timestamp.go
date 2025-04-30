// https://github.com/dcsaorg/DCSA-JIT/blob/master/jit-application/src/main/java/org/dcsa/jit/controller/TimestampController.java

package jitcontrollers

import (
	"encoding/json"
	"net/http"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)

// TimestampController - Create Timestamp Controller
type TimestampController struct {
	log                    *zap.Logger
	UserServiceClient      partyproto.UserServiceClient
	TimestampServiceClient jitproto.TimestampServiceClient
	wfHelper               common.WfHelper
	workflowClient         client.Client
	ServerOpt              *config.ServerOptions
}

// NewTimestampController - Create Timestamp Handler
func NewTimestampController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, timestampServiceClient jitproto.TimestampServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *TimestampController {
	return &TimestampController{
		log:                    log,
		UserServiceClient:      userServiceClient,
		TimestampServiceClient: timestampServiceClient,
		wfHelper:               wfHelper,
		workflowClient:         workflowClient,
		ServerOpt:              serverOpt,
	}
}

// CreateTimestamp - Create Timestamp
func (ts *TimestampController) CreateTimestamp(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"ts:cud"}, ts.ServerOpt.Auth0Audience, ts.ServerOpt.Auth0Domain, ts.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	form := jitproto.CreateTimestampRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		ts.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	timestamp, err := ts.TimestampServiceClient.CreateTimestamp(ctx, &form)
	if err != nil {
		ts.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, timestamp)
}
