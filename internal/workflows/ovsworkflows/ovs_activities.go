package ovsworkflows

import (
	"context"

	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type ServiceScheduleActivities struct {
	ServiceScheduleServiceClient ovsproto.ServiceScheduleServiceClient
}

// CreateServiceScheduleActivity - Create Transport Document activity
func (ss *ServiceScheduleActivities) CreateServiceScheduleActivity(ctx context.Context, form *ovsproto.CreateServiceScheduleRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*ovsproto.CreateServiceScheduleResponse, error) {
	serviceScheduleServiceClient := ss.ServiceScheduleServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	serviceSchedule, err := serviceScheduleServiceClient.CreateServiceSchedule(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return serviceSchedule, nil
}

// UpdateServiceScheduleActivity - update ServiceSchedule activity
func (ss *ServiceScheduleActivities) UpdateServiceScheduleActivity(ctx context.Context, form *ovsproto.UpdateServiceScheduleByUniversalServiceReferenceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	serviceScheduleServiceClient := ss.ServiceScheduleServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := serviceScheduleServiceClient.UpdateServiceScheduleByUniversalServiceReference(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}
