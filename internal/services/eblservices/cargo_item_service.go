package eblservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eblstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ebl/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// InsertCargoItemSQL - Insert CargoItemSQL Query
const InsertCargoItemSQL = `insert into cargo_items
	  (
uuid4,
consignment_item_id,
weight,
volume,
weight_unit,
volume_unit,
number_of_packages,
package_code,
utilized_transport_equipment_id,
package_name_on_bl,
     status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (
:uuid4,
:consignment_item_id,
:weight,
:volume,
:weight_unit,
:volume_unit,
:number_of_packages,
:package_code,
:utilized_transport_equipment_id,
:package_name_on_bl,
     :status_code,
  :created_by_user_id,
  :updated_by_user_id,
  :created_at,
  :updated_at);`

// selectCargoItemSQL - select CargoItemSQL Query
/*const selectCargoItemsSQL = `select
  id,
  uuid4,
  consignment_item_id,
  weight,
  volume,
  weight_unit,
  volume_unit,
  number_of_packages,
  package_code,
  utilized_transport_equipment_id,
  package_name_on_bl from cargo_items`*/

// InsertCargoLineItemSQL - Insert CargoLineItemSQL Query
const InsertCargoLineItemSQL = `insert into cargo_line_items
	  (
uuid4,
cargo_item_id,
shipping_marks,
     status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (
:uuid4,
:cargo_item_id,
:shipping_marks,
     :status_code,
  :created_by_user_id,
  :updated_by_user_id,
  :created_at,
  :updated_at);`

// selectCargoLineItemSQL - select CargoLineItemSQL Query
/*const selectCargoLineItemsSQL = `select
  id,
  uuid4,
  cargo_item_id,
  shipping_marks from cargo_line_items`*/

// CreateCargoItem - CreateCargoItem
func (cis *ConsignmentItemService) CreateCargoItem(ctx context.Context, in *eblproto.CreateCargoItemRequest) (*eblproto.CargoItem, error) {
	cargoItem, err := cis.ProcessCargoItemRequest(ctx, in)
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = cis.insertCargoItem(ctx, InsertCargoItemSQL, cargoItem, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return cargoItem, nil
}

// ProcessCargoItemRequest - ProcessCargoItemRequest
func (cis *ConsignmentItemService) ProcessCargoItemRequest(ctx context.Context, in *eblproto.CreateCargoItemRequest) (*eblproto.CargoItem, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, cis.UserServiceClient)
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	cargoItemD := eblproto.CargoItemD{}
	cargoItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	cargoItemD.ConsignmentItemId = in.ConsignmentItemId
	cargoItemD.Weight = in.Weight
	cargoItemD.Volume = in.Volume
	cargoItemD.WeightUnit = in.WeightUnit
	cargoItemD.VolumeUnit = in.VolumeUnit
	cargoItemD.NumberOfPackages = in.NumberOfPackages
	cargoItemD.PackageCode = in.PackageCode
	cargoItemD.UtilizedTransportEquipmentId = in.UtilizedTransportEquipmentId
	cargoItemD.PackageNameOnBl = in.PackageNameOnBl

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	cargoItem := eblproto.CargoItem{CargoItemD: &cargoItemD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	cargoLineItems := []*eblproto.CargoLineItem{}
	for _, cargoLineItem := range in.CargoLineItems {
		cargoLineItem.UserId = in.UserId
		cargoLineItem.UserEmail = in.UserEmail
		cargoLineItem.RequestId = in.RequestId
		cargoLineItem, err := cis.ProcessCargoLineItemRequest(ctx, cargoLineItem)
		if err != nil {
			cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		cargoLineItems = append(cargoLineItems, cargoLineItem)
	}
	cargoItem.CargoLineItems = cargoLineItems

	return &cargoItem, nil
}

// insertCargoItem - Insert CargoItem
func (cis *ConsignmentItemService) insertCargoItem(ctx context.Context, insertCargoItemSQL string, cargoItem *eblproto.CargoItem, userEmail string, requestID string) error {
	cargoItemTmp, err := cis.crCargoItemStruct(ctx, cargoItem, userEmail, requestID)
	if err != nil {
		cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = cis.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertCargoItemSQL, cargoItemTmp)
		if err != nil {
			cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		cargoItem.CargoItemD.Id = uint32(uID)

		for _, cargoItemLineTmp := range cargoItemTmp.CargoLineItems {
			cargoItemLineTmp.CargoItemId = uint32(uID)
			_, err = tx.NamedExecContext(ctx, InsertCargoLineItemSQL, cargoItemLineTmp)
			if err != nil {
				cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		uuid4Str, err := common.UUIDBytesToStr(cargoItem.CargoItemD.Uuid4)
		if err != nil {
			cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		cargoItem.CargoItemD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crCargoItemStruct - process CargoItem details
func (cis *ConsignmentItemService) crCargoItemStruct(ctx context.Context, cargoItem *eblproto.CargoItem, userEmail string, requestID string) (*eblstruct.CargoItem, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(cargoItem.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(cargoItem.CrUpdTime.UpdatedAt)

	cargoItemTmp := eblstruct.CargoItem{CargoItemD: cargoItem.CargoItemD, CrUpdUser: cargoItem.CrUpdUser, CrUpdTime: crUpdTime}

	cargoLineItems := []*eblstruct.CargoLineItem{}
	for _, cargoLineItem := range cargoItem.CargoLineItems {
		crUpdTime := new(commonstruct.CrUpdTime)
		crUpdTime.CreatedAt = common.TimestampToTime(cargoLineItem.CrUpdTime.CreatedAt)
		crUpdTime.UpdatedAt = common.TimestampToTime(cargoLineItem.CrUpdTime.UpdatedAt)
		cargoLineItemTmp := eblstruct.CargoLineItem{CargoLineItemD: cargoLineItem.CargoLineItemD, CrUpdUser: cargoLineItem.CrUpdUser, CrUpdTime: crUpdTime}
		cargoLineItems = append(cargoLineItems, &cargoLineItemTmp)
	}

	cargoItemTmp.CargoLineItems = cargoLineItems

	return &cargoItemTmp, nil
}

// ProcessCargoLineItemRequest - ProcessCargoLineItemRequest
func (cis *ConsignmentItemService) ProcessCargoLineItemRequest(ctx context.Context, in *eblproto.CreateCargoLineItemRequest) (*eblproto.CargoLineItem, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, cis.UserServiceClient)
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	cargoLineItemD := eblproto.CargoLineItemD{}
	cargoLineItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	cargoLineItemD.CargoItemId = in.CargoItemId
	cargoLineItemD.ShippingMarks = in.ShippingMarks

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	cargoLineItem := eblproto.CargoLineItem{CargoLineItemD: &cargoLineItemD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &cargoLineItem, nil
}
