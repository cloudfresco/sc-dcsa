package tntworkflows

import (
	"context"

	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type EventSubscriptionActivities struct {
	EventSubscriptionServiceClient tntproto.EventSubscriptionServiceClient
}

// CreateEventSubscriptionActivity - Create Transport Document activity
func (ss *EventSubscriptionActivities) CreateEventSubscriptionActivity(ctx context.Context, form *tntproto.CreateEventSubscriptionRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*tntproto.CreateEventSubscriptionResponse, error) {
	eventSubscriptionServiceClient := ss.EventSubscriptionServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	eventSubscription, err := eventSubscriptionServiceClient.CreateEventSubscription(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return eventSubscription, nil
}
