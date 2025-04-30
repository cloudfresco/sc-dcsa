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

func TestSurrenderRequestService_CreateTransactionParty(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	transactionParty := eblproto.CreateTransactionPartyRequest{}
	transactionParty.EblPlatformIdentifier = "BOLE"
	transactionParty.LegalName = "Digital Container Shipping Association"
	transactionParty.RegistrationNumber = "74567837"
	transactionParty.LocationOfRegistration = "NL"
	transactionParty.TaxReference = "NL859951480B01"
	transactionParty.UserId = "auth0|66fd06d0bfea78a82bb42459"
	transactionParty.UserEmail = "sprov300@gmail.com"
	transactionParty.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateTransactionPartyRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.CreateTransactionPartyResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx: ctx,
				in:  &transactionParty,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transactionPartyResp, err := tt.ss.CreateTransactionParty(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.CreateTransactionParty() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, transactionPartyResp)
		transactionPartyResult := transactionPartyResp.TransactionParty
		assert.Equal(t, transactionPartyResult.TransactionPartyD.EblPlatformIdentifier, "BOLE", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.LegalName, "Digital Container Shipping Association", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.RegistrationNumber, "74567837", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.LocationOfRegistration, "NL", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.TaxReference, "NL859951480B01", "they should be equal")
	}
}

func TestSurrenderRequestService_GetTransactionParties(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	transactionParties := []*eblproto.TransactionParty{}

	transactionParty1, err := GetTransactionParty(uint32(1), []byte{194, 104, 62, 218, 117, 253, 72, 253, 134, 218, 146, 170, 68, 56, 27, 14}, "c2683eda-75fd-48fd-86da-92aa44381b0e", "BOLE", "Digital Container Shipping Association", "74567837", "NL", "NL859951480B01", "2023-03-07T12:12:12Z", "2023-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	transactionParties = append(transactionParties, transactionParty1)

	form := eblproto.GetTransactionPartiesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	transactionPartiesResponse := eblproto.GetTransactionPartiesResponse{TransactionParties: transactionParties, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *eblproto.GetTransactionPartiesRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.GetTransactionPartiesResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &transactionPartiesResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transactionPartyResp, err := tt.ss.GetTransactionParties(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.GetTransactionParties() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transactionPartyResp, tt.want) {
			t.Errorf("SurrenderRequestService.GetTransactionParties() = %v, want %v", transactionPartyResp, tt.want)
		}
		assert.NotNil(t, transactionPartyResp)
		transactionPartyResult := transactionPartyResp.TransactionParties[0]
		assert.Equal(t, transactionPartyResult.TransactionPartyD.EblPlatformIdentifier, "BOLE", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.LegalName, "Digital Container Shipping Association", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.RegistrationNumber, "74567837", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.LocationOfRegistration, "NL", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.TaxReference, "NL859951480B01", "they should be equal")
	}
}

func TestSurrenderRequestService_GetTransactionParty(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	transactionParty, err := GetTransactionParty(uint32(1), []byte{194, 104, 62, 218, 117, 253, 72, 253, 134, 218, 146, 170, 68, 56, 27, 14}, "c2683eda-75fd-48fd-86da-92aa44381b0e", "BOLE", "Digital Container Shipping Association", "74567837", "NL", "NL859951480B01", "2023-03-07T12:12:12Z", "2023-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	form := eblproto.GetTransactionPartyRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "c2683eda-75fd-48fd-86da-92aa44381b0e"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	transactionPartyResponse := eblproto.GetTransactionPartyResponse{}
	transactionPartyResponse.TransactionParty = transactionParty

	type args struct {
		ctx   context.Context
		inReq *eblproto.GetTransactionPartyRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.GetTransactionPartyResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &transactionPartyResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transactionPartyResp, err := tt.ss.GetTransactionParty(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.GetTransactionParty() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transactionPartyResp, tt.want) {
			t.Errorf("SurrenderRequestService.GetTransactionParty() = %v, want %v", transactionPartyResp, tt.want)
		}

		assert.NotNil(t, transactionPartyResp)
		transactionPartyResult := transactionPartyResp.TransactionParty
		assert.Equal(t, transactionPartyResult.TransactionPartyD.EblPlatformIdentifier, "BOLE", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.LegalName, "Digital Container Shipping Association", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.RegistrationNumber, "74567837", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.LocationOfRegistration, "NL", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.TaxReference, "NL859951480B01", "they should be equal")

	}
}

func TestSurrenderRequestService_GetTransactionPartyByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	transactionParty, err := GetTransactionParty(uint32(1), []byte{194, 104, 62, 218, 117, 253, 72, 253, 134, 218, 146, 170, 68, 56, 27, 14}, "c2683eda-75fd-48fd-86da-92aa44381b0e", "BOLE", "Digital Container Shipping Association", "74567837", "NL", "NL859951480B01", "2023-03-07T12:12:12Z", "2023-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	transactionPartyResponse := eblproto.GetTransactionPartyByPkResponse{}
	transactionPartyResponse.TransactionParty = transactionParty

	form := eblproto.GetTransactionPartyByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(1)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *eblproto.GetTransactionPartyByPkRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.GetTransactionPartyByPkResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &transactionPartyResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transactionPartyResp, err := tt.ss.GetTransactionPartyByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.GetTransactionPartyByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transactionPartyResp, tt.want) {
			t.Errorf("SurrenderRequestService.GetTransactionPartyByPk() = %v, want %v", transactionPartyResp, tt.want)
		}
		assert.NotNil(t, transactionPartyResp)
		transactionPartyResult := transactionPartyResp.TransactionParty
		assert.Equal(t, transactionPartyResult.TransactionPartyD.EblPlatformIdentifier, "BOLE", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.LegalName, "Digital Container Shipping Association", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.RegistrationNumber, "74567837", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.LocationOfRegistration, "NL", "they should be equal")
		assert.Equal(t, transactionPartyResult.TransactionPartyD.TaxReference, "NL859951480B01", "they should be equal")
	}
}

