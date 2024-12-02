// https://github.com/dcsaorg/DCSA-JIT/blob/master/jit-application/src/main/java/org/dcsa/jit/controller/TimestampController.java

package jitcontrollers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// TimestampController - Create Timestamp Controller
type TimestampController struct {
	log                    *zap.Logger
	UserServiceClient      partyproto.UserServiceClient
	TimestampServiceClient jitproto.TimestampServiceClient
	wfHelper               common.WfHelper
	workflowClient         client.Client
}

// NewTimestampController - Create Timestamp Handler
func NewTimestampController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, timestampServiceClient jitproto.TimestampServiceClient, wfHelper common.WfHelper, workflowClient client.Client) *TimestampController {
	return &TimestampController{
		log:                    log,
		UserServiceClient:      userServiceClient,
		TimestampServiceClient: timestampServiceClient,
		wfHelper:               wfHelper,
		workflowClient:         workflowClient,
	}
}

// ServeHTTP - parse url and call controller action
func (cc *TimestampController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	case http.MethodPatch:
		cc.processPatch(ctx, w, r, user, pathParts)
	case http.MethodDelete:
		cc.processDelete(ctx, w, r, user, pathParts)
	default:
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processGet - Parse URL for all the GET paths and call the controller action

func (cc *TimestampController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
}

// processPost - Parse URL for all the POST paths and call the controller action
/*
	POST "/v1/timestamps"

*/
func (cc *TimestampController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
	if (len(pathParts) == 2) && (pathParts[1] == "timestamps") {
		cc.CreateTimestamp(ctx, w, r, user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPut - Parse URL for all the put paths and call the controller action

func (cc *TimestampController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// processPatch - Parse URL for all the patch paths and call the controller action

func (cc *TimestampController) processPatch(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// processDelete - Parse URL for all the delete paths and call the controller action

func (cc *TimestampController) processDelete(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// CreateTimestamp - Create Timestamp Header
func (cc *TimestampController) CreateTimestamp(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse) {
	form := jitproto.CreateTimestampRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		cc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	if err != nil {
		cc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	timestamp, err := cc.TimestampServiceClient.CreateTimestamp(ctx, &form)
	if err != nil {
		cc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, timestamp)
}
