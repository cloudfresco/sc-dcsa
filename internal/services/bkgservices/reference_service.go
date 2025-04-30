// https://github.com/dcsaorg/DCSA-Event-Core/blob/master/src/main/java/org/dcsa/core/events/service/ReferenceService.java
// https://github.com/dcsaorg/DCSA-Event-Core/blob/master/src/main/java/org/dcsa/core/events/service/impl/ReferenceServiceImpl.java
package bkgservices

import (
	"context"
	"errors"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	bkgstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/bkg/v2"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// InsertReferenceSQL - Insert ReferenceSQL Query
const InsertReferenceSQL = `insert into reference1
	  (
reference_type_code,
reference_value,
shipment_id,
shipping_instruction_id,
booking_id,
consignment_item_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (
:reference_type_code,
:reference_value,
:shipment_id,
:shipping_instruction_id,
:booking_id,
:consignment_item_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectReferenceSQL - select ReferenceSQL Query
const selectReferencesSQL = `select 
  id,
  reference_type_code,
  reference_value,
  shipment_id,
  shipping_instruction_id,
  booking_id,
  consignment_item_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from reference1`

// ReferenceService - For accessing Reference services
type ReferenceService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	bkgproto.UnimplementedReferenceServiceServer
}

// NewReferenceService - Create Reference service
func NewReferenceService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *ReferenceService {
	return &ReferenceService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// CreateReference - CreateReference
func (rs *ReferenceService) CreateReference(ctx context.Context, in *bkgproto.CreateReferenceRequest) (*bkgproto.CreateReferenceResponse, error) {
	reference, err := rs.ProcessReferenceRequest(ctx, in)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = rs.insertReference(ctx, InsertReferenceSQL, reference, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	refResponse := bkgproto.CreateReferenceResponse{}
	refResponse.Reference1 = reference
	return &refResponse, nil
}

// CreateReferencesByBookingIdAndTOs - CreateReferencesByBookingIdAndTOs
func (rs *ReferenceService) CreateReferencesByBookingIdAndTOs(ctx context.Context, inReq *bkgproto.CreateReferencesByBookingIdAndTOsRequest) (*bkgproto.CreateReferencesByBookingIdAndTOsResponse, error) {
	in := inReq.CreateReferenceRequest
	if in == nil {
		err := errors.New("Form Data should not be empty")
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	if in.BookingId == 0 {
		err := errors.New("Booking should not be null")
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	reference, err := rs.ProcessReferenceRequest(ctx, in)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = rs.insertReference(ctx, InsertReferenceSQL, reference, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	refResponse := bkgproto.CreateReferencesByBookingIdAndTOsResponse{}
	refResponse.Reference1 = reference
	return &refResponse, nil
}

// CreateReferencesByShippingInstructionIdAndTOs - CreateReferencesByShippingInstructionIdAndTOs
func (rs *ReferenceService) CreateReferencesByShippingInstructionIdAndTOs(ctx context.Context, inReq *bkgproto.CreateReferencesByShippingInstructionIdAndTOsRequest) (*bkgproto.CreateReferencesByShippingInstructionIdAndTOsResponse, error) {
	in := inReq.CreateReferenceRequest
	if in == nil {
		err := errors.New("Form Data should not be empty")
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	if in.ShippingInstructionId == 0 {
		err := errors.New("ShippingInstruction should not be null")
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	reference, err := rs.ProcessReferenceRequest(ctx, in)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = rs.insertReference(ctx, InsertReferenceSQL, reference, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	refResponse := bkgproto.CreateReferencesByShippingInstructionIdAndTOsResponse{}
	refResponse.Reference1 = reference
	return &refResponse, nil
}

// CreateReferencesByShippingInstructionReferenceAndConsignmentIdAndTOs - CreateReferencesByShippingInstructionReferenceAndConsignmentIdAndTOs
func (rs *ReferenceService) CreateReferencesByShippingInstructionReferenceAndConsignmentIdAndTOs(ctx context.Context, inReq *bkgproto.CreateReferencesByShippingInstructionReferenceAndConsignmentIdAndTOsRequest) (*bkgproto.CreateReferencesByShippingInstructionReferenceAndConsignmentIdAndTOsResponse, error) {
	in := inReq.CreateReferenceRequest
	reference, err := rs.ProcessReferenceRequest(ctx, in)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = rs.insertReference(ctx, InsertReferenceSQL, reference, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	refResponse := bkgproto.CreateReferencesByShippingInstructionReferenceAndConsignmentIdAndTOsResponse{}
	refResponse.Reference1 = reference

	return &refResponse, nil
}

// ProcessReferenceRequest - ProcessReferenceRequest
func (rs *ReferenceService) ProcessReferenceRequest(ctx context.Context, in *bkgproto.CreateReferenceRequest) (*bkgproto.Reference1, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, rs.UserServiceClient)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	referenceD := bkgproto.Reference1D{}

	referenceD.ReferenceTypeCode = in.ReferenceTypeCode
	referenceD.ReferenceValue = in.ReferenceValue
	referenceD.ShipmentId = in.ShipmentId
	referenceD.ShippingInstructionId = in.ShippingInstructionId
	referenceD.BookingId = in.BookingId
	referenceD.ConsignmentItemId = in.ConsignmentItemId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	reference := bkgproto.Reference1{Reference1D: &referenceD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}
	return &reference, nil
}

// insertReference - Insert Reference
func (rs *ReferenceService) insertReference(ctx context.Context, insertReferenceSQL string, reference *bkgproto.Reference1, userEmail string, requestID string) error {
	referenceTmp, err := rs.crReferenceStruct(ctx, reference, userEmail, requestID)
	if err != nil {
		rs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = rs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertReferenceSQL, referenceTmp)
		if err != nil {
			rs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uID, err := res.LastInsertId()
		if err != nil {
			rs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		reference.Reference1D.Id = uint32(uID)

		return nil
	})
	if err != nil {
		rs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crReferenceStruct - process Reference details
func (rs *ReferenceService) crReferenceStruct(ctx context.Context, reference *bkgproto.Reference1, userEmail string, requestID string) (*bkgstruct.Reference1, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(reference.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(reference.CrUpdTime.UpdatedAt)
	referenceTmp := bkgstruct.Reference1{Reference1D: reference.Reference1D, CrUpdUser: reference.CrUpdUser, CrUpdTime: crUpdTime}
	return &referenceTmp, nil
}

// FindByBookingId - FindByBookingId
func (rs *ReferenceService) FindByBookingId(ctx context.Context, in *bkgproto.FindByBookingIdRequest) (*bkgproto.FindByBookingIdResponse, error) {
	nselectReferencesSQL := selectReferencesSQL + ` where booking_id = ?;`
	row := rs.DBService.DB.QueryRowxContext(ctx, nselectReferencesSQL, in.BookingId)
	referenceTmp := bkgstruct.Reference1{}
	err := row.StructScan(&referenceTmp)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	reference, err := rs.getReferenceStruct(ctx, &getRequest, referenceTmp)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	refResponse := bkgproto.FindByBookingIdResponse{}
	refResponse.Reference1 = reference
	return &refResponse, nil
}

// FindByShippingInstructionId - FindByShippingInstructionId
func (rs *ReferenceService) FindByShippingInstructionId(ctx context.Context, in *bkgproto.FindByShippingInstructionIdRequest) (*bkgproto.FindByShippingInstructionIdResponse, error) {
	nselectReferencesSQL := selectReferencesSQL + ` where shipping_instruction_id = ?;`
	row := rs.DBService.DB.QueryRowxContext(ctx, nselectReferencesSQL, in.ShippingInstructionId)
	referenceTmp := bkgstruct.Reference1{}
	err := row.StructScan(&referenceTmp)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	reference, err := rs.getReferenceStruct(ctx, &getRequest, referenceTmp)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	refResponse := bkgproto.FindByShippingInstructionIdResponse{}
	refResponse.Reference1 = reference
	return &refResponse, nil
}

// ResolveReferencesForBookingId - ResolveReferencesForBookingId
func (rs *ReferenceService) ResolveReferencesForBookingId(ctx context.Context, in *bkgproto.ResolveReferencesForBookingIdRequest) (*bkgproto.ResolveReferencesForBookingIdResponse, error) {
	return &bkgproto.ResolveReferencesForBookingIdResponse{}, nil
}

// ResolveReferencesForShippingInstructionReference - ResolveReferencesForShippingInstructionReference
func (rs *ReferenceService) ResolveReferencesForShippingInstructionReference(ctx context.Context, in *bkgproto.ResolveReferencesForShippingInstructionReferenceRequest) (*bkgproto.ResolveReferencesForShippingInstructionReferenceResponse, error) {
	return &bkgproto.ResolveReferencesForShippingInstructionReferenceResponse{}, nil
}

// FindShipmentIDsByShippingInstructionId - findShipmentIDsByShippingInstructionId
func (rs *ReferenceService) FindShipmentIDsByShippingInstructionId(ctx context.Context, in *commonproto.GetByIdRequest) (*bkgproto.ShipmentIds, error) {
	nselectReferencesSQL := selectReferencesSQL + ` where shipping_instruction_id = ?;`

	rows, err := rs.DBService.DB.QueryxContext(ctx, nselectReferencesSQL, in.Id)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	shipmentIds := []uint32{}
	for rows.Next() {

		referenceTmp := bkgstruct.Reference1{}
		err := rows.StructScan(&referenceTmp)
		if err != nil {
			rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		reference, err := rs.getReferenceStruct(ctx, &getRequest, referenceTmp)
		if err != nil {
			rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		shipmentIds = append(shipmentIds, reference.Reference1D.ShipmentId)

	}

	shipIds := bkgproto.ShipmentIds{}
	shipIds.ShipmentIds = shipmentIds
	return &shipIds, nil
}

// FindByShipmentId - FindByShipmentId
func (rs *ReferenceService) FindByShipmentId(ctx context.Context, in *bkgproto.FindByShipmentIdRequest) (*bkgproto.FindByShipmentIdResponse, error) {
	nselectReferencesSQL := selectReferencesSQL + `si WHERE shipment_id = ?
		OR shipping_instruction_id IN ( SELECT si.id from shipping_instruction si
		inner join consignment_item ci ON ci.shipping_instruction_id = si.id
		WHERE ci.shipment_id = ? );`

	row := rs.DBService.DB.QueryRowxContext(ctx, nselectReferencesSQL, in.ShipmentId, in.ShipmentId)
	referenceTmp := bkgstruct.Reference1{}
	err := row.StructScan(&referenceTmp)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	reference, err := rs.getReferenceStruct(ctx, &getRequest, referenceTmp)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	refResponse := bkgproto.FindByShipmentIdResponse{}
	refResponse.Reference1 = reference
	return &refResponse, nil
}

// FindByTransportDocumentReference - FindByTransportDocumentReference
func (rs *ReferenceService) FindByTransportDocumentReference(ctx context.Context, in *bkgproto.FindByTransportDocumentReferenceRequest) (*bkgproto.FindByTransportDocumentReferenceResponse, error) {
	nselectReferencesSQL := selectReferencesSQL + `si inner join shipping_instruction si ON shipping_instruction_id = si.id
		inner join transport_document td ON td.shipping_instruction_id = si.id
		WHERE  td.transport_document_reference = ?
		OR shipment_id IN ( SELECT s.id from shipment s
		inner join consignment_item ci ON ci.shipment_id = s.id
		inner join shipping_instruction si ON ci.shipping_instruction_id = si.id
		inner join transport_document td2 ON  td2.shipping_instruction_id = si.id
		WHERE td2.transport_document_reference = ? );`

	row := rs.DBService.DB.QueryRowxContext(ctx, nselectReferencesSQL, in.TransportDocumentReference, in.TransportDocumentReference)
	referenceTmp := bkgstruct.Reference1{}
	err := row.StructScan(&referenceTmp)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	reference, err := rs.getReferenceStruct(ctx, &getRequest, referenceTmp)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	refResponse := bkgproto.FindByTransportDocumentReferenceResponse{}
	refResponse.Reference1 = reference
	return &refResponse, nil
}

// FindByCarrierBookingReference - FindByCarrierBookingReference
func (rs *ReferenceService) FindByCarrierBookingReference(ctx context.Context, in *bkgproto.FindByCarrierBookingReferenceRequest) (*bkgproto.FindByCarrierBookingReferenceResponse, error) {
	nselectReferencesSQL := selectReferencesSQL + `s inner join shipping_instruction si ON shipping_instruction_id = si.id left inner join shipment s ON shipment_id = s.id
		WHERE s.carrier_booking_reference = ?
		OR shipping_instruction_id IN ( SELECT si.id from shipping_instruction si
		inner join consignment_item ci ON ci.shipping_instruction_id = si.id
		inner join shipment s ON ci.shipment_id = s.id
		WHERE s.carrier_booking_reference = ? );`

	row := rs.DBService.DB.QueryRowxContext(ctx, nselectReferencesSQL, in.CarrierBookingReference, in.CarrierBookingReference)
	referenceTmp := bkgstruct.Reference1{}
	err := row.StructScan(&referenceTmp)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	reference, err := rs.getReferenceStruct(ctx, &getRequest, referenceTmp)
	if err != nil {
		rs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	refResponse := bkgproto.FindByCarrierBookingReferenceResponse{}
	refResponse.Reference1 = reference
	return &refResponse, nil
}

// getReferenceStruct - Get reference
func (rs *ReferenceService) getReferenceStruct(ctx context.Context, in *commonproto.GetRequest, referenceTmp bkgstruct.Reference1) (*bkgproto.Reference1, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(referenceTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(referenceTmp.CrUpdTime.UpdatedAt)

	reference := bkgproto.Reference1{Reference1D: referenceTmp.Reference1D, CrUpdUser: referenceTmp.CrUpdUser, CrUpdTime: crUpdTime}
	return &reference, nil
}
