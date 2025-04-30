package eblservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	bkgservice "github.com/cloudfresco/sc-dcsa/internal/services/bkgservices"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eblstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ebl/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// ShippingService - For accessing Shipping services
type ShippingService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eblproto.UnimplementedShippingServiceServer
}

// NewShippingService - Create Shipping service
func NewShippingService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *ShippingService {
	return &ShippingService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertShippingSQL - Insert ShippingSQL Query
const insertShippingSQL = `insert into shipments
	  ( 
  uuid4,
  booking_id,
  carrier_id,
  carrier_booking_reference,
  terms_and_conditions,
  confirmation_datetime,
  updated_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id, 
  created_at,
  updated_at)
  values (:uuid4,
:booking_id,
:carrier_id,
:carrier_booking_reference,
:terms_and_conditions,
:confirmation_datetime,
:updated_date_time,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectShippingsSQL - select Shippings Query
const selectShippingsSQL = `select 
  id,
  uuid4,
  booking_id,
  carrier_id,
  carrier_booking_reference,
  terms_and_conditions,
  confirmation_datetime,
  updated_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id, 
  created_at,
  updated_at from shipments`

// CreateShipment - Create  Shipping
func (ss *ShippingService) CreateShipment(ctx context.Context, in *eblproto.CreateShipmentRequest) (*eblproto.CreateShipmentResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ss.UserServiceClient)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	confirmationDateTime, err := time.Parse(common.Layout, in.ConfirmationDatetime)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	updatedDateTime, err := time.Parse(common.Layout, in.UpdatedDateTime)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipmentD := eblproto.ShipmentD{}
	shipmentD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	shipmentD.BookingId = in.BookingId
	shipmentD.CarrierId = in.CarrierId
	shipmentD.CarrierBookingReference = in.CarrierBookingReference
	shipmentD.TermsAndConditions = in.TermsAndConditions

	shipmentT := eblproto.ShipmentT{}
	shipmentT.ConfirmationDatetime = common.TimeToTimestamp(confirmationDateTime.UTC().Truncate(time.Second))
	shipmentT.UpdatedDateTime = common.TimeToTimestamp(updatedDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	shipment := eblproto.Shipment{ShipmentD: &shipmentD, ShipmentT: &shipmentT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ss.insertShipping(ctx, insertShippingSQL, &shipment, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipResponse := eblproto.CreateShipmentResponse{}
	shipResponse.Shipment = &shipment
	return &shipResponse, nil
}

// insertShipping - Insert Shipping
func (ss *ShippingService) insertShipping(ctx context.Context, insertShippingSQL string, shipment *eblproto.Shipment, userEmail string, requestID string) error {
	shipmentTmp, err := ss.crShipmentStruct(ctx, shipment, userEmail, requestID)
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = ss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertShippingSQL, shipmentTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		shipment.ShipmentD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(shipment.ShipmentD.Uuid4)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		shipment.ShipmentD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crShipmentStruct - process ShippingC details
func (ss *ShippingService) crShipmentStruct(ctx context.Context, shipment *eblproto.Shipment, userEmail string, requestID string) (*eblstruct.Shipment, error) {
	shipmentT := new(eblstruct.ShipmentT)
	shipmentT.ConfirmationDatetime = common.TimestampToTime(shipment.ShipmentT.ConfirmationDatetime)
	shipmentT.UpdatedDateTime = common.TimestampToTime(shipment.ShipmentT.UpdatedDateTime)

	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(shipment.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(shipment.CrUpdTime.UpdatedAt)

	shipmentTmp := eblstruct.Shipment{ShipmentD: shipment.ShipmentD, ShipmentT: shipmentT, CrUpdUser: shipment.CrUpdUser, CrUpdTime: crUpdTime}

	return &shipmentTmp, nil
}

// GetShipments - Get  Shippings
func (ss *ShippingService) GetShipments(ctx context.Context, in *eblproto.GetShipmentsRequest) (*eblproto.GetShipmentsResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = ss.DBService.LimitSQLRows
	}
	query := "status_code = ?"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	shipments := []*eblproto.Shipment{}

	nselectShippingsSQL := selectShippingsSQL + ` where ` + query

	rows, err := ss.DBService.DB.QueryxContext(ctx, nselectShippingsSQL, "active")
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		shipmentTmp := eblstruct.Shipment{}
		err = rows.StructScan(&shipmentTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		shipment, err := ss.getShipmentStruct(ctx, &getRequest, shipmentTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		shipments = append(shipments, shipment)

	}

	shipmentsResponse := eblproto.GetShipmentsResponse{}
	if len(shipments) != 0 {
		next := shipments[len(shipments)-1].ShipmentD.Id
		next--
		nextc := common.EncodeCursor(next)
		shipmentsResponse = eblproto.GetShipmentsResponse{Shipments: shipments, NextCursor: nextc}
	} else {
		shipmentsResponse = eblproto.GetShipmentsResponse{Shipments: shipments, NextCursor: "0"}
	}
	return &shipmentsResponse, nil
}

// GetShipment - Get Shipping
func (ss *ShippingService) GetShipment(ctx context.Context, inReq *eblproto.GetShipmentRequest) (*eblproto.GetShipmentResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectShippingsSQL := selectShippingsSQL + ` where uuid4 = ? and status_code = ?;`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectShippingsSQL, uuid4byte, "active")
	shipmentTmp := eblstruct.Shipment{}
	err = row.StructScan(&shipmentTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipment, err := ss.getShipmentStruct(ctx, in, shipmentTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipResponse := eblproto.GetShipmentResponse{}
	shipResponse.Shipment = shipment

	return &shipResponse, nil
}

// GetShipmentByPk - Get Shipping By Primary key(Id)
func (ss *ShippingService) GetShipmentByPk(ctx context.Context, inReq *eblproto.GetShipmentByPkRequest) (*eblproto.GetShipmentByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectShippingsSQL := selectShippingsSQL + ` where id = ? and status_code = ?;`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectShippingsSQL, in.Id, "active")
	shipmentTmp := eblstruct.Shipment{}
	err := row.StructScan(&shipmentTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	shipment, err := ss.getShipmentStruct(ctx, &getRequest, shipmentTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	shipResponse := eblproto.GetShipmentByPkResponse{}
	shipResponse.Shipment = shipment

	return &shipResponse, nil
}

// getShipmentStruct - Get shipment header
func (ss *ShippingService) getShipmentStruct(ctx context.Context, in *commonproto.GetRequest, shipmentTmp eblstruct.Shipment) (*eblproto.Shipment, error) {
	shipmentT := new(eblproto.ShipmentT)
	shipmentT.ConfirmationDatetime = common.TimeToTimestamp(shipmentTmp.ShipmentT.ConfirmationDatetime)
	shipmentT.UpdatedDateTime = common.TimeToTimestamp(shipmentTmp.ShipmentT.UpdatedDateTime)

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(shipmentTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(shipmentTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(shipmentTmp.ShipmentD.Uuid4)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipmentTmp.ShipmentD.IdS = uuid4Str

	shipment := eblproto.Shipment{ShipmentD: shipmentTmp.ShipmentD, ShipmentT: shipmentT, CrUpdUser: shipmentTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &shipment, nil
}

// FindCarrierBookingReferenceByShippingInstructionId - FindCarrierBookingReferenceByShippingInstructionId
func (ss *ShippingService) FindCarrierBookingReferenceByShippingInstructionId(ctx context.Context, inReq *eblproto.FindCarrierBookingReferenceByShippingInstructionIdRequest) (*eblproto.FindCarrierBookingReferenceByShippingInstructionIdResponse, error) {
	in := inReq.GetByIdRequest
	referenceServ := &bkgservice.ReferenceService{DBService: ss.DBService, RedisService: ss.RedisService, UserServiceClient: ss.UserServiceClient}
	findByShippingInstructionIdRequest := bkgproto.FindByShippingInstructionIdRequest{}
	findByShippingInstructionIdRequest.ShippingInstructionId = in.Id
	findByShippingInstructionIdRequest.UserEmail = in.UserEmail
	findByShippingInstructionIdRequest.RequestId = in.RequestId

	refResponse, err := referenceServ.FindByShippingInstructionId(ctx, &findByShippingInstructionIdRequest)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	reference := refResponse.Reference1

	getByIdRequest := commonproto.GetByIdRequest{}
	getByIdRequest.Id = reference.Reference1D.ShipmentId
	getByIdRequest.UserEmail = in.UserEmail
	getByIdRequest.RequestId = in.RequestId

	getShipByPkRequest := eblproto.GetShipmentByPkRequest{}
	getShipByPkRequest.GetByIdRequest = &getByIdRequest
	shipmentResponse, err := ss.GetShipmentByPk(ctx, &getShipByPkRequest)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	shipment := shipmentResponse.Shipment
	carBkgReference := eblproto.FindCarrierBookingReferenceByShippingInstructionIdResponse{}
	carBkgReference.CarrierBookingReference = shipment.ShipmentD.CarrierBookingReference
	return &carBkgReference, nil
}
