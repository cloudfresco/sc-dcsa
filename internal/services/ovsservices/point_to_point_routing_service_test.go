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

func TestLegService_CreatePointToPointRouting(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	legService := NewLegService(log, dbService, redisService, userServiceClient)
	pointToPointRouting := ovsproto.CreatePointToPointRoutingRequest{}
	pointToPointRouting.SequenceNumber = int32(30)
	pointToPointRouting.PlaceOfReceiptId = uint32(3)
	pointToPointRouting.PlaceOfDeliveryId = uint32(3)
	pointToPointRouting.UserId = "auth0|66fd06d0bfea78a82bb42459"
	pointToPointRouting.UserEmail = "sprov300@gmail.com"
	pointToPointRouting.RequestId = "bks1m1g91jau4nkks2f0"

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

	legs := []*ovsproto.CreateLegRequest{}
	legs = append(legs, &leg)
	pointToPointRouting.Legs = legs

	type args struct {
		ctx context.Context
		in  *ovsproto.CreatePointToPointRoutingRequest
	}
	tests := []struct {
		lg      *LegService
		args    args
		want    *ovsproto.CreatePointToPointRoutingResponse
		wantErr bool
	}{
		{
			lg: legService,
			args: args{
				ctx: ctx,
				in:  &pointToPointRouting,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		pointToPointRoutingResp, err := tt.lg.CreatePointToPointRouting(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("LegService.CreatePointToPointRouting() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		pointToPointRoutingResult := pointToPointRoutingResp.PointToPointRouting
		assert.NotNil(t, pointToPointRoutingResult)
		assert.Equal(t, pointToPointRoutingResult.PointToPointRoutingD.SequenceNumber, int32(30), "they should be equal")
	}
}

func TestLegService_GetPointToPointRoutings(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	legService := NewLegService(log, dbService, redisService, userServiceClient)

	pointToPointRouting1, err := GetPointToPointRouting(uint32(2), []byte{72, 35, 14, 219, 43, 135, 79, 38, 169, 100, 21, 9, 95, 239, 215, 164}, "48230edb-2b87-4f26-a964-15095fefd7a4", 20, uint32(2), uint32(2), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	pointToPointRouting2, err := GetPointToPointRouting(uint32(1), []byte{223, 157, 107, 204, 249, 131, 77, 138, 174, 98, 55, 186, 93, 87, 44, 76}, "df9d6bcc-f983-4d8a-ae62-37ba5d572c4c", 10, uint32(1), uint32(1), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	pointToPointRoutings := []*ovsproto.PointToPointRouting{}
	pointToPointRoutings = append(pointToPointRoutings, pointToPointRouting1, pointToPointRouting2)

	form := ovsproto.GetPointToPointRoutingsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	pointToPointRoutingResp := ovsproto.GetPointToPointRoutingsResponse{PointToPointRoutings: pointToPointRoutings, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *ovsproto.GetPointToPointRoutingsRequest
	}
	tests := []struct {
		lg      *LegService
		args    args
		want    *ovsproto.GetPointToPointRoutingsResponse
		wantErr bool
	}{
		{
			lg: legService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &pointToPointRoutingResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		pointToPointRoutingResponse, err := tt.lg.GetPointToPointRoutings(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("LegService.GetPointToPointRoutings() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(pointToPointRoutingResponse, tt.want) {
			t.Errorf("LegService.GetPointToPointRoutings() = %v, want %v", pointToPointRoutingResponse, tt.want)
		}
		pointToPointRoutingResult := pointToPointRoutingResponse.PointToPointRoutings[0]
		assert.NotNil(t, pointToPointRoutingResult)
		assert.Equal(t, pointToPointRoutingResult.PointToPointRoutingD.SequenceNumber, int32(20), "they should be equal")
	}
}

func TestLegService_GetPointToPointRouting(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	legService := NewLegService(log, dbService, redisService, userServiceClient)

	pointToPointRouting, err := GetPointToPointRouting(uint32(1), []byte{223, 157, 107, 204, 249, 131, 77, 138, 174, 98, 55, 186, 93, 87, 44, 76}, "df9d6bcc-f983-4d8a-ae62-37ba5d572c4c", 10, uint32(1), uint32(1), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	pointToPointRoutingResp := ovsproto.GetPointToPointRoutingResponse{}
	pointToPointRoutingResp.PointToPointRouting = pointToPointRouting

	gform := commonproto.GetRequest{}
	gform.Id = "df9d6bcc-f983-4d8a-ae62-37ba5d572c4c"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := ovsproto.GetPointToPointRoutingRequest{}
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *ovsproto.GetPointToPointRoutingRequest
	}
	tests := []struct {
		lg      *LegService
		args    args
		want    *ovsproto.GetPointToPointRoutingResponse
		wantErr bool
	}{
		{
			lg: legService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &pointToPointRoutingResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		pointToPointRoutingResponse, err := tt.lg.GetPointToPointRouting(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("LegService.GetPointToPointRouting() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(pointToPointRoutingResponse, tt.want) {
			t.Errorf("LegService.GetPointToPointRouting() = %v, want %v", pointToPointRoutingResponse, tt.want)
		}
		pointToPointRoutingResult := pointToPointRoutingResponse.PointToPointRouting
		assert.NotNil(t, pointToPointRoutingResult)
		assert.Equal(t, pointToPointRoutingResult.PointToPointRoutingD.SequenceNumber, int32(10), "they should be equal")

	}
}

func TestLegService_GetPointToPointRoutingByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	legService := NewLegService(log, dbService, redisService, userServiceClient)

	pointToPointRouting, err := GetPointToPointRouting(uint32(2), []byte{72, 35, 14, 219, 43, 135, 79, 38, 169, 100, 21, 9, 95, 239, 215, 164}, "48230edb-2b87-4f26-a964-15095fefd7a4", 20, uint32(2), uint32(2), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	pointToPointRoutingResp := ovsproto.GetPointToPointRoutingByPkResponse{}
	pointToPointRoutingResp.PointToPointRouting = pointToPointRouting

	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(2)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := ovsproto.GetPointToPointRoutingByPkRequest{}
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *ovsproto.GetPointToPointRoutingByPkRequest
	}
	tests := []struct {
		lg      *LegService
		args    args
		want    *ovsproto.GetPointToPointRoutingByPkResponse
		wantErr bool
	}{
		{
			lg: legService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &pointToPointRoutingResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		pointToPointRoutingResponse, err := tt.lg.GetPointToPointRoutingByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("LegService.GetPointToPointRoutingByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(pointToPointRoutingResponse, tt.want) {
			t.Errorf("LegService.GetPointToPointRoutingByPk() = %v, want %v", pointToPointRoutingResponse, tt.want)
		}
		pointToPointRoutingResult := pointToPointRoutingResponse.PointToPointRouting
		assert.NotNil(t, pointToPointRoutingResult)
		assert.Equal(t, pointToPointRoutingResult.PointToPointRoutingD.SequenceNumber, int32(20), "they should be equal")

	}
}

func GetPointToPointRouting(id uint32, uuid4 []byte, idS string, sequenceNumber int32, placeOfReceiptId uint32, placeOfDeliveryId uint32, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*ovsproto.PointToPointRouting, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	pointToPointRoutingD := new(ovsproto.PointToPointRoutingD)
	pointToPointRoutingD.Id = id
	pointToPointRoutingD.Uuid4 = uuid4
	pointToPointRoutingD.IdS = idS
	pointToPointRoutingD.SequenceNumber = sequenceNumber
	pointToPointRoutingD.PlaceOfReceiptId = placeOfReceiptId
	pointToPointRoutingD.PlaceOfDeliveryId = placeOfDeliveryId

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	pointToPointRouting := ovsproto.PointToPointRouting{PointToPointRoutingD: pointToPointRoutingD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &pointToPointRouting, nil
}
