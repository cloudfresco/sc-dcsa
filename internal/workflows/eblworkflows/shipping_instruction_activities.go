package eblworkflows

import (
	"context"
	"fmt"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type ShippingInstructionActivities struct {
	ShippingInstructionServiceClient eblproto.ShippingInstructionServiceClient
}

// CreateShippingInstructionActivity - Create ShippingInstruction activity
func (s *ShippingInstructionActivities) CreateShippingInstructionActivity(ctx context.Context, form *eblproto.CreateShippingInstructionRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*eblproto.CreateShippingInstructionResponse, error) {
	shippingInstructionServiceClient := s.ShippingInstructionServiceClient
	fmt.Println("CreateShippingInstructionActivity shippingInstructionServiceClient", shippingInstructionServiceClient)
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	shippingInstruction, err := shippingInstructionServiceClient.CreateShippingInstruction(ctxNew, form)
	fmt.Println("CreateShippingInstructionActivity err", err)
	fmt.Println("CreateShippingInstructionActivity shippingInstruction", shippingInstruction)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return shippingInstruction, nil
}

// UpdateShippingInstructionActivity - update ShippingInstruction activity
func (s *ShippingInstructionActivities) UpdateShippingInstructionActivity(ctx context.Context, form *eblproto.UpdateShippingInstructionByShippingInstructionReferenceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	shippingInstructionServiceClient := s.ShippingInstructionServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := shippingInstructionServiceClient.UpdateShippingInstructionByShippingInstructionReference(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}