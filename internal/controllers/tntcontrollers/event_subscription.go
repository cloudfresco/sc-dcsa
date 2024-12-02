// https://github.com/dcsaorg/DCSA-TNT/blob/master/tnt-application/src/main/java/org/dcsa/tnt/controller/EventSubscriptionController.java

package tntcontrollers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	tntworkflows "github.com/cloudfresco/sc-dcsa/internal/workflows/tntworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// EventSubscriptionController - Create EventSubscription Controller
type EventSubscriptionController struct {
	log                            *zap.Logger
	UserServiceClient              partyproto.UserServiceClient
	EventSubscriptionServiceClient tntproto.EventSubscriptionServiceClient
	wfHelper                       common.WfHelper
	workflowClient                 client.Client
}

// NewEventSubscriptionController - Create EventSubscription Handler
func NewEventSubscriptionController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, eventSubscriptionServiceClient tntproto.EventSubscriptionServiceClient, wfHelper common.WfHelper, workflowClient client.Client) *EventSubscriptionController {
	return &EventSubscriptionController{
		log:                            log,
		UserServiceClient:              userServiceClient,
		EventSubscriptionServiceClient: eventSubscriptionServiceClient,
		wfHelper:                       wfHelper,
		workflowClient:                 workflowClient,
	}
}

// ServeHTTP - parse url and call controller action
func (esc *EventSubscriptionController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := common.GetAuthData(r)
	cdata := partyproto.GetAuthUserDetailsRequest{}
	cdata.TokenString = data.TokenString
	cdata.Email = data.Email

	md := metadata.Pairs("authorization", "Bearer "+cdata.TokenString)
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	user, err := esc.UserServiceClient.GetAuthUserDetails(ctx, &cdata)
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
		esc.processGet(ctx, w, r, user, pathParts, queryString)
	case http.MethodPost:
		esc.processPost(ctx, w, r, user, pathParts, data.TokenString)
	case http.MethodPut:
		esc.processPut(ctx, w, r, user, pathParts)
	case http.MethodDelete:
		esc.processDelete(ctx, w, r, user, pathParts)
	default:
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processGet - Parse URL for all the GET paths and call the controller action
/*
  GET  "/v3/event-subscriptions/"
  GET  "/v3/event-subscriptions/{id}"
*/

func (esc *EventSubscriptionController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
	if (len(pathParts) == 2) && (pathParts[1] == "event-subscriptions") {
		limit := queryString.Get("limit")
		cursor := queryString.Get("cursor")
		esc.GetEventSubscriptions(ctx, w, r, limit, cursor, user)
	} else if len(pathParts) == 3 {
		if pathParts[1] == "event-subscriptions" {
			esc.GetEventSubscriptionByCarrierEventSubscriptionRequestReference(ctx, w, r, pathParts[2], user)
		} else {
			common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
			return
		}
	}
}

// processPost - Parse URL for all the POST paths and call the controller action
/*
	POST "/v3/event-subscriptions"

*/
func (esc *EventSubscriptionController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
	if (len(pathParts) == 2) && (pathParts[1] == "event-subscriptions") {
		esc.CreateEventSubscription(ctx, w, r, user, tokenString)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPut - Parse URL for all the put paths and call the controller action

func (esc *EventSubscriptionController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// processDelete - Parse URL for all the delete paths and call the controller action
/*
 DELETE  "/v3/event-subscriptions/{id}"
*/

func (esc *EventSubscriptionController) processDelete(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
	if (len(pathParts) == 3) && (pathParts[1] == "event-subscriptions") {
		esc.DeleteEventSubscription(ctx, w, r, pathParts[2], user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// GetEventSubscriptions - used to view all EventSubscription
func (esc *EventSubscriptionController) GetEventSubscriptions(ctx context.Context, w http.ResponseWriter, r *http.Request, limit string, cursor string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		eventSubscriptions, err := esc.EventSubscriptionServiceClient.GetEventSubscriptions(ctx, &tntproto.GetEventSubscriptionsRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
		if err != nil {
			esc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4016", err.Error(), 402, user.RequestId)
			return
		}
		common.RenderJSON(w, eventSubscriptions)
	}
}

// CreateEventSubscription - Create EventSubscription Header
func (esc *EventSubscriptionController) CreateEventSubscription(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, tokenString string) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        tntworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := tntproto.CreateEventSubscriptionRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		esc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := esc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, tntworkflows.CreateEventSubscriptionWorkflow, &form, tokenString, user, esc.log)
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
func (esc *EventSubscriptionController) GetEventSubscriptionByCarrierEventSubscriptionRequestReference(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		// inReq := tntproto.FindEventSubscriptionByID{}
		// inReq.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}
		c, err := esc.EventSubscriptionServiceClient.FindEventSubscriptionByID(ctx, &tntproto.FindEventSubscriptionByIDRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
		if err != nil {
			esc.log.Error("Error",
				zap.String("reqid", user.RequestId),
				zap.Error(err))
			common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
			return
		}
		common.RenderJSON(w, c)
	}
}

// DeleteEventSubscription - DeleteEventSubscription
func (esc *EventSubscriptionController) DeleteEventSubscription(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	c, err := esc.EventSubscriptionServiceClient.DeleteEventSubscriptionByID(ctx, &tntproto.DeleteEventSubscriptionByIDRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		esc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, c)
}
