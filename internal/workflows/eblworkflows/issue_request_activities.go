package eblworkflows

import (
	"context"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type IssueRequestActivities struct {
	IssueRequestServiceClient eblproto.IssueRequestServiceClient
}

// CreateIssuePartyActivity - Create IssueParty activity
func (i *IssueRequestActivities) CreateIssuePartyActivity(ctx context.Context, form *eblproto.CreateIssuePartyRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateIssuePartyResponse, error) {
	issueRequestServiceClient := i.IssueRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	issueParty, err := issueRequestServiceClient.CreateIssueParty(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return issueParty, nil
}

// UpdateIssuePartyActivity - update IssueParty activity
func (i *IssueRequestActivities) UpdateIssuePartyActivity(ctx context.Context, form *eblproto.UpdateIssuePartyRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	issueRequestServiceClient := i.IssueRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := issueRequestServiceClient.UpdateIssueParty(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}

// CreateIssuePartySupportingCodeActivity - Create IssuePartySupportingCode activity
func (i *IssueRequestActivities) CreateIssuePartySupportingCodeActivity(ctx context.Context, form *eblproto.CreateIssuePartySupportingCodeRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateIssuePartySupportingCodeResponse, error) {
	issueRequestServiceClient := i.IssueRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	issuePartySupportingCode, err := issueRequestServiceClient.CreateIssuePartySupportingCode(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return issuePartySupportingCode, nil
}

// UpdateIssuePartySupportingCodeActivity - update IssuePartySupportingCode activity
func (i *IssueRequestActivities) UpdateIssuePartySupportingCodeActivity(ctx context.Context, form *eblproto.UpdateIssuePartySupportingCodeRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	issueRequestServiceClient := i.IssueRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := issueRequestServiceClient.UpdateIssuePartySupportingCode(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}

// CreateIssuanceRequestActivity - Create IssuanceRequest activity
func (i *IssueRequestActivities) CreateIssuanceRequestActivity(ctx context.Context, form *eblproto.CreateIssuanceRequestRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateIssuanceRequestResponse, error) {
	issueRequestServiceClient := i.IssueRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	issuanceRequest, err := issueRequestServiceClient.CreateIssuanceRequest(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return issuanceRequest, nil
}

// UpdateIssuanceRequestActivity - update IssuanceRequest activity
func (i *IssueRequestActivities) UpdateIssuanceRequestActivity(ctx context.Context, form *eblproto.UpdateIssuanceRequestRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	issueRequestServiceClient := i.IssueRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := issueRequestServiceClient.UpdateIssuanceRequest(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}

// CreateEblVisualizationActivity - Create EblVisualization activity
func (i *IssueRequestActivities) CreateEblVisualizationActivity(ctx context.Context, form *eblproto.CreateEblVisualizationRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateEblVisualizationResponse, error) {
	issueRequestServiceClient := i.IssueRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	eblVisualization, err := issueRequestServiceClient.CreateEblVisualization(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return eblVisualization, nil
}

// UpdateEblVisualizationActivity - update EblVisualization activity
func (i *IssueRequestActivities) UpdateEblVisualizationActivity(ctx context.Context, form *eblproto.UpdateEblVisualizationRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	issueRequestServiceClient := i.IssueRequestServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := issueRequestServiceClient.UpdateEblVisualization(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}
