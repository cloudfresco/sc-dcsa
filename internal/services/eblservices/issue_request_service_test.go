package eblservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestIssueRequestService_GetIssueParties(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	issueParties := []*eblproto.IssueParty{}

	issueParty1, err := GetIssueParty(uint32(1), []byte{126, 51, 121, 158, 146, 48, 76, 94, 178, 9, 92, 21, 239, 206, 19, 3}, "7e33799e-9230-4c5e-b209-5c15efce1303", "BOLE", "Digital Container Shipping Association", "74567837", "NL", "NL859951480B01", "2023-03-07T12:12:12Z", "2023-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	issueParties = append(issueParties, issueParty1)

	form := eblproto.GetIssuePartiesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	issuePartiesResponse := eblproto.GetIssuePartiesResponse{IssueParties: issueParties, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *eblproto.GetIssuePartiesRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		want    *eblproto.GetIssuePartiesResponse
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &issuePartiesResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		issuePartyResp, err := tt.is.GetIssueParties(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.GetIssueParties() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(issuePartyResp, tt.want) {
			t.Errorf("IssueRequestService.GetIssueParties() = %v, want %v", issuePartyResp, tt.want)
		}
		assert.NotNil(t, issuePartyResp)
		issuePartyResult := issuePartyResp.IssueParties[0]
		assert.Equal(t, issuePartyResult.IssuePartyD.EblPlatformIdentifier, "BOLE", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.LegalName, "Digital Container Shipping Association", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.RegistrationNumber, "74567837", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.LocationOfRegistration, "NL", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.TaxReference, "NL859951480B01", "they should be equal")
	}
}

func TestIssueRequestService_GetIssueParty(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	issueParty, err := GetIssueParty(uint32(1), []byte{126, 51, 121, 158, 146, 48, 76, 94, 178, 9, 92, 21, 239, 206, 19, 3}, "7e33799e-9230-4c5e-b209-5c15efce1303", "BOLE", "Digital Container Shipping Association", "74567837", "NL", "NL859951480B01", "2023-03-07T12:12:12Z", "2023-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	form := eblproto.GetIssuePartyRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "7e33799e-9230-4c5e-b209-5c15efce1303"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	issuePartyResponse := eblproto.GetIssuePartyResponse{}
	issuePartyResponse.IssueParty = issueParty

	type args struct {
		ctx   context.Context
		inReq *eblproto.GetIssuePartyRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		want    *eblproto.GetIssuePartyResponse
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &issuePartyResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		issuePartyResp, err := tt.is.GetIssueParty(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.GetIssueParty() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(issuePartyResp, tt.want) {
			t.Errorf("IssueRequestService.GetIssueParty() = %v, want %v", issuePartyResp, tt.want)
		}
		assert.NotNil(t, issuePartyResp)
		issuePartyResult := issuePartyResp.IssueParty
		assert.Equal(t, issuePartyResult.IssuePartyD.EblPlatformIdentifier, "BOLE", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.LegalName, "Digital Container Shipping Association", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.RegistrationNumber, "74567837", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.LocationOfRegistration, "NL", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.TaxReference, "NL859951480B01", "they should be equal")
	}
}

func TestIssueRequestService_GetIssuePartyByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	issueParty, err := GetIssueParty(uint32(1), []byte{126, 51, 121, 158, 146, 48, 76, 94, 178, 9, 92, 21, 239, 206, 19, 3}, "7e33799e-9230-4c5e-b209-5c15efce1303", "BOLE", "Digital Container Shipping Association", "74567837", "NL", "NL859951480B01", "2023-03-07T12:12:12Z", "2023-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	issuePartyResponse := eblproto.GetIssuePartyByPkResponse{}
	issuePartyResponse.IssueParty = issueParty

	form := eblproto.GetIssuePartyByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(1)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *eblproto.GetIssuePartyByPkRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		want    *eblproto.GetIssuePartyByPkResponse
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &issuePartyResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		issuePartyResp, err := tt.is.GetIssuePartyByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.GetIssuePartyByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(issuePartyResp, tt.want) {
			t.Errorf("IssueRequestService.GetIssuePartyByPk() = %v, want %v", issuePartyResp, tt.want)
		}
		assert.NotNil(t, issuePartyResp)
		issuePartyResult := issuePartyResp.IssueParty
		assert.Equal(t, issuePartyResult.IssuePartyD.EblPlatformIdentifier, "BOLE", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.LegalName, "Digital Container Shipping Association", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.RegistrationNumber, "74567837", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.LocationOfRegistration, "NL", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.TaxReference, "NL859951480B01", "they should be equal")
	}
}

func TestIssueRequestService_CreateIssueParty(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	issueParty := eblproto.CreateIssuePartyRequest{}
	issueParty.EblPlatformIdentifier = "BOLE"
	issueParty.LegalName = "Digital Container Shipping Association"
	issueParty.RegistrationNumber = "74567837"
	issueParty.LocationOfRegistration = "NL"
	issueParty.TaxReference = "NL859951480B01"
	issueParty.UserId = "auth0|66fd06d0bfea78a82bb42459"
	issueParty.UserEmail = "sprov300@gmail.com"
	issueParty.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateIssuePartyRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx: ctx,
				in:  &issueParty,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		issuePartyResp, err := tt.is.CreateIssueParty(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.CreateIssueParty() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		assert.NotNil(t, issuePartyResp)
		issuePartyResult := issuePartyResp.IssueParty
		assert.Equal(t, issuePartyResult.IssuePartyD.EblPlatformIdentifier, "BOLE", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.LegalName, "Digital Container Shipping Association", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.RegistrationNumber, "74567837", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.LocationOfRegistration, "NL", "they should be equal")
		assert.Equal(t, issuePartyResult.IssuePartyD.TaxReference, "NL859951480B01", "they should be equal")
	}
}

func TestIssueRequestService_UpdateIssueParty(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	issueParty := eblproto.UpdateIssuePartyRequest{}
	issueParty.EblPlatformIdentifier = "BOLE"
	issueParty.LegalName = "Digital Container"
	issueParty.RegistrationNumber = "74567877"
	issueParty.LocationOfRegistration = "ML"
	issueParty.Id = "7e33799e-9230-4c5e-b209-5c15efce1303"
	issueParty.UserId = "auth0|66fd06d0bfea78a82bb42459"
	issueParty.UserEmail = "sprov300@gmail.com"
	issueParty.RequestId = "bks1m1g91jau4nkks2f0"

	updateResponse := eblproto.UpdateIssuePartyResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateIssuePartyRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		want    *eblproto.UpdateIssuePartyResponse
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx: ctx,
				in:  &issueParty,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.is.UpdateIssueParty(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.UpdateIssueParty() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("IssueRequestService.UpdateIssueParty() = %v, want %v", got, tt.want)
		}
	}
}

func TestIssueRequestService_CreateIssuePartySupportingCode(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	issuePartySupportingCode := eblproto.CreateIssuePartySupportingCodeRequest{}
	issuePartySupportingCode.IssuePartyId = uint32(1)
	issuePartySupportingCode.PartyCode = "529900T8BM49AURSDO55"
	issuePartySupportingCode.PartyCodeListProvider = "EPIU"
	issuePartySupportingCode.UserId = "auth0|66fd06d0bfea78a82bb42459"
	issuePartySupportingCode.UserEmail = "sprov300@gmail.com"
	issuePartySupportingCode.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateIssuePartySupportingCodeRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		want    *eblproto.CreateIssuePartySupportingCodeResponse
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx: ctx,
				in:  &issuePartySupportingCode,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		issuePartySupportingCodeResp, err := tt.is.CreateIssuePartySupportingCode(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.CreateIssuePartySupportingCode() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, issuePartySupportingCodeResp)
		issuePartySupportingCodeResult := issuePartySupportingCodeResp.IssuePartySupportingCode
		assert.Equal(t, issuePartySupportingCodeResult.IssuePartySupportingCodeD.PartyCode, "529900T8BM49AURSDO55", "they should be equal")
		assert.Equal(t, issuePartySupportingCodeResult.IssuePartySupportingCodeD.PartyCodeListProvider, "EPIU", "they should be equal")
	}
}

func TestIssueRequestService_UpdateIssuePartySupportingCode(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	issuePartySupportingCode := eblproto.UpdateIssuePartySupportingCodeRequest{}
	issuePartySupportingCode.IssuePartyId = uint32(2)
	issuePartySupportingCode.PartyCode = "859900T8BM49AURSDO55"
	issuePartySupportingCode.PartyCodeListProvider = "PPIU"
	issuePartySupportingCode.Id = "1be42473-a5a8-4dcf-abfe-53fc0ed01ba8"
	issuePartySupportingCode.UserId = "auth0|66fd06d0bfea78a82bb42459"
	issuePartySupportingCode.UserEmail = "sprov300@gmail.com"
	issuePartySupportingCode.RequestId = "bks1m1g91jau4nkks2f0"

	updateResponse := eblproto.UpdateIssuePartySupportingCodeResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateIssuePartySupportingCodeRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		want    *eblproto.UpdateIssuePartySupportingCodeResponse
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx: ctx,
				in:  &issuePartySupportingCode,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.is.UpdateIssuePartySupportingCode(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.UpdateIssuePartySupportingCode() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("IssueRequestService.UpdateIssuePartySupportingCode() = %v, want %v", got, tt.want)
		}
	}
}

