package partyservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestPartyService_GetParties(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()
	parties := []*partyproto.Party{}
	partyService := NewPartyService(log, dbService, redisService, userServiceClient)
	party1, err := GetParty(uint32(1), []byte{78, 68, 143, 38, 64, 53, 17, 235, 164, 157, 127, 158, 185, 188, 141, 217}, "4e448f26-4035-11eb-a49d-7f9eb9bc8dd9", "Malwart", "", "", "", uint32(1), "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	party2, err := GetParty(uint32(2), []byte{141, 217, 164, 196, 64, 57, 17, 235, 135, 112, 11, 43, 25, 132, 127, 171}, "8dd9a4c4-4039-11eb-8770-0b2b19847fab", "Malwart Düsseldorf", "", "", "", uint32(2), "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	party3, err := GetParty(uint32(3), []byte{157, 217, 164, 196, 64, 57, 17, 235, 135, 112, 11, 43, 25, 132, 127, 171}, "9dd9a4c4-4039-11eb-8770-0b2b19847fab", "Malwart Lyngy", "", "", "", uint32(3), "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	party4, err := GetParty(uint32(4), []byte{73, 153, 24, 162, 209, 45, 77, 246, 132, 12, 221, 146, 53, 112, 2, 223}, "499918a2-d12d-4df6-840c-dd92357002df", "FTL International", "", "", "", uint32(4), "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	party5, err := GetParty(uint32(5), []byte{142, 70, 58, 132, 10, 45, 71, 205, 147, 50, 81, 230, 203, 54, 182, 53}, "8e463a84-0a2d-47cd-9332-51e6cb36b635", "Superdæk Albertslund", "", "", "", uint32(5), "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	party6, err := GetParty(uint32(6), []byte{196, 158, 162, 214, 56, 6, 70, 200, 132, 144, 41, 74, 255, 199, 18, 134}, "c49ea2d6-3806-46c8-8490-294affc71286", "FDM Quality Control", "", "", "", uint32(6), "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	party7, err := GetParty(uint32(7), []byte{123, 246, 244, 40, 88, 240, 67, 71, 156, 232, 214, 190, 47, 93, 87, 69}, "7bf6f428-58f0-4347-9ce8-d6be2f5d5745", "Hapag Lloyd", "CVR-25645774", "CVR-25645774", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkFzaW", uint32(7), "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	party8, err := GetParty(uint32(8), []byte{190, 91, 194, 144, 123, 172, 72, 187, 162, 17, 243, 250, 90, 58, 179, 174}, "be5bc290-7bac-48bb-a211-f3fa5a3ab3ae", "Asseco Denmark", "CVR-25645774", "CVR-25645774", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkFzaW", uint32(8), "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	parties = append(parties, party8, party7, party6, party5, party4, party3, party2, party1)

	nextc := "MA=="
	partyResponse := partyproto.GetPartiesResponse{Parties: parties, NextCursor: nextc}

	form := partyproto.GetPartiesRequest{}
	form.Limit = "8"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *partyproto.GetPartiesRequest
	}
	tests := []struct {
		ps      *PartyService
		args    args
		want    *partyproto.GetPartiesResponse
		wantErr bool
	}{
		{
			ps: partyService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &partyResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		partyResp, err := tt.ps.GetParties(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("PartyService.GetParties() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(partyResp, tt.want) {
			t.Errorf("PartyService.GetParties() = %v, want %v", partyResp, tt.want)
		}
		assert.NotNil(t, partyResp)
		partyResult := partyResp.Parties[1]
		assert.Equal(t, partyResult.PartyD.PartyName, "Hapag Lloyd", "they should be equal")
		assert.Equal(t, partyResult.PartyD.TaxReference1, "CVR-25645774", "they should be equal")
	}
}

func TestPartyService_GetParty(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	partyService := NewPartyService(log, dbService, redisService, userServiceClient)
	party, err := GetParty(uint32(8), []byte{190, 91, 194, 144, 123, 172, 72, 187, 162, 17, 243, 250, 90, 58, 179, 174}, "be5bc290-7bac-48bb-a211-f3fa5a3ab3ae", "Asseco Denmark", "CVR-25645774", "CVR-25645774", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkFzaW", uint32(8), "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	partyResponse := partyproto.GetPartyResponse{}
	partyResponse.Party = party

	form := partyproto.GetPartyRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "be5bc290-7bac-48bb-a211-f3fa5a3ab3ae"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx context.Context
		in  *partyproto.GetPartyRequest
	}
	tests := []struct {
		ps      *PartyService
		args    args
		want    *partyproto.GetPartyResponse
		wantErr bool
	}{
		{
			ps: partyService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &partyResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		partyResp, err := tt.ps.GetParty(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("PartyService.GetParty() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(partyResp, tt.want) {
			t.Errorf("PartyService.GetParty() = %v, want %v", partyResp, tt.want)
		}
		assert.NotNil(t, partyResp)
		partyResult := partyResp.Party
		assert.Equal(t, partyResult.PartyD.PartyName, "Asseco Denmark", "they should be equal")
		assert.Equal(t, partyResult.PartyD.TaxReference1, "CVR-25645774", "they should be equal")
	}
}

func TestPartyService_GetPartyByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	partyService := NewPartyService(log, dbService, redisService, userServiceClient)
	party, err := GetParty(uint32(8), []byte{190, 91, 194, 144, 123, 172, 72, 187, 162, 17, 243, 250, 90, 58, 179, 174}, "be5bc290-7bac-48bb-a211-f3fa5a3ab3ae", "Asseco Denmark", "CVR-25645774", "CVR-25645774", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkFzaW", uint32(8), "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}
	partyResponse := partyproto.GetPartyByPkResponse{}
	partyResponse.Party = party

	form := partyproto.GetPartyByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(8)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx context.Context
		in  *partyproto.GetPartyByPkRequest
	}
	tests := []struct {
		ps      *PartyService
		args    args
		want    *partyproto.GetPartyByPkResponse
		wantErr bool
	}{
		{
			ps: partyService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &partyResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		partyResp, err := tt.ps.GetPartyByPk(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("PartyService.GetPartyByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(partyResp, tt.want) {
			t.Errorf("PartyService.GetPartyByPk() = %v, want %v", partyResp, tt.want)
		}
		assert.NotNil(t, partyResp)
		partyResult := partyResp.Party
		assert.Equal(t, partyResult.PartyD.PartyName, "Asseco Denmark", "they should be equal")
		assert.Equal(t, partyResult.PartyD.TaxReference1, "CVR-25645774", "they should be equal")
	}
}

func TestPartyService_UpdateParty(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	partyService := NewPartyService(log, dbService, redisService, userServiceClient)
	form := partyproto.UpdatePartyRequest{}
	form.PartyName = "Consortial New"
	form.Id = "be5bc290-7bac-48bb-a211-f3fa5a3ab3ae"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	partyResponse := partyproto.UpdatePartyResponse{}

	type args struct {
		ctx context.Context
		in  *partyproto.UpdatePartyRequest
	}
	tests := []struct {
		ps      *PartyService
		args    args
		want    *partyproto.UpdatePartyResponse
		wantErr bool
	}{
		{
			ps: partyService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &partyResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.ps.UpdateParty(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("PartyService.UpdateParty() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("PartyService.UpdateParty() = %v, want %v", got, tt.want)
		}
	}
}

func TestPartyService_DeleteParty(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	partyService := NewPartyService(log, dbService, redisService, userServiceClient)

	form := partyproto.DeletePartyRequest{}

	gform := commonproto.GetRequest{}
	gform.Id = "be5bc290-7bac-48bb-a211-f3fa5a3ab3ae"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	partyResponse := partyproto.DeletePartyResponse{}

	type args struct {
		ctx context.Context
		in  *partyproto.DeletePartyRequest
	}
	tests := []struct {
		ps      *PartyService
		args    args
		want    *partyproto.DeletePartyResponse
		wantErr bool
	}{
		{
			ps: partyService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &partyResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.ps.DeleteParty(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("PartyService.DeleteParty() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("PartyService.DeleteParty() = %v, want %v", got, tt.want)
		}
	}
}

func TestPartyService_GetPartyContactDetail(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	partyService := NewPartyService(log, dbService, redisService, userServiceClient)

	partyContactDetail, err := GetPartyContactDetail(uint32(1), []byte{178, 77, 9, 158, 166, 246, 64, 78, 176, 130, 119, 111, 127, 88, 144, 35}, "b24d099e-a6f6-404e-b082-776f7f589023", uint32(1), "DCSA", "info@dcsa.org", "+31123456789", "https://www.dcsa.org", "active", "2019-07-23T10:04:26Z", "2019-07-23T10:04:26Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	partyContactDetailResponse := partyproto.GetPartyContactDetailResponse{}
	partyContactDetailResponse.PartyContactDetail = partyContactDetail

	form := partyproto.GetPartyContactDetailRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "b24d099e-a6f6-404e-b082-776f7f589023"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx context.Context
		in  *partyproto.GetPartyContactDetailRequest
	}
	tests := []struct {
		ps      *PartyService
		args    args
		want    *partyproto.GetPartyContactDetailResponse
		wantErr bool
	}{
		{
			ps: partyService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &partyContactDetailResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		partyContactDetailResp, err := tt.ps.GetPartyContactDetail(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("PartyService.GetPartyContact() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(partyContactDetailResp, tt.want) {
			t.Errorf("PartyService.GetPartyContact() = %v, want %v", partyContactDetailResp, tt.want)
		}
		assert.NotNil(t, partyContactDetailResp)
		partyContactDetailResult := partyContactDetailResp.PartyContactDetail
		assert.Equal(t, partyContactDetailResult.PartyContactDetailD.Name, "DCSA", "they should be equal")
		assert.Equal(t, partyContactDetailResult.PartyContactDetailD.Email, "info@dcsa.org", "they should be equal")
		assert.Equal(t, partyContactDetailResult.PartyContactDetailD.Phone, "+31123456789", "they should be equal")
		assert.Equal(t, partyContactDetailResult.PartyContactDetailD.Url, "https://www.dcsa.org", "they should be equal")
	}
}

func TestPartyService_CreateParty(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	partyService := NewPartyService(log, dbService, redisService, userServiceClient)

	party := partyproto.CreatePartyRequest{}
	party.PartyName = "Asseco Denmark"
	party.TaxReference1 = "CVR-25645774"
	party.TaxReference2 = "CVR-25645774"
	party.PublicKey = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkFzaW"
	party.Name1 = "Henrik"
	party.Street = "Kronprincessegade"
	party.StreetNumber = "54"
	party.Floor1 = "1"
	party.PostalCode = "København"
	party.City = "Oslo"
	party.StateRegion = "State1"
	party.CountryName = "Norway"
	party.UserId = "auth0|66fd06d0bfea78a82bb42459"
	party.UserEmail = "sprov300@gmail.com"
	party.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *partyproto.CreatePartyRequest
	}
	tests := []struct {
		ps      *PartyService
		args    args
		wantErr bool
	}{
		{
			ps: partyService,
			args: args{
				ctx: ctx,
				in:  &party,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		partyResp, err := tt.ps.CreateParty(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("PartyService.CreateParty() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, partyResp)
		partyResult := partyResp.Party
		assert.Equal(t, partyResult.PartyD.PartyName, "Asseco Denmark", "they should be equal")
		assert.Equal(t, partyResult.PartyD.TaxReference1, "CVR-25645774", "they should be equal")
	}
}

func GetParty(id uint32, uuid4 []byte, idS string, partyName string, taxReference1 string, taxReference2 string, publicKey string, addressId uint32, statusCode string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*partyproto.Party, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	partyD := new(partyproto.PartyD)
	partyD.Id = id
	partyD.Uuid4 = uuid4
	partyD.IdS = idS
	partyD.PartyName = partyName
	partyD.TaxReference1 = taxReference1
	partyD.TaxReference2 = taxReference2
	partyD.PublicKey = publicKey
	partyD.AddressId = addressId

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	party := partyproto.Party{PartyD: partyD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}
	return &party, nil
}

func GetPartyContactDetail(id uint32, uuid4 []byte, idS string, partyId uint32, name string, email string, phone string, url string, statusCode string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*partyproto.PartyContactDetail, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	partyContactDetailD := new(partyproto.PartyContactDetailD)
	partyContactDetailD.Id = id
	partyContactDetailD.Uuid4 = uuid4
	partyContactDetailD.IdS = idS
	partyContactDetailD.PartyId = partyId
	partyContactDetailD.Name = name
	partyContactDetailD.Email = email
	partyContactDetailD.Phone = phone
	partyContactDetailD.Url = url

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	partyContactDetail := partyproto.PartyContactDetail{PartyContactDetailD: partyContactDetailD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &partyContactDetail, nil
}
