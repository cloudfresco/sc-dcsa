package ovsworkflows

import (
	"time"

	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

const (
	// ApplicationName is the task list
	ApplicationName = "dcsa"
)

// CreateServiceScheduleWorkflow - Create ServiceSchedule workflow
func CreateServiceScheduleWorkflow(ctx workflow.Context, form *ovsproto.CreateServiceScheduleRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*ovsproto.CreateServiceScheduleResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *ServiceScheduleActivities
	var serviceSchedule ovsproto.CreateServiceScheduleResponse
	err := workflow.ExecuteActivity(ctx, sa.CreateServiceScheduleActivity, form, tokenString, user, log).Get(ctx, &serviceSchedule)
	if err != nil {
		logger.Error("Failed to CreateServiceScheduleWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &serviceSchedule, nil
}

// UpdateServiceScheduleWorkflow - update ServiceSchedule workflow
func UpdateServiceScheduleWorkflow(ctx workflow.Context, form *ovsproto.UpdateServiceScheduleByUniversalServiceReferenceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *ServiceScheduleActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, sa.UpdateServiceScheduleActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateServiceScheduleWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
