package eblworkflows

import (
	"time"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

const (
	// ApplicationName is the task list
	ApplicationName = "dcsa"
)

// CreateShippingInstructionWorkflow - Create ShippingInstruction workflow
func CreateShippingInstructionWorkflow(ctx workflow.Context, form *eblproto.CreateShippingInstructionRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateShippingInstructionResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var s *ShippingInstructionActivities

	var shippingInstruction eblproto.CreateShippingInstructionResponse

	err := workflow.ExecuteActivity(ctx, s.CreateShippingInstructionActivity, form, tokenString, user, log).Get(ctx, &shippingInstruction)
	if err != nil {
		logger.Error("Failed to CreateShippingInstructionWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &shippingInstruction, nil
}

// UpdateShippingInstructionWorkflow - update ShippingInstruction workflow
func UpdateShippingInstructionWorkflow(ctx workflow.Context, form *eblproto.UpdateShippingInstructionByShippingInstructionReferenceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var s *ShippingInstructionActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, s.UpdateShippingInstructionActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateShippingInstructionWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
