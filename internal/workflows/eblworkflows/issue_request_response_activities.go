package eblworkflows

import (
	"context"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type IssueRequestResponseActivities struct {
	IssueRequestResponseServiceClient eblproto.IssueRequestResponseServiceClient
}

// CreateIssuanceRequestResponseActivity - Create IssuanceRequestResponse activity
func (i *IssueRequestResponseActivities) CreateIssuanceRequestResponseActivity(ctx context.Context, form *eblproto.CreateIssuanceRequestResponseRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateIssuanceRequestResponseResponse, error) {
	issueRequestResponseServiceClient := i.IssueRequestResponseServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	issuanceRequestResponse, err := issueRequestResponseServiceClient.CreateIssuanceRequestResponse(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return issuanceRequestResponse, nil
}

// UpdateIssuanceRequestResponseActivity - update IssuanceRequestResponse activity
func (i *IssueRequestResponseActivities) UpdateIssuanceRequestResponseActivity(ctx context.Context, form *eblproto.UpdateIssuanceRequestResponseRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	issueRequestResponseServiceClient := i.IssueRequestResponseServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := issueRequestResponseServiceClient.UpdateIssuanceRequestResponse(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}
