package eblservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	eblstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ebl/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// TransportDocumentService - For accessing Transport Document services
type TransportDocumentService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	CurrencyService   *common.CurrencyService
	eblproto.UnimplementedTransportDocumentServiceServer
}

// NewTransportDocumentService - Create Transport Document service
func NewTransportDocumentService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient, currency *common.CurrencyService) *TransportDocumentService {
	return &TransportDocumentService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
		CurrencyService:   currency,
	}
}

// InsertTransportDocumentSQL - Insert TransportDocumentSQL Query
const insertTransportDocumentSQL = `insert into transport_documents
	  ( 
  uuid4,
  transport_document_reference,
  location_id,
  number_of_originals,
  carrier_id,
  shipping_instruction_id,
  declared_value_currency,
  declared_value,
  number_of_rider_pages,
  issuing_party,
  issue_date,
  shipped_onboard_date,
  received_for_shipment_date,
  created_date_time,
  updated_date_time
  )
  values (:uuid4,
:transport_document_reference,
:location_id,
:number_of_originals,
:carrier_id,
:shipping_instruction_id,
:declared_value_currency,
:declared_value,
:number_of_rider_pages,
:issuing_party,
:issue_date,
:shipped_onboard_date,
:received_for_shipment_date,
:created_date_time,
:updated_date_time);`

// selectTransportDocumentsSQL - select TransportDocumentsSQL Query
const selectTransportDocumentsSQL = `select 
  id,
  uuid4,
  transport_document_reference,
  location_id,
  number_of_originals,
  carrier_id,
  shipping_instruction_id,
  declared_value_currency,
  declared_value,
  number_of_rider_pages,
  issuing_party,
  issue_date,
  shipped_onboard_date,
  received_for_shipment_date,
  created_date_time,
  updated_date_time from transport_documents`