func TestIssueRequestService_CreateEblVisualization(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	eblVisualization := eblproto.CreateEblVisualizationRequest{}
	eblVisualization.Name = "Carrier rendered copy of the EBL.pdf"
	eblVisualization.Content = "string"
	eblVisualization.UserId = "auth0|66fd06d0bfea78a82bb42459"
	eblVisualization.UserEmail = "sprov300@gmail.com"
	eblVisualization.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateEblVisualizationRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		want    *eblproto.CreateEblVisualizationResponse
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx: ctx,
				in:  &eblVisualization,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		eblVisualizationResp, err := tt.is.CreateEblVisualization(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.CreateEblVisualization() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, eblVisualizationResp)
		eblVisualizationResult := eblVisualizationResp.EblVisualization
		assert.Equal(t, eblVisualizationResult.EblVisualizationD.Name, "Carrier rendered copy of the EBL.pdf", "they should be equal")
		assert.Equal(t, eblVisualizationResult.EblVisualizationD.Content, "string", "they should be equal")
	}
}

func TestIssueRequestService_UpdateEblVisualization(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	eblVisualization := eblproto.UpdateEblVisualizationRequest{}
	eblVisualization.Name = "Carrier rendered EBL.pdf"
	eblVisualization.Content = "ws"
	eblVisualization.Id = "1270cfd4-4402-45c5-997b-63cc92ed0a9d"
	eblVisualization.UserId = "auth0|66fd06d0bfea78a82bb42459"
	eblVisualization.UserEmail = "sprov300@gmail.com"
	eblVisualization.RequestId = "bks1m1g91jau4nkks2f0"

	updateResponse := eblproto.UpdateEblVisualizationResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateEblVisualizationRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		want    *eblproto.UpdateEblVisualizationResponse
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx: ctx,
				in:  &eblVisualization,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.is.UpdateEblVisualization(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.UpdateEblVisualization() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("IssueRequestService.UpdateEblVisualization() = %v, want %v", got, tt.want)
		}
	}
}

