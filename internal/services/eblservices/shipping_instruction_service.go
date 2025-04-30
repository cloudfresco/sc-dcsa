package eblservices

// https://github.com/dcsaorg/DCSA-EBL/blob/master/src/main/java/org/dcsa/ebl/service/ShippingInstructionService.java
import (
	"context"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	bkgservice "github.com/cloudfresco/sc-dcsa/internal/services/bkgservices"
	eventcoreservice "github.com/cloudfresco/sc-dcsa/internal/services/eventcoreservices"
	eblstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ebl/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// ShippingInstructionService - For accessing Shipping Instruction services
type ShippingInstructionService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eblproto.UnimplementedShippingInstructionServiceServer
}

// NewShippingInstructionService - Create Shipping Instruction service
func NewShippingInstructionService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *ShippingInstructionService {
	return &ShippingInstructionService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertShippingInstructionSQL - Insert ShippingInstructionSQL Query
const insertShippingInstructionSQL = `insert into shipping_instructions
	  ( 
  uuid4,
  shipping_instruction_reference,
  document_status,
  is_shipped_onboard_type,
  number_of_copies,
  number_of_originals,
  is_electronic,
  is_to_order,
  are_charges_displayed_on_originals,
  are_charges_displayed_on_copies,
  location_id,
  transport_document_type_code,
  displayed_name_for_place_of_receipt,
  displayed_name_for_port_of_load,
  displayed_name_for_port_of_discharge,
  displayed_name_for_place_of_delivery,
  amend_to_transport_document,
  created_date_time,
  updated_date_time
  )
  values (:uuid4,
  :shipping_instruction_reference,
  :document_status,
  :is_shipped_onboard_type,
  :number_of_copies,
  :number_of_originals,
  :is_electronic,
  :is_to_order,
  :are_charges_displayed_on_originals,
  :are_charges_displayed_on_copies,
  :location_id,
  :transport_document_type_code,
  :displayed_name_for_place_of_receipt,
  :displayed_name_for_port_of_load,
  :displayed_name_for_port_of_discharge,
  :displayed_name_for_place_of_delivery,
  :amend_to_transport_document,
  :created_date_time,
  :updated_date_time);`

// selectShippingInstructionsSQL - select ShippingInstructionsSQL Query
const selectShippingInstructionsSQL = `select 
  id,
  uuid4,
  shipping_instruction_reference,
  document_status,
  is_shipped_onboard_type,
  number_of_copies,
  number_of_originals,
  is_electronic,
  is_to_order,
  are_charges_displayed_on_originals,
  are_charges_displayed_on_copies,
  location_id,
  transport_document_type_code,
  displayed_name_for_place_of_receipt,
  displayed_name_for_port_of_load,
  displayed_name_for_port_of_discharge,
  displayed_name_for_place_of_delivery,
  amend_to_transport_document,
  created_date_time,
  updated_date_time from shipping_instructions`

// updateShippingInstructionSQL - update ShippingInstructionsSQL Query
const updateShippingInstructionSQL = `update shipping_instructions set 
    document_status = ?,
    transport_document_type_code = ?,
    displayed_name_for_place_of_receipt = ?,
		updated_date_time = ? where shipping_instruction_reference = ?;`

// CreateShippingInstruction - Create ShippingInstruction
func (sis *ShippingInstructionService) CreateShippingInstruction(ctx context.Context, in *eblproto.CreateShippingInstructionRequest) (*eblproto.CreateShippingInstructionResponse, error) {
	createdDateTime, err := time.Parse(common.Layout, in.CreatedDateTime)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	updatedDateTime, err := time.Parse(common.Layout, in.UpdatedDateTime)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shippingInstructionD := eblproto.ShippingInstructionD{}
	shippingInstructionD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shippingInstructionD.ShippingInstructionReference = in.ShippingInstructionReference
	shippingInstructionD.DocumentStatus = in.DocumentStatus
	shippingInstructionD.IsShippedOnboardType = in.IsShippedOnboardType
	shippingInstructionD.NumberOfCopies = in.NumberOfCopies
	shippingInstructionD.NumberOfOriginals = in.NumberOfOriginals
	shippingInstructionD.IsElectronic = in.IsElectronic
	shippingInstructionD.IsToOrder = in.IsToOrder
	shippingInstructionD.AreChargesDisplayedOnOriginals = in.AreChargesDisplayedOnOriginals
	shippingInstructionD.AreChargesDisplayedOnCopies = in.AreChargesDisplayedOnCopies
	shippingInstructionD.LocationId = in.LocationId
	shippingInstructionD.TransportDocumentTypeCode = in.TransportDocumentTypeCode
	shippingInstructionD.DisplayedNameForPlaceOfReceipt = in.DisplayedNameForPlaceOfReceipt
	shippingInstructionD.DisplayedNameForPortOfLoad = in.DisplayedNameForPortOfLoad
	shippingInstructionD.DisplayedNameForPortOfDischarge = in.DisplayedNameForPortOfDischarge
	shippingInstructionD.DisplayedNameForPlaceOfDelivery = in.DisplayedNameForPlaceOfDelivery
	shippingInstructionD.AmendToTransportDocument = in.AmendToTransportDocument

	shippingInstructionT := eblproto.ShippingInstructionT{}
	shippingInstructionT.CreatedDateTime = common.TimeToTimestamp(createdDateTime.UTC().Truncate(time.Second))
	shippingInstructionT.UpdatedDateTime = common.TimeToTimestamp(updatedDateTime.UTC().Truncate(time.Second))

	shippingInstruction := eblproto.ShippingInstruction{ShippingInstructionD: &shippingInstructionD, ShippingInstructionT: &shippingInstructionT}

	uequipments := []*eventcoreproto.UtilizedTransportEquipment{}
	// we will do for loop on references, UtilizedTransportEquipments, which is comes from client form
	for _, uequipment := range in.UtilizedTransportEquipments {
		uequipment.UserId = in.UserId
		uequipment.UserEmail = in.UserEmail
		uequipment.RequestId = in.RequestId
		utilizedTransportEquipmentServ := &eventcoreservice.UtilizedTransportEquipmentService{DBService: sis.DBService, RedisService: sis.RedisService, UserServiceClient: sis.UserServiceClient}
		uequipment, err := utilizedTransportEquipmentServ.ProcessUtilizedTransportEquipmentRequest(ctx, uequipment)
		if err != nil {
			sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		uequipments = append(uequipments, uequipment)
	}

	consignmentItems := []*eblproto.ConsignmentItem{}
	consignmentItemServ := &ConsignmentItemService{DBService: sis.DBService, RedisService: sis.RedisService, UserServiceClient: sis.UserServiceClient}
	for _, consignmentItem := range in.ConsignmentItems {
		consignmentItem.UserId = in.UserId
		consignmentItem.UserEmail = in.UserEmail
		consignmentItem.RequestId = in.RequestId
		consignmentItem, err := consignmentItemServ.ProcessConsignmentItemRequest(ctx, consignmentItem)
		if err != nil {
			sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		consignmentItems = append(consignmentItems, consignmentItem)
	}

	references := []*bkgproto.Reference1{}
	referenceServ := &bkgservice.ReferenceService{DBService: sis.DBService, RedisService: sis.RedisService, UserServiceClient: sis.UserServiceClient}
	for _, reference := range in.References {
		reference.UserId = in.UserId
		reference.UserEmail = in.UserEmail
		reference.RequestId = in.RequestId
		reference, err := referenceServ.ProcessReferenceRequest(ctx, reference)
		if err != nil {
			sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		references = append(references, reference)
	}

	err = sis.insertShippingInstruction(ctx, insertShippingInstructionSQL, &shippingInstruction, eventcoreservice.InsertUtilizedTransportEquipmentSQL, uequipments, bkgservice.InsertReferenceSQL, references, InsertConsignmentItemSQL, consignmentItems, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	shipInstResponse := eblproto.CreateShippingInstructionResponse{}
	shipInstResponse.ShippingInstruction = &shippingInstruction
	return &shipInstResponse, nil
}

// insertShippingInstruction - Insert ShippingInstruction
func (sis *ShippingInstructionService) insertShippingInstruction(ctx context.Context, insertShippingInstructionSQL string, shippingInstruction *eblproto.ShippingInstruction, insertUtilizedTransportEquipmentSQL string, uequipments []*eventcoreproto.UtilizedTransportEquipment, insertReferenceSQL string, references []*bkgproto.Reference1, insertConsignmentItemSQL string, consignmentItems []*eblproto.ConsignmentItem, userEmail string, requestID string) error {
	shippingInstructionTmp, err := sis.crShippingInstructionStruct(ctx, shippingInstruction, userEmail, requestID)
	if err != nil {
		sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = sis.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertShippingInstructionSQL, shippingInstructionTmp)
		if err != nil {
			sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		shippingInstruction.ShippingInstructionD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(shippingInstruction.ShippingInstructionD.Uuid4)
		if err != nil {
			sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		shippingInstruction.ShippingInstructionD.IdS = uuid4Str
		for _, uequipment := range uequipments {
			_, err = tx.NamedExecContext(ctx, insertUtilizedTransportEquipmentSQL, uequipment)
			if err != nil {
				sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}

			_, err = tx.NamedExecContext(ctx, eventcoreservice.InsertEquipmentSQL, uequipment.Equipment)
			if err != nil {
				sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}

		}
		for _, reference := range references {
			reference.Reference1D.ShippingInstructionId = shippingInstruction.ShippingInstructionD.Id
			_, err = tx.NamedExecContext(ctx, insertReferenceSQL, reference)
			if err != nil {
				sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		consignmentItemServ := &ConsignmentItemService{DBService: sis.DBService, RedisService: sis.RedisService, UserServiceClient: sis.UserServiceClient}
		for _, consignmentItem := range consignmentItems {
			consignmentItem.ConsignmentItemD.ShippingInstructionId = shippingInstruction.ShippingInstructionD.Id
			consignmentItemTmp, err := consignmentItemServ.CrConsignmentItemStruct(ctx, consignmentItem, userEmail, requestID)
			if err != nil {
				sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			res1, err := tx.NamedExecContext(ctx, insertConsignmentItemSQL, consignmentItemTmp)
			if err != nil {
				sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			uID1, err := res1.LastInsertId()
			if err != nil {
				sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}

			for _, cargoItem := range consignmentItem.CargoItems {
				cargoItem.CargoItemD.ConsignmentItemId = uint32(uID1)
				resp2, err := tx.NamedExecContext(ctx, InsertCargoItemSQL, cargoItem)
				if err != nil {
					sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
					return err
				}
				uID2, err := resp2.LastInsertId()
				if err != nil {
					sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
					return err
				}

				for _, cargoLineItem := range cargoItem.CargoLineItems {
					cargoLineItem.CargoLineItemD.CargoItemId = uint32(uID2)
					_, err = tx.NamedExecContext(ctx, InsertCargoLineItemSQL, cargoLineItem)
					if err != nil {
						sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
						return err
					}
				}
			}

		}

		return nil
	})

	if err != nil {
		sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crShippingInstructionStruct - process ShippingInstruction details
func (sis *ShippingInstructionService) crShippingInstructionStruct(ctx context.Context, shippingInstruction *eblproto.ShippingInstruction, userEmail string, requestID string) (*eblstruct.ShippingInstruction, error) {
	shippingInstructionT := new(eblstruct.ShippingInstructionT)
	shippingInstructionT.CreatedDateTime = common.TimestampToTime(shippingInstruction.ShippingInstructionT.CreatedDateTime)
	shippingInstructionT.UpdatedDateTime = common.TimestampToTime(shippingInstruction.ShippingInstructionT.UpdatedDateTime)

	shippingInstructionTmp := eblstruct.ShippingInstruction{ShippingInstructionD: shippingInstruction.ShippingInstructionD, ShippingInstructionT: shippingInstructionT}

	return &shippingInstructionTmp, nil
}

// GetShippingInstructions - Get ShippingInstructions
func (sis *ShippingInstructionService) GetShippingInstructions(ctx context.Context, in *eblproto.GetShippingInstructionsRequest) (*eblproto.GetShippingInstructionsResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = sis.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	shippingInstructions := []*eblproto.ShippingInstruction{}

	nselectShippingInstructionsSQL := selectShippingInstructionsSQL + query

	rows, err := sis.DBService.DB.QueryxContext(ctx, nselectShippingInstructionsSQL)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		shippingInstructionTmp := eblstruct.ShippingInstruction{}
		err = rows.StructScan(&shippingInstructionTmp)
		if err != nil {
			sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		shippingInstruction, err := sis.getShippingInstructionStruct(ctx, &getRequest, shippingInstructionTmp)
		if err != nil {
			sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		shippingInstructions = append(shippingInstructions, shippingInstruction)

	}

	shipInstsResponse := eblproto.GetShippingInstructionsResponse{}
	if len(shippingInstructions) != 0 {
		next := shippingInstructions[len(shippingInstructions)-1].ShippingInstructionD.Id
		next--
		nextc := common.EncodeCursor(next)
		shipInstsResponse = eblproto.GetShippingInstructionsResponse{ShippingInstructions: shippingInstructions, NextCursor: nextc}
	} else {
		shipInstsResponse = eblproto.GetShippingInstructionsResponse{ShippingInstructions: shippingInstructions, NextCursor: "0"}
	}
	return &shipInstsResponse, nil
}

// FindById - Get FindById
func (sis *ShippingInstructionService) FindById(ctx context.Context, inReq *eblproto.FindByIdRequest) (*eblproto.FindByIdResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectShippingInstructionsSQL := selectShippingInstructionsSQL + ` where uuid4 = ?;`
	row := sis.DBService.DB.QueryRowxContext(ctx, nselectShippingInstructionsSQL, uuid4byte)
	shippingInstructionTmp := eblstruct.ShippingInstruction{}
	err = row.StructScan(&shippingInstructionTmp)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shippingInstruction, err := sis.getShippingInstructionStruct(ctx, in, shippingInstructionTmp)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipInstResponse := eblproto.FindByIdResponse{}
	shipInstResponse.ShippingInstruction = shippingInstruction

	return &shipInstResponse, nil
}

// GetShippingInstructionByPk - Get ShippingInstruction By Primary key(Id)
func (sis *ShippingInstructionService) GetShippingInstructionByPk(ctx context.Context, inReq *eblproto.GetShippingInstructionByPkRequest) (*eblproto.GetShippingInstructionByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectShippingInstructionsSQL := selectShippingInstructionsSQL + ` where id = ?;`
	row := sis.DBService.DB.QueryRowxContext(ctx, nselectShippingInstructionsSQL, in.Id)
	shippingInstructionTmp := eblstruct.ShippingInstruction{}
	err := row.StructScan(&shippingInstructionTmp)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	shippingInstruction, err := sis.getShippingInstructionStruct(ctx, &getRequest, shippingInstructionTmp)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	shipInstResponse := eblproto.GetShippingInstructionByPkResponse{}
	shipInstResponse.ShippingInstruction = shippingInstruction

	return &shipInstResponse, nil
}

// getShippingInstructionStruct - Get shippingInstruction header
func (sis *ShippingInstructionService) getShippingInstructionStruct(ctx context.Context, in *commonproto.GetRequest, shippingInstructionTmp eblstruct.ShippingInstruction) (*eblproto.ShippingInstruction, error) {
	shippingInstructionT := new(eblproto.ShippingInstructionT)
	shippingInstructionT.CreatedDateTime = common.TimeToTimestamp(shippingInstructionTmp.ShippingInstructionT.CreatedDateTime)
	shippingInstructionT.UpdatedDateTime = common.TimeToTimestamp(shippingInstructionTmp.ShippingInstructionT.UpdatedDateTime)

	uuid4Str, err := common.UUIDBytesToStr(shippingInstructionTmp.ShippingInstructionD.Uuid4)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shippingInstructionTmp.ShippingInstructionD.IdS = uuid4Str

	shippingInstruction := eblproto.ShippingInstruction{ShippingInstructionD: shippingInstructionTmp.ShippingInstructionD, ShippingInstructionT: shippingInstructionT}

	return &shippingInstruction, nil
}

// FindByReference - Get FindByReference
func (sis *ShippingInstructionService) FindByReference(ctx context.Context, in *eblproto.FindByReferenceRequest) (*eblproto.FindByReferenceResponse, error) {
	nselectShippingInstructionsSQL := selectShippingInstructionsSQL + ` where shipping_instruction_reference = ?;`
	row := sis.DBService.DB.QueryRowxContext(ctx, nselectShippingInstructionsSQL, in.ShippingInstructionReference)
	shippingInstructionTmp := eblstruct.ShippingInstruction{}
	err := row.StructScan(&shippingInstructionTmp)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId

	shippingInstruction, err := sis.getShippingInstructionStruct(ctx, &getRequest, shippingInstructionTmp)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipInstResponse := eblproto.FindByReferenceResponse{}
	shipInstResponse.ShippingInstruction = shippingInstruction

	return &shipInstResponse, nil
}

// UpdateShippingInstructionByShippingInstructionReference - UpdateShippingInstructionByShippingInstructionReference
func (sis *ShippingInstructionService) UpdateShippingInstructionByShippingInstructionReference(ctx context.Context, in *eblproto.UpdateShippingInstructionByShippingInstructionReferenceRequest) (*eblproto.UpdateShippingInstructionByShippingInstructionReferenceResponse, error) {
	db := sis.DBService.DB
	tn := common.GetTimeDetails()

	stmt, err := db.PreparexContext(ctx, updateShippingInstructionSQL)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = sis.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.DocumentStatus,
			in.TransportDocumentTypeCode,
			in.DisplayedNameForPlaceOfReceipt,
			tn,
			in.ShippingInstructionReference)
		if err != nil {
			sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return nil
			}
			return nil
		}
		return nil
	})

	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateShippingInstructionByShippingInstructionReferenceResponse{}, nil
}
