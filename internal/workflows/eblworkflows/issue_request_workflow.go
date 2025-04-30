package eblworkflows

import (
	"time"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// CreateIssuePartyWorkflow - Create IssueParty workflow
func CreateIssuePartyWorkflow(ctx workflow.Context, form *eblproto.CreateIssuePartyRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateIssuePartyResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var i *IssueRequestActivities
	var issueParty eblproto.CreateIssuePartyResponse
	err := workflow.ExecuteActivity(ctx, i.CreateIssuePartyActivity, form, tokenString, user, log).Get(ctx, &issueParty)
	if err != nil {
		logger.Error("Failed to CreateIssuePartyWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &issueParty, nil
}

// UpdateIssuePartyWorkflow - update IssueParty workflow
func UpdateIssuePartyWorkflow(ctx workflow.Context, form *eblproto.UpdateIssuePartyRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var i *IssueRequestActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, i.UpdateIssuePartyActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateIssuePartyWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}

// CreateIssuePartySupportingCodeWorkflow - Create IssuePartySupportingCode workflow
func CreateIssuePartySupportingCodeWorkflow(ctx workflow.Context, form *eblproto.CreateIssuePartySupportingCodeRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateIssuePartySupportingCodeResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var i *IssueRequestActivities
	var issuePartySupportingCode eblproto.CreateIssuePartySupportingCodeResponse
	err := workflow.ExecuteActivity(ctx, i.CreateIssuePartySupportingCodeActivity, form, tokenString, user, log).Get(ctx, &issuePartySupportingCode)
	if err != nil {
		logger.Error("Failed to CreateIssuePartySupportingCodeWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &issuePartySupportingCode, nil
}

// UpdateIssuePartySupportingCodeWorkflow - update IssuePartySupportingCode workflow
func UpdateIssuePartySupportingCodeWorkflow(ctx workflow.Context, form *eblproto.UpdateIssuePartySupportingCodeRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var i *IssueRequestActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, i.UpdateIssuePartySupportingCodeActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateIssuePartySupportingCodeWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}

// CreateIssuanceRequestWorkflow - Create IssuanceRequest workflow
func CreateIssuanceRequestWorkflow(ctx workflow.Context, form *eblproto.CreateIssuanceRequestRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateIssuanceRequestResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var i *IssueRequestActivities
	var issuanceRequest eblproto.CreateIssuanceRequestResponse
	err := workflow.ExecuteActivity(ctx, i.CreateIssuanceRequestActivity, form, tokenString, user, log).Get(ctx, &issuanceRequest)
	if err != nil {
		logger.Error("Failed to CreateIssuanceRequestWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &issuanceRequest, nil
}

// UpdateIssuanceRequestWorkflow - update IssuanceRequest workflow
func UpdateIssuanceRequestWorkflow(ctx workflow.Context, form *eblproto.UpdateIssuanceRequestRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var i *IssueRequestActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, i.UpdateIssuanceRequestActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateIssuanceRequestWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}

// CreateEblVisualizationWorkflow - Create EblVisualization workflow
func CreateEblVisualizationWorkflow(ctx workflow.Context, form *eblproto.CreateEblVisualizationRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateEblVisualizationResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var i *IssueRequestActivities
	var eblVisualization eblproto.CreateEblVisualizationResponse
	err := workflow.ExecuteActivity(ctx, i.CreateEblVisualizationActivity, form, tokenString, user, log).Get(ctx, &eblVisualization)
	if err != nil {
		logger.Error("Failed to CreateEblVisualizationWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &eblVisualization, nil
}

// UpdateEblVisualizationWorkflow - update EblVisualization workflow
func UpdateEblVisualizationWorkflow(ctx workflow.Context, form *eblproto.UpdateEblVisualizationRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var i *IssueRequestActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, i.UpdateEblVisualizationActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateEblVisualizationWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
