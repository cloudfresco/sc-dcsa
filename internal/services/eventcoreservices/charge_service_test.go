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

func TestChargeService_Create(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	chargeService := NewChargeService(log, dbService, redisService, userServiceClient, currencyService)

	charge := eventcoreproto.CreateChargeRequest{}
	charge.TransportDocumentId = uint32(4)
	charge.ShipmentId = uint32(11)
	charge.ChargeType = "TBD"
	charge.Amount = "12.12"
	charge.AmountCurrency = "EUR"
	charge.PaymentTermCode = "PRE"
	charge.CalculationBasis = "WHAT"
	charge.UnitPrice = "12.12"
	charge.UnitPriceCurrency = "EUR"
	charge.Quantity = float64(123.321)
	charge.UserId = "auth0|66fd06d0bfea78a82bb42459"
	charge.UserEmail = "sprov300@gmail.com"
	charge.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eventcoreproto.CreateChargeRequest
	}
	tests := []struct {
		cs      *ChargeService
		args    args
		wantErr bool
	}{
		{
			cs: chargeService,
			args: args{
				ctx: ctx,
				in:  &charge,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		chargeResp, err := tt.cs.Create(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ChargeService.Create() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		chargeResult := chargeResp.Charge
		assert.NotNil(t, chargeResult)
		assert.Equal(t, chargeResult.ChargeD.ChargeType, "TBD", "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.Amount, int64(1212), "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.AmountCurrency, "EUR", "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.PaymentTermCode, "PRE", "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.CalculationBasis, "WHAT", "they should be equal")
		assert.NotNil(t, chargeResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, chargeResult.CrUpdTime.UpdatedAt)

	}
}

func TestChargeService_FetchChargesByTransportDocumentId(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	chargeService := NewChargeService(log, dbService, redisService, userServiceClient, currencyService)

	charge, err := GetCharge(uint32(1), []byte{249, 211, 201, 174, 137, 193, 67, 148, 165, 252, 142, 115, 83, 138, 170, 196}, "f9d3c9ae-89c1-4394-a5fc-8e73538aaac4", uint32(6), uint32(19), "TBD", int64(1212), "EUR", "12.12", "PRE", "WHAT", int64(1212), "EUR", "12.12", float64(123.321), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	charges := []*eventcoreproto.Charge{}
	charges = append(charges, charge)

	form := eventcoreproto.FetchChargesByTransportDocumentIdRequest{}
	form.TransportDocumentId = uint32(6)
	form.UserId = "auth0|66fd06d0bfea78a82bb42459"
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	chargeResp := eventcoreproto.FetchChargesByTransportDocumentIdResponse{Charges: charges, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *eventcoreproto.FetchChargesByTransportDocumentIdRequest
	}
	tests := []struct {
		cs      *ChargeService
		args    args
		want    *eventcoreproto.FetchChargesByTransportDocumentIdResponse
		wantErr bool
	}{
		{
			cs: chargeService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &chargeResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		chargeResponse, err := tt.cs.FetchChargesByTransportDocumentId(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ChargeService.FetchChargesByTransportDocumentId() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(chargeResponse, tt.want) {
			t.Errorf("ChargeService.FetchChargesByTransportDocumentId() = %v, want %v", chargeResponse, tt.want)
		}
		chargeResult := chargeResponse.Charges[0]
		assert.NotNil(t, chargeResult)
		assert.Equal(t, chargeResult.ChargeD.ChargeType, "TBD", "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.Amount, int64(1212), "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.AmountCurrency, "EUR", "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.PaymentTermCode, "PRE", "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.CalculationBasis, "WHAT", "they should be equal")
		assert.NotNil(t, chargeResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, chargeResult.CrUpdTime.UpdatedAt)
	}
}

func TestChargeService_FetchChargesByShipmentId(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	chargeService := NewChargeService(log, dbService, redisService, userServiceClient, currencyService)
	charge, err := GetCharge(uint32(1), []byte{249, 211, 201, 174, 137, 193, 67, 148, 165, 252, 142, 115, 83, 138, 170, 196}, "f9d3c9ae-89c1-4394-a5fc-8e73538aaac4", uint32(6), uint32(19), "TBD", int64(1212), "EUR", "12.12", "PRE", "WHAT", int64(1212), "EUR", "12.12", float64(123.321), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	charges := []*eventcoreproto.Charge{}
	charges = append(charges, charge)

	form := eventcoreproto.FetchChargesByShipmentIdRequest{}
	form.ShipmentId = uint32(19)
	form.UserId = "auth0|66fd06d0bfea78a82bb42459"
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	chargeResp := eventcoreproto.FetchChargesByShipmentIdResponse{Charges: charges, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *eventcoreproto.FetchChargesByShipmentIdRequest
	}
	tests := []struct {
		cs      *ChargeService
		args    args
		want    *eventcoreproto.FetchChargesByShipmentIdResponse
		wantErr bool
	}{
		{
			cs: chargeService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &chargeResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		chargeResponse, err := tt.cs.FetchChargesByShipmentId(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ChargeService.FetchChargesByShipmentId() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(chargeResponse, tt.want) {
			t.Errorf("ChargeService.FetchChargesByShipmentId() = %v, want %v", chargeResponse, tt.want)
		}
		chargeResult := chargeResponse.Charges[0]
		assert.NotNil(t, chargeResult)
		assert.Equal(t, chargeResult.ChargeD.ChargeType, "TBD", "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.Amount, int64(1212), "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.AmountCurrency, "EUR", "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.PaymentTermCode, "PRE", "they should be equal")
		assert.Equal(t, chargeResult.ChargeD.CalculationBasis, "WHAT", "they should be equal")
		assert.NotNil(t, chargeResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, chargeResult.CrUpdTime.UpdatedAt)
	}
}

func GetCharge(id uint32, uuid4 []byte, idS string, transportDocumentId uint32, shipmentId uint32, chargeType string, amount int64, amountCurrency string, amountString string, paymentTermCode string, calculationBasis string, unitPrice int64, unitPriceCurrency string, unitPriceString string, quantity float64, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*eventcoreproto.Charge, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	chargeD := new(eventcoreproto.ChargeD)
	chargeD.Id = id
	chargeD.Uuid4 = uuid4
	chargeD.IdS = idS
	chargeD.TransportDocumentId = transportDocumentId
	chargeD.ShipmentId = shipmentId
	chargeD.ChargeType = chargeType
	chargeD.Amount = amount
	chargeD.AmountCurrency = amountCurrency
	chargeD.AmountString = amountString
	chargeD.PaymentTermCode = paymentTermCode
	chargeD.CalculationBasis = calculationBasis
	chargeD.UnitPrice = unitPrice
	chargeD.UnitPriceCurrency = unitPriceCurrency
	chargeD.UnitPriceString = unitPriceString
	chargeD.Quantity = quantity

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	charge := eventcoreproto.Charge{ChargeD: chargeD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &charge, nil
}
