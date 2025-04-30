package eblworkflows

import (
	"time"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// CreateTransactionPartyWorkflow - Create TransactionParty workflow
func CreateTransactionPartyWorkflow(ctx workflow.Context, form *eblproto.CreateTransactionPartyRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateTransactionPartyResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *SurrenderRequestActivities
	var transactionParty eblproto.CreateTransactionPartyResponse
	err := workflow.ExecuteActivity(ctx, sa.CreateTransactionPartyActivity, form, tokenString, user, log).Get(ctx, &transactionParty)
	if err != nil {
		logger.Error("Failed to CreateTransactionPartyWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &transactionParty, nil
}

// UpdateTransactionPartyWorkflow - update TransactionParty workflow
func UpdateTransactionPartyWorkflow(ctx workflow.Context, form *eblproto.UpdateTransactionPartyRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *SurrenderRequestActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, sa.UpdateTransactionPartyActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateTransactionPartyWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}

// CreateTransactionPartySupportingCodeWorkflow - Create TransactionPartySupportingCode workflow
func CreateTransactionPartySupportingCodeWorkflow(ctx workflow.Context, form *eblproto.CreateTransactionPartySupportingCodeRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateTransactionPartySupportingCodeResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *SurrenderRequestActivities
	var transactionPartySupportingCode eblproto.CreateTransactionPartySupportingCodeResponse
	err := workflow.ExecuteActivity(ctx, sa.CreateTransactionPartySupportingCodeActivity, form, tokenString, user, log).Get(ctx, &transactionPartySupportingCode)
	if err != nil {
		logger.Error("Failed to CreateTransactionPartySupportingCodeWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &transactionPartySupportingCode, nil
}

// UpdateTransactionPartySupportingCodeWorkflow - update TransactionPartySupportingCode workflow
func UpdateTransactionPartySupportingCodeWorkflow(ctx workflow.Context, form *eblproto.UpdateTransactionPartySupportingCodeRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *SurrenderRequestActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, sa.UpdateTransactionPartySupportingCodeActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateTransactionPartySupportingCodeWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}

// CreateSurrenderRequestWorkflow - Create SurrenderRequest workflow
func CreateSurrenderRequestWorkflow(ctx workflow.Context, form *eblproto.CreateSurrenderRequestRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateSurrenderRequestResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *SurrenderRequestActivities
	var surrenderRequest eblproto.CreateSurrenderRequestResponse
	err := workflow.ExecuteActivity(ctx, sa.CreateSurrenderRequestActivity, form, tokenString, user, log).Get(ctx, &surrenderRequest)
	if err != nil {
		logger.Error("Failed to CreateSurrenderRequestWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &surrenderRequest, nil
}

// UpdateSurrenderRequestWorkflow - update SurrenderRequest workflow
func UpdateSurrenderRequestWorkflow(ctx workflow.Context, form *eblproto.UpdateSurrenderRequestRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *SurrenderRequestActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, sa.UpdateSurrenderRequestActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateSurrenderRequestWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}

// CreateEndorsementChainLinkWorkflow - Create EndorsementChainLink workflow
func CreateEndorsementChainLinkWorkflow(ctx workflow.Context, form *eblproto.CreateEndorsementChainLinkRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateEndorsementChainLinkResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *SurrenderRequestActivities
	var endorsementChainLink eblproto.CreateEndorsementChainLinkResponse
	err := workflow.ExecuteActivity(ctx, sa.CreateEndorsementChainLinkActivity, form, tokenString, user, log).Get(ctx, &endorsementChainLink)
	if err != nil {
		logger.Error("Failed to CreateEndorsementChainLinkWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &endorsementChainLink, nil
}

// UpdateEndorsementChainLinkWorkflow - update EndorsementChainLink workflow
func UpdateEndorsementChainLinkWorkflow(ctx workflow.Context, form *eblproto.UpdateEndorsementChainLinkRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *SurrenderRequestActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, sa.UpdateEndorsementChainLinkActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateEndorsementChainLinkWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
