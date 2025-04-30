package tntworkflows

import (
	"context"

	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type EventActivities struct {
	EventServiceClient tntproto.EventServiceClient
}

// CreateEquipmentEventActivity - Create EquipmentEvent activity
func (ea *EventActivities) CreateEquipmentEventActivity(ctx context.Context, form *tntproto.CreateEquipmentEventRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*tntproto.CreateEquipmentEventResponse, error) {
	eventServiceClient := ea.EventServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	equipmentEvent, err := eventServiceClient.CreateEquipmentEvent(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return equipmentEvent, nil
}

// CreateOperationsEventActivity - Create OperationsEvent activity
func (ea *EventActivities) CreateOperationsEventActivity(ctx context.Context, form *tntproto.CreateOperationsEventRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*tntproto.CreateOperationsEventResponse, error) {
	eventServiceClient := ea.EventServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	operationsEvent, err := eventServiceClient.CreateOperationsEvent(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return operationsEvent, nil
}

// CreateShipmentEventActivity - Create ShipmentEvent activity
func (ea *EventActivities) CreateShipmentEventActivity(ctx context.Context, form *tntproto.CreateShipmentEventRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*tntproto.CreateShipmentEventResponse, error) {
	eventServiceClient := ea.EventServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	shipmentEvent, err := eventServiceClient.CreateShipmentEvent(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return shipmentEvent, nil
}

// CreateTransportEventActivity - Create TransportEvent activity
func (ea *EventActivities) CreateTransportEventActivity(ctx context.Context, form *tntproto.CreateTransportEventRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*tntproto.CreateTransportEventResponse, error) {
	eventServiceClient := ea.EventServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	transportEvent, err := eventServiceClient.CreateTransportEvent(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return transportEvent, nil
}