// CreateTransportDocument - Create TransportDocument
func (tds *TransportDocumentService) CreateTransportDocument(ctx context.Context, in *eblproto.CreateTransportDocumentRequest) (*eblproto.CreateTransportDocumentResponse, error) {
	issueDate, err := time.Parse(common.Layout, in.IssueDate)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shippedOnboardDate, err := time.Parse(common.Layout, in.ShippedOnboardDate)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivedForShipmentDate, err := time.Parse(common.Layout, in.ReceivedForShipmentDate)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	createdDateTime, err := time.Parse(common.Layout, in.CreatedDateTime)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	updatedDateTime, err := time.Parse(common.Layout, in.UpdatedDateTime)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocumentD := eblproto.TransportDocumentD{}
	transportDocumentD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	transportDocumentD.TransportDocumentReference = in.TransportDocumentReference
	transportDocumentD.LocationId = in.LocationId
	transportDocumentD.NumberOfOriginals = in.NumberOfOriginals
	transportDocumentD.CarrierId = in.CarrierId
	transportDocumentD.ShippingInstructionId = in.ShippingInstructionId

	declaredValueCurrency, err := tds.CurrencyService.GetCurrency(ctx, in.DeclaredValueCurrency)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	declaredValueMinor, err := common.ParseAmountString(in.DeclaredValue, declaredValueCurrency)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocumentD.DeclaredValueCurrency = declaredValueCurrency.Code
	transportDocumentD.DeclaredValue = declaredValueMinor
	transportDocumentD.DeclaredValueString = common.FormatAmountString(declaredValueMinor, declaredValueCurrency)

	transportDocumentD.NumberOfRiderPages = in.NumberOfRiderPages
	transportDocumentD.IssuingParty = in.IssuingParty

	transportDocumentT := eblproto.TransportDocumentT{}
	transportDocumentT.IssueDate = common.TimeToTimestamp(issueDate.UTC().Truncate(time.Second))
	transportDocumentT.ShippedOnboardDate = common.TimeToTimestamp(shippedOnboardDate.UTC().Truncate(time.Second))
	transportDocumentT.ReceivedForShipmentDate = common.TimeToTimestamp(receivedForShipmentDate.UTC().Truncate(time.Second))
	transportDocumentT.CreatedDateTime = common.TimeToTimestamp(createdDateTime.UTC().Truncate(time.Second))
	transportDocumentT.UpdatedDateTime = common.TimeToTimestamp(updatedDateTime.UTC().Truncate(time.Second))

	transportDocument := eblproto.TransportDocument{TransportDocumentD: &transportDocumentD, TransportDocumentT: &transportDocumentT}

	err = tds.insertTransportDocument(ctx, insertTransportDocumentSQL, &transportDocument, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	tranDocResponse := eblproto.CreateTransportDocumentResponse{}
	tranDocResponse.TransportDocument = &transportDocument
	return &tranDocResponse, nil
}

// CreateTransportDocumentFromShippingInstructionTO - Create TransportDocumentFromShippingInstructionTO
func (tds *TransportDocumentService) CreateTransportDocumentFromShippingInstructionTO(ctx context.Context, inReq *eblproto.CreateTransportDocumentFromShippingInstructionTORequest) (*eblproto.CreateTransportDocumentFromShippingInstructionTOResponse, error) {
	in := inReq.CreateTransportDocumentRequest
	shippingInstructionServ := &ShippingInstructionService{DBService: tds.DBService, RedisService: tds.RedisService, UserServiceClient: tds.UserServiceClient}
	getShippingInstructionByPkRequest := eblproto.GetShippingInstructionByPkRequest{}
	getByIdRequest := commonproto.GetByIdRequest{}
	getByIdRequest.Id = in.ShippingInstructionId
	getByIdRequest.UserEmail = in.UserEmail
	getByIdRequest.RequestId = in.RequestId
	getShippingInstructionByPkRequest.GetByIdRequest = &getByIdRequest

	shippingInstructionResp, err := shippingInstructionServ.GetShippingInstructionByPk(ctx, &getShippingInstructionByPkRequest)
	shippingInstruction := shippingInstructionResp.ShippingInstruction
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issueDate, err := time.Parse(common.Layout, in.IssueDate)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shippedOnboardDate, err := time.Parse(common.Layout, in.ShippedOnboardDate)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivedForShipmentDate, err := time.Parse(common.Layout, in.ReceivedForShipmentDate)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	createdDateTime, err := time.Parse(common.Layout, in.CreatedDateTime)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	updatedDateTime, err := time.Parse(common.Layout, in.UpdatedDateTime)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocumentD := eblproto.TransportDocumentD{}
	transportDocumentD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocumentD.TransportDocumentReference = in.TransportDocumentReference
	transportDocumentD.LocationId = shippingInstruction.ShippingInstructionD.LocationId
	transportDocumentD.NumberOfOriginals = in.NumberOfOriginals
	transportDocumentD.CarrierId = in.CarrierId
	transportDocumentD.ShippingInstructionId = shippingInstruction.ShippingInstructionD.Id

	declaredValueCurrency, err := tds.CurrencyService.GetCurrency(ctx, in.DeclaredValueCurrency)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	declaredValueMinor, err := common.ParseAmountString(in.DeclaredValue, declaredValueCurrency)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocumentD.DeclaredValueCurrency = declaredValueCurrency.Code
	transportDocumentD.DeclaredValue = declaredValueMinor
	transportDocumentD.DeclaredValueString = common.FormatAmountString(declaredValueMinor, declaredValueCurrency)
	transportDocumentD.NumberOfRiderPages = in.NumberOfRiderPages
	transportDocumentD.IssuingParty = in.IssuingParty

	transportDocumentT := eblproto.TransportDocumentT{}
	transportDocumentT.IssueDate = common.TimeToTimestamp(issueDate.UTC().Truncate(time.Second))
	transportDocumentT.ShippedOnboardDate = common.TimeToTimestamp(shippedOnboardDate.UTC().Truncate(time.Second))
	transportDocumentT.ReceivedForShipmentDate = common.TimeToTimestamp(receivedForShipmentDate.UTC().Truncate(time.Second))
	transportDocumentT.CreatedDateTime = common.TimeToTimestamp(createdDateTime.UTC().Truncate(time.Second))
	transportDocumentT.UpdatedDateTime = common.TimeToTimestamp(updatedDateTime.UTC().Truncate(time.Second))

	transportDocument := eblproto.TransportDocument{TransportDocumentD: &transportDocumentD, TransportDocumentT: &transportDocumentT}

	err = tds.insertTransportDocument(ctx, insertTransportDocumentSQL, &transportDocument, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	tranDocResponse := eblproto.CreateTransportDocumentFromShippingInstructionTOResponse{}
	tranDocResponse.TransportDocument = &transportDocument
	return &tranDocResponse, nil
}

// insertTransportDocument - Insert Transport Document
func (tds *TransportDocumentService) insertTransportDocument(ctx context.Context, insertTransportDocumentSQL string, transportDocument *eblproto.TransportDocument, userEmail string, requestID string) error {
	transportDocumentTmp, err := tds.CrTransportDocumentStruct(ctx, transportDocument, userEmail, requestID)
	if err != nil {
		tds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = tds.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransportDocumentSQL, transportDocumentTmp)
		if err != nil {
			tds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			tds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transportDocument.TransportDocumentD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transportDocument.TransportDocumentD.Uuid4)
		if err != nil {
			tds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transportDocument.TransportDocumentD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		tds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrTransportDocumentStruct - process TransportDocument details
func (tds *TransportDocumentService) CrTransportDocumentStruct(ctx context.Context, transportDocument *eblproto.TransportDocument, userEmail string, requestID string) (*eblstruct.TransportDocument, error) {
	transportDocumentT := new(eblstruct.TransportDocumentT)
	transportDocumentT.CreatedDateTime = common.TimestampToTime(transportDocument.TransportDocumentT.CreatedDateTime)
	transportDocumentT.UpdatedDateTime = common.TimestampToTime(transportDocument.TransportDocumentT.UpdatedDateTime)
	transportDocumentT.IssueDate = common.TimestampToTime(transportDocument.TransportDocumentT.IssueDate)
	transportDocumentT.ShippedOnboardDate = common.TimestampToTime(transportDocument.TransportDocumentT.ShippedOnboardDate)
	transportDocumentT.ReceivedForShipmentDate = common.TimestampToTime(transportDocument.TransportDocumentT.ReceivedForShipmentDate)

	transportDocumentTmp := eblstruct.TransportDocument{TransportDocumentD: transportDocument.TransportDocumentD, TransportDocumentT: transportDocumentT}
	return &transportDocumentTmp, nil
}

// GetTransportDocuments - Get TransportDocuments
func (tds *TransportDocumentService) GetTransportDocuments(ctx context.Context, in *eblproto.GetTransportDocumentsRequest) (*eblproto.GetTransportDocumentsResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = tds.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	transportDocuments := []*eblproto.TransportDocument{}

	nselectTransportDocumentsSQL := selectTransportDocumentsSQL + query

	rows, err := tds.DBService.DB.QueryxContext(ctx, nselectTransportDocumentsSQL)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		transportDocumentTmp := eblstruct.TransportDocument{}
		err = rows.StructScan(&transportDocumentTmp)
		if err != nil {
			tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		transportDocument, err := tds.GetTransportDocumentStruct(ctx, &getRequest, transportDocumentTmp)
		if err != nil {
			tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		declaredValueCurrency, err := tds.CurrencyService.GetCurrency(ctx, transportDocument.TransportDocumentD.DeclaredValueCurrency)
		if err != nil {
			tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		transportDocument.TransportDocumentD.DeclaredValueString = common.FormatAmountString(transportDocument.TransportDocumentD.DeclaredValue, declaredValueCurrency)

		transportDocuments = append(transportDocuments, transportDocument)

	}

	trasDocsResponse := eblproto.GetTransportDocumentsResponse{}
	if len(transportDocuments) != 0 {
		next := transportDocuments[len(transportDocuments)-1].TransportDocumentD.Id
		next--
		nextc := common.EncodeCursor(next)
		trasDocsResponse = eblproto.GetTransportDocumentsResponse{TransportDocuments: transportDocuments, NextCursor: nextc}
	} else {
		trasDocsResponse = eblproto.GetTransportDocumentsResponse{TransportDocuments: transportDocuments, NextCursor: "0"}
	}
	return &trasDocsResponse, nil
}

// FindTransportDocumentById - FindTransportDocumentById
func (tds *TransportDocumentService) FindTransportDocumentById(ctx context.Context, inReq *eblproto.FindTransportDocumentByIdRequest) (*eblproto.FindTransportDocumentByIdResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectTransportDocumentsSQL := selectTransportDocumentsSQL + ` where uuid4 = ?;`
	row := tds.DBService.DB.QueryRowxContext(ctx, nselectTransportDocumentsSQL, uuid4byte)
	transportDocumentTmp := eblstruct.TransportDocument{}
	err = row.StructScan(&transportDocumentTmp)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocument, err := tds.GetTransportDocumentStruct(ctx, in, transportDocumentTmp)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	declaredValueCurrency, err := tds.CurrencyService.GetCurrency(ctx, transportDocument.TransportDocumentD.DeclaredValueCurrency)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocument.TransportDocumentD.DeclaredValueString = common.FormatAmountString(transportDocument.TransportDocumentD.DeclaredValue, declaredValueCurrency)

	tranDocResponse := eblproto.FindTransportDocumentByIdResponse{}
	tranDocResponse.TransportDocument = transportDocument

	return &tranDocResponse, nil
}

// GetTransportDocumentByPk - Get TransportDocument By Primary key(Id)
func (tds *TransportDocumentService) GetTransportDocumentByPk(ctx context.Context, inReq *eblproto.GetTransportDocumentByPkRequest) (*eblproto.GetTransportDocumentByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectTransportDocumentsSQL := selectTransportDocumentsSQL + ` where id = ?;`
	row := tds.DBService.DB.QueryRowxContext(ctx, nselectTransportDocumentsSQL, in.Id)
	transportDocumentTmp := eblstruct.TransportDocument{}
	err := row.StructScan(&transportDocumentTmp)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	transportDocument, err := tds.GetTransportDocumentStruct(ctx, &getRequest, transportDocumentTmp)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	declaredValueCurrency, err := tds.CurrencyService.GetCurrency(ctx, transportDocument.TransportDocumentD.DeclaredValueCurrency)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocument.TransportDocumentD.DeclaredValueString = common.FormatAmountString(transportDocument.TransportDocumentD.DeclaredValue, declaredValueCurrency)

	tranDocResponse := eblproto.GetTransportDocumentByPkResponse{}
	tranDocResponse.TransportDocument = transportDocument

	return &tranDocResponse, nil
}

// GetTransportDocumentStruct - Get TransportDocument header
func (tds *TransportDocumentService) GetTransportDocumentStruct(ctx context.Context, in *commonproto.GetRequest, transportDocumentTmp eblstruct.TransportDocument) (*eblproto.TransportDocument, error) {
	transportDocumentT := new(eblproto.TransportDocumentT)
	transportDocumentT.CreatedDateTime = common.TimeToTimestamp(transportDocumentTmp.TransportDocumentT.CreatedDateTime)
	transportDocumentT.UpdatedDateTime = common.TimeToTimestamp(transportDocumentTmp.TransportDocumentT.UpdatedDateTime)
	transportDocumentT.IssueDate = common.TimeToTimestamp(transportDocumentTmp.TransportDocumentT.IssueDate)
	transportDocumentT.ShippedOnboardDate = common.TimeToTimestamp(transportDocumentTmp.TransportDocumentT.ShippedOnboardDate)
	transportDocumentT.ReceivedForShipmentDate = common.TimeToTimestamp(transportDocumentTmp.TransportDocumentT.ReceivedForShipmentDate)

	uuid4Str, err := common.UUIDBytesToStr(transportDocumentTmp.TransportDocumentD.Uuid4)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocumentTmp.TransportDocumentD.IdS = uuid4Str

	transportDocument := eblproto.TransportDocument{TransportDocumentD: transportDocumentTmp.TransportDocumentD, TransportDocumentT: transportDocumentT}

	return &transportDocument, nil
}

// FindByTransportDocumentReference - Find By TransportDocumentReference
func (tds *TransportDocumentService) FindByTransportDocumentReference(ctx context.Context, in *eblproto.FindByTransportDocumentReferenceRequest) (*eblproto.FindByTransportDocumentReferenceResponse, error) {
	nselectTransportDocumentsSQL := selectTransportDocumentsSQL + ` where transport_document_reference = ?;`
	row := tds.DBService.DB.QueryRowxContext(ctx, nselectTransportDocumentsSQL, in.TransportDocumentReference)
	transportDocumentTmp := eblstruct.TransportDocument{}
	err := row.StructScan(&transportDocumentTmp)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId

	transportDocument, err := tds.GetTransportDocumentStruct(ctx, &getRequest, transportDocumentTmp)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	declaredValueCurrency, err := tds.CurrencyService.GetCurrency(ctx, transportDocument.TransportDocumentD.DeclaredValueCurrency)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocument.TransportDocumentD.DeclaredValueString = common.FormatAmountString(transportDocument.TransportDocumentD.DeclaredValue, declaredValueCurrency)

	tranDocResponse := eblproto.FindByTransportDocumentReferenceResponse{}
	tranDocResponse.TransportDocument = transportDocument

	return &tranDocResponse, nil
}

// ApproveTransportDocument - Approve TransportDocument
func (tds *TransportDocumentService) ApproveTransportDocument(ctx context.Context, in *eblproto.ApproveTransportDocumentRequest) (*eblproto.ApproveTransportDocumentResponse, error) {
	return &eblproto.ApproveTransportDocumentResponse{}, nil
}
