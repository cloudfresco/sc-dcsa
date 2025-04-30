package eventcoreservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eventcorestruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/eventcore/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// InsertUtilizedTransportEquipmentSQL - insert UtilizedTransportEquipmentSQL query
const InsertUtilizedTransportEquipmentSQL = `insert into utilized_transport_equipments
	  (
uuid4,
equipment_reference,
cargo_gross_weight,
cargo_gross_weight_unit,
is_shipper_owned,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (
:uuid4,
:equipment_reference,
:cargo_gross_weight,
:cargo_gross_weight_unit,
:is_shipper_owned,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectUtilizedTransportEquipmentsSQL - select UtilizedTransportEquipmentsSQL query
/*const selectUtilizedTransportEquipmentsSQL = `select
  id,
  uuid4,
  equipment_reference,
  cargo_gross_weight,
  cargo_gross_weight_unit,
  is_shipper_owned,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from utilized_transport_equipments`*/

// InsertEquipmentSQL - insert EquipmentSQL query
const InsertEquipmentSQL = `insert into equipment
	  (
uuid4,
equipment_reference,
iso_equipment_code,
tare_weight,
weight_unit,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (
:uuid4,
:equipment_reference,
:iso_equipment_code,
:tare_weight,
:weight_unit,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectEquipmentsSQL - select EquipmentsSQL query
/*const selectEquipmentsSQL = `select
  id,
  uuid4,
  equipment_reference,
  iso_equipment_code,
  tare_weight,
  weight_unit,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from equipment`*/

// UtilizedTransportEquipmentService - For accessing UtilizedTransportEquipment services
type UtilizedTransportEquipmentService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eventcoreproto.UnimplementedUtilizedTransportEquipmentServiceServer
}

// NewUtilizedTransportEquipmentService - Create UtilizedTransportEquipment service
func NewUtilizedTransportEquipmentService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *UtilizedTransportEquipmentService {
	return &UtilizedTransportEquipmentService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// CreateUtilizedTransportEquipment - CreateUtilizedTransportEquipment
func (uts *UtilizedTransportEquipmentService) CreateUtilizedTransportEquipment(ctx context.Context, in *eventcoreproto.CreateUtilizedTransportEquipmentRequest) (*eventcoreproto.CreateUtilizedTransportEquipmentResponse, error) {
	utilizedTransportEquipment, err := uts.ProcessUtilizedTransportEquipmentRequest(ctx, in)
	if err != nil {
		uts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = uts.insertUtilizedTransportEquipment(ctx, InsertUtilizedTransportEquipmentSQL, utilizedTransportEquipment, InsertEquipmentSQL, utilizedTransportEquipment.Equipment, in.GetUserEmail(), in.GetRequestId())

	if err != nil {
		uts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	utilTransEquipResponse := eventcoreproto.CreateUtilizedTransportEquipmentResponse{}
	utilTransEquipResponse.UtilizedTransportEquipment = utilizedTransportEquipment
	return &utilTransEquipResponse, nil
}

// ProcessUtilizedTransportEquipmentRequest - ProcessUtilizedTransportEquipmentRequest
func (uts *UtilizedTransportEquipmentService) ProcessUtilizedTransportEquipmentRequest(ctx context.Context, in *eventcoreproto.CreateUtilizedTransportEquipmentRequest) (*eventcoreproto.UtilizedTransportEquipment, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, uts.UserServiceClient)
	if err != nil {
		uts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	utilizedTransportEquipmentD := eventcoreproto.UtilizedTransportEquipmentD{}
	utilizedTransportEquipmentD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		uts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	utilizedTransportEquipmentD.EquipmentReference = in.EquipmentReference
	utilizedTransportEquipmentD.CargoGrossWeight = in.CargoGrossWeight
	utilizedTransportEquipmentD.CargoGrossWeightUnit = in.CargoGrossWeightUnit
	utilizedTransportEquipmentD.IsShipperOwned = in.IsShipperOwned

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	equipmentD := eventcoreproto.EquipmentD{}
	equipmentD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		uts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	equipmentD.EquipmentReference = in.EquipmentReference
	equipmentD.IsoEquipmentCode = in.Equipment.IsoEquipmentCode
	equipmentD.TareWeight = in.Equipment.TareWeight
	equipmentD.WeightUnit = in.Equipment.WeightUnit

	equipment := eventcoreproto.Equipment{EquipmentD: &equipmentD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	utilizedTransportEquipment := eventcoreproto.UtilizedTransportEquipment{UtilizedTransportEquipmentD: &utilizedTransportEquipmentD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime, Equipment: &equipment}

	return &utilizedTransportEquipment, nil
}

// insertUtilizedTransportEquipment - Insert UtilizedTransportEquipment
func (uts *UtilizedTransportEquipmentService) insertUtilizedTransportEquipment(ctx context.Context, insertUtilizedTransportEquipmentSQL string, utilizedTransportEquipment *eventcoreproto.UtilizedTransportEquipment, insertEquipmentSQL string, equipment *eventcoreproto.Equipment, userEmail string, requestID string) error {
	utilizedTransportEquipmentTmp, err := uts.crUtilizedTransportEquipmentStruct(ctx, utilizedTransportEquipment, userEmail, requestID)
	if err != nil {
		uts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	equipmentTmp, err := uts.crEquipmentStruct(ctx, equipment, userEmail, requestID)
	if err != nil {
		uts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = uts.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertUtilizedTransportEquipmentSQL, utilizedTransportEquipmentTmp)
		if err != nil {
			uts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			uts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		_, err = tx.NamedExecContext(ctx, insertEquipmentSQL, equipmentTmp)
		if err != nil {
			uts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		utilizedTransportEquipment.UtilizedTransportEquipmentD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(utilizedTransportEquipment.UtilizedTransportEquipmentD.Uuid4)
		if err != nil {
			uts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		utilizedTransportEquipment.UtilizedTransportEquipmentD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		uts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crUtilizedTransportEquipmentStruct - process UtilizedTransportEquipment details
func (uts *UtilizedTransportEquipmentService) crUtilizedTransportEquipmentStruct(ctx context.Context, utilizedTransportEquipment *eventcoreproto.UtilizedTransportEquipment, userEmail string, requestID string) (*eventcorestruct.UtilizedTransportEquipment, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(utilizedTransportEquipment.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(utilizedTransportEquipment.CrUpdTime.UpdatedAt)

	utilizedTransportEquipmentTmp := eventcorestruct.UtilizedTransportEquipment{UtilizedTransportEquipmentD: utilizedTransportEquipment.UtilizedTransportEquipmentD, CrUpdUser: utilizedTransportEquipment.CrUpdUser, CrUpdTime: crUpdTime}
	return &utilizedTransportEquipmentTmp, nil
}

// crEquipmentStruct - process Equipment details
func (uts *UtilizedTransportEquipmentService) crEquipmentStruct(ctx context.Context, equipment *eventcoreproto.Equipment, userEmail string, requestID string) (*eventcorestruct.Equipment, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(equipment.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(equipment.CrUpdTime.UpdatedAt)

	equipmentTmp := eventcorestruct.Equipment{EquipmentD: equipment.EquipmentD, CrUpdUser: equipment.CrUpdUser, CrUpdTime: crUpdTime}
	return &equipmentTmp, nil
}
