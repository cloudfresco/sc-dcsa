package eblworkflows

import (
	"time"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// CreateTransportDocumentWorkflow - Create TransportDocument workflow
func CreateTransportDocumentWorkflow(ctx workflow.Context, form *eblproto.CreateTransportDocumentRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateTransportDocumentResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var t *TransportDocumentActivities
	var transportDocument eblproto.CreateTransportDocumentResponse
	err := workflow.ExecuteActivity(ctx, t.CreateTransportDocumentActivity, form, tokenString, user, log).Get(ctx, &transportDocument)
	if err != nil {
		logger.Error("Failed to CreateTransportDocumentWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &transportDocument, nil
}

// ApproveTransportDocumentWorkflow - update TransportDocument workflow
func ApproveTransportDocumentWorkflow(ctx workflow.Context, form *eblproto.ApproveTransportDocumentRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var s *TransportDocumentActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, s.ApproveTransportDocumentActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to ApproveTransportDocumentWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
