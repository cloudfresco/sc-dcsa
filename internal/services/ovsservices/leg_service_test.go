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

func TestLegService_CreateLeg(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	legService := NewLegService(log, dbService, redisService, userServiceClient)

	leg := ovsproto.CreateLegRequest{}
	leg.SequenceNumber = int32(3)
	leg.ModeOfTransport = "VESSEL"
	leg.VesselOperatorSmdgLinerCode = "MSC"
	leg.VesselImoNumber = "9930038"
	leg.VesselName = "MSC TESSA"
	leg.CarrierServiceName = "Singapore Express"
	leg.UniversalServiceReference = "SR12365B"
	leg.CarrierServiceCode = "AE55"
	leg.UniversalImportVoyageReference = "2401E"
	leg.UniversalExportVoyageReference = "2401E"
	leg.CarrierImportVoyageNumber = "401E"
	leg.CarrierExportVoyageNumber = "401E"
	leg.DepartureId = uint32(1)
	leg.ArrivalId = uint32(1)
	leg.PointToPointRoutingId = uint32(2)
	leg.UserId = "auth0|66fd06d0bfea78a82bb42459"
	leg.UserEmail = "sprov300@gmail.com"
	leg.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *ovsproto.CreateLegRequest
	}
	tests := []struct {
		lg      *LegService
		args    args
		want    *ovsproto.CreateLegResponse
		wantErr bool
	}{
		{
			lg: legService,
			args: args{
				ctx: ctx,
				in:  &leg,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		legResp, err := tt.lg.CreateLeg(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("LegService.CreateLeg() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		legResult := legResp.Leg
		assert.NotNil(t, legResult)
		assert.Equal(t, legResult.LegD.ModeOfTransport, "VESSEL", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselOperatorSmdgLinerCode, "MSC", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselImoNumber, "9930038", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselName, "MSC TESSA", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierServiceName, "Singapore Express", "they should be equal")
		assert.Equal(t, legResult.LegD.UniversalServiceReference, "SR12365B", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierServiceCode, "AE55", "they should be equal")
		assert.Equal(t, legResult.LegD.UniversalImportVoyageReference, "2401E", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierImportVoyageNumber, "401E", "they should be equal")
	}
}

func TestLegService_GetLegs(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	legService := NewLegService(log, dbService, redisService, userServiceClient)
	leg1, err := GetLeg(uint32(3), []byte{17, 123, 72, 86, 97, 22, 76, 170, 159, 223, 239, 161, 106, 184, 147, 0}, "117b4856-6116-4caa-9fdf-efa16ab89300", 2, "VESSEL", "MSC", "9930038", "MSC TESSA", "Singapore Express", "SR12365B", "AE55", "2401E", "2401E", "401E", "401E", uint32(1), uint32(1), uint32(2), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	leg2, err := GetLeg(uint32(2), []byte{81, 48, 52, 235, 253, 144, 69, 230, 155, 25, 115, 237, 161, 75, 81, 56}, "513034eb-fd90-45e6-9b19-73eda14b5138", 1, "BARGE", "", "", "", "", "", "", "", "", "", "", uint32(1), uint32(1), uint32(2), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	legs := []*ovsproto.Leg{}
	legs = append(legs, leg1, leg2)

	form := ovsproto.GetLegsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MQ=="
	legResp := ovsproto.GetLegsResponse{Legs: legs, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *ovsproto.GetLegsRequest
	}
	tests := []struct {
		lg      *LegService
		args    args
		want    *ovsproto.GetLegsResponse
		wantErr bool
	}{
		{
			lg: legService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &legResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		legResponse, err := tt.lg.GetLegs(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("LegService.GetLegs() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(legResponse, tt.want) {
			t.Errorf("LegService.GetLegs() = %v, want %v", legResponse, tt.want)
		}

		legResult := legResponse.Legs[0]
		assert.NotNil(t, legResult)
		assert.Equal(t, legResult.LegD.ModeOfTransport, "VESSEL", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselOperatorSmdgLinerCode, "MSC", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselImoNumber, "9930038", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselName, "MSC TESSA", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierServiceName, "Singapore Express", "they should be equal")
		assert.Equal(t, legResult.LegD.UniversalServiceReference, "SR12365B", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierServiceCode, "AE55", "they should be equal")
		assert.Equal(t, legResult.LegD.UniversalImportVoyageReference, "2401E", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierImportVoyageNumber, "401E", "they should be equal")

	}
}

func TestLegService_GetLeg(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	legService := NewLegService(log, dbService, redisService, userServiceClient)

	leg, err := GetLeg(uint32(1), []byte{90, 126, 16, 146, 24, 40, 70, 44, 141, 44, 115, 90, 108, 181, 199, 115}, "5a7e1092-1828-462c-8d2c-735a6cb5c773", 1, "VESSEL", "HLC", "9321483", "King of the Seas", "Great Lion Service", "SR12345A", "FE1", "2103N", "2103N", "2103S", "2103N", uint32(1), uint32(1), uint32(1), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	legResp := ovsproto.GetLegResponse{}
	legResp.Leg = leg

	gform := commonproto.GetRequest{}
	gform.Id = "5a7e1092-1828-462c-8d2c-735a6cb5c773"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := ovsproto.GetLegRequest{}
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *ovsproto.GetLegRequest
	}
	tests := []struct {
		lg      *LegService
		args    args
		want    *ovsproto.GetLegResponse
		wantErr bool
	}{
		{
			lg: legService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &legResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		legResponse, err := tt.lg.GetLeg(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("LegService.GetLeg() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(legResponse, tt.want) {
			t.Errorf("LegService.GetLeg() = %v, want %v", legResponse, tt.want)
		}
		legResult := legResponse.Leg
		assert.NotNil(t, legResult)
		assert.Equal(t, legResult.LegD.ModeOfTransport, "VESSEL", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselOperatorSmdgLinerCode, "HLC", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselImoNumber, "9321483", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselName, "King of the Seas", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierServiceName, "Great Lion Service", "they should be equal")
		assert.Equal(t, legResult.LegD.UniversalServiceReference, "SR12345A", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierServiceCode, "FE1", "they should be equal")
		assert.Equal(t, legResult.LegD.UniversalImportVoyageReference, "2103N", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierImportVoyageNumber, "2103S", "they should be equal")

	}
}

func TestLegService_GetLegByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	legService := NewLegService(log, dbService, redisService, userServiceClient)

	leg, err := GetLeg(uint32(1), []byte{90, 126, 16, 146, 24, 40, 70, 44, 141, 44, 115, 90, 108, 181, 199, 115}, "5a7e1092-1828-462c-8d2c-735a6cb5c773", 1, "VESSEL", "HLC", "9321483", "King of the Seas", "Great Lion Service", "SR12345A", "FE1", "2103N", "2103N", "2103S", "2103N", uint32(1), uint32(1), uint32(1), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	legResp := ovsproto.GetLegByPkResponse{}
	legResp.Leg = leg

	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(1)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := ovsproto.GetLegByPkRequest{}
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *ovsproto.GetLegByPkRequest
	}
	tests := []struct {
		lg      *LegService
		args    args
		want    *ovsproto.GetLegByPkResponse
		wantErr bool
	}{
		{
			lg: legService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &legResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		legResponse, err := tt.lg.GetLegByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("LegService.GetLegByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(legResponse, tt.want) {
			t.Errorf("LegService.GetLegByPk() = %v, want %v", legResponse, tt.want)
		}
		legResult := legResponse.Leg
		assert.NotNil(t, legResult)
		assert.Equal(t, legResult.LegD.ModeOfTransport, "VESSEL", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselOperatorSmdgLinerCode, "HLC", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselImoNumber, "9321483", "they should be equal")
		assert.Equal(t, legResult.LegD.VesselName, "King of the Seas", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierServiceName, "Great Lion Service", "they should be equal")
		assert.Equal(t, legResult.LegD.UniversalServiceReference, "SR12345A", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierServiceCode, "FE1", "they should be equal")
		assert.Equal(t, legResult.LegD.UniversalImportVoyageReference, "2103N", "they should be equal")
		assert.Equal(t, legResult.LegD.CarrierImportVoyageNumber, "2103S", "they should be equal")
	}
}

func GetLeg(id uint32, uuid4 []byte, idS string, sequenceNumber int32, modeOfTransport string, vesselOperatorSmdgLinerCode string, vesselImoNumber string, vesselName string, carrierServiceName string, universalServiceReference string, carrierServiceCode string, universalImportVoyageReference string, universalExportVoyageReference string, carrierImportVoyageNumber string, carrierExportVoyageNumber string, departureId uint32, arrivalId uint32, pointToPointRoutingId uint32, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*ovsproto.Leg, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	legD := new(ovsproto.LegD)
	legD.Id = id
	legD.Uuid4 = uuid4
	legD.IdS = idS
	legD.SequenceNumber = sequenceNumber
	legD.ModeOfTransport = modeOfTransport
	legD.VesselOperatorSmdgLinerCode = vesselOperatorSmdgLinerCode
	legD.VesselImoNumber = vesselImoNumber
	legD.VesselName = vesselName
	legD.CarrierServiceName = carrierServiceName
	legD.UniversalServiceReference = universalServiceReference
	legD.CarrierServiceCode = carrierServiceCode
	legD.UniversalImportVoyageReference = universalImportVoyageReference
	legD.UniversalExportVoyageReference = universalExportVoyageReference
	legD.CarrierImportVoyageNumber = carrierImportVoyageNumber
	legD.CarrierExportVoyageNumber = carrierExportVoyageNumber
	legD.DepartureId = departureId
	legD.ArrivalId = arrivalId
	legD.PointToPointRoutingId = pointToPointRoutingId

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	leg := ovsproto.Leg{LegD: legD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &leg, nil
}
