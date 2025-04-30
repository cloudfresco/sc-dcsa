// https://github.com/dcsaorg/DCSA-TNT/blob/master/tnt-application/src/main/java/org/dcsa/tnt/controller/EventSubscriptionController.java

package tntcontrollers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	tntworkflows "github.com/cloudfresco/sc-dcsa/internal/workflows/tntworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)

// EventSubscriptionController - Create EventSubscription Controller
type EventSubscriptionController struct {
	log                            *zap.Logger
	UserServiceClient              partyproto.UserServiceClient
	EventSubscriptionServiceClient tntproto.EventSubscriptionServiceClient
	wfHelper                       common.WfHelper
	workflowClient                 client.Client
	ServerOpt                      *config.ServerOptions
}

// NewEventSubscriptionController - Create EventSubscription Handler
func NewEventSubscriptionController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, eventSubscriptionServiceClient tntproto.EventSubscriptionServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *EventSubscriptionController {
	return &EventSubscriptionController{
		log:                            log,
		UserServiceClient:              userServiceClient,
		EventSubscriptionServiceClient: eventSubscriptionServiceClient,
		wfHelper:                       wfHelper,
		workflowClient:                 workflowClient,
		ServerOpt:                      serverOpt,
	}
}

// GetEventSubscriptions - used to view all EventSubscription
func (esc *EventSubscriptionController) GetEventSubscriptions(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"eventsub:read"}, esc.ServerOpt.Auth0Audience, esc.ServerOpt.Auth0Domain, esc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")

	eventSubscriptions, err := esc.EventSubscriptionServiceClient.GetEventSubscriptions(ctx, &tntproto.GetEventSubscriptionsRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		esc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4016", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, eventSubscriptions)
}

// CreateEventSubscription - Create EventSubscription Header
func (esc *EventSubscriptionController) CreateEventSubscription(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"eventsub:cud"}, esc.ServerOpt.Auth0Audience, esc.ServerOpt.Auth0Domain, esc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        tntworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := tntproto.CreateEventSubscriptionRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		esc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := esc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, tntworkflows.CreateEventSubscriptionWorkflow, &form, token, user, esc.log)
	workflowClient := esc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var eventSubscription tntproto.CreateEventSubscriptionResponse
	err = workflowRun.Get(ctx, &eventSubscription)
	if err != nil {
		esc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, eventSubscription)
}

// GetEventSubscriptionByCarrierEventSubscriptionRequestReference - GetEventSubscriptionByCarrierEventSubscriptionRequestReference EventSubscription
func (esc *EventSubscriptionController) GetEventSubscriptionByCarrierEventSubscriptionRequestReference(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"eventsub:read"}, esc.ServerOpt.Auth0Audience, esc.ServerOpt.Auth0Domain, esc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	eventSubscription, err := esc.EventSubscriptionServiceClient.FindEventSubscriptionByID(ctx, &tntproto.FindEventSubscriptionByIDRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		esc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, eventSubscription)
}

// DeleteEventSubscription - DeleteEventSubscription
func (esc *EventSubscriptionController) DeleteEventSubscription(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"eventsub:cud"}, esc.ServerOpt.Auth0Audience, esc.ServerOpt.Auth0Domain, esc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	eventSubscription, err := esc.EventSubscriptionServiceClient.DeleteEventSubscriptionByID(ctx, &tntproto.DeleteEventSubscriptionByIDRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		esc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, eventSubscription)
}
