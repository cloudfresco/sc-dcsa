package partyservices

import (
	"context"
	"net"
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	partystruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/party/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// PartyService - For accessing Party services
type PartyService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	partyproto.UnimplementedPartyServiceServer
}

// NewPartyService - Create Party service
func NewPartyService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *PartyService {
	return &PartyService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// StartPartyServer - Start Party server
func StartPartyServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
	common.SetJWTOpt(jwtOpt)

	creds, err := common.GetSrvCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		os.Exit(1)
	}

	userCreds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		os.Exit(1)
	}

	var srvOpts []grpc.ServerOption

	userConn, err := grpc.NewClient(grpcServerOpt.GrpcUserServerPort, grpc.WithTransportCredentials(userCreds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srvOpts = append(srvOpts, grpc.Creds(creds))

	srvOpts = append(srvOpts, grpc.StatsHandler(otelgrpc.NewServerHandler()))

	uc := partyproto.NewUserServiceClient(userConn)
	partyService := NewPartyService(log, dbService, redisService, uc)
	documentPartyService := NewDocumentPartyService(log, dbService, redisService, uc)
	locationService := NewLocationService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcPartyServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	partyproto.RegisterPartyServiceServer(srv, partyService)
	partyproto.RegisterDocumentPartyServiceServer(srv, documentPartyService)
	partyproto.RegisterLocationServiceServer(srv, locationService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

// insertPartySQL - insert PartySQL query
const insertPartySQL = `insert into parties
	  ( 
    uuid4,
    party_name,
    tax_reference1,
    tax_reference2,
    public_key,
    address_id,
    status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (:uuid4,
    :party_name,
    :tax_reference1,
    :tax_reference2,
    :public_key,
    :address_id,
    :status_code,
    :created_by_user_id,
    :updated_by_user_id,
    :created_at,
    :updated_at);`

// updatePartySQL - update PartySQL query
const updatePartySQL = `update parties set 
		  party_name = ?,
			updated_at = ? where id = ? and status_code = ?;`

// deletePartySQL - delete PartySQL query
const deletePartySQL = `update parties set 
		  status_code = ?,
			updated_at = ? where uuid4= ?;`

// selectPartiesSQL - select Parties query
const selectPartiesSQL = `select 
      id, 
      uuid4,
      party_name,
      tax_reference1,
      tax_reference2,
      public_key,
      address_id,
      status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from parties`

// insertPartyContactDetailSQL - insert PartyContactDetail query
const insertPartyContactDetailSQL = `insert into party_contact_details
	  (
      uuid4,
      party_id,
      name,
      email,
      phone,
      url,
     status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (:uuid4,
      :party_id,
      :name,
      :email,
      :phone,
      :url,
      :status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// updatePartyContactDetailSQL - update PartyContactDetailSQL query
const updatePartyContactDetailSQL = `update party_contact_details set 
		  name = ?,
			updated_at = ? where id = ? and status_code = ?;`

// deletePartyContactDetailSQL - delete PartyContactDetailSQL query
const deletePartyContactDetailSQL = `update party_contact_details set 
		  status_code = ?,
			updated_at = ? where uuid4= ?;`

// selectPartyContactDetailsSQL - select PartyContactDetailsSQL query
const selectPartyContactDetailsSQL = `select 
  id,
  uuid4,
  party_id,
  name,
  email,
  phone,
  url,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from party_contact_details`

// CreateParty - Create Party
func (ps *PartyService) CreateParty(ctx context.Context, in *partyproto.CreatePartyRequest) (*partyproto.CreatePartyResponse, error) {
	user, err := GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ps.UserServiceClient)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	partyD := partyproto.PartyD{}
	partyD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	partyD.PartyName = in.PartyName

	partyD.TaxReference1 = in.TaxReference1
	partyD.TaxReference2 = in.TaxReference2
	partyD.PublicKey = in.PublicKey

	address, err := ps.processAddress(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	party := partyproto.Party{PartyD: &partyD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ps.insertParty(ctx, insertPartySQL, &party, address, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	partyResponse := partyproto.CreatePartyResponse{}
	partyResponse.Party = &party
	return &partyResponse, nil
}

// insertParty - Insert party details into database
func (ps *PartyService) insertParty(ctx context.Context, insertPartySQL string, party *partyproto.Party, address *commonproto.Address, userEmail string, requestID string) error {
	partyTmp, err := ps.crPartyStruct(ctx, party, userEmail, requestID)
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		addr, err := common.InsertAddress(ctx, tx, address, userEmail, requestID)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		party.PartyD.AddressId = addr.Id
		partyTmp.PartyD.AddressId = addr.Id

		res, err := tx.NamedExecContext(ctx, insertPartySQL, partyTmp)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		party.PartyD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(party.PartyD.Uuid4)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		party.PartyD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crPartyStruct - process Party details
func (ps *PartyService) crPartyStruct(ctx context.Context, party *partyproto.Party, userEmail string, requestID string) (*partystruct.Party, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(party.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(party.CrUpdTime.UpdatedAt)

	partyTmp := partystruct.Party{PartyD: party.PartyD, CrUpdUser: party.CrUpdUser, CrUpdTime: crUpdTime}

	return &partyTmp, nil
}

// GetParties - Get Parties
func (ps *PartyService) GetParties(ctx context.Context, in *partyproto.GetPartiesRequest) (*partyproto.GetPartiesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = ps.DBService.LimitSQLRows
	}
	query := "status_code = ?"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	parties := []*partyproto.Party{}

	nselectPartiesSQL := selectPartiesSQL + ` where ` + query

	rows, err := ps.DBService.DB.QueryxContext(ctx, nselectPartiesSQL, "active")
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		partyTmp := partystruct.Party{}
		err = rows.StructScan(&partyTmp)
		if err != nil {
			ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		party, err := ps.GetPartyStruct(ctx, &getRequest, partyTmp)
		if err != nil {
			ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		parties = append(parties, party)

	}

	partiesResponse := partyproto.GetPartiesResponse{}
	if len(parties) != 0 {
		next := parties[len(parties)-1].PartyD.Id
		next--
		nextc := common.EncodeCursor(next)
		partiesResponse = partyproto.GetPartiesResponse{Parties: parties, NextCursor: nextc}
	} else {
		partiesResponse = partyproto.GetPartiesResponse{Parties: parties, NextCursor: "0"}
	}
	return &partiesResponse, nil
}

// GetParty - Get Party
func (ps *PartyService) GetParty(ctx context.Context, inReq *partyproto.GetPartyRequest) (*partyproto.GetPartyResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectPartiesSQL := selectPartiesSQL + ` where uuid4 = ? and status_code = ?;`
	row := ps.DBService.DB.QueryRowxContext(ctx, nselectPartiesSQL, uuid4byte, "active")

	partyTmp := partystruct.Party{}
	err = row.StructScan(&partyTmp)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	party, err := ps.GetPartyStruct(ctx, in, partyTmp)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	partyResponse := partyproto.GetPartyResponse{}
	partyResponse.Party = party
	return &partyResponse, nil
}

// GetPartyByPk - Get Party By Primary key(Id)
func (ps *PartyService) GetPartyByPk(ctx context.Context, inReq *partyproto.GetPartyByPkRequest) (*partyproto.GetPartyByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectPartiesSQL := selectPartiesSQL + ` where id = ? and status_code = ?;`
	row := ps.DBService.DB.QueryRowxContext(ctx, nselectPartiesSQL, in.Id, "active")

	partyTmp := partystruct.Party{}
	err := row.StructScan(&partyTmp)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	party, err := ps.GetPartyStruct(ctx, &getRequest, partyTmp)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	partyResponse := partyproto.GetPartyByPkResponse{}
	partyResponse.Party = party
	return &partyResponse, nil
}

// GetPartyStruct - Get party
func (ps *PartyService) GetPartyStruct(ctx context.Context, in *commonproto.GetRequest, partyTmp partystruct.Party) (*partyproto.Party, error) {
	uuid4Str, err := common.UUIDBytesToStr(partyTmp.PartyD.Uuid4)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	partyTmp.PartyD.IdS = uuid4Str

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(partyTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(partyTmp.CrUpdTime.UpdatedAt)

	party := partyproto.Party{PartyD: partyTmp.PartyD, CrUpdUser: partyTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &party, nil
}

// UpdateParty - Update party
func (ps *PartyService) UpdateParty(ctx context.Context, in *partyproto.UpdatePartyRequest) (*partyproto.UpdatePartyResponse, error) {
	getRequest := commonproto.GetRequest{}
	getRequest.Id = in.Id
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	form := partyproto.GetPartyRequest{}
	form.GetRequest = &getRequest

	partyResponse, err := ps.GetParty(ctx, &form)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	party := partyResponse.Party

	db := ps.DBService.DB
	tn := common.GetTimeDetails()
	stmt, err := db.PreparexContext(ctx, updatePartySQL)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.PartyName,
			tn,
			party.PartyD.Id,
			"active")
		if err != nil {
			ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &partyproto.UpdatePartyResponse{}, nil
}

// DeleteParty - Delete party
func (ps *PartyService) DeleteParty(ctx context.Context, inReq *partyproto.DeletePartyRequest) (*partyproto.DeletePartyResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	db := ps.DBService.DB
	tn := common.GetTimeDetails()
	stmt, err := db.PreparexContext(ctx, deletePartySQL)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			"inactive",
			tn,
			uuid4byte)

		if err != nil {
			ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err != nil {
				ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &partyproto.DeletePartyResponse{}, nil
}

// CreatePartyContactDetail - Create PartyContactDetail
func (ps *PartyService) CreatePartyContactDetail(ctx context.Context, in *partyproto.CreatePartyContactDetailRequest) (*partyproto.CreatePartyContactDetailResponse, error) {
	user, err := GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ps.UserServiceClient)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getByIDRequest := commonproto.GetByIdRequest{}
	getByIDRequest.Id = in.PartyId
	getByIDRequest.UserEmail = in.UserEmail
	getByIDRequest.RequestId = in.RequestId
	form := partyproto.GetPartyByPkRequest{}
	form.GetByIdRequest = &getByIDRequest

	partyResponse, err := ps.GetPartyByPk(ctx, &form)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	party := partyResponse.Party

	partyContactDetail, err := ps.createPartyContactDetail(ctx, insertPartyContactDetailSQL, in, user.Id, party.PartyD.Id, party.PartyD.PartyName, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	partyContactDetailResponse := partyproto.CreatePartyContactDetailResponse{}
	partyContactDetailResponse.PartyContactDetail = partyContactDetail
	return &partyContactDetailResponse, nil
}

// create PartyContactDetail - create PartyContactDetail
func (ps *PartyService) createPartyContactDetail(ctx context.Context, insertPartyContactDetailSQL string, form *partyproto.CreatePartyContactDetailRequest, userID string, partyID uint32, partyName string, userEmail string, requestID string) (*partyproto.PartyContactDetail, error) {
	var err error
	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	partyContactDetailD := partyproto.PartyContactDetailD{}
	partyContactDetailD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return nil, err
	}
	partyContactDetailD.PartyId = form.PartyId
	partyContactDetailD.Name = form.Name
	partyContactDetailD.Email = form.Email
	partyContactDetailD.Phone = form.Phone
	partyContactDetailD.Url = form.Url

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = userID
	crUpdUser.UpdatedByUserId = userID

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	partyContactDetail := partyproto.PartyContactDetail{PartyContactDetailD: &partyContactDetailD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ps.insertPartyContactDetail(ctx, insertPartyContactDetailSQL, &partyContactDetail, userEmail, requestID)
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return nil, err
	}
	return &partyContactDetail, nil
}

// insertPartyContactDetail - Insert partyContactDetail details into database
func (ps *PartyService) insertPartyContactDetail(ctx context.Context, insertPartyContactDetailSQL string, partyContactDetail *partyproto.PartyContactDetail, userEmail string, requestID string) error {
	partyContactDetailTmp, err := ps.crPartyContactDetailStruct(ctx, partyContactDetail, userEmail, requestID)
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertPartyContactDetailSQL, partyContactDetailTmp)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		partyContactDetail.PartyContactDetailD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(partyContactDetail.PartyContactDetailD.Uuid4)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		partyContactDetail.PartyContactDetailD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crPartyContactDetailStruct - process Party details
func (ps *PartyService) crPartyContactDetailStruct(ctx context.Context, partyContactDetail *partyproto.PartyContactDetail, userEmail string, requestID string) (*partystruct.PartyContactDetail, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(partyContactDetail.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(partyContactDetail.CrUpdTime.UpdatedAt)

	partyContactDetailTmp := partystruct.PartyContactDetail{PartyContactDetailD: partyContactDetail.PartyContactDetailD, CrUpdUser: partyContactDetail.CrUpdUser, CrUpdTime: crUpdTime}

	return &partyContactDetailTmp, nil
}

// GetPartyContactDetail - Get Party Contact
func (ps *PartyService) GetPartyContactDetail(ctx context.Context, inReq *partyproto.GetPartyContactDetailRequest) (*partyproto.GetPartyContactDetailResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectPartyContactDetailsSQL := selectPartyContactDetailsSQL + ` where uuid4 = ? and status_code = ?;`
	row := ps.DBService.DB.QueryRowxContext(ctx, nselectPartyContactDetailsSQL, uuid4byte, "active")

	partyContactDetailTmp := partystruct.PartyContactDetail{}
	err = row.StructScan(&partyContactDetailTmp)

	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	partyContactDetail, err := ps.GetPartyContactDetailStruct(ctx, &getRequest, partyContactDetailTmp)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	partyContactDetailResponse := partyproto.GetPartyContactDetailResponse{}
	partyContactDetailResponse.PartyContactDetail = partyContactDetail
	return &partyContactDetailResponse, nil
}

// GetPartyContactDetailStruct - Get partyContactDetail
func (ps *PartyService) GetPartyContactDetailStruct(ctx context.Context, in *commonproto.GetRequest, partyContactDetailTmp partystruct.PartyContactDetail) (*partyproto.PartyContactDetail, error) {
	uuid4Str, err := common.UUIDBytesToStr(partyContactDetailTmp.PartyContactDetailD.Uuid4)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	partyContactDetailTmp.PartyContactDetailD.IdS = uuid4Str

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(partyContactDetailTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(partyContactDetailTmp.CrUpdTime.UpdatedAt)

	partyContactDetail := partyproto.PartyContactDetail{PartyContactDetailD: partyContactDetailTmp.PartyContactDetailD, CrUpdUser: partyContactDetailTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &partyContactDetail, nil
}

// UpdatePartyContactDetail - Update User In Parties
func (ps *PartyService) UpdatePartyContactDetail(ctx context.Context, in *partyproto.UpdatePartyContactDetailRequest) (*partyproto.UpdatePartyContactDetailResponse, error) {
	getRequest := commonproto.GetRequest{}
	getRequest.Id = in.PartyContactDetailId
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	form := partyproto.GetPartyContactDetailRequest{}
	form.GetRequest = &getRequest

	partyContactDetailResponse, err := ps.GetPartyContactDetail(ctx, &form)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	partyContactDetail := partyContactDetailResponse.PartyContactDetail

	db := ps.DBService.DB
	tn := common.GetTimeDetails()
	stmt, err := db.PreparexContext(ctx, updatePartyContactDetailSQL)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	err = ps.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.Name,
			tn,
			partyContactDetail.PartyContactDetailD.Id,
			"active")
		if err != nil {
			ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	return &partyproto.UpdatePartyContactDetailResponse{}, nil
}

// DeletePartyContactDetail - Delete party contact
func (ps *PartyService) DeletePartyContactDetail(ctx context.Context, inReq *partyproto.DeletePartyContactDetailRequest) (*partyproto.DeletePartyContactDetailResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	db := ps.DBService.DB
	tn := common.GetTimeDetails()
	stmt, err := db.PreparexContext(ctx, deletePartyContactDetailSQL)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			"inactive",
			tn,
			uuid4byte)

		if err != nil {
			ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}

		return nil
	})

	err = stmt.Close()
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	return &partyproto.DeletePartyContactDetailResponse{}, nil
}

// processAddress - process Address
func (ps *PartyService) processAddress(ctx context.Context, in *partyproto.CreatePartyRequest) (*commonproto.Address, error) {
	var err error
	address := commonproto.Address{}
	address.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	/*  Addr  */
	address.Name1 = in.Name1
	address.Name1 = in.Name1
	address.Street = in.Street
	address.StreetNumber = in.StreetNumber
	address.Floor1 = in.Floor1
	address.PostalCode = in.PostalCode
	address.City = in.City
	address.StateRegion = in.StateRegion
	address.CountryName = in.CountryName
	return &address, nil
}
