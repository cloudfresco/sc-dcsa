package eblworkflows

import (
	"context"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type SurrenderRequestActivities struct {
	SurrenderRequestServiceClient eblproto.SurrenderRequestServiceClient
}

// CreateTransactionPartyActivity - Create TransactionParty activity
func (s *SurrenderRequestActivities) CreateTransactionPartyActivity(ctx context.Context, form *eblproto.CreateTransactionPartyRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateTransactionPartyResponse, error) {
	surrenderRequestServiceClient := s.SurrenderRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	transactionParty, err := surrenderRequestServiceClient.CreateTransactionParty(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return transactionParty, nil
}

// UpdateTransactionPartyActivity - update TransactionParty activity
func (s *SurrenderRequestActivities) UpdateTransactionPartyActivity(ctx context.Context, form *eblproto.UpdateTransactionPartyRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	surrenderRequestServiceClient := s.SurrenderRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := surrenderRequestServiceClient.UpdateTransactionParty(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}

// CreateTransactionPartySupportingCodeActivity - Create TransactionPartySupportingCode activity
func (s *SurrenderRequestActivities) CreateTransactionPartySupportingCodeActivity(ctx context.Context, form *eblproto.CreateTransactionPartySupportingCodeRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateTransactionPartySupportingCodeResponse, error) {
	surrenderRequestServiceClient := s.SurrenderRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	transactionPartySupportingCode, err := surrenderRequestServiceClient.CreateTransactionPartySupportingCode(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return transactionPartySupportingCode, nil
}

// UpdateTransactionPartySupportingCodeActivity - update TransactionPartySupportingCode activity
func (s *SurrenderRequestActivities) UpdateTransactionPartySupportingCodeActivity(ctx context.Context, form *eblproto.UpdateTransactionPartySupportingCodeRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	surrenderRequestServiceClient := s.SurrenderRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := surrenderRequestServiceClient.UpdateTransactionPartySupportingCode(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}

// CreateSurrenderRequestActivity - Create SurrenderRequest activity
func (s *SurrenderRequestActivities) CreateSurrenderRequestActivity(ctx context.Context, form *eblproto.CreateSurrenderRequestRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateSurrenderRequestResponse, error) {
	surrenderRequestServiceClient := s.SurrenderRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	surrenderRequest, err := surrenderRequestServiceClient.CreateSurrenderRequest(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return surrenderRequest, nil
}

// UpdateSurrenderRequestActivity - update SurrenderRequest activity
func (s *SurrenderRequestActivities) UpdateSurrenderRequestActivity(ctx context.Context, form *eblproto.UpdateSurrenderRequestRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	surrenderRequestServiceClient := s.SurrenderRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := surrenderRequestServiceClient.UpdateSurrenderRequest(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}

// CreateEndorsementChainLinkActivity - Create EndorsementChainLink activity
func (s *SurrenderRequestActivities) CreateEndorsementChainLinkActivity(ctx context.Context, form *eblproto.CreateEndorsementChainLinkRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateEndorsementChainLinkResponse, error) {
	surrenderRequestServiceClient := s.SurrenderRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	endorsementChainLink, err := surrenderRequestServiceClient.CreateEndorsementChainLink(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return endorsementChainLink, nil
}

// UpdateEndorsementChainLinkActivity - update EndorsementChainLink activity
func (s *SurrenderRequestActivities) UpdateEndorsementChainLinkActivity(ctx context.Context, form *eblproto.UpdateEndorsementChainLinkRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	surrenderRequestServiceClient := s.SurrenderRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := surrenderRequestServiceClient.UpdateEndorsementChainLink(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}
