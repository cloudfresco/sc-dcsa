package eventcoreservices

import (
	"context"
	"net"
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eventcorestruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/eventcore/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ChargeService - For accessing Charge services
type ChargeService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eventcoreproto.UnimplementedChargeServiceServer
}

// NewChargeService - Create Charge service
func NewChargeService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *ChargeService {
	return &ChargeService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertChargeSQL - insert ChargeSQL query
const insertChargeSQL = `insert into charges
	  (
  uuid4,
  transport_document_id,
  shipment_id,
  charge_type,
  currency_amount,
  currency_code,
  payment_term_code,
  calculation_basis,
  unit_price,
  quantity,
  status_code,
  created_by_user_id,
  updated_by_user_id, 
  created_at,
  updated_at
  )
  values (:uuid4,
:transport_document_id,
:shipment_id,
:charge_type,
:currency_amount,
:currency_code,
:payment_term_code,
:calculation_basis,
:unit_price,
:quantity,
  :status_code,
  :created_by_user_id,
  :updated_by_user_id,
  :created_at,
  :updated_at);`

// selectChargesSQL - select ChargesSQL query
const selectChargesSQL = `select 
  id,
  uuid4,
  transport_document_id,
  shipment_id,
  charge_type,
  currency_amount,
  currency_code,
  payment_term_code,
  calculation_basis,
  unit_price,
  quantity,
  status_code,
  created_by_user_id,
  updated_by_user_id, 
  created_at,
  updated_at from charges`

// StartEventCoreServer - Start EventCore server
func StartEventCoreServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
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
	chargeService := NewChargeService(log, dbService, redisService, uc)
	transportCallService := NewTransportCallService(log, dbService, redisService, uc)
	utilizedTransportEquipmentService := NewUtilizedTransportEquipmentService(log, dbService, redisService, uc)
	voyageService := NewVoyageService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcEventCoreServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	eventcoreproto.RegisterChargeServiceServer(srv, chargeService)
	eventcoreproto.RegisterTransportCallServiceServer(srv, transportCallService)
	eventcoreproto.RegisterUtilizedTransportEquipmentServiceServer(srv, utilizedTransportEquipmentService)
	eventcoreproto.RegisterVoyageServiceServer(srv, voyageService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

// Create - Create Charge
func (cs *ChargeService) Create(ctx context.Context, in *eventcoreproto.CreateChargeRequest) (*eventcoreproto.CreateChargeResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, cs.UserServiceClient)
	if err != nil {
		cs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	chargeD := eventcoreproto.ChargeD{}
	chargeD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		cs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	chargeD.TransportDocumentId = in.TransportDocumentId
	chargeD.ShipmentId = in.ShipmentId
	chargeD.ChargeType = in.ChargeType
	chargeD.CurrencyAmount = in.CurrencyAmount
	chargeD.CurrencyCode = in.CurrencyCode
	chargeD.PaymentTermCode = in.PaymentTermCode
	chargeD.CalculationBasis = in.CalculationBasis
	chargeD.UnitPrice = in.UnitPrice
	chargeD.Quantity = in.Quantity

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	charge := eventcoreproto.Charge{ChargeD: &chargeD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = cs.insertCharge(ctx, insertChargeSQL, &charge, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		cs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	chargeResponse := eventcoreproto.CreateChargeResponse{}
	chargeResponse.Charge = &charge
	return &chargeResponse, nil
}

// insertCharge - Insert Charge
func (cs *ChargeService) insertCharge(ctx context.Context, insertChargeSQL string, charge *eventcoreproto.Charge, userEmail string, requestID string) error {
	chargeTmp, err := cs.crChargeStruct(ctx, charge, userEmail, requestID)
	if err != nil {
		cs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = cs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertChargeSQL, chargeTmp)
		if err != nil {
			cs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			cs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		charge.ChargeD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(charge.ChargeD.Uuid4)
		if err != nil {
			cs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		charge.ChargeD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		cs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crChargeStruct - process Charge details
func (cs *ChargeService) crChargeStruct(ctx context.Context, charge *eventcoreproto.Charge, userEmail string, requestID string) (*eventcorestruct.Charge, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(charge.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(charge.CrUpdTime.UpdatedAt)
	chargeTmp := eventcorestruct.Charge{ChargeD: charge.ChargeD, CrUpdUser: charge.CrUpdUser, CrUpdTime: crUpdTime}

	return &chargeTmp, nil
}

// FetchChargesByTransportDocumentId - FetchChargesByTransportDocumentId
func (cs *ChargeService) FetchChargesByTransportDocumentId(ctx context.Context, in *eventcoreproto.FetchChargesByTransportDocumentIdRequest) (*eventcoreproto.FetchChargesByTransportDocumentIdResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = cs.DBService.LimitSQLRows
	}
	query := "transport_document_id = ?"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	charges := []*eventcoreproto.Charge{}

	nselectChargesSQL := selectChargesSQL + ` where ` + query

	rows, err := cs.DBService.DB.QueryxContext(ctx, nselectChargesSQL, in.TransportDocumentId)
	if err != nil {
		cs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		chargeTmp := eventcorestruct.Charge{}
		err = rows.StructScan(&chargeTmp)
		if err != nil {
			cs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		charge, err := cs.getChargeStruct(ctx, &getRequest, chargeTmp)
		if err != nil {
			cs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		charges = append(charges, charge)

	}

	chargeResponse := eventcoreproto.FetchChargesByTransportDocumentIdResponse{}
	if len(charges) != 0 {
		next := charges[len(charges)-1].ChargeD.Id
		next--
		nextc := common.EncodeCursor(next)
		chargeResponse = eventcoreproto.FetchChargesByTransportDocumentIdResponse{Charges: charges, NextCursor: nextc}
	} else {
		chargeResponse = eventcoreproto.FetchChargesByTransportDocumentIdResponse{Charges: charges, NextCursor: "0"}
	}
	return &chargeResponse, nil
}

// FetchChargesByShipmentId - FetchChargesByShipmentId
func (cs *ChargeService) FetchChargesByShipmentId(ctx context.Context, in *eventcoreproto.FetchChargesByShipmentIdRequest) (*eventcoreproto.FetchChargesByShipmentIdResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = cs.DBService.LimitSQLRows
	}
	query := "shipment_id = ?"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	charges := []*eventcoreproto.Charge{}

	nselectChargesSQL := selectChargesSQL + ` where ` + query

	rows, err := cs.DBService.DB.QueryxContext(ctx, nselectChargesSQL, in.ShipmentId)
	if err != nil {
		cs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		chargeTmp := eventcorestruct.Charge{}
		err = rows.StructScan(&chargeTmp)
		if err != nil {
			cs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		charge, err := cs.getChargeStruct(ctx, &getRequest, chargeTmp)
		if err != nil {
			cs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		charges = append(charges, charge)

	}

	chargeResponse := eventcoreproto.FetchChargesByShipmentIdResponse{}
	if len(charges) != 0 {
		next := charges[len(charges)-1].ChargeD.Id
		next--
		nextc := common.EncodeCursor(next)
		chargeResponse = eventcoreproto.FetchChargesByShipmentIdResponse{Charges: charges, NextCursor: nextc}
	} else {
		chargeResponse = eventcoreproto.FetchChargesByShipmentIdResponse{Charges: charges, NextCursor: "0"}
	}
	return &chargeResponse, nil
}

// getChargeStruct - Get Charge struct
func (cs *ChargeService) getChargeStruct(ctx context.Context, in *commonproto.GetRequest, chargeTmp eventcorestruct.Charge) (*eventcoreproto.Charge, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(chargeTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(chargeTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(chargeTmp.ChargeD.Uuid4)
	if err != nil {
		cs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	chargeTmp.ChargeD.IdS = uuid4Str
	charge := eventcoreproto.Charge{ChargeD: chargeTmp.ChargeD, CrUpdUser: chargeTmp.CrUpdUser, CrUpdTime: crUpdTime}
	return &charge, nil
}
