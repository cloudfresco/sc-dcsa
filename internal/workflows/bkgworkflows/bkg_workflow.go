package bkgworkflows

import (
	"time"

	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

const (
	// ApplicationName is the task list
	ApplicationName = "dcsa"
)

// CreateBookingWorkflow - Create Booking workflow
func CreateBookingWorkflow(ctx workflow.Context, form *bkgproto.CreateBookingRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*bkgproto.CreateBookingResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var b *BkgActivities
	var booking bkgproto.CreateBookingResponse
	err := workflow.ExecuteActivity(ctx, b.CreateBookingActivity, form, tokenString, user, log).Get(ctx, &booking)
	if err != nil {
		logger.Error("Failed to CreateBookingWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &booking, nil
}

// UpdateBookingWorkflow - update Booking workflow
func UpdateBookingWorkflow(ctx workflow.Context, form *bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var b *BkgActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, b.UpdateBookingActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateBookingWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}

// CancelBookingByCarrierBookingReferenceWorkflow - Cancel Booking By CarrierBookingReference workflow
func CancelBookingByCarrierBookingReferenceWorkflow(ctx workflow.Context, form *bkgproto.CancelBookingByCarrierBookingReferenceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*bkgproto.CancelBookingByCarrierBookingReferenceResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var b *BkgActivities
	var resp bkgproto.CancelBookingByCarrierBookingReferenceResponse
	err := workflow.ExecuteActivity(ctx, b.CancelBookingByCarrierBookingReferenceActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to CancelBookingByCarrierBookingReferenceWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &resp, nil
}
