package eblworkflows

import (
	"context"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type TransportDocumentActivities struct {
	TransportDocumentServiceClient eblproto.TransportDocumentServiceClient
}

// CreateTransportDocumentActivity - Create Transport Document activity
func (td *TransportDocumentActivities) CreateTransportDocumentActivity(ctx context.Context, form *eblproto.CreateTransportDocumentRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateTransportDocumentResponse, error) {
	transportDocumentServiceClient := td.TransportDocumentServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	transportDocument, err := transportDocumentServiceClient.CreateTransportDocument(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return transportDocument, nil
}

// ApproveTransportDocumentActivity - approve TransportDocument activity
func (td *TransportDocumentActivities) ApproveTransportDocumentActivity(ctx context.Context, form *eblproto.ApproveTransportDocumentRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	transportDocumentServiceClient := td.TransportDocumentServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := transportDocumentServiceClient.ApproveTransportDocument(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Approved Successfully", nil
}