func TestSurrenderRequestService_UpdateTransactionParty(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	transactionParty := eblproto.UpdateTransactionPartyRequest{}
	transactionParty.EblPlatformIdentifier = "BOLE"
	transactionParty.LegalName = "Digital Container"
	transactionParty.RegistrationNumber = "74567877"
	transactionParty.LocationOfRegistration = "ML"
	transactionParty.Id = "c2683eda-75fd-48fd-86da-92aa44381b0e"
	transactionParty.UserId = "auth0|66fd06d0bfea78a82bb42459"
	transactionParty.UserEmail = "sprov300@gmail.com"
	transactionParty.RequestId = "bks1m1g91jau4nkks2f0"

	updateResponse := eblproto.UpdateTransactionPartyResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateTransactionPartyRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.UpdateTransactionPartyResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx: ctx,
				in:  &transactionParty,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.ss.UpdateTransactionParty(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.UpdateTransactionParty() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SurrenderRequestService.UpdateTransactionParty() = %v, want %v", got, tt.want)
		}
	}
}

func TestSurrenderRequestService_CreateTransactionPartySupportingCode(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	transactionPartySupportingCode := eblproto.CreateTransactionPartySupportingCodeRequest{}
	transactionPartySupportingCode.TransactionPartyId = uint32(1)
	transactionPartySupportingCode.PartyCode = "990052T8BM49ARSUDO55"
	transactionPartySupportingCode.PartyCodeListProvider = "SPIU"
	transactionPartySupportingCode.UserId = "auth0|66fd06d0bfea78a82bb42459"
	transactionPartySupportingCode.UserEmail = "sprov300@gmail.com"
	transactionPartySupportingCode.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateTransactionPartySupportingCodeRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.CreateTransactionPartySupportingCodeResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx: ctx,
				in:  &transactionPartySupportingCode,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transactionPartySupportingCodeResp, err := tt.ss.CreateTransactionPartySupportingCode(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.CreateTransactionPartySupportingCode() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, transactionPartySupportingCodeResp)
		transactionPartySupportingCodeResult := transactionPartySupportingCodeResp.TransactionPartySupportingCode
		assert.Equal(t, transactionPartySupportingCodeResult.TransactionPartySupportingCodeD.PartyCode, "990052T8BM49ARSUDO55", "they should be equal")
		assert.Equal(t, transactionPartySupportingCodeResult.TransactionPartySupportingCodeD.PartyCodeListProvider, "SPIU", "they should be equal")
	}
}

func TestSurrenderRequestService_UpdateTransactionPartySupportingCode(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	transactionPartySupportingCode := eblproto.UpdateTransactionPartySupportingCodeRequest{}
	transactionPartySupportingCode.TransactionPartyId = uint32(2)
	transactionPartySupportingCode.PartyCode = "80900T8BM49AURSDO55"
	transactionPartySupportingCode.PartyCodeListProvider = "RPIU"
	transactionPartySupportingCode.Id = "8839213f-8092-429e-8b7b-ca1c718cd140"
	transactionPartySupportingCode.UserId = "auth0|66fd06d0bfea78a82bb42459"
	transactionPartySupportingCode.UserEmail = "sprov300@gmail.com"
	transactionPartySupportingCode.RequestId = "bks1m1g91jau4nkks2f0"

	updateResponse := eblproto.UpdateTransactionPartySupportingCodeResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateTransactionPartySupportingCodeRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.UpdateTransactionPartySupportingCodeResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx: ctx,
				in:  &transactionPartySupportingCode,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.ss.UpdateTransactionPartySupportingCode(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.UpdateTransactionPartySupportingCode() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SurrenderRequestService.UpdateTransactionPartySupportingCode() = %v, want %v", got, tt.want)
		}
	}
}

