package tntworkflows

import (
	"time"

	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// CreateEquipmentEventWorkflow - Create EquipmentEvent workflow
func CreateEquipmentEventWorkflow(ctx workflow.Context, form *tntproto.CreateEquipmentEventRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*tntproto.CreateEquipmentEventResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ea *EventActivities
	var equipmentEvent tntproto.CreateEquipmentEventResponse
	err := workflow.ExecuteActivity(ctx, ea.CreateEquipmentEventActivity, form, tokenString, user, log).Get(ctx, &equipmentEvent)
	if err != nil {
		logger.Error("Failed to CreateEquipmentEventWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &equipmentEvent, nil
}

// CreateOperationsEventWorkflow - Create OperationsEvent workflow
func CreateOperationsEventWorkflow(ctx workflow.Context, form *tntproto.CreateOperationsEventRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*tntproto.CreateOperationsEventResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ea *EventActivities
	var operationsEvent tntproto.CreateOperationsEventResponse
	err := workflow.ExecuteActivity(ctx, ea.CreateOperationsEventActivity, form, tokenString, user, log).Get(ctx, &operationsEvent)
	if err != nil {
		logger.Error("Failed to CreateOperationsEventWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &operationsEvent, nil
}

// CreateShipmentEventWorkflow - Create ShipmentEvent workflow
func CreateShipmentEventWorkflow(ctx workflow.Context, form *tntproto.CreateShipmentEventRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*tntproto.CreateShipmentEventResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ea *EventActivities
	var shipmentEvent tntproto.CreateShipmentEventResponse
	err := workflow.ExecuteActivity(ctx, ea.CreateShipmentEventActivity, form, tokenString, user, log).Get(ctx, &shipmentEvent)
	if err != nil {
		logger.Error("Failed to CreateShipmentEventWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &shipmentEvent, nil
}

// CreateTransportEventWorkflow - Create TransportEvent workflow
func CreateTransportEventWorkflow(ctx workflow.Context, form *tntproto.CreateTransportEventRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*tntproto.CreateTransportEventResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ea *EventActivities
	var transportEvent tntproto.CreateTransportEventResponse
	err := workflow.ExecuteActivity(ctx, ea.CreateTransportEventActivity, form, tokenString, user, log).Get(ctx, &transportEvent)
	if err != nil {
		logger.Error("Failed to CreateTransportEventWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &transportEvent, nil
}
