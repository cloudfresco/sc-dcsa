package ovsservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestVesselScheduleService_CreateVessel(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	vesselService := NewVesselScheduleService(log, dbService, redisService, userServiceClient)
	vessel := ovsproto.CreateVesselRequest{}
	vessel.VesselImoNumber = "9811000"
	vessel.VesselName = "Ever Given"
	vessel.VesselFlag = "PA"
	vessel.VesselCallSign = "H3RC"
	vessel.IsDummyVessel = false
	vessel.VesselOperatorCarrierCode = "EMC"
	vessel.VesselOperatorCarrierCodeListProvider = ""
	vessel.VesselLength = float64(0)
	vessel.VesselWidth = float64(0)
	vessel.DimensionUnit = ""
	vessel.UserId = "auth0|66fd06d0bfea78a82bb42459"
	vessel.UserEmail = "sprov300@gmail.com"
	vessel.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *ovsproto.CreateVesselRequest
	}
	tests := []struct {
		vs      *VesselScheduleService
		args    args
		want    *ovsproto.CreateVesselResponse
		wantErr bool
	}{
		{
			vs: vesselService,
			args: args{
				ctx: ctx,
				in:  &vessel,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		vesselResponse, err := tt.vs.CreateVessel(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("VesselScheduleService.CreateVessel() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		vesselResult := vesselResponse.Vessel
		assert.NotNil(t, vesselResult)
		assert.Equal(t, vesselResult.VesselD.VesselImoNumber, "9811000", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselFlag, "PA", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselOperatorCarrierCode, "EMC", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselName, "Ever Given", "they should be equal")
		assert.NotNil(t, vesselResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, vesselResult.CrUpdTime.UpdatedAt)
	}
}

func TestVesselScheduleService_GetVessels(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	vesselService := NewVesselScheduleService(log, dbService, redisService, userServiceClient)
	vessel1, err := GetVessel(uint32(8), []byte{248, 169, 122, 57, 135, 148, 68, 41, 157, 203, 85, 9, 182, 166, 203, 185}, "f8a97a39-8794-4429-9dcb-5509b6a6cbb9", "9372872", "BURGUNDY", "DK", "", false, "CMA", "", float64(0), float64(0), "", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	vessel2, err := GetVessel(uint32(7), []byte{76, 125, 40, 176, 41, 41, 66, 190, 175, 157, 160, 155, 33, 30, 43, 149}, "4c7d28b0-2929-42be-af9d-a09b211e2b95", "9776418", "CC ANTOINE DE ST EXUPERY", "DK", "", false, "CMA", "", float64(0), float64(0), "", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	vessels := []*ovsproto.Vessel{}
	vessels = append(vessels, vessel1, vessel2)

	form := ovsproto.GetVesselsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "Ng=="
	vesselResp := ovsproto.GetVesselsResponse{Vessels: vessels, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *ovsproto.GetVesselsRequest
	}
	tests := []struct {
		vs      *VesselScheduleService
		args    args
		want    *ovsproto.GetVesselsResponse
		wantErr bool
	}{
		{
			vs: vesselService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &vesselResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		vesselResponse, err := tt.vs.GetVessels(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("VesselScheduleService.GetVessels() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(vesselResponse, tt.want) {
			t.Errorf("VesselScheduleService.GetVessels() = %v, want %v", vesselResponse, tt.want)
		}
		vesselResult := vesselResponse.Vessels[0]
		assert.NotNil(t, vesselResult)
		assert.Equal(t, vesselResult.VesselD.VesselImoNumber, "9372872", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselFlag, "DK", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselOperatorCarrierCode, "CMA", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselName, "BURGUNDY", "they should be equal")
		assert.NotNil(t, vesselResult.CrUpdTime.CreatedAt)

	}
}

func TestVesselScheduleService_GetVessel(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	vesselService := NewVesselScheduleService(log, dbService, redisService, userServiceClient)
	vessel, err := GetVessel(uint32(7), []byte{76, 125, 40, 176, 41, 41, 66, 190, 175, 157, 160, 155, 33, 30, 43, 149}, "4c7d28b0-2929-42be-af9d-a09b211e2b95", "9776418", "CC ANTOINE DE ST EXUPERY", "DK", "", false, "CMA", "", float64(0), float64(0), "", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	vesselResp := ovsproto.GetVesselResponse{}
	vesselResp.Vessel = vessel

	gform := commonproto.GetRequest{}
	gform.Id = "4c7d28b0-2929-42be-af9d-a09b211e2b95"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := ovsproto.GetVesselRequest{}
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *ovsproto.GetVesselRequest
	}
	tests := []struct {
		vs      *VesselScheduleService
		args    args
		want    *ovsproto.GetVesselResponse
		wantErr bool
	}{
		{
			vs: vesselService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &vesselResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		vesselResponse, err := tt.vs.GetVessel(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("VesselScheduleService.GetVessel() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(vesselResponse, tt.want) {
			t.Errorf("VesselScheduleService.GetVessel() = %v, want %v", vesselResponse, tt.want)
		}
		vesselResult := vesselResponse.Vessel
		assert.NotNil(t, vesselResult)
		assert.Equal(t, vesselResult.VesselD.VesselImoNumber, "9776418", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselFlag, "DK", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselOperatorCarrierCode, "CMA", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselName, "CC ANTOINE DE ST EXUPERY", "they should be equal")
		assert.NotNil(t, vesselResult.CrUpdTime.CreatedAt)
	}
}

func TestVesselScheduleService_GetVesselByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	vesselService := NewVesselScheduleService(log, dbService, redisService, userServiceClient)
	vessel, err := GetVessel(uint32(7), []byte{76, 125, 40, 176, 41, 41, 66, 190, 175, 157, 160, 155, 33, 30, 43, 149}, "4c7d28b0-2929-42be-af9d-a09b211e2b95", "9776418", "CC ANTOINE DE ST EXUPERY", "DK", "", false, "CMA", "", float64(0), float64(0), "", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	vesselResp := ovsproto.GetVesselByPkResponse{}
	vesselResp.Vessel = vessel

	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(7)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := ovsproto.GetVesselByPkRequest{}
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *ovsproto.GetVesselByPkRequest
	}
	tests := []struct {
		vs      *VesselScheduleService
		args    args
		want    *ovsproto.GetVesselByPkResponse
		wantErr bool
	}{
		{
			vs: vesselService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &vesselResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		vesselResponse, err := tt.vs.GetVesselByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("VesselScheduleService.GetVesselByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(vesselResponse, tt.want) {
			t.Errorf("VesselScheduleService.GetVesselByPk() = %v, want %v", vesselResponse, tt.want)
		}
		vesselResult := vesselResponse.Vessel
		assert.NotNil(t, vesselResult)
		assert.Equal(t, vesselResult.VesselD.VesselImoNumber, "9776418", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselFlag, "DK", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselOperatorCarrierCode, "CMA", "they should be equal")
		assert.Equal(t, vesselResult.VesselD.VesselName, "CC ANTOINE DE ST EXUPERY", "they should be equal")
		assert.NotNil(t, vesselResult.CrUpdTime.CreatedAt)
	}
}

func GetVessel(id uint32, uuid4 []byte, idS string, vesselImoNumber string, vesselName string, vesselFlag string, vesselCallSign string, isDummyVessel bool, vesselOperatorCarrierCode string, vesselOperatorCarrierCodeListProvider string, vesselLength float64, vesselWidth float64, dimensionUnit string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*ovsproto.Vessel, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	vesselD := new(ovsproto.VesselD)
	vesselD.Id = id
	vesselD.Uuid4 = uuid4
	vesselD.IdS = idS
	vesselD.VesselImoNumber = vesselImoNumber
	vesselD.VesselName = vesselName
	vesselD.VesselFlag = vesselFlag
	vesselD.VesselCallSign = vesselCallSign
	vesselD.IsDummyVessel = isDummyVessel
	vesselD.VesselOperatorCarrierCode = vesselOperatorCarrierCode
	vesselD.VesselOperatorCarrierCodeListProvider = vesselOperatorCarrierCodeListProvider
	vesselD.VesselLength = vesselLength
	vesselD.VesselWidth = vesselWidth
	vesselD.DimensionUnit = dimensionUnit

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	vessel := ovsproto.Vessel{VesselD: vesselD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &vessel, nil
}
