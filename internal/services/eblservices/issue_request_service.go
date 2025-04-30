package eblservices

import (
	"context"

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

// IssueRequestService - For accessing Issuance services
type IssueRequestService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eblproto.UnimplementedIssueRequestServiceServer
}

// NewIssueRequestService - Create Shipping service
func NewIssueRequestService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *IssueRequestService {
	return &IssueRequestService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertIssuePartySQL - Insert IssuePartySQL Query
const insertIssuePartySQL = `insert into issue_parties
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

// selectIssuePartiesSQL - select IssuePartySQL Query
const selectIssuePartiesSQL = `select 
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
  updated_at from issue_parties`

// updateIssuePartySQL - update IssuePartySQL query
const updateIssuePartySQL = `update issue_parties set 
  ebl_platform_identifier = ?,
  legal_name = ?,
  registration_number = ?,
  location_of_registration = ?,
  updated_at = ? where uuid4 = ?;`

// insertIssuePartySupportingCodeSQL - Insert IssuePartySupportingCodeSQL Query
const insertIssuePartySupportingCodeSQL = `insert into issue_party_supporting_codes
	  ( 
  uuid4,
  issue_party_id,
  party_code,
  party_code_list_provider,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:issue_party_id,
:party_code,
:party_code_list_provider,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// updateIssuePartySupportingCodeSQL - update IssuePartySupportingCodeSQL query
const updateIssuePartySupportingCodeSQL = `update issue_party_supporting_codes set 
  issue_party_id = ?,
  party_code = ?,
  party_code_list_provider = ?,
  updated_at = ? where uuid4 = ?;`

// insertIssuanceRequestSQL - Insert IssuanceRequestSQL Query
const insertIssuanceRequestSQL = `insert into issuance_requests
	  ( 
  uuid4,
  transport_document_reference,
  issuance_request_state,
  issue_to,
  ebl_visualization_id,
  transport_document_json,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:transport_document_reference,
:issuance_request_state,
:issue_to,
:ebl_visualization_id,
:transport_document_json,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// updateIssuanceRequestSQL - update IssuanceRequestSQL query
const updateIssuanceRequestSQL = `update issuance_requests set 
  transport_document_reference = ?,
  issuance_request_state = ?,
  updated_at = ? where uuid4 = ?;`

// insertEblVisualizationSQL - Insert EblVisualizationSQL Query
const insertEblVisualizationSQL = `insert into ebl_visualizations
	  ( 
  uuid4,
  name,
  content,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:name,
:content,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// updateEblVisualizationSQL - update EblVisualizationSQL query
const updateEblVisualizationSQL = `update ebl_visualizations set 
  name = ?,
  content = ?,
  updated_at = ? where uuid4 = ?;`

// CreateIssueParty - Create  IssueParty
func (is *IssueRequestService) CreateIssueParty(ctx context.Context, in *eblproto.CreateIssuePartyRequest) (*eblproto.CreateIssuePartyResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, is.UserServiceClient)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	issuePartyD := eblproto.IssuePartyD{}
	issuePartyD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuePartyD.EblPlatformIdentifier = in.EblPlatformIdentifier
	issuePartyD.LegalName = in.LegalName
	issuePartyD.RegistrationNumber = in.RegistrationNumber
	issuePartyD.LocationOfRegistration = in.LocationOfRegistration
	issuePartyD.TaxReference = in.TaxReference

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	issueParty := eblproto.IssueParty{IssuePartyD: &issuePartyD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = is.insertIssueParty(ctx, insertIssuePartySQL, &issueParty, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuePartyResponse := eblproto.CreateIssuePartyResponse{}
	issuePartyResponse.IssueParty = &issueParty
	return &issuePartyResponse, nil
}

// insertIssueParty - Insert IssueParty
func (is *IssueRequestService) insertIssueParty(ctx context.Context, insertIssuePartySQL string, issueParty *eblproto.IssueParty, userEmail string, requestID string) error {
	issuePartyTmp, err := is.CrIssuePartyStruct(ctx, issueParty, userEmail, requestID)
	if err != nil {
		is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = is.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertIssuePartySQL, issuePartyTmp)
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		issueParty.IssuePartyD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(issueParty.IssuePartyD.Uuid4)
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		issueParty.IssuePartyD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrIssuePartyStruct - process IssueParty details
func (is *IssueRequestService) CrIssuePartyStruct(ctx context.Context, issueParty *eblproto.IssueParty, userEmail string, requestID string) (*eblstruct.IssueParty, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(issueParty.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(issueParty.CrUpdTime.UpdatedAt)

	issuePartyTmp := eblstruct.IssueParty{IssuePartyD: issueParty.IssuePartyD, CrUpdUser: issueParty.CrUpdUser, CrUpdTime: crUpdTime}

	return &issuePartyTmp, nil
}

// GetIssueParties - Get  IssueParties
func (is *IssueRequestService) GetIssueParties(ctx context.Context, in *eblproto.GetIssuePartiesRequest) (*eblproto.GetIssuePartiesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = is.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	issueParties := []*eblproto.IssueParty{}

	nselectIssuePartiesSQL := selectIssuePartiesSQL + query

	rows, err := is.DBService.DB.QueryxContext(ctx, nselectIssuePartiesSQL)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		issuePartyTmp := eblstruct.IssueParty{}
		err = rows.StructScan(&issuePartyTmp)
		if err != nil {
			is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		issueParty, err := is.getIssuePartyStruct(ctx, &getRequest, issuePartyTmp)
		if err != nil {
			is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		issueParties = append(issueParties, issueParty)

	}

	issuePartiesResponse := eblproto.GetIssuePartiesResponse{}
	if len(issueParties) != 0 {
		next := issueParties[len(issueParties)-1].IssuePartyD.Id
		next--
		nextc := common.EncodeCursor(next)
		issuePartiesResponse = eblproto.GetIssuePartiesResponse{IssueParties: issueParties, NextCursor: nextc}
	} else {
		issuePartiesResponse = eblproto.GetIssuePartiesResponse{IssueParties: issueParties, NextCursor: "0"}
	}
	return &issuePartiesResponse, nil
}

// GetIssueParty - Get IssueParty
func (is *IssueRequestService) GetIssueParty(ctx context.Context, inReq *eblproto.GetIssuePartyRequest) (*eblproto.GetIssuePartyResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectIssuePartiesSQL := selectIssuePartiesSQL + ` where uuid4 = ?;`
	row := is.DBService.DB.QueryRowxContext(ctx, nselectIssuePartiesSQL, uuid4byte)
	issuePartyTmp := eblstruct.IssueParty{}
	err = row.StructScan(&issuePartyTmp)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	issueParty, err := is.getIssuePartyStruct(ctx, &getRequest, issuePartyTmp)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuePartyResponse := eblproto.GetIssuePartyResponse{}
	issuePartyResponse.IssueParty = issueParty
	return &issuePartyResponse, nil
}

// GetIssuePartyByPk - Get IssueParty By Primary key(Id)
func (is *IssueRequestService) GetIssuePartyByPk(ctx context.Context, inReq *eblproto.GetIssuePartyByPkRequest) (*eblproto.GetIssuePartyByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectIssuePartiesSQL := selectIssuePartiesSQL + ` where id = ?;`
	row := is.DBService.DB.QueryRowxContext(ctx, nselectIssuePartiesSQL, in.Id)
	issuePartyTmp := eblstruct.IssueParty{}
	err := row.StructScan(&issuePartyTmp)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	issueParty, err := is.getIssuePartyStruct(ctx, &getRequest, issuePartyTmp)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuePartyResponse := eblproto.GetIssuePartyByPkResponse{}
	issuePartyResponse.IssueParty = issueParty
	return &issuePartyResponse, nil
}

// GetIssuePartyStruct - Get IssueParty header
func (is *IssueRequestService) getIssuePartyStruct(ctx context.Context, in *commonproto.GetRequest, issuePartyTmp eblstruct.IssueParty) (*eblproto.IssueParty, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(issuePartyTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(issuePartyTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(issuePartyTmp.IssuePartyD.Uuid4)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuePartyTmp.IssuePartyD.IdS = uuid4Str

	issueParty := eblproto.IssueParty{IssuePartyD: issuePartyTmp.IssuePartyD, CrUpdUser: issuePartyTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &issueParty, nil
}

// UpdateIssueParty - Update IssueParty
func (is *IssueRequestService) UpdateIssueParty(ctx context.Context, in *eblproto.UpdateIssuePartyRequest) (*eblproto.UpdateIssuePartyResponse, error) {
	db := is.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateIssuePartySQL)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = is.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.EblPlatformIdentifier,
			in.LegalName,
			in.RegistrationNumber,
			in.LocationOfRegistration,
			tn,
			uuid4byte)
		if err != nil {
			is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateIssuePartyResponse{}, nil
}

// CreateIssuePartySupportingCode - Create  IssuePartySupportingCode
func (is *IssueRequestService) CreateIssuePartySupportingCode(ctx context.Context, in *eblproto.CreateIssuePartySupportingCodeRequest) (*eblproto.CreateIssuePartySupportingCodeResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, is.UserServiceClient)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	issuePartySupportingCodeD := eblproto.IssuePartySupportingCodeD{}
	issuePartySupportingCodeD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuePartySupportingCodeD.IssuePartyId = in.IssuePartyId
	issuePartySupportingCodeD.PartyCode = in.PartyCode
	issuePartySupportingCodeD.PartyCodeListProvider = in.PartyCodeListProvider

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	issuePartySupportingCode := eblproto.IssuePartySupportingCode{IssuePartySupportingCodeD: &issuePartySupportingCodeD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = is.insertIssuePartySupportingCode(ctx, insertIssuePartySupportingCodeSQL, &issuePartySupportingCode, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuePartySupportingCodeResponse := eblproto.CreateIssuePartySupportingCodeResponse{}
	issuePartySupportingCodeResponse.IssuePartySupportingCode = &issuePartySupportingCode
	return &issuePartySupportingCodeResponse, nil
}

// insertIssuePartySupportingCode - Insert IssuePartySupportingCode
func (is *IssueRequestService) insertIssuePartySupportingCode(ctx context.Context, insertIssuePartySupportingCodeSQL string, issuePartySupportingCode *eblproto.IssuePartySupportingCode, userEmail string, requestID string) error {
	issuePartySupportingCodeTmp, err := is.CrIssuePartySupportingCodeStruct(ctx, issuePartySupportingCode, userEmail, requestID)
	if err != nil {
		is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = is.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertIssuePartySupportingCodeSQL, issuePartySupportingCodeTmp)
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		issuePartySupportingCode.IssuePartySupportingCodeD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(issuePartySupportingCode.IssuePartySupportingCodeD.Uuid4)
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		issuePartySupportingCode.IssuePartySupportingCodeD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrIssuePartySupportingCodeStruct - process IssuePartySupportingCode details
func (is *IssueRequestService) CrIssuePartySupportingCodeStruct(ctx context.Context, issuePartySupportingCode *eblproto.IssuePartySupportingCode, userEmail string, requestID string) (*eblstruct.IssuePartySupportingCode, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(issuePartySupportingCode.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(issuePartySupportingCode.CrUpdTime.UpdatedAt)

	issuePartySupportingCodeTmp := eblstruct.IssuePartySupportingCode{IssuePartySupportingCodeD: issuePartySupportingCode.IssuePartySupportingCodeD, CrUpdUser: issuePartySupportingCode.CrUpdUser, CrUpdTime: crUpdTime}

	return &issuePartySupportingCodeTmp, nil
}

// UpdateIssuePartySupportingCode - Update IssuePartySupportingCode
func (is *IssueRequestService) UpdateIssuePartySupportingCode(ctx context.Context, in *eblproto.UpdateIssuePartySupportingCodeRequest) (*eblproto.UpdateIssuePartySupportingCodeResponse, error) {
	db := is.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateIssuePartySupportingCodeSQL)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = is.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.IssuePartyId,
			in.PartyCode,
			in.PartyCodeListProvider,
			tn,
			uuid4byte)
		if err != nil {
			is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateIssuePartySupportingCodeResponse{}, nil
}

// CreateEblVisualization - Create  EblVisualization
func (is *IssueRequestService) CreateEblVisualization(ctx context.Context, in *eblproto.CreateEblVisualizationRequest) (*eblproto.CreateEblVisualizationResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, is.UserServiceClient)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	eblVisualizationD := eblproto.EblVisualizationD{}
	eblVisualizationD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	eblVisualizationD.Name = in.Name
	eblVisualizationD.Content = in.Content

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	eblVisualization := eblproto.EblVisualization{EblVisualizationD: &eblVisualizationD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = is.insertEblVisualization(ctx, insertEblVisualizationSQL, &eblVisualization, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	eblVisualizationResponse := eblproto.CreateEblVisualizationResponse{}
	eblVisualizationResponse.EblVisualization = &eblVisualization
	return &eblVisualizationResponse, nil
}

// insertEblVisualization - Insert EblVisualization
func (is *IssueRequestService) insertEblVisualization(ctx context.Context, insertEblVisualizationSQL string, eblVisualization *eblproto.EblVisualization, userEmail string, requestID string) error {
	eblVisualizationTmp, err := is.CrEblVisualizationStruct(ctx, eblVisualization, userEmail, requestID)
	if err != nil {
		is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = is.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertEblVisualizationSQL, eblVisualizationTmp)
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		eblVisualization.EblVisualizationD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(eblVisualization.EblVisualizationD.Uuid4)
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		eblVisualization.EblVisualizationD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrEblVisualizationStruct - process EblVisualization details
func (is *IssueRequestService) CrEblVisualizationStruct(ctx context.Context, eblVisualization *eblproto.EblVisualization, userEmail string, requestID string) (*eblstruct.EblVisualization, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(eblVisualization.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(eblVisualization.CrUpdTime.UpdatedAt)

	eblVisualizationTmp := eblstruct.EblVisualization{EblVisualizationD: eblVisualization.EblVisualizationD, CrUpdUser: eblVisualization.CrUpdUser, CrUpdTime: crUpdTime}

	return &eblVisualizationTmp, nil
}

// UpdateEblVisualization - Update EblVisualization
func (is *IssueRequestService) UpdateEblVisualization(ctx context.Context, in *eblproto.UpdateEblVisualizationRequest) (*eblproto.UpdateEblVisualizationResponse, error) {
	db := is.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateEblVisualizationSQL)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = is.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.Name,
			in.Content,
			tn,
			uuid4byte)
		if err != nil {
			is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateEblVisualizationResponse{}, nil
}

// CreateIssuanceRequest - Create  IssuanceRequest
func (is *IssueRequestService) CreateIssuanceRequest(ctx context.Context, in *eblproto.CreateIssuanceRequestRequest) (*eblproto.CreateIssuanceRequestResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, is.UserServiceClient)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	issuanceRequestD := eblproto.IssuanceRequestD{}
	issuanceRequestD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuanceRequestD.TransportDocumentReference = in.TransportDocumentReference
	issuanceRequestD.IssuanceRequestState = in.IssuanceRequestState
	issuanceRequestD.IssueTo = in.IssueTo
	issuanceRequestD.EblVisualizationId = in.EblVisualizationId
	issuanceRequestD.TransportDocumentJson = in.TransportDocumentJson

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	issuanceRequest := eblproto.IssuanceRequest{IssuanceRequestD: &issuanceRequestD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = is.insertIssuanceRequest(ctx, insertIssuanceRequestSQL, &issuanceRequest, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuanceRequestResponse := eblproto.CreateIssuanceRequestResponse{}
	issuanceRequestResponse.IssuanceRequest = &issuanceRequest
	return &issuanceRequestResponse, nil
}

// insertIssuanceRequest - Insert IssuanceRequest
func (is *IssueRequestService) insertIssuanceRequest(ctx context.Context, insertIssuanceRequestSQL string, issuanceRequest *eblproto.IssuanceRequest, userEmail string, requestID string) error {
	issuanceRequestTmp, err := is.CrIssuanceRequestStruct(ctx, issuanceRequest, userEmail, requestID)
	if err != nil {
		is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = is.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertIssuanceRequestSQL, issuanceRequestTmp)
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		issuanceRequest.IssuanceRequestD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(issuanceRequest.IssuanceRequestD.Uuid4)
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		issuanceRequest.IssuanceRequestD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrIssuanceRequestStruct - process IssuanceRequest details
func (is *IssueRequestService) CrIssuanceRequestStruct(ctx context.Context, issuanceRequest *eblproto.IssuanceRequest, userEmail string, requestID string) (*eblstruct.IssuanceRequest, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(issuanceRequest.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(issuanceRequest.CrUpdTime.UpdatedAt)

	issuanceRequestTmp := eblstruct.IssuanceRequest{IssuanceRequestD: issuanceRequest.IssuanceRequestD, CrUpdUser: issuanceRequest.CrUpdUser, CrUpdTime: crUpdTime}

	return &issuanceRequestTmp, nil
}

// UpdateIssuanceRequest - Update IssuanceRequest
func (is *IssueRequestService) UpdateIssuanceRequest(ctx context.Context, in *eblproto.UpdateIssuanceRequestRequest) (*eblproto.UpdateIssuanceRequestResponse, error) {
	db := is.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateIssuanceRequestSQL)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = is.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.TransportDocumentReference,
			in.IssuanceRequestState,
			tn,
			uuid4byte)
		if err != nil {
			is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateIssuanceRequestResponse{}, nil
}
