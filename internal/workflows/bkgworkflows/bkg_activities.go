package bkgworkflows

import (
	"context"

	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type BkgActivities struct {
	BkgServiceClient bkgproto.BkgServiceClient
}

// CreateBookingActivity - Create Booking activity
func (b *BkgActivities) CreateBookingActivity(ctx context.Context, form *bkgproto.CreateBookingRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*bkgproto.CreateBookingResponse, error) {
	bkgServiceClient := b.BkgServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	booking, err := bkgServiceClient.CreateBooking(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return booking, nil
}

// UpdateBookingActivity - update Booking activity
func (b *BkgActivities) UpdateBookingActivity(ctx context.Context, form *bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	bkgServiceClient := b.BkgServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := bkgServiceClient.UpdateBookingByReferenceCarrierBookingRequestReference(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}

// CancelBookingByCarrierBookingReferenceActivity - Cancel Booking By CarrierBookingRequestReference activity
func (b *BkgActivities) CancelBookingByCarrierBookingReferenceActivity(ctx context.Context, form *bkgproto.CancelBookingByCarrierBookingReferenceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*bkgproto.CancelBookingByCarrierBookingReferenceResponse, error) {
	bkgServiceClient := b.BkgServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	resp, err := bkgServiceClient.CancelBookingByCarrierBookingReference(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return resp, nil
}
