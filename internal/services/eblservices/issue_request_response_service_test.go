package eblservices

import (
	"context"
	"reflect"
	"testing"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
)

func TestIssueRequestResponseService_CreateIssuanceRequestResponse(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestResponseService := NewIssueRequestResponseService(log, dbService, redisService, userServiceClient)

	issuanceRequestResponse := eblproto.CreateIssuanceRequestResponseRequest{}
	issuanceRequestResponse.TransportDocumentReference = "HHL71800000"
	issuanceRequestResponse.IssuanceResponseCode = "ISSU"
	issuanceRequestResponse.Reason = "null"
	issuanceRequestResponse.IssuanceRequestId = uint32(1)
	issuanceRequestResponse.CreatedDateTime = "03/22/2024"
	issuanceRequestResponse.UserId = "auth0|66fd06d0bfea78a82bb42459"
	issuanceRequestResponse.UserEmail = "sprov300@gmail.com"
	issuanceRequestResponse.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateIssuanceRequestResponseRequest
	}
	tests := []struct {
		is      *IssueRequestResponseService
		args    args
		want    *eblproto.CreateIssuanceRequestResponseResponse
		wantErr bool
	}{
		{
			is: issueRequestResponseService,
			args: args{
				ctx: ctx,
				in:  &issuanceRequestResponse,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		issuanceRequestResponse, err := tt.is.CreateIssuanceRequestResponse(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestResponseService.CreateIssuanceRequestResponse() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, issuanceRequestResponse)
		issuanceRequestRespResult := issuanceRequestResponse.IssuanceRequestResponse
		assert.Equal(t, issuanceRequestRespResult.IssuanceRequestResponseD.TransportDocumentReference, "HHL71800000", "they should be equal")
		assert.Equal(t, issuanceRequestRespResult.IssuanceRequestResponseD.IssuanceResponseCode, "ISSU", "they should be equal")
	}
}

func TestIssueRequestResponseService_UpdateIssuanceRequestResponse(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestResponseService := NewIssueRequestResponseService(log, dbService, redisService, userServiceClient)

	issuanceRequestResponse := eblproto.UpdateIssuanceRequestResponseRequest{}
	issuanceRequestResponse.TransportDocumentReference = "AAL71800000"
	issuanceRequestResponse.IssuanceResponseCode = "ESSU"
	issuanceRequestResponse.Reason = "emergency"
	issuanceRequestResponse.Id = "311ca76d-56a9-42af-94a9-6602af9ed683"
	issuanceRequestResponse.UserId = "auth0|66fd06d0bfea78a82bb42459"
	issuanceRequestResponse.UserEmail = "sprov300@gmail.com"
	issuanceRequestResponse.RequestId = "bks1m1g91jau4nkks2f0"

	updateResponse := eblproto.UpdateIssuanceRequestResponseResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateIssuanceRequestResponseRequest
	}
	tests := []struct {
		is      *IssueRequestResponseService
		args    args
		want    *eblproto.UpdateIssuanceRequestResponseResponse
		wantErr bool
	}{
		{
			is: issueRequestResponseService,
			args: args{
				ctx: ctx,
				in:  &issuanceRequestResponse,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.is.UpdateIssuanceRequestResponse(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestResponseService.UpdateIssuanceRequestResponse() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("IssueRequestResponseService.UpdateIssuanceRequestResponse() = %v, want %v", got, tt.want)
		}
	}
}
