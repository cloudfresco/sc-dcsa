package eblservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eblstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ebl/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// SurrenderRequestService - For accessing Surrender services
type SurrenderRequestService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eblproto.UnimplementedSurrenderRequestServiceServer
}

// NewSurrenderRequestService - Create Shipping service
func NewSurrenderRequestService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *SurrenderRequestService {
	return &SurrenderRequestService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertTransactionPartySQL - Insert TransactionPartySQL Query
const insertTransactionPartySQL = `insert into transaction_parties
	  ( 
  uuid4,
  ebl_platform_identifier,
  legal_name,
  registration_number,
  location_of_registration,
  tax_reference,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:ebl_platform_identifier,
:legal_name,
:registration_number,
:location_of_registration,
:tax_reference,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectTransactionPartiesSQL - select TransactionPartySQL Query
const selectTransactionPartiesSQL = `select 
  id,
  uuid4,
  ebl_platform_identifier,
  legal_name,
  registration_number,
  location_of_registration,
  tax_reference,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from transaction_parties`

// updateTransactionPartySQL - update TransactionPartySQL query
const updateTransactionPartySQL = `update transaction_parties set 
  ebl_platform_identifier = ?,
  legal_name = ?,
  registration_number = ?,
  location_of_registration = ?,
  updated_at = ? where uuid4 = ?;`

// insertTransactionPartySupportingCodeSQL - Insert TransactionPartySupportingCodeSQL Query
const insertTransactionPartySupportingCodeSQL = `insert into transaction_party_supporting_codes
	  ( 
  uuid4,
  transaction_party_id,
  party_code,
  party_code_list_provider,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:transaction_party_id,
:party_code,
:party_code_list_provider,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// updateTransactionPartySupportingCodeSQL - update TransactionPartySupportingCodeSQL query
const updateTransactionPartySupportingCodeSQL = `update transaction_party_supporting_codes set 
  transaction_party_id = ?,
  party_code = ?,
  party_code_list_provider = ?,
  updated_at = ? where uuid4 = ?;`

// insertSurrenderRequestSQL - Insert SurrenderRequestSQL Query
const insertSurrenderRequestSQL = `insert into surrender_requests
	  ( 
  uuid4,
  surrender_request_reference,
  transport_document_reference,
  surrender_request_code,
  comments,
  surrender_requested_by,
  created_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:surrender_request_reference,
:transport_document_reference,
:surrender_request_code,
:comments,
:surrender_requested_by,
:created_date_time,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// updateSurrenderRequestSQL - update SurrenderRequestSQL query
const updateSurrenderRequestSQL = `update surrender_requests set 
  surrender_request_reference = ?,
  transport_document_reference = ?,
  surrender_request_code = ?,
  comments = ?,
  updated_at = ? where uuid4 = ?;`

// insertEndorsementChainLinkSQL - Insert EndorsementChainLinkSQL Query
const insertEndorsementChainLinkSQL = `insert into endorsement_chain_links
	  ( 
  uuid4,
  entry_order,
  action_date_time,
  actor,
  recipient,
  surrender_request_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:entry_order,
:action_date_time,
:actor,
:recipient,
:surrender_request_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// updateEndorsementChainLinkSQL - update EndorsementChainLinkSQL query
const updateEndorsementChainLinkSQL = `update endorsement_chain_links set 
  entry_order = ?,
  actor = ?,
  recipient = ?,
  updated_at = ? where uuid4 = ?;`

// CreateTransactionParty - Create  TransactionParty
func (ss *SurrenderRequestService) CreateTransactionParty(ctx context.Context, in *eblproto.CreateTransactionPartyRequest) (*eblproto.CreateTransactionPartyResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ss.UserServiceClient)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	transactionPartyD := eblproto.TransactionPartyD{}
	transactionPartyD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionPartyD.EblPlatformIdentifier = in.EblPlatformIdentifier
	transactionPartyD.LegalName = in.LegalName
	transactionPartyD.RegistrationNumber = in.RegistrationNumber
	transactionPartyD.LocationOfRegistration = in.LocationOfRegistration
	transactionPartyD.TaxReference = in.TaxReference

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	transactionParty := eblproto.TransactionParty{TransactionPartyD: &transactionPartyD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ss.insertTransactionParty(ctx, insertTransactionPartySQL, &transactionParty, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionPartyResponse := eblproto.CreateTransactionPartyResponse{}
	transactionPartyResponse.TransactionParty = &transactionParty
	return &transactionPartyResponse, nil
}

// insertTransactionParty - Insert TransactionParty
func (ss *SurrenderRequestService) insertTransactionParty(ctx context.Context, insertTransactionPartySQL string, transactionParty *eblproto.TransactionParty, userEmail string, requestID string) error {
	transactionPartyTmp, err := ss.CrTransactionPartyStruct(ctx, transactionParty, userEmail, requestID)
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = ss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransactionPartySQL, transactionPartyTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionParty.TransactionPartyD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transactionParty.TransactionPartyD.Uuid4)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionParty.TransactionPartyD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrTransactionPartyStruct - process TransactionParty details
func (ss *SurrenderRequestService) CrTransactionPartyStruct(ctx context.Context, transactionParty *eblproto.TransactionParty, userEmail string, requestID string) (*eblstruct.TransactionParty, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(transactionParty.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(transactionParty.CrUpdTime.UpdatedAt)

	transactionPartyTmp := eblstruct.TransactionParty{TransactionPartyD: transactionParty.TransactionPartyD, CrUpdUser: transactionParty.CrUpdUser, CrUpdTime: crUpdTime}

	return &transactionPartyTmp, nil
}

// GetTransactionParties - Get  TransactionParties
func (ss *SurrenderRequestService) GetTransactionParties(ctx context.Context, in *eblproto.GetTransactionPartiesRequest) (*eblproto.GetTransactionPartiesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = ss.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	transactionParties := []*eblproto.TransactionParty{}

	nselectTransactionPartiesSQL := selectTransactionPartiesSQL + query

	rows, err := ss.DBService.DB.QueryxContext(ctx, nselectTransactionPartiesSQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		transactionPartyTmp := eblstruct.TransactionParty{}
		err = rows.StructScan(&transactionPartyTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		transactionParty, err := ss.getTransactionPartyStruct(ctx, &getRequest, transactionPartyTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		transactionParties = append(transactionParties, transactionParty)

	}

	transactionPartiesResponse := eblproto.GetTransactionPartiesResponse{}
	if len(transactionParties) != 0 {
		next := transactionParties[len(transactionParties)-1].TransactionPartyD.Id
		next--
		nextc := common.EncodeCursor(next)
		transactionPartiesResponse = eblproto.GetTransactionPartiesResponse{TransactionParties: transactionParties, NextCursor: nextc}
	} else {
		transactionPartiesResponse = eblproto.GetTransactionPartiesResponse{TransactionParties: transactionParties, NextCursor: "0"}
	}
	return &transactionPartiesResponse, nil
}

// GetTransactionParty - Get TransactionParty
func (ss *SurrenderRequestService) GetTransactionParty(ctx context.Context, inReq *eblproto.GetTransactionPartyRequest) (*eblproto.GetTransactionPartyResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectTransactionPartiesSQL := selectTransactionPartiesSQL + ` where uuid4 = ?;`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectTransactionPartiesSQL, uuid4byte)
	transactionPartyTmp := eblstruct.TransactionParty{}
	err = row.StructScan(&transactionPartyTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	transactionParty, err := ss.getTransactionPartyStruct(ctx, &getRequest, transactionPartyTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionPartyResponse := eblproto.GetTransactionPartyResponse{}
	transactionPartyResponse.TransactionParty = transactionParty
	return &transactionPartyResponse, nil
}

// GetTransactionPartyByPk - Get TransactionParty By Primary key(Id)
func (ss *SurrenderRequestService) GetTransactionPartyByPk(ctx context.Context, inReq *eblproto.GetTransactionPartyByPkRequest) (*eblproto.GetTransactionPartyByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectTransactionPartiesSQL := selectTransactionPartiesSQL + ` where id = ?;`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectTransactionPartiesSQL, in.Id)
	transactionPartyTmp := eblstruct.TransactionParty{}
	err := row.StructScan(&transactionPartyTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	transactionParty, err := ss.getTransactionPartyStruct(ctx, &getRequest, transactionPartyTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionPartyResponse := eblproto.GetTransactionPartyByPkResponse{}
	transactionPartyResponse.TransactionParty = transactionParty
	return &transactionPartyResponse, nil
}

// GetTransactionPartyStruct - Get TransactionParty header
func (ss *SurrenderRequestService) getTransactionPartyStruct(ctx context.Context, in *commonproto.GetRequest, transactionPartyTmp eblstruct.TransactionParty) (*eblproto.TransactionParty, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(transactionPartyTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(transactionPartyTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(transactionPartyTmp.TransactionPartyD.Uuid4)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionPartyTmp.TransactionPartyD.IdS = uuid4Str

	transactionParty := eblproto.TransactionParty{TransactionPartyD: transactionPartyTmp.TransactionPartyD, CrUpdUser: transactionPartyTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &transactionParty, nil
}

// UpdateTransactionParty - Update TransactionParty
func (ss *SurrenderRequestService) UpdateTransactionParty(ctx context.Context, in *eblproto.UpdateTransactionPartyRequest) (*eblproto.UpdateTransactionPartyResponse, error) {
	db := ss.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateTransactionPartySQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ss.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.EblPlatformIdentifier,
			in.LegalName,
			in.RegistrationNumber,
			in.LocationOfRegistration,
			tn,
			uuid4byte)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateTransactionPartyResponse{}, nil
}

// CreateTransactionPartySupportingCode - Create  TransactionPartySupportingCode
func (ss *SurrenderRequestService) CreateTransactionPartySupportingCode(ctx context.Context, in *eblproto.CreateTransactionPartySupportingCodeRequest) (*eblproto.CreateTransactionPartySupportingCodeResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ss.UserServiceClient)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	transactionPartySupportingCodeD := eblproto.TransactionPartySupportingCodeD{}
	transactionPartySupportingCodeD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionPartySupportingCodeD.TransactionPartyId = in.TransactionPartyId
	transactionPartySupportingCodeD.PartyCode = in.PartyCode
	transactionPartySupportingCodeD.PartyCodeListProvider = in.PartyCodeListProvider

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	transactionPartySupportingCode := eblproto.TransactionPartySupportingCode{TransactionPartySupportingCodeD: &transactionPartySupportingCodeD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ss.insertTransactionPartySupportingCode(ctx, insertTransactionPartySupportingCodeSQL, &transactionPartySupportingCode, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionPartySupportingCodeResponse := eblproto.CreateTransactionPartySupportingCodeResponse{}
	transactionPartySupportingCodeResponse.TransactionPartySupportingCode = &transactionPartySupportingCode
	return &transactionPartySupportingCodeResponse, nil
}

// insertTransactionPartySupportingCode - Insert TransactionPartySupportingCode
func (ss *SurrenderRequestService) insertTransactionPartySupportingCode(ctx context.Context, insertTransactionPartySupportingCodeSQL string, transactionPartySupportingCode *eblproto.TransactionPartySupportingCode, userEmail string, requestID string) error {
	transactionPartySupportingCodeTmp, err := ss.CrTransactionPartySupportingCodeStruct(ctx, transactionPartySupportingCode, userEmail, requestID)
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = ss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransactionPartySupportingCodeSQL, transactionPartySupportingCodeTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionPartySupportingCode.TransactionPartySupportingCodeD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transactionPartySupportingCode.TransactionPartySupportingCodeD.Uuid4)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionPartySupportingCode.TransactionPartySupportingCodeD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrTransactionPartySupportingCodeStruct - process TransactionPartySupportingCode details
func (ss *SurrenderRequestService) CrTransactionPartySupportingCodeStruct(ctx context.Context, transactionPartySupportingCode *eblproto.TransactionPartySupportingCode, userEmail string, requestID string) (*eblstruct.TransactionPartySupportingCode, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(transactionPartySupportingCode.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(transactionPartySupportingCode.CrUpdTime.UpdatedAt)

	transactionPartySupportingCodeTmp := eblstruct.TransactionPartySupportingCode{TransactionPartySupportingCodeD: transactionPartySupportingCode.TransactionPartySupportingCodeD, CrUpdUser: transactionPartySupportingCode.CrUpdUser, CrUpdTime: crUpdTime}

	return &transactionPartySupportingCodeTmp, nil
}

// UpdateTransactionPartySupportingCode - Update TransactionPartySupportingCode
func (ss *SurrenderRequestService) UpdateTransactionPartySupportingCode(ctx context.Context, in *eblproto.UpdateTransactionPartySupportingCodeRequest) (*eblproto.UpdateTransactionPartySupportingCodeResponse, error) {
	db := ss.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateTransactionPartySupportingCodeSQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ss.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.TransactionPartyId,
			in.PartyCode,
			in.PartyCodeListProvider,
			tn,
			uuid4byte)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateTransactionPartySupportingCodeResponse{}, nil
}

// CreateEndorsementChainLink - Create  EndorsementChainLink
func (ss *SurrenderRequestService) CreateEndorsementChainLink(ctx context.Context, in *eblproto.CreateEndorsementChainLinkRequest) (*eblproto.CreateEndorsementChainLinkResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ss.UserServiceClient)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	actionDateTime, err := time.Parse(common.Layout, in.ActionDateTime)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	endorsementChainLinkD := eblproto.EndorsementChainLinkD{}
	endorsementChainLinkD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	endorsementChainLinkD.EntryOrder = in.EntryOrder
	endorsementChainLinkD.Actor = in.Actor
	endorsementChainLinkD.Recipient = in.Recipient
	endorsementChainLinkD.SurrenderRequestId = in.SurrenderRequestId

	endorsementChainLinkT := eblproto.EndorsementChainLinkT{}
	endorsementChainLinkT.ActionDateTime = common.TimeToTimestamp(actionDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	endorsementChainLink := eblproto.EndorsementChainLink{EndorsementChainLinkD: &endorsementChainLinkD, EndorsementChainLinkT: &endorsementChainLinkT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ss.insertEndorsementChainLink(ctx, insertEndorsementChainLinkSQL, &endorsementChainLink, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	endorsementChainLinkResponse := eblproto.CreateEndorsementChainLinkResponse{}
	endorsementChainLinkResponse.EndorsementChainLink = &endorsementChainLink
	return &endorsementChainLinkResponse, nil
}

// insertEndorsementChainLink - Insert EndorsementChainLink
func (ss *SurrenderRequestService) insertEndorsementChainLink(ctx context.Context, insertEndorsementChainLinkSQL string, endorsementChainLink *eblproto.EndorsementChainLink, userEmail string, requestID string) error {
	endorsementChainLinkTmp, err := ss.CrEndorsementChainLinkStruct(ctx, endorsementChainLink, userEmail, requestID)
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = ss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertEndorsementChainLinkSQL, endorsementChainLinkTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		endorsementChainLink.EndorsementChainLinkD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(endorsementChainLink.EndorsementChainLinkD.Uuid4)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		endorsementChainLink.EndorsementChainLinkD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrEndorsementChainLinkStruct - process EndorsementChainLink details
func (ss *SurrenderRequestService) CrEndorsementChainLinkStruct(ctx context.Context, endorsementChainLink *eblproto.EndorsementChainLink, userEmail string, requestID string) (*eblstruct.EndorsementChainLink, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(endorsementChainLink.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(endorsementChainLink.CrUpdTime.UpdatedAt)

	endorsementChainLinkT := new(eblstruct.EndorsementChainLinkT)
	endorsementChainLinkT.ActionDateTime = common.TimestampToTime(endorsementChainLink.EndorsementChainLinkT.ActionDateTime)

	endorsementChainLinkTmp := eblstruct.EndorsementChainLink{EndorsementChainLinkD: endorsementChainLink.EndorsementChainLinkD, EndorsementChainLinkT: endorsementChainLinkT, CrUpdUser: endorsementChainLink.CrUpdUser, CrUpdTime: crUpdTime}

	return &endorsementChainLinkTmp, nil
}

// UpdateEndorsementChainLink - Update EndorsementChainLink
func (ss *SurrenderRequestService) UpdateEndorsementChainLink(ctx context.Context, in *eblproto.UpdateEndorsementChainLinkRequest) (*eblproto.UpdateEndorsementChainLinkResponse, error) {
	db := ss.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateEndorsementChainLinkSQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ss.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.EntryOrder,
			in.Actor,
			in.Recipient,
			tn,
			uuid4byte)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateEndorsementChainLinkResponse{}, nil
}

// CreateSurrenderRequest - Create  SurrenderRequest
func (ss *SurrenderRequestService) CreateSurrenderRequest(ctx context.Context, in *eblproto.CreateSurrenderRequestRequest) (*eblproto.CreateSurrenderRequestResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ss.UserServiceClient)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	createdDateTime, err := time.Parse(common.Layout, in.CreatedDateTime)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	surrenderRequestD := eblproto.SurrenderRequestD{}
	surrenderRequestD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	surrenderRequestD.SurrenderRequestReference = in.SurrenderRequestReference
	surrenderRequestD.TransportDocumentReference = in.TransportDocumentReference
	surrenderRequestD.SurrenderRequestCode = in.SurrenderRequestCode
	surrenderRequestD.Comments = in.Comments
	surrenderRequestD.SurrenderRequestedBy = in.SurrenderRequestedBy

	surrenderRequestT := eblproto.SurrenderRequestT{}
	surrenderRequestT.CreatedDateTime = common.TimeToTimestamp(createdDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	surrenderRequest := eblproto.SurrenderRequest{SurrenderRequestD: &surrenderRequestD, SurrenderRequestT: &surrenderRequestT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ss.insertSurrenderRequest(ctx, insertSurrenderRequestSQL, &surrenderRequest, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	surrenderRequestResponse := eblproto.CreateSurrenderRequestResponse{}
	surrenderRequestResponse.SurrenderRequest = &surrenderRequest
	return &surrenderRequestResponse, nil
}

// insertSurrenderRequest - Insert SurrenderRequest
func (ss *SurrenderRequestService) insertSurrenderRequest(ctx context.Context, insertSurrenderRequestSQL string, surrenderRequest *eblproto.SurrenderRequest, userEmail string, requestID string) error {
	surrenderRequestTmp, err := ss.CrSurrenderRequestStruct(ctx, surrenderRequest, userEmail, requestID)
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = ss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertSurrenderRequestSQL, surrenderRequestTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		surrenderRequest.SurrenderRequestD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(surrenderRequest.SurrenderRequestD.Uuid4)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		surrenderRequest.SurrenderRequestD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrSurrenderRequestStruct - process SurrenderRequest details
func (ss *SurrenderRequestService) CrSurrenderRequestStruct(ctx context.Context, surrenderRequest *eblproto.SurrenderRequest, userEmail string, requestID string) (*eblstruct.SurrenderRequest, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(surrenderRequest.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(surrenderRequest.CrUpdTime.UpdatedAt)

	surrenderRequestT := new(eblstruct.SurrenderRequestT)
	surrenderRequestT.CreatedDateTime = common.TimestampToTime(surrenderRequest.SurrenderRequestT.CreatedDateTime)

	surrenderRequestTmp := eblstruct.SurrenderRequest{SurrenderRequestD: surrenderRequest.SurrenderRequestD, SurrenderRequestT: surrenderRequestT, CrUpdUser: surrenderRequest.CrUpdUser, CrUpdTime: crUpdTime}

	return &surrenderRequestTmp, nil
}

// UpdateSurrenderRequest - Update SurrenderRequest
func (ss *SurrenderRequestService) UpdateSurrenderRequest(ctx context.Context, in *eblproto.UpdateSurrenderRequestRequest) (*eblproto.UpdateSurrenderRequestResponse, error) {
	db := ss.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateSurrenderRequestSQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ss.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.SurrenderRequestReference,
			in.TransportDocumentReference,
			in.SurrenderRequestCode,
			in.Comments,
			tn,
			uuid4byte)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateSurrenderRequestResponse{}, nil
}
