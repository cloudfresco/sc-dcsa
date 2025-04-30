package bkgservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	bkgstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/bkg/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// insertShipmentLocationSQL - insert ShipmentLocationSQL query
const insertShipmentLocationSQL = `insert into shipment_locations
	  (
shipment_id,
booking_id,
location_id,
shipment_location_type_code,
displayed_name,
event_date_time)
  values (
:shipment_id,
:booking_id,
:location_id,
:shipment_location_type_code,
:displayed_name,
:event_date_time);`

// selectShipmentLocationsSQL - select ShipmentLocationsSQL query
const selectShipmentLocationsSQL = `select 
  shipment_id,
  booking_id,
  location_id,
  shipment_location_type_code,
  displayed_name,
  event_date_time from shipment_locations`

// CreateShipmentLocation - CreateShipmentLocation
func (bs *BkgService) CreateShipmentLocation(ctx context.Context, in *bkgproto.CreateShipmentLocationRequest) (*bkgproto.CreateShipmentLocationResponse, error) {
	shipmentLocation, err := bs.ProcessShipmentLocationRequest(ctx, in)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipmentLocation, err = bs.insertShipmentLocation(ctx, insertShipmentLocationSQL, shipmentLocation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	shipmentLocationResponse := bkgproto.CreateShipmentLocationResponse{}
	shipmentLocationResponse.ShipmentLocation = shipmentLocation
	return &shipmentLocationResponse, nil
}

// CreateShipmentLocationsByBookingIDAndTOs - CreateShipmentLocationsByBookingIDAndTOs
func (bs *BkgService) CreateShipmentLocationsByBookingIDAndTOs(ctx context.Context, req *bkgproto.CreateShipmentLocationsByBookingIDAndTOsRequest) (*bkgproto.CreateShipmentLocationsByBookingIDAndTOsResponse, error) {
	in := req.CreateShipmentLocationRequest
	shipmentLocation, err := bs.ProcessShipmentLocationRequest(ctx, in)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipmentLocation, err = bs.insertShipmentLocation(ctx, insertShipmentLocationSQL, shipmentLocation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	shipmentLocationResponse := bkgproto.CreateShipmentLocationsByBookingIDAndTOsResponse{}
	shipmentLocationResponse.ShipmentLocation = shipmentLocation
	return &shipmentLocationResponse, nil
}

// ProcessShipmentLocationRequest - ProcessShipmentLocationRequest
func (bs *BkgService) ProcessShipmentLocationRequest(ctx context.Context, in *bkgproto.CreateShipmentLocationRequest) (*bkgproto.ShipmentLocation, error) {
	eventDateTime, err := time.Parse(common.Layout, in.EventDateTime)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipmentLocationD := bkgproto.ShipmentLocationD{}

	shipmentLocationD.ShipmentId = in.ShipmentId
	shipmentLocationD.BookingId = in.BookingId
	shipmentLocationD.LocationId = in.LocationId
	shipmentLocationD.ShipmentLocationTypeCode = in.ShipmentLocationTypeCode
	shipmentLocationD.DisplayedName = in.DisplayedName

	shipmentLocationT := bkgproto.ShipmentLocationT{}
	shipmentLocationT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	shipmentLocation := bkgproto.ShipmentLocation{ShipmentLocationD: &shipmentLocationD, ShipmentLocationT: &shipmentLocationT}

	return &shipmentLocation, nil
}

// insertShipmentLocation - Insert ShipmentLocation
func (bs *BkgService) insertShipmentLocation(ctx context.Context, insertShipmentLocationSQL string, shipmentLocation *bkgproto.ShipmentLocation, userEmail string, requestID string) (*bkgproto.ShipmentLocation, error) {
	shipmentLocationTmp, err := bs.CrShipmentLocationStruct(ctx, shipmentLocation, userEmail, requestID)
	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return nil, err
	}

	err = bs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertShipmentLocationSQL, shipmentLocationTmp)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return nil, err
	}
	return shipmentLocation, nil
}

// CrShipmentLocationStruct - process ShipmentLocation details
func (bs *BkgService) CrShipmentLocationStruct(ctx context.Context, shipmentLocation *bkgproto.ShipmentLocation, userEmail string, requestID string) (*bkgstruct.ShipmentLocation, error) {
	shipmentLocationT := new(bkgstruct.ShipmentLocationT)
	shipmentLocationT.EventDateTime = common.TimestampToTime(shipmentLocation.ShipmentLocationT.EventDateTime)

	shipmentLocationTmp := bkgstruct.ShipmentLocation{ShipmentLocationD: shipmentLocation.ShipmentLocationD, ShipmentLocationT: shipmentLocationT}
	return &shipmentLocationTmp, nil
}

// FetchShipmentLocationsByBookingID - Get FetchShipmentLocationsByBookingID
func (bs *BkgService) FetchShipmentLocationsByBookingID(ctx context.Context, in *bkgproto.FetchShipmentLocationsByBookingIDRequest) (*bkgproto.FetchShipmentLocationsByBookingIDResponse, error) {
	nselectShipmentLocationsSQL := selectShipmentLocationsSQL + ` where booking_id = ?;`
	row := bs.DBService.DB.QueryRowxContext(ctx, nselectShipmentLocationsSQL, in.BookingId)
	shipmentLocationTmp := bkgstruct.ShipmentLocation{}
	err := row.StructScan(&shipmentLocationTmp)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	shipmentLocation, err := bs.getShipmentLocationStruct(ctx, &getRequest, shipmentLocationTmp)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	shipmentLocationResponse := bkgproto.FetchShipmentLocationsByBookingIDResponse{}
	shipmentLocationResponse.ShipmentLocation = shipmentLocation
	return &shipmentLocationResponse, nil
}

// getShipmentLocationStruct - Get shipmentLocation
func (bs *BkgService) getShipmentLocationStruct(ctx context.Context, in *commonproto.GetRequest, shipmentLocationTmp bkgstruct.ShipmentLocation) (*bkgproto.ShipmentLocation, error) {
	shipmentLocationT := new(bkgproto.ShipmentLocationT)
	shipmentLocationT.EventDateTime = common.TimeToTimestamp(shipmentLocationTmp.ShipmentLocationT.EventDateTime)

	shipmentLocation := bkgproto.ShipmentLocation{ShipmentLocationD: shipmentLocationTmp.ShipmentLocationD, ShipmentLocationT: shipmentLocationT}

	return &shipmentLocation, nil
}