func GetTransactionParty(id uint32, uuid4 []byte, idS string, eblPlatformIdentifier string, legalName string, registrationNumber string, locationOfRegistration string, taxReference string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*eblproto.TransactionParty, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	transactionPartyD := eblproto.TransactionPartyD{}
	transactionPartyD.Id = id
	transactionPartyD.Uuid4 = uuid4
	transactionPartyD.IdS = idS
	transactionPartyD.EblPlatformIdentifier = eblPlatformIdentifier
	transactionPartyD.LegalName = legalName
	transactionPartyD.RegistrationNumber = registrationNumber
	transactionPartyD.LocationOfRegistration = locationOfRegistration
	transactionPartyD.TaxReference = taxReference

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	transactionParty := eblproto.TransactionParty{TransactionPartyD: &transactionPartyD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &transactionParty, nil
}

func TestSurrenderRequestService_CreateEndorsementChainLink(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	endorsementChainLink := eblproto.CreateEndorsementChainLinkRequest{}
	endorsementChainLink.EntryOrder = int32(1)
	endorsementChainLink.Actor = uint32(1)
	endorsementChainLink.Recipient = uint32(1)
	endorsementChainLink.SurrenderRequestId = uint32(1)
	endorsementChainLink.ActionDateTime = "03/22/2024"
	endorsementChainLink.UserId = "auth0|66fd06d0bfea78a82bb42459"
	endorsementChainLink.UserEmail = "sprov300@gmail.com"
	endorsementChainLink.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateEndorsementChainLinkRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.CreateEndorsementChainLinkResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx: ctx,
				in:  &endorsementChainLink,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		endorsementChainLinkResponse, err := tt.ss.CreateEndorsementChainLink(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.CreateEndorsementChainLink() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, endorsementChainLinkResponse)
	}
}

func TestSurrenderRequestService_UpdateEndorsementChainLink(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	endorsementChainLink := eblproto.UpdateEndorsementChainLinkRequest{}
	endorsementChainLink.EntryOrder = int32(2)
	endorsementChainLink.Actor = uint32(2)
	endorsementChainLink.Recipient = uint32(2)
	endorsementChainLink.Id = "9d16c5f9-dd0a-444d-a9e4-4b252852cbf3"
	endorsementChainLink.UserId = "auth0|66fd06d0bfea78a82bb42459"
	endorsementChainLink.UserEmail = "sprov300@gmail.com"
	endorsementChainLink.RequestId = "bks1m1g91jau4nkks2f0"

	updateResponse := eblproto.UpdateEndorsementChainLinkResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateEndorsementChainLinkRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.UpdateEndorsementChainLinkResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx: ctx,
				in:  &endorsementChainLink,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.ss.UpdateEndorsementChainLink(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.UpdateEndorsementChainLink() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SurrenderRequestService.UpdateEndorsementChainLink() = %v, want %v", got, tt.want)
		}
	}
}

func TestSurrenderRequestService_CreateSurrenderRequest(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	surrenderRequest := eblproto.CreateSurrenderRequestRequest{}
	surrenderRequest.SurrenderRequestReference = "Z12345"
	surrenderRequest.TransportDocumentReference = "string"
	surrenderRequest.SurrenderRequestCode = "SREQ"
	surrenderRequest.Comments = "string"
	surrenderRequest.SurrenderRequestedBy = uint32(1)
	surrenderRequest.CreatedDateTime = "03/22/2024"
	surrenderRequest.UserId = "auth0|66fd06d0bfea78a82bb42459"
	surrenderRequest.UserEmail = "sprov300@gmail.com"
	surrenderRequest.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateSurrenderRequestRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.CreateSurrenderRequestResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx: ctx,
				in:  &surrenderRequest,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		surrenderRequestResp, err := tt.ss.CreateSurrenderRequest(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.CreateSurrenderRequest() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, surrenderRequestResp)
		surrenderRequestResult := surrenderRequestResp.SurrenderRequest
		assert.Equal(t, surrenderRequestResult.SurrenderRequestD.SurrenderRequestReference, "Z12345", "they should be equal")
		assert.Equal(t, surrenderRequestResult.SurrenderRequestD.SurrenderRequestCode, "SREQ", "they should be equal")
	}
}

func TestSurrenderRequestService_UpdateSurrenderRequest(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, userServiceClient)

	surrenderRequest := eblproto.UpdateSurrenderRequestRequest{}
	surrenderRequest.SurrenderRequestReference = "Y12345"
	surrenderRequest.TransportDocumentReference = "X3456"
	surrenderRequest.SurrenderRequestCode = "PREQ"
	surrenderRequest.Comments = "comments"
	surrenderRequest.Id = "e40fc6f6-c8ce-4ef4-97c0-bf112b70d3f2"
	surrenderRequest.UserId = "auth0|66fd06d0bfea78a82bb42459"
	surrenderRequest.UserEmail = "sprov300@gmail.com"
	surrenderRequest.RequestId = "bks1m1g91jau4nkks2f0"

	updateResponse := eblproto.UpdateSurrenderRequestResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateSurrenderRequestRequest
	}
	tests := []struct {
		ss      *SurrenderRequestService
		args    args
		want    *eblproto.UpdateSurrenderRequestResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestService,
			args: args{
				ctx: ctx,
				in:  &surrenderRequest,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.ss.UpdateSurrenderRequest(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestService.UpdateSurrenderRequest() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SurrenderRequestService.UpdateSurrenderRequest() = %v, want %v", got, tt.want)
		}
	}
}
