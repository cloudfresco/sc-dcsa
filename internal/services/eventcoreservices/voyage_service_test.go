package eventcoreservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestVoyageService_CreateVoyage(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	voyageService := NewVoyageService(log, dbService, redisService, userServiceClient)

	voyage := eventcoreproto.CreateVoyageRequest{}
	voyage.CarrierVoyageNumber = "2219E"
	voyage.UniversalVoyageReference = "UVR02"
	voyage.ServiceId = uint32(2)
	voyage.UserId = "auth0|66fd06d0bfea78a82bb42459"
	voyage.UserEmail = "sprov300@gmail.com"
	voyage.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eventcoreproto.CreateVoyageRequest
	}
	tests := []struct {
		vs      *VoyageService
		args    args
		want    *eventcoreproto.CreateVoyageResponse
		wantErr bool
	}{
		{
			vs: voyageService,
			args: args{
				ctx: ctx,
				in:  &voyage,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		voyageResp, err := tt.vs.CreateVoyage(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("VoyageService.CreateVoyage() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		voyageResult := voyageResp.Voyage
		assert.NotNil(t, voyageResult)
		assert.Equal(t, voyageResult.VoyageD.CarrierVoyageNumber, "2219E", "they should be equal")
		assert.Equal(t, voyageResult.VoyageD.UniversalVoyageReference, "UVR02", "they should be equal")
		assert.Equal(t, voyageResult.VoyageD.ServiceId, uint32(2), "they should be equal")
		assert.NotNil(t, voyageResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, voyageResult.CrUpdTime.UpdatedAt)
	}
}

func TestVoyageService_GetVoyages(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	voyageService := NewVoyageService(log, dbService, redisService, userServiceClient)

	voyage1, err := GetVoyage(uint32(13), []byte{111, 3, 76, 150, 151, 228, 71, 198, 161, 64, 196, 78, 155, 229, 6, 9}, "6f034c96-97e4-47c6-a140-c44e9be50609", "4420E", "UVR02", uint32(2), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	voyage2, err := GetVoyage(uint32(12), []byte{63, 176, 185, 25, 243, 140, 65, 152, 182, 31, 176, 140, 54, 24, 88, 247}, "3fb0b919-f38c-4198-b61f-b08c361858f7", "4419W", "UVR01", uint32(2), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	voyages := []*eventcoreproto.Voyage{}
	voyages = append(voyages, voyage1, voyage2)

	form := eventcoreproto.GetVoyagesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MTE="
	voyageResponse := eventcoreproto.GetVoyagesResponse{Voyages: voyages, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *eventcoreproto.GetVoyagesRequest
	}
	tests := []struct {
		vs      *VoyageService
		args    args
		want    *eventcoreproto.GetVoyagesResponse
		wantErr bool
	}{
		{
			vs: voyageService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &voyageResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		voyageResp, err := tt.vs.GetVoyages(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("VoyageService.GetVoyages() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(voyageResp, tt.want) {
			t.Errorf("VoyageService.GetVoyages() = %v, want %v", voyageResp, tt.want)
		}

		voyageResult := voyageResp.Voyages[1]
		assert.NotNil(t, voyageResult)
		assert.Equal(t, voyageResult.VoyageD.CarrierVoyageNumber, "4419W", "they should be equal")
		assert.Equal(t, voyageResult.VoyageD.UniversalVoyageReference, "UVR01", "they should be equal")
		assert.Equal(t, voyageResult.VoyageD.ServiceId, uint32(2), "they should be equal")
		assert.NotNil(t, voyageResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, voyageResult.CrUpdTime.UpdatedAt)
	}
}

func TestVoyageService_FindByCarrierVoyageNumberAndServiceId(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	voyageService := NewVoyageService(log, dbService, redisService, userServiceClient)

	voyage, err := GetVoyage(uint32(13), []byte{111, 3, 76, 150, 151, 228, 71, 198, 161, 64, 196, 78, 155, 229, 6, 9}, "6f034c96-97e4-47c6-a140-c44e9be50609", "4420E", "UVR02", uint32(2), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	form := eventcoreproto.FindByCarrierVoyageNumberAndServiceIdRequest{}
	form.CarrierVoyageNumber = "4420E"
	form.ServiceId = uint32(2)
	form.UserId = "auth0|66fd06d0bfea78a82bb42459"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	voyageResponse := eventcoreproto.FindByCarrierVoyageNumberAndServiceIdResponse{}
	voyageResponse.Voyage = voyage

	type args struct {
		ctx context.Context
		in  *eventcoreproto.FindByCarrierVoyageNumberAndServiceIdRequest
	}
	tests := []struct {
		vs      *VoyageService
		args    args
		want    *eventcoreproto.FindByCarrierVoyageNumberAndServiceIdResponse
		wantErr bool
	}{
		{
			vs: voyageService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &voyageResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		voyageResp, err := tt.vs.FindByCarrierVoyageNumberAndServiceId(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("VoyageService.FindByCarrierVoyageNumberAndServiceId() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(voyageResp, tt.want) {
			t.Errorf("VoyageService.FindByCarrierVoyageNumberAndServiceId() = %v, want %v", voyageResp, tt.want)
		}
		voyageResult := voyageResp.Voyage
		assert.NotNil(t, voyageResult)
		assert.Equal(t, voyageResult.VoyageD.CarrierVoyageNumber, "4420E", "they should be equal")
		assert.Equal(t, voyageResult.VoyageD.UniversalVoyageReference, "UVR02", "they should be equal")
		assert.Equal(t, voyageResult.VoyageD.ServiceId, uint32(2), "they should be equal")
		assert.NotNil(t, voyageResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, voyageResult.CrUpdTime.UpdatedAt)
	}
}

func TestVoyageService_FindByCarrierVoyageNumber(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	voyageService := NewVoyageService(log, dbService, redisService, userServiceClient)

	voyage, err := GetVoyage(uint32(13), []byte{111, 3, 76, 150, 151, 228, 71, 198, 161, 64, 196, 78, 155, 229, 6, 9}, "6f034c96-97e4-47c6-a140-c44e9be50609", "4420E", "UVR02", uint32(2), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	form := eventcoreproto.FindByCarrierVoyageNumberRequest{}
	form.CarrierVoyageNumber = "4420E"
	form.UserId = "auth0|66fd06d0bfea78a82bb42459"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	voyageResponse := eventcoreproto.FindByCarrierVoyageNumberResponse{}
	voyageResponse.Voyage = voyage

	type args struct {
		ctx context.Context
		in  *eventcoreproto.FindByCarrierVoyageNumberRequest
	}
	tests := []struct {
		vs      *VoyageService
		args    args
		want    *eventcoreproto.FindByCarrierVoyageNumberResponse
		wantErr bool
	}{
		{
			vs: voyageService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &voyageResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		voyageResp, err := tt.vs.FindByCarrierVoyageNumber(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("VoyageService.FindByCarrierVoyageNumber() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(voyageResp, tt.want) {
			t.Errorf("VoyageService.FindByCarrierVoyageNumber() = %v, want %v", voyageResp, tt.want)
		}
		voyageResult := voyageResp.Voyage
		assert.NotNil(t, voyageResult)
		assert.Equal(t, voyageResult.VoyageD.CarrierVoyageNumber, "4420E", "they should be equal")
		assert.Equal(t, voyageResult.VoyageD.UniversalVoyageReference, "UVR02", "they should be equal")
		assert.Equal(t, voyageResult.VoyageD.ServiceId, uint32(2), "they should be equal")
		assert.NotNil(t, voyageResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, voyageResult.CrUpdTime.UpdatedAt)
	}
}

func GetVoyage(id uint32, uuid4 []byte, idS string, carrierVoyageNumber string, universalVoyageReference string, serviceId uint32, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*eventcoreproto.Voyage, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	voyageD := new(eventcoreproto.VoyageD)
	voyageD.Id = id
	voyageD.Uuid4 = uuid4
	voyageD.IdS = idS
	voyageD.CarrierVoyageNumber = carrierVoyageNumber
	voyageD.UniversalVoyageReference = universalVoyageReference
	voyageD.ServiceId = serviceId

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	voyage := eventcoreproto.Voyage{VoyageD: voyageD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &voyage, nil
}
