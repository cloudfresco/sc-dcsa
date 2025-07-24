package bkgservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestBkgService_FetchShipmentLocationsByBookingID(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	bookingService := NewBkgService(log, dbService, redisService, userServiceClient, currencyService)
	shipmentLocation, err := GetShipmentLocation(uint32(15), uint32(10), uint32(1), "POL", "Hamburg", "2020-03-07T00:00:00Z")
	if err != nil {
		t.Error(err)
		return
	}
	form := bkgproto.FetchShipmentLocationsByBookingIDRequest{}
	form.BookingId = uint32(10)
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	shipmentLocationResponse := bkgproto.FetchShipmentLocationsByBookingIDResponse{}
	shipmentLocationResponse.ShipmentLocation = shipmentLocation

	type args struct {
		ctx context.Context
		in  *bkgproto.FetchShipmentLocationsByBookingIDRequest
	}
	tests := []struct {
		bs      *BkgService
		args    args
		want    *bkgproto.FetchShipmentLocationsByBookingIDResponse
		wantErr bool
	}{
		{
			bs: bookingService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &shipmentLocationResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		shipmentLocationResp, err := tt.bs.FetchShipmentLocationsByBookingID(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("BkgService.FetchShipmentLocationsByBookingID() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(shipmentLocationResp, tt.want) {
			t.Errorf("BkgService.FetchShipmentLocationsByBookingID() = %v, want %v", shipmentLocationResp, tt.want)
		}
		shipmentLocationResult := shipmentLocationResp.ShipmentLocation
		assert.NotNil(t, shipmentLocationResult)
		assert.Equal(t, shipmentLocationResult.ShipmentLocationD.ShipmentLocationTypeCode, "POL", "they should be equal")
		assert.Equal(t, shipmentLocationResult.ShipmentLocationD.DisplayedName, "Hamburg", "they should be equal")
	}
}

func TestBkgService_CreateShipmentLocation(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}
	ctx := LoginUser()

	bkgService := NewBkgService(log, dbService, redisService, userServiceClient, currencyService)
	shipmentLocation := bkgproto.CreateShipmentLocationRequest{}
	shipmentLocation.ShipmentId = uint32(7)
	shipmentLocation.BookingId = uint32(1)
	shipmentLocation.LocationId = uint32(7)
	shipmentLocation.ShipmentLocationTypeCode = "PDE"
	shipmentLocation.DisplayedName = "PDE"
	shipmentLocation.EventDateTime = "07/03/2020"

	type args struct {
		ctx context.Context
		in  *bkgproto.CreateShipmentLocationRequest
	}
	tests := []struct {
		bs      *BkgService
		args    args
		wantErr bool
	}{
		{
			bs: bkgService,
			args: args{
				ctx: ctx,
				in:  &shipmentLocation,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		shipmentLocationResp, err := tt.bs.CreateShipmentLocation(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("BkgService.CreateShipmentLocation() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		shipmentLocationResult := shipmentLocationResp.ShipmentLocation
		assert.NotNil(t, shipmentLocationResult)
		assert.Equal(t, shipmentLocationResult.ShipmentLocationD.ShipmentLocationTypeCode, "PDE", "they should be equal")
	}
}

func GetShipmentLocation(shipmentId uint32, bookingId uint32, locationId uint32, shipmentLocationTypeCode string, displayedName string, eventDateTime string) (*bkgproto.ShipmentLocation, error) {
	eventDateTime1, err := common.ConvertTimeToTimestamp(Layout, eventDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	shipmentLocationD := new(bkgproto.ShipmentLocationD)
	shipmentLocationD.ShipmentId = shipmentId
	shipmentLocationD.BookingId = bookingId
	shipmentLocationD.LocationId = locationId
	shipmentLocationD.ShipmentLocationTypeCode = shipmentLocationTypeCode
	shipmentLocationD.DisplayedName = displayedName

	shipmentLocationT := new(bkgproto.ShipmentLocationT)
	shipmentLocationT.EventDateTime = eventDateTime1

	shipmentLocation := bkgproto.ShipmentLocation{ShipmentLocationD: shipmentLocationD, ShipmentLocationT: shipmentLocationT}
	return &shipmentLocation, nil
}