func TestIssueRequestService_CreateIssuanceRequest(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	issuanceRequest := eblproto.CreateIssuanceRequestRequest{}
	issuanceRequest.TransportDocumentReference = "HHL71800000"
	issuanceRequest.IssuanceRequestState = "DR"
	issuanceRequest.IssueTo = uint32(1)
	issuanceRequest.EblVisualizationId = uint32(1)
	issuanceRequest.TransportDocumentJson = "SWB"
	issuanceRequest.UserId = "auth0|66fd06d0bfea78a82bb42459"
	issuanceRequest.UserEmail = "sprov300@gmail.com"
	issuanceRequest.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateIssuanceRequestRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		want    *eblproto.CreateIssuanceRequestResponse
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx: ctx,
				in:  &issuanceRequest,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		issuanceRequestResp, err := tt.is.CreateIssuanceRequest(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.CreateIssuanceRequest() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, issuanceRequestResp)
		issuanceRequestResult := issuanceRequestResp.IssuanceRequest
		assert.Equal(t, issuanceRequestResult.IssuanceRequestD.TransportDocumentReference, "HHL71800000", "they should be equal")
		assert.Equal(t, issuanceRequestResult.IssuanceRequestD.IssuanceRequestState, "DR", "they should be equal")
	}
}

func TestIssueRequestService_UpdateIssuanceRequest(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	issueRequestService := NewIssueRequestService(log, dbService, redisService, userServiceClient)

	issuanceRequest := eblproto.UpdateIssuanceRequestRequest{}
	issuanceRequest.TransportDocumentReference = "AAL71801000"
	issuanceRequest.IssuanceRequestState = "AR"
	issuanceRequest.Id = "f40939a3-c05b-4a34-bc67-a8e3676cdc80"
	issuanceRequest.UserId = "auth0|66fd06d0bfea78a82bb42459"
	issuanceRequest.UserEmail = "sprov300@gmail.com"
	issuanceRequest.RequestId = "bks1m1g91jau4nkks2f0"

	updateResponse := eblproto.UpdateIssuanceRequestResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateIssuanceRequestRequest
	}
	tests := []struct {
		is      *IssueRequestService
		args    args
		want    *eblproto.UpdateIssuanceRequestResponse
		wantErr bool
	}{
		{
			is: issueRequestService,
			args: args{
				ctx: ctx,
				in:  &issuanceRequest,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.is.UpdateIssuanceRequest(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("IssueRequestService.UpdateIssuanceRequest() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("IssueRequestService.UpdateIssuanceRequest() = %v, want %v", got, tt.want)
		}
	}
}

func GetIssueParty(id uint32, uuid4 []byte, idS string, eblPlatformIdentifier string, legalName string, registrationNumber string, locationOfRegistration string, taxReference string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*eblproto.IssueParty, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	issuePartyD := eblproto.IssuePartyD{}
	issuePartyD.Id = id
	issuePartyD.Uuid4 = uuid4
	issuePartyD.IdS = idS
	issuePartyD.EblPlatformIdentifier = eblPlatformIdentifier
	issuePartyD.LegalName = legalName
	issuePartyD.RegistrationNumber = registrationNumber
	issuePartyD.LocationOfRegistration = locationOfRegistration
	issuePartyD.TaxReference = taxReference

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	issueParty := eblproto.IssueParty{IssuePartyD: &issuePartyD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &issueParty, nil
}
