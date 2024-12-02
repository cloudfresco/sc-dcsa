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

func TestShippingService_GetShipments(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	shippingService := NewShippingService(log, dbService, redisService, userServiceClient)
	shipment1, err := GetShipment(uint32(19), []byte{112, 221, 204, 99, 88, 194, 76, 32, 174, 107, 98, 80, 143, 7, 29, 137}, "70ddcc63-58c2-4c20-ae6b-62508f071d89", uint32(2), uint32(5), "ABC123123123", "TERMS AND CONDITIONS!", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	shipment2, err := GetShipment(uint32(18), []byte{19, 22, 172, 59, 211, 211, 70, 34, 154, 54, 168, 98, 34, 123, 179, 133}, "1316ac3b-d3d3-4622-9a36-a862227bb385", uint32(1), uint32(5), "BR1239719971", "TERMS AND CONDITIONS!", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	shipments := []*eblproto.Shipment{}
	shipments = append(shipments, shipment1, shipment2)

	form := eblproto.GetShipmentsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MTc="
	c1 := eblproto.GetShipmentsResponse{Shipments: shipments, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *eblproto.GetShipmentsRequest
	}
	tests := []struct {
		ss      *ShippingService
		args    args
		want    *eblproto.GetShipmentsResponse
		wantErr bool
	}{
		{
			ss: shippingService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &c1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		shipmentsResp, err := tt.ss.GetShipments(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ShippingService.GetShipments() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(shipmentsResp, tt.want) {
			t.Errorf("ShippingService.GetShipments() = %v, want %v", shipmentsResp, tt.want)
		}
		assert.NotNil(t, shipmentsResp)
		shipmentResult := shipmentsResp.Shipments[1]
		assert.Equal(t, shipmentResult.ShipmentD.CarrierBookingReference, "BR1239719971", "they should be equal")
		assert.Equal(t, shipmentResult.ShipmentD.TermsAndConditions, "TERMS AND CONDITIONS!", "they should be equal")
	}
}

func TestShippingService_GetShipment(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	shippingService := NewShippingService(log, dbService, redisService, userServiceClient)
	shipment, err := GetShipment(uint32(19), []byte{112, 221, 204, 99, 88, 194, 76, 32, 174, 107, 98, 80, 143, 7, 29, 137}, "70ddcc63-58c2-4c20-ae6b-62508f071d89", uint32(2), uint32(5), "ABC123123123", "TERMS AND CONDITIONS!", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	shipResponse := eblproto.GetShipmentResponse{}
	shipResponse.Shipment = shipment

	form := eblproto.GetShipmentRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "70ddcc63-58c2-4c20-ae6b-62508f071d89"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx context.Context
		in  *eblproto.GetShipmentRequest
	}
	tests := []struct {
		ss      *ShippingService
		args    args
		want    *eblproto.GetShipmentResponse
		wantErr bool
	}{
		{
			ss: shippingService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &shipResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		shipmentResp, err := tt.ss.GetShipment(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ShippingService.GetShipment() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(shipmentResp, tt.want) {
			t.Errorf("ShippingService.GetShipment() = %v, want %v", shipmentResp, tt.want)
		}
		assert.NotNil(t, shipmentResp)
		shipmentResult := shipmentResp.Shipment
		assert.Equal(t, shipmentResult.ShipmentD.CarrierBookingReference, "ABC123123123", "they should be equal")
		assert.Equal(t, shipmentResult.ShipmentD.TermsAndConditions, "TERMS AND CONDITIONS!", "they should be equal")
	}
}

func TestShippingService_GetShipmentByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	shippingService := NewShippingService(log, dbService, redisService, userServiceClient)
	shipment, err := GetShipment(uint32(19), []byte{112, 221, 204, 99, 88, 194, 76, 32, 174, 107, 98, 80, 143, 7, 29, 137}, "70ddcc63-58c2-4c20-ae6b-62508f071d89", uint32(2), uint32(5), "ABC123123123", "TERMS AND CONDITIONS!", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "2021-12-12T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	shipResponse := eblproto.GetShipmentByPkResponse{}
	shipResponse.Shipment = shipment

	form := eblproto.GetShipmentByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(19)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx context.Context
		in  *eblproto.GetShipmentByPkRequest
	}
	tests := []struct {
		ss      *ShippingService
		args    args
		want    *eblproto.GetShipmentByPkResponse
		wantErr bool
	}{
		{
			ss: shippingService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &shipResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		shipmentResp, err := tt.ss.GetShipmentByPk(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ShippingService.GetShipmentByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(shipmentResp, tt.want) {
			t.Errorf("ShippingService.GetShipmentByPk() = %v, want %v", shipmentResp, tt.want)
		}
		assert.NotNil(t, shipmentResp)
		shipmentResult := shipmentResp.Shipment
		assert.Equal(t, shipmentResult.ShipmentD.CarrierBookingReference, "ABC123123123", "they should be equal")
		assert.Equal(t, shipmentResult.ShipmentD.TermsAndConditions, "TERMS AND CONDITIONS!", "they should be equal")
	}
}

func TestShippingService_FindCarrierBookingReferenceByShippingInstructionId(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	shippingService := NewShippingService(log, dbService, redisService, userServiceClient)

	form := eblproto.FindCarrierBookingReferenceByShippingInstructionIdRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(1)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	carBkgReference := eblproto.FindCarrierBookingReferenceByShippingInstructionIdResponse{}
	carBkgReference.CarrierBookingReference = "DCR987876762"

	type args struct {
		ctx context.Context
		in  *eblproto.FindCarrierBookingReferenceByShippingInstructionIdRequest
	}
	tests := []struct {
		ss      *ShippingService
		args    args
		want    *eblproto.FindCarrierBookingReferenceByShippingInstructionIdResponse
		wantErr bool
	}{
		{
			ss: shippingService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &carBkgReference,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		carBkgReferenceResp, err := tt.ss.FindCarrierBookingReferenceByShippingInstructionId(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ShippingService.FindCarrierBookingReferenceByShippingInstructionId() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(carBkgReferenceResp, tt.want) {
			t.Errorf("ShippingService.FindCarrierBookingReferenceByShippingInstructionId() = %v, want %v", carBkgReferenceResp, tt.want)
		}
		assert.NotNil(t, carBkgReferenceResp)
		assert.Equal(t, carBkgReferenceResp.CarrierBookingReference, "DCR987876762", "they should be equal")
	}
}

func GetShipment(id uint32, uuid4 []byte, idS string, bookingId uint32, carrierId uint32, carrierBookingReference string, termsAndConditions string, ConfirmationDatetime string, updatedDateTime string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*eblproto.Shipment, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	ConfirmationDatetime1, err := common.ConvertTimeToTimestamp(Layout, ConfirmationDatetime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedDateTime1, err := common.ConvertTimeToTimestamp(Layout, updatedDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	shipmentD := new(eblproto.ShipmentD)
	shipmentD.Id = id
	shipmentD.Uuid4 = uuid4
	shipmentD.IdS = idS
	shipmentD.BookingId = bookingId
	shipmentD.CarrierId = carrierId
	shipmentD.CarrierBookingReference = carrierBookingReference
	shipmentD.TermsAndConditions = termsAndConditions

	shipmentT := new(eblproto.ShipmentT)
	shipmentT.ConfirmationDatetime = ConfirmationDatetime1
	shipmentT.UpdatedDateTime = updatedDateTime1

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	shipment := eblproto.Shipment{ShipmentD: shipmentD, ShipmentT: shipmentT, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &shipment, nil
}
