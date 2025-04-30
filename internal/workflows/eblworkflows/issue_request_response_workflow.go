package eblworkflows

import (
	"time"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// CreateIssuanceRequestResponseWorkflow - Create IssuanceRequestResp workflow
func CreateIssuanceRequestResponseWorkflow(ctx workflow.Context, form *eblproto.CreateIssuanceRequestResponseRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateIssuanceRequestResponseResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var i *IssueRequestResponseActivities
	var issuanceRequestResponse eblproto.CreateIssuanceRequestResponseResponse
	err := workflow.ExecuteActivity(ctx, i.CreateIssuanceRequestResponseActivity, form, tokenString, user, log).Get(ctx, &issuanceRequestResponse)
	if err != nil {
		logger.Error("Failed to CreateIssuanceRequestResponseWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &issuanceRequestResponse, nil
}

// UpdateIssuanceRequestResponseWorkflow - update IssuanceRequestResponse workflow
func UpdateIssuanceRequestResponseWorkflow(ctx workflow.Context, form *eblproto.UpdateIssuanceRequestResponseRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var i *IssueRequestResponseActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, i.UpdateIssuanceRequestResponseActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateIssuanceRequestResponseWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
