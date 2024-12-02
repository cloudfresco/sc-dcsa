package jitworkflows

import (
	"time"

	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

const (
	// ApplicationName is the task list
	ApplicationName = "dcsa"
)

// CreateTimestampWorkflow - Create Timestamp workflow
func CreateTimestampWorkflow(ctx workflow.Context, form *jitproto.CreateTimestampRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*jitproto.CreateTimestampResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var t *TimestampActivities
	var tstamp jitproto.CreateTimestampResponse
	err := workflow.ExecuteActivity(ctx, t.CreateTimestampActivity, form, tokenString, user, log).Get(ctx, &tstamp)
	if err != nil {
		logger.Error("Failed to CreateTimestampWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &tstamp, nil
}
