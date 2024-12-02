// https://github.com/dcsaorg/DCSA-Edocumentation/blob/master/edocumentation-service/src/test/java/org/dcsa/edocumentation/service/ReferenceServiceTest.java
package bkgservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestReferenceService_FindByBookingId(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	referenceService := NewReferenceService(log, dbService, redisService, userServiceClient)
	reference, err := GetReference(uint32(5), "FF", "test", uint32(15), uint32(1), uint32(10), uint32(0), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	refer := bkgproto.FindByBookingIdResponse{}
	refer.Reference1 = reference

	form := bkgproto.FindByBookingIdRequest{}
	form.BookingId = uint32(10)
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *bkgproto.FindByBookingIdRequest
	}
	tests := []struct {
		rs      *ReferenceService
		args    args
		want    *bkgproto.FindByBookingIdResponse
		wantErr bool
	}{
		{
			rs: referenceService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &refer,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		refResponse, err := tt.rs.FindByBookingId(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReferenceService.FindByBookingId() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(refResponse, tt.want) {
			t.Errorf("ReferenceService.FindByBookingId() = %v, want %v", refResponse, tt.want)
		}
		refResult := refResponse.Reference1
		assert.NotNil(t, refResult)
		assert.Equal(t, refResult.Reference1D.ReferenceValue, "test", "they should be equal")
		assert.NotNil(t, refResult.Reference1D.ReferenceTypeCode)
	}
}

func TestReferenceService_FindByShippingInstructionId(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	referenceService := NewReferenceService(log, dbService, redisService, userServiceClient)
	reference, err := GetReference(uint32(5), "FF", "test", uint32(15), uint32(1), uint32(10), uint32(0), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}
	refer := bkgproto.FindByShippingInstructionIdResponse{}
	refer.Reference1 = reference

	form := bkgproto.FindByShippingInstructionIdRequest{}
	form.ShippingInstructionId = uint32(1)
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *bkgproto.FindByShippingInstructionIdRequest
	}
	tests := []struct {
		rs      *ReferenceService
		args    args
		want    *bkgproto.FindByShippingInstructionIdResponse
		wantErr bool
	}{
		{
			rs: referenceService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &refer,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		refResponse, err := tt.rs.FindByShippingInstructionId(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReferenceService.FindByShippingInstructionId() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(refResponse, tt.want) {
			t.Errorf("ReferenceService.FindByShippingInstructionId() = %v, want %v", refResponse, tt.want)
		}
		refResult := refResponse.Reference1
		assert.NotNil(t, refResult)
		assert.Equal(t, refResult.Reference1D.ReferenceValue, "test", "they should be equal")
		assert.NotNil(t, refResult.Reference1D.ReferenceTypeCode)
	}
}

func GetReference(id uint32, referenceTypeCode string, referenceValue string, shipmentId uint32, shippingInstructionId uint32, bookingId uint32, consignmentItemId uint32, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*bkgproto.Reference1, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	referenceD := new(bkgproto.Reference1D)
	referenceD.Id = id
	referenceD.ReferenceTypeCode = referenceTypeCode
	referenceD.ReferenceValue = referenceValue
	referenceD.ShipmentId = shipmentId
	referenceD.ShippingInstructionId = shippingInstructionId
	referenceD.BookingId = bookingId
	referenceD.ConsignmentItemId = consignmentItemId

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	reference := bkgproto.Reference1{Reference1D: referenceD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}
	return &reference, nil
}

func TestReferenceService_CreateReferencesByBookingIdAndTOs(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	referenceService := NewReferenceService(log, dbService, redisService, userServiceClient)

	refer := bkgproto.CreateReferencesByBookingIdAndTOsRequest{}
	reference := bkgproto.CreateReferenceRequest{}
	reference.ReferenceTypeCode = "AAO"
	reference.ReferenceValue = "ref value"
	reference.BookingId = uint32(10)
	reference.UserId = "auth0|66fd06d0bfea78a82bb42459"
	reference.UserEmail = "sprov300@gmail.com"
	reference.RequestId = "bks1m1g91jau4nkks2f0"
	refer.CreateReferenceRequest = &reference
	type args struct {
		ctx context.Context
		in  *bkgproto.CreateReferencesByBookingIdAndTOsRequest
	}
	tests := []struct {
		rs      *ReferenceService
		args    args
		wantErr bool
	}{
		{
			rs: referenceService,
			args: args{
				ctx: ctx,
				in:  &refer,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		refResponse, err := tt.rs.CreateReferencesByBookingIdAndTOs(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReferenceService.CreateReferencesByBookingIdAndTOs() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		refResult := refResponse.Reference1
		assert.NotNil(t, refResult)
		assert.Equal(t, refResult.Reference1D.ReferenceTypeCode, "AAO", "they should be equal")
		assert.NotNil(t, refResult.Reference1D.ReferenceValue)
	}
}

func TestReferenceService_CreateReferencesByShippingInstructionIdAndTOs(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	referenceService := NewReferenceService(log, dbService, redisService, userServiceClient)
	refer := bkgproto.CreateReferencesByShippingInstructionIdAndTOsRequest{}
	reference := bkgproto.CreateReferenceRequest{}
	reference.ReferenceTypeCode = "AAO"
	reference.ReferenceValue = "ref value2"
	reference.ShippingInstructionId = uint32(1)
	reference.UserId = "auth0|66fd06d0bfea78a82bb42459"
	reference.UserEmail = "sprov300@gmail.com"
	reference.RequestId = "bks1m1g91jau4nkks2f0"
	refer.CreateReferenceRequest = &reference
	type args struct {
		ctx context.Context
		in  *bkgproto.CreateReferencesByShippingInstructionIdAndTOsRequest
	}
	tests := []struct {
		rs      *ReferenceService
		args    args
		wantErr bool
	}{
		{
			rs: referenceService,
			args: args{
				ctx: ctx,
				in:  &refer,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		refResponse, err := tt.rs.CreateReferencesByShippingInstructionIdAndTOs(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReferenceService.CreateReferencesByShippingInstructionIdAndTOs() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		refResult := refResponse.Reference1
		assert.NotNil(t, refResult)
		assert.Equal(t, refResult.Reference1D.ReferenceValue, "ref value2", "they should be equal")
		assert.NotNil(t, refResult.Reference1D.ReferenceTypeCode)
	}
}

func TestReferenceService_CreateReferencesWithEmptyBooking(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	referenceService := NewReferenceService(log, dbService, redisService, userServiceClient)

	refer := bkgproto.CreateReferencesByBookingIdAndTOsRequest{}

	type args struct {
		ctx context.Context
		in  *bkgproto.CreateReferencesByBookingIdAndTOsRequest
	}
	tests := []struct {
		rs      *ReferenceService
		args    args
		wantErr bool
	}{
		{
			rs: referenceService,
			args: args{
				ctx: ctx,
				in:  &refer,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		refResponse, err := tt.rs.CreateReferencesByBookingIdAndTOs(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReferenceService.CreateReferencesByBookingIdAndTOs() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.Nil(t, refResponse)
	}
}

func TestReferenceService_CreateReferenceWithNullBooking(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	referenceService := NewReferenceService(log, dbService, redisService, userServiceClient)

	refer := bkgproto.CreateReferencesByBookingIdAndTOsRequest{}
	reference := bkgproto.CreateReferenceRequest{}
	reference.ReferenceTypeCode = "AAO"
	reference.ReferenceValue = "ref value2"
	reference.UserId = "auth0|66fd06d0bfea78a82bb42459"
	reference.UserEmail = "sprov300@gmail.com"
	reference.RequestId = "bks1m1g91jau4nkks2f0"
	refer.CreateReferenceRequest = &reference

	type args struct {
		ctx context.Context
		in  *bkgproto.CreateReferencesByBookingIdAndTOsRequest
	}
	tests := []struct {
		rs      *ReferenceService
		args    args
		wantErr bool
	}{
		{
			rs: referenceService,
			args: args{
				ctx: ctx,
				in:  &refer,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		refResponse, err := tt.rs.CreateReferencesByBookingIdAndTOs(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReferenceService.CreateReferencesByBookingIdAndTOs() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.Nil(t, refResponse)
	}
}

func TestReferenceService_CreateReferenceWithNullShippingInstruction(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	referenceService := NewReferenceService(log, dbService, redisService, userServiceClient)

	refer := bkgproto.CreateReferencesByShippingInstructionIdAndTOsRequest{}
	reference := bkgproto.CreateReferenceRequest{}
	reference.ReferenceTypeCode = "AAO"
	reference.ReferenceValue = "ref value2"
	reference.UserId = "auth0|66fd06d0bfea78a82bb42459"
	reference.UserEmail = "sprov300@gmail.com"
	reference.RequestId = "bks1m1g91jau4nkks2f0"
	refer.CreateReferenceRequest = &reference
	type args struct {
		ctx context.Context
		in  *bkgproto.CreateReferencesByShippingInstructionIdAndTOsRequest
	}
	tests := []struct {
		rs      *ReferenceService
		args    args
		wantErr bool
	}{
		{
			rs: referenceService,
			args: args{
				ctx: ctx,
				in:  &refer,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		refResponse, err := tt.rs.CreateReferencesByShippingInstructionIdAndTOs(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReferenceService.CreateReferencesByShippingInstructionIdAndTOs() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.Nil(t, refResponse)
	}
}

func TestReferenceService_CreateReferenceWithEmptyShippingInstruction(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	referenceService := NewReferenceService(log, dbService, redisService, userServiceClient)

	refer := bkgproto.CreateReferencesByShippingInstructionIdAndTOsRequest{}

	type args struct {
		ctx context.Context
		in  *bkgproto.CreateReferencesByShippingInstructionIdAndTOsRequest
	}
	tests := []struct {
		rs      *ReferenceService
		args    args
		wantErr bool
	}{
		{
			rs: referenceService,
			args: args{
				ctx: ctx,
				in:  &refer,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		refResponse, err := tt.rs.CreateReferencesByShippingInstructionIdAndTOs(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReferenceService.CreateReferencesByShippingInstructionIdAndTOs() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.Nil(t, refResponse)
	}
}
