package tntworkflows

import (
	"time"

	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

const (
	// ApplicationName is the task list
	ApplicationName = "dcsa"
)

// CreateEventSubscriptionWorkflow - Create EventSubscription workflow
func CreateEventSubscriptionWorkflow(ctx workflow.Context, form *tntproto.CreateEventSubscriptionRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*tntproto.CreateEventSubscriptionResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var esa *EventSubscriptionActivities
	var eventSubscription tntproto.CreateEventSubscriptionResponse
	err := workflow.ExecuteActivity(ctx, esa.CreateEventSubscriptionActivity, form, tokenString, user, log).Get(ctx, &eventSubscription)
	if err != nil {
		logger.Error("Failed to CreateEventSubscriptionWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &eventSubscription, nil
}
