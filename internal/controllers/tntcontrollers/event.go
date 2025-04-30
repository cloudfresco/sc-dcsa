package tntcontrollers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	tntworkflows "github.com/cloudfresco/sc-dcsa/internal/workflows/tntworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)

// EventController - Create Event Controller
type EventController struct {
	log                *zap.Logger
	UserServiceClient  partyproto.UserServiceClient
	EventServiceClient tntproto.EventServiceClient
	wfHelper           common.WfHelper
	workflowClient     client.Client
	ServerOpt          *config.ServerOptions
}

// NewEventController - Create Event Handler
func NewEventController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, eventServiceClient tntproto.EventServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *EventController {
	return &EventController{
		log:                log,
		UserServiceClient:  userServiceClient,
		EventServiceClient: eventServiceClient,
		wfHelper:           wfHelper,
		workflowClient:     workflowClient,
		ServerOpt:          serverOpt,
	}
}

// CreateEquipmentEvent - Create EquipmentEvent
func (ec *EventController) CreateEquipmentEvent(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"event:cud"}, ec.ServerOpt.Auth0Audience, ec.ServerOpt.Auth0Domain, ec.UserServiceClient)
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

	form := tntproto.CreateEquipmentEventRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		ec.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := ec.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, tntworkflows.CreateEquipmentEventWorkflow, &form, token, user, ec.log)
	workflowClient := ec.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var equipmentEvent tntproto.CreateEquipmentEventResponse
	err = workflowRun.Get(ctx, &equipmentEvent)
	if err != nil {
		ec.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, equipmentEvent)
}

// CreateOperationsEvent - Create OperationsEvent
func (ec *EventController) CreateOperationsEvent(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"event:cud"}, ec.ServerOpt.Auth0Audience, ec.ServerOpt.Auth0Domain, ec.UserServiceClient)
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

	form := tntproto.CreateOperationsEventRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		ec.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := ec.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, tntworkflows.CreateOperationsEventWorkflow, &form, token, user, ec.log)
	workflowClient := ec.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var operationsEvent tntproto.CreateOperationsEventResponse
	err = workflowRun.Get(ctx, &operationsEvent)
	if err != nil {
		ec.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, operationsEvent)
}

// CreateShipmentEvent - Create ShipmentEvent
func (ec *EventController) CreateShipmentEvent(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"event:cud"}, ec.ServerOpt.Auth0Audience, ec.ServerOpt.Auth0Domain, ec.UserServiceClient)
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

	form := tntproto.CreateShipmentEventRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		ec.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := ec.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, tntworkflows.CreateShipmentEventWorkflow, &form, token, user, ec.log)
	workflowClient := ec.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var shipmentEvent tntproto.CreateShipmentEventResponse
	err = workflowRun.Get(ctx, &shipmentEvent)
	if err != nil {
		ec.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, shipmentEvent)
}

// CreateTransportEvent - Create TransportEvent
func (ec *EventController) CreateTransportEvent(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"event:cud"}, ec.ServerOpt.Auth0Audience, ec.ServerOpt.Auth0Domain, ec.UserServiceClient)
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

	form := tntproto.CreateTransportEventRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		ec.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := ec.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, tntworkflows.CreateTransportEventWorkflow, &form, token, user, ec.log)
	workflowClient := ec.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var transportEvent tntproto.CreateTransportEventResponse
	err = workflowRun.Get(ctx, &transportEvent)
	if err != nil {
		ec.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, transportEvent)
}
