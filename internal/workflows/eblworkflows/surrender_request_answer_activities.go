package eblworkflows

import (
	"context"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type SurrenderRequestAnswerActivities struct {
	SurrenderRequestAnswerServiceClient eblproto.SurrenderRequestAnswerServiceClient
}

// CreateSurrenderRequestAnswerActivity - Create SurrenderRequestAnswer activity
func (s *SurrenderRequestAnswerActivities) CreateSurrenderRequestAnswerActivity(ctx context.Context, form *eblproto.CreateSurrenderRequestAnswerRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateSurrenderRequestAnswerResponse, error) {
	surrenderRequestAnswerServiceClient := s.SurrenderRequestAnswerServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	surrenderRequestAnswer, err := surrenderRequestAnswerServiceClient.CreateSurrenderRequestAnswer(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return surrenderRequestAnswer, nil
}

// UpdateSurrenderRequestAnswerActivity - update SurrenderRequestAnswer activity
func (s *SurrenderRequestAnswerActivities) UpdateSurrenderRequestAnswerActivity(ctx context.Context, form *eblproto.UpdateSurrenderRequestAnswerRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	surrenderRequestAnswerServiceClient := s.SurrenderRequestAnswerServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := surrenderRequestAnswerServiceClient.UpdateSurrenderRequestAnswer(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}
