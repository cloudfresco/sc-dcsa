package eblworkflows

import (
	"time"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// CreateSurrenderRequestAnswerWorkflow - Create SurrenderRequestAnswer workflow
func CreateSurrenderRequestAnswerWorkflow(ctx workflow.Context, form *eblproto.CreateSurrenderRequestAnswerRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateSurrenderRequestAnswerResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *SurrenderRequestAnswerActivities
	var surrenderRequestAnswer eblproto.CreateSurrenderRequestAnswerResponse
	err := workflow.ExecuteActivity(ctx, sa.CreateSurrenderRequestAnswerActivity, form, tokenString, user, log).Get(ctx, &surrenderRequestAnswer)
	if err != nil {
		logger.Error("Failed to CreateSurrenderRequestAnswerWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &surrenderRequestAnswer, nil
}

// UpdateSurrenderRequestAnswerWorkflow - update SurrenderRequestAnswer workflow
func UpdateSurrenderRequestAnswerWorkflow(ctx workflow.Context, form *eblproto.UpdateSurrenderRequestAnswerRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var sa *SurrenderRequestAnswerActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, sa.UpdateSurrenderRequestAnswerActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateSurrenderRequestAnswerWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
