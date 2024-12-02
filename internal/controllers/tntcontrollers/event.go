// https://github.com/dcsaorg/DCSA-TNT/blob/master/tnt-application/src/main/java/org/dcsa/tnt/controller/EventController.java

package tntcontrollers

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/tnt/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// EventController - Create Event Controller
type EventController struct {
	log                *zap.Logger
	UserServiceClient  partyproto.UserServiceClient
	EventServiceClient tntproto.EventServiceClient
}

// NewEventController - Create Event Handler
func NewEventController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, eventServiceClient tntproto.EventServiceClient) *EventController {
	return &EventController{
		log:                log,
		UserServiceClient:  userServiceClient,
		EventServiceClient: eventServiceClient,
	}
}

// ServeHTTP - parse url and call controller action
func (ec *EventController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := common.GetAuthData(r)
	cdata := partyproto.GetAuthUserDetailsRequest{}
	cdata.TokenString = data.TokenString
	cdata.Email = data.Email

	md := metadata.Pairs("authorization", "Bearer "+cdata.TokenString)
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	user, err := ec.UserServiceClient.GetAuthUserDetails(ctx, &cdata)
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
		ec.processGet(ctx, w, r, user, pathParts, queryString)
	case http.MethodPost:
		ec.processPost(ctx, w, r, user, pathParts)
	case http.MethodPut:
		ec.processPut(ctx, w, r, user, pathParts)
	default:
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processGet - Parse URL for all the GET paths and call the controller action
/*
	GET    "/v3/events"
  GET    "/v3/events/{id}"
*/

func (ec *EventController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
	if len(pathParts) == 2 && (pathParts[1] == "events") {
		// GET    "/v3/events"
		limit := queryString.Get("limit")
		cursor := queryString.Get("cursor")
		ec.GetEvents(ctx, w, r, limit, cursor, user)
	} else if len(pathParts) == 3 {
		if pathParts[1] == "events" {
			ec.GetEvent(ctx, w, r, pathParts[2], user)
		} else {
			common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
			return
		}
	}
}

// processPost - Parse URL for all the POST paths and call the controller action

func (ec *EventController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// processPut - Parse URL for all the put paths and call the controller action

func (ec *EventController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
}

// GetEvents - list Event Headers
func (ec *EventController) GetEvents(ctx context.Context, w http.ResponseWriter, r *http.Request, limit string, cursor string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		events, err := ec.EventServiceClient.GetEvents(ctx,
			&tntproto.GetEventsRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
		if err != nil {
			ec.log.Error("Error",
				zap.String("user", user.Email),
				zap.String("reqid", user.RequestId),
				zap.Error(err))
			common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
			return
		}

		common.RenderJSON(w, events)
	}
}

func (ec *EventController) GetEvent(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		event, err := ec.EventServiceClient.GetEvent(ctx, &tntproto.GetEventRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
		if err != nil {
			ec.log.Error("Error",
				zap.String("user", user.Email),
				zap.String("reqid", user.RequestId),
				zap.Error(err))
			common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
			return
		}
		common.RenderJSON(w, event)
	}
}
