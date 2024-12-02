package bkgservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	bkgstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/bkg/v2"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// insertRequestedEquipmentSQL - insert RequestedEquipmentSQL query
const insertRequestedEquipmentSQL = `insert into requested_equipments
	  (
uuid4,
booking_id,
shipment_id,
requested_equipment_sizetype,
requested_equipment_units,
confirmed_equipment_sizetype,
confirmed_equipment_units,
is_shipper_owned,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (
:uuid4,
:booking_id,
:shipment_id,
:requested_equipment_sizetype,
:requested_equipment_units,
:confirmed_equipment_sizetype,
:confirmed_equipment_units,
:is_shipper_owned,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectRequestedEquipmentsSQL - select RequestedEquipmentsSQL query
/*const selectRequestedEquipmentsSQL = `select
  id,
  uuid4,
  booking_id,
  shipment_id,
  requested_equipment_sizetype,
  requested_equipment_units,
  confirmed_equipment_sizetype,
  confirmed_equipment_units,
  is_shipper_owned,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from requested_equipments`*/

// CreateRequestedEquipment - CreateRequestedEquipment
func (bs *BkgService) CreateRequestedEquipment(ctx context.Context, in *bkgproto.CreateRequestedEquipmentRequest) (*bkgproto.CreateRequestedEquipmentResponse, error) {
	requestedEquipment, err := bs.ProcessRequestedEquipmentRequest(ctx, in)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = bs.insertRequestedEquipment(ctx, insertRequestedEquipmentSQL, requestedEquipment, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	requestedEquipmentResponse := bkgproto.CreateRequestedEquipmentResponse{}
	requestedEquipmentResponse.RequestedEquipment = requestedEquipment

	return &requestedEquipmentResponse, nil
}

// ProcessRequestedEquipmentRequest - ProcessRequestedEquipmentRequest
func (bs *BkgService) ProcessRequestedEquipmentRequest(ctx context.Context, in *bkgproto.CreateRequestedEquipmentRequest) (*bkgproto.RequestedEquipment, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, bs.UserServiceClient)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)
	requestedEquipmentD := bkgproto.RequestedEquipmentD{}
	requestedEquipmentD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	requestedEquipmentD.BookingId = in.BookingId
	requestedEquipmentD.ShipmentId = in.ShipmentId
	requestedEquipmentD.RequestedEquipmentSizetype = in.RequestedEquipmentSizetype
	requestedEquipmentD.RequestedEquipmentUnits = in.RequestedEquipmentUnits
	requestedEquipmentD.ConfirmedEquipmentSizetype = in.ConfirmedEquipmentSizetype
	requestedEquipmentD.ConfirmedEquipmentUnits = in.ConfirmedEquipmentUnits
	requestedEquipmentD.IsShipperOwned = in.IsShipperOwned

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	requestedEquipment := bkgproto.RequestedEquipment{RequestedEquipmentD: &requestedEquipmentD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &requestedEquipment, nil
}

// insertRequestedEquipment - Insert RequestedEquipment
func (bs *BkgService) insertRequestedEquipment(ctx context.Context, insertRequestedEquipmentSQL string, requestedEquipment *bkgproto.RequestedEquipment, userEmail string, requestID string) error {
	requestedEquipmentTmp, err := bs.CrRequestedEquipmentStruct(ctx, requestedEquipment, userEmail, requestID)
	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = bs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertRequestedEquipmentSQL, requestedEquipmentTmp)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		requestedEquipment.RequestedEquipmentD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(requestedEquipment.RequestedEquipmentD.Uuid4)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		requestedEquipment.RequestedEquipmentD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	return nil
}

// CrRequestedEquipmentStruct - process RequestedEquipment details
func (bs *BkgService) CrRequestedEquipmentStruct(ctx context.Context, requestedEquipment *bkgproto.RequestedEquipment, userEmail string, requestID string) (*bkgstruct.RequestedEquipment, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(requestedEquipment.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(requestedEquipment.CrUpdTime.UpdatedAt)

	requestedEquipmentTmp := bkgstruct.RequestedEquipment{RequestedEquipmentD: requestedEquipment.RequestedEquipmentD, CrUpdUser: requestedEquipment.CrUpdUser, CrUpdTime: crUpdTime}

	return &requestedEquipmentTmp, nil
}
