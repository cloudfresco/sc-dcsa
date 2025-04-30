package jitworkflows

import (
	"context"

	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type TimestampActivities struct {
	TimestampServiceClient jitproto.TimestampServiceClient
}

// CreateTimestampActivity - Create Transport Document activity
func (t *TimestampActivities) CreateTimestampActivity(ctx context.Context, form *jitproto.CreateTimestampRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*jitproto.CreateTimestampResponse, error) {
	timestampServiceClient := t.TimestampServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	tstamp, err := timestampServiceClient.CreateTimestamp(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return tstamp, nil
}
