package eblservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eblstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ebl/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// insertTransportDocumentSummarySQL - Insert TransportDocumentSummarySQL Query
const insertTransportDocumentSummarySQL = `insert into transport_document_summaries
	  ( 
uuid4,
transport_document_reference,
number_of_originals,
carrier_code,
carrier_code_list_provider,
number_of_rider_pages,
shipping_instruction_reference,
document_status,
transport_document_created_date_time,
transport_document_updated_date_time,
issue_date,
shipped_onboard_date,
received_for_shipment_date,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:transport_document_reference,
:number_of_originals,
:carrier_code,
:carrier_code_list_provider,
:number_of_rider_pages,
:shipping_instruction_reference,
:document_status,
:transport_document_created_date_time,
:transport_document_updated_date_time,
:issue_date,
:shipped_onboard_date,
:received_for_shipment_date,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectTransportDocumentSummariesSQL - select TransportDocumentSummariesSQL Query
const selectTransportDocumentSummariesSQL = `select 
  id,
  uuid4,
  transport_document_reference,
  number_of_originals,
  carrier_code,
  carrier_code_list_provider,
  number_of_rider_pages,
  shipping_instruction_reference,
  document_status,  
  transport_document_created_date_time,
  transport_document_updated_date_time,
  issue_date,
  shipped_onboard_date,
  received_for_shipment_date,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from transport_document_summaries`

// CreateTransportDocumentSummary - Create TransportDocumentSummary
func (tds *TransportDocumentService) CreateTransportDocumentSummary(ctx context.Context, in *eblproto.CreateTransportDocumentSummaryRequest) (*eblproto.CreateTransportDocumentSummaryResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, tds.UserServiceClient)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

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

	transportDocumentCreatedDateTime, err := time.Parse(common.Layout, in.TransportDocumentCreatedDateTime)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocumentUpdatedDateTime, err := time.Parse(common.Layout, in.TransportDocumentUpdatedDateTime)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocumentSummaryD := eblproto.TransportDocumentSummaryD{}
	transportDocumentSummaryD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocumentSummaryD.TransportDocumentReference = in.TransportDocumentReference
	transportDocumentSummaryD.NumberOfOriginals = in.NumberOfOriginals
	transportDocumentSummaryD.CarrierCode = in.CarrierCode
	transportDocumentSummaryD.CarrierCodeListProvider = in.CarrierCodeListProvider
	transportDocumentSummaryD.NumberOfRiderPages = in.NumberOfRiderPages
	transportDocumentSummaryD.ShippingInstructionReference = in.ShippingInstructionReference
	transportDocumentSummaryD.DocumentStatus = in.DocumentStatus

	transportDocumentSummaryT := eblproto.TransportDocumentSummaryT{}
	transportDocumentSummaryT.TransportDocumentCreatedDateTime = common.TimeToTimestamp(transportDocumentCreatedDateTime.UTC().Truncate(time.Second))
	transportDocumentSummaryT.TransportDocumentUpdatedDateTime = common.TimeToTimestamp(transportDocumentUpdatedDateTime.UTC().Truncate(time.Second))
	transportDocumentSummaryT.IssueDate = common.TimeToTimestamp(issueDate.UTC().Truncate(time.Second))
	transportDocumentSummaryT.ShippedOnboardDate = common.TimeToTimestamp(shippedOnboardDate.UTC().Truncate(time.Second))
	transportDocumentSummaryT.ReceivedForShipmentDate = common.TimeToTimestamp(receivedForShipmentDate.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	transportDocumentSummary := eblproto.TransportDocumentSummary{TransportDocumentSummaryD: &transportDocumentSummaryD, TransportDocumentSummaryT: &transportDocumentSummaryT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = tds.insertTransportDocumentSummary(ctx, insertTransportDocumentSummarySQL, &transportDocumentSummary, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	tranDocSumResponse := eblproto.CreateTransportDocumentSummaryResponse{}
	tranDocSumResponse.TransportDocumentSummary = &transportDocumentSummary

	return &tranDocSumResponse, nil
}

// insertTransportDocumentSummary - Insert TransportDocumentSummary
func (tds *TransportDocumentService) insertTransportDocumentSummary(ctx context.Context, insertTransportDocumentSummarySQL string, transportDocumentSummary *eblproto.TransportDocumentSummary, userEmail string, requestID string) error {
	transportDocumentSummaryTmp, err := tds.CrTransportDocumentSummaryStruct(ctx, transportDocumentSummary, userEmail, requestID)
	if err != nil {
		tds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = tds.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransportDocumentSummarySQL, transportDocumentSummaryTmp)
		if err != nil {
			tds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			tds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transportDocumentSummary.TransportDocumentSummaryD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transportDocumentSummary.TransportDocumentSummaryD.Uuid4)
		if err != nil {
			tds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transportDocumentSummary.TransportDocumentSummaryD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		tds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrTransportDocumentSummaryStruct - process TransportDocumentSummary details
func (tds *TransportDocumentService) CrTransportDocumentSummaryStruct(ctx context.Context, transportDocumentSummary *eblproto.TransportDocumentSummary, userEmail string, requestID string) (*eblstruct.TransportDocumentSummary, error) {
	transportDocumentSummaryT := new(eblstruct.TransportDocumentSummaryT)
	transportDocumentSummaryT.TransportDocumentCreatedDateTime = common.TimestampToTime(transportDocumentSummary.TransportDocumentSummaryT.TransportDocumentCreatedDateTime)
	transportDocumentSummaryT.TransportDocumentUpdatedDateTime = common.TimestampToTime(transportDocumentSummary.TransportDocumentSummaryT.TransportDocumentUpdatedDateTime)
	transportDocumentSummaryT.IssueDate = common.TimestampToTime(transportDocumentSummary.TransportDocumentSummaryT.IssueDate)
	transportDocumentSummaryT.ShippedOnboardDate = common.TimestampToTime(transportDocumentSummary.TransportDocumentSummaryT.ShippedOnboardDate)
	transportDocumentSummaryT.ReceivedForShipmentDate = common.TimestampToTime(transportDocumentSummary.TransportDocumentSummaryT.ReceivedForShipmentDate)

	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(transportDocumentSummary.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(transportDocumentSummary.CrUpdTime.UpdatedAt)

	transportDocumentSummaryTmp := eblstruct.TransportDocumentSummary{TransportDocumentSummaryD: transportDocumentSummary.TransportDocumentSummaryD, TransportDocumentSummaryT: transportDocumentSummaryT, CrUpdUser: transportDocumentSummary.CrUpdUser, CrUpdTime: crUpdTime}

	return &transportDocumentSummaryTmp, nil
}

// GetTransportDocumentSummaries - Get TransportDocumentSummaries
func (tds *TransportDocumentService) GetTransportDocumentSummaries(ctx context.Context, in *eblproto.GetTransportDocumentSummariesRequest) (*eblproto.GetTransportDocumentSummariesResponse, error) {
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

	transportDocumentSummaries := []*eblproto.TransportDocumentSummary{}

	nselectTransportDocumentSummariesSQL := selectTransportDocumentSummariesSQL + query

	rows, err := tds.DBService.DB.QueryxContext(ctx, nselectTransportDocumentSummariesSQL)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		transportDocumentSummaryTmp := eblstruct.TransportDocumentSummary{}
		err = rows.StructScan(&transportDocumentSummaryTmp)
		if err != nil {
			tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		transportDocumentSummary, err := tds.GetTransportDocumentSummaryStruct(ctx, &getRequest, transportDocumentSummaryTmp)
		if err != nil {
			tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		transportDocumentSummaries = append(transportDocumentSummaries, transportDocumentSummary)

	}

	trasDocSumsResponse := eblproto.GetTransportDocumentSummariesResponse{}
	if len(transportDocumentSummaries) != 0 {
		next := transportDocumentSummaries[len(transportDocumentSummaries)-1].TransportDocumentSummaryD.Id
		next--
		nextc := common.EncodeCursor(next)
		trasDocSumsResponse = eblproto.GetTransportDocumentSummariesResponse{TransportDocumentSummaries: transportDocumentSummaries, NextCursor: nextc}
	} else {
		trasDocSumsResponse = eblproto.GetTransportDocumentSummariesResponse{TransportDocumentSummaries: transportDocumentSummaries, NextCursor: "0"}
	}
	return &trasDocSumsResponse, nil
}

// GetTransportDocumentSummaryByPk - Get TransportDocumentSummary By Primary key(Id)
func (tds *TransportDocumentService) GetTransportDocumentSummaryByPk(ctx context.Context, inReq *eblproto.GetTransportDocumentSummaryByPkRequest) (*eblproto.GetTransportDocumentSummaryByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectTransportDocumentSummariesSQL := selectTransportDocumentSummariesSQL + ` where id = ?;`
	row := tds.DBService.DB.QueryRowxContext(ctx, nselectTransportDocumentSummariesSQL, in.Id)
	transportDocumentSummaryTmp := eblstruct.TransportDocumentSummary{}
	err := row.StructScan(&transportDocumentSummaryTmp)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	transportDocumentSummary, err := tds.GetTransportDocumentSummaryStruct(ctx, &getRequest, transportDocumentSummaryTmp)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	tranDocSumResponse := eblproto.GetTransportDocumentSummaryByPkResponse{}
	tranDocSumResponse.TransportDocumentSummary = transportDocumentSummary

	return &tranDocSumResponse, nil
}

// GetTransportDocumentSummaryStruct - Get TransportDocumentSummary header
func (tds *TransportDocumentService) GetTransportDocumentSummaryStruct(ctx context.Context, in *commonproto.GetRequest, transportDocumentSummaryTmp eblstruct.TransportDocumentSummary) (*eblproto.TransportDocumentSummary, error) {
	transportDocumentSummaryT := new(eblproto.TransportDocumentSummaryT)
	transportDocumentSummaryT.TransportDocumentCreatedDateTime = common.TimeToTimestamp(transportDocumentSummaryTmp.TransportDocumentSummaryT.TransportDocumentCreatedDateTime)
	transportDocumentSummaryT.TransportDocumentUpdatedDateTime = common.TimeToTimestamp(transportDocumentSummaryTmp.TransportDocumentSummaryT.TransportDocumentUpdatedDateTime)
	transportDocumentSummaryT.IssueDate = common.TimeToTimestamp(transportDocumentSummaryTmp.TransportDocumentSummaryT.IssueDate)
	transportDocumentSummaryT.ShippedOnboardDate = common.TimeToTimestamp(transportDocumentSummaryTmp.TransportDocumentSummaryT.ShippedOnboardDate)
	transportDocumentSummaryT.ReceivedForShipmentDate = common.TimeToTimestamp(transportDocumentSummaryTmp.TransportDocumentSummaryT.ReceivedForShipmentDate)

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(transportDocumentSummaryTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(transportDocumentSummaryTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(transportDocumentSummaryTmp.TransportDocumentSummaryD.Uuid4)
	if err != nil {
		tds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportDocumentSummaryTmp.TransportDocumentSummaryD.IdS = uuid4Str

	transportDocumentSummary := eblproto.TransportDocumentSummary{TransportDocumentSummaryD: transportDocumentSummaryTmp.TransportDocumentSummaryD, TransportDocumentSummaryT: transportDocumentSummaryT, CrUpdUser: transportDocumentSummaryTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &transportDocumentSummary, nil
}
