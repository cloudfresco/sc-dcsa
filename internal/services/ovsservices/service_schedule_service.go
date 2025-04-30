package ovsservices

import (
	"context"
	"net"
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	ovsstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ovs/v3"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ServiceScheduleService - For accessing ServiceSchedule services
type ServiceScheduleService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	ovsproto.UnimplementedServiceScheduleServiceServer
}

// NewServiceScheduleService - Create ServiceSchedule service
func NewServiceScheduleService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *ServiceScheduleService {
	return &ServiceScheduleService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertServiceScheduleSQL - insert ServiceScheduleSQL Query
const insertServiceScheduleSQL = `insert into service_schedules
	  ( 
  uuid4,
  carrier_service_name,
  carrier_service_code,
  universal_service_reference,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:carrier_service_name,
:carrier_service_code,
:universal_service_reference,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectServiceSchedulesSQL - select ServiceSchedulesSQL Query
const selectServiceSchedulesSQL = `select 
  id,
  uuid4,
  carrier_service_name,
  carrier_service_code,
  universal_service_reference,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from service_schedules`

// updateServiceSchedulesSQL - update ServiceSchedulesSQL query
const updateServiceScheduleSQL = `update service_schedules set
	  carrier_service_name = ?,
    carrier_service_code = ?,
		updated_at = ? where universal_service_reference = ?;`

// StartOvsServer - Start Ovs server
func StartOvsServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
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
	serviceScheduleService := NewServiceScheduleService(log, dbService, redisService, uc)
	vesselScheduleService := NewVesselScheduleService(log, dbService, redisService, uc)
	legScheduleService := NewLegService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcOvsServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	ovsproto.RegisterServiceScheduleServiceServer(srv, serviceScheduleService)
	ovsproto.RegisterVesselScheduleServiceServer(srv, vesselScheduleService)
	ovsproto.RegisterLegServiceServer(srv, legScheduleService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

// CreateServiceSchedule - Create  ServiceSchedule
func (ss *ServiceScheduleService) CreateServiceSchedule(ctx context.Context, in *ovsproto.CreateServiceScheduleRequest) (*ovsproto.CreateServiceScheduleResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ss.UserServiceClient)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	serviceScheduleD := ovsproto.ServiceScheduleD{}
	serviceScheduleD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	serviceScheduleD.CarrierServiceName = in.CarrierServiceName
	serviceScheduleD.CarrierServiceCode = in.CarrierServiceCode
	serviceScheduleD.UniversalServiceReference = in.UniversalServiceReference

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	serviceSchedule := ovsproto.ServiceSchedule{ServiceScheduleD: &serviceScheduleD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ss.insertServiceSchedule(ctx, insertServiceScheduleSQL, &serviceSchedule, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	serviceScheduleResponse := ovsproto.CreateServiceScheduleResponse{}
	serviceScheduleResponse.ServiceSchedule = &serviceSchedule
	return &serviceScheduleResponse, nil
}

// insertServiceSchedule - Insert ServiceSchedule
func (ss *ServiceScheduleService) insertServiceSchedule(ctx context.Context, insertServiceScheduleSQL string, serviceSchedule *ovsproto.ServiceSchedule, userEmail string, requestID string) error {
	serviceScheduleTmp, err := ss.crServiceScheduleStruct(ctx, serviceSchedule, userEmail, requestID)
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertServiceScheduleSQL, serviceScheduleTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		serviceSchedule.ServiceScheduleD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(serviceSchedule.ServiceScheduleD.Uuid4)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		serviceSchedule.ServiceScheduleD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crServiceScheduleStruct - process ServiceSchedule details
func (ss *ServiceScheduleService) crServiceScheduleStruct(ctx context.Context, serviceSchedule *ovsproto.ServiceSchedule, userEmail string, requestID string) (*ovsstruct.ServiceSchedule, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(serviceSchedule.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(serviceSchedule.CrUpdTime.UpdatedAt)
	serviceScheduleTmp := ovsstruct.ServiceSchedule{ServiceScheduleD: serviceSchedule.ServiceScheduleD, CrUpdUser: serviceSchedule.CrUpdUser, CrUpdTime: crUpdTime}

	return &serviceScheduleTmp, nil
}

// GetServiceSchedules - Get  ServiceSchedules
func (ss *ServiceScheduleService) GetServiceSchedules(ctx context.Context, in *ovsproto.GetServiceSchedulesRequest) (*ovsproto.GetServiceSchedulesResponse, error) {
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

	serviceSchedules := []*ovsproto.ServiceSchedule{}

	nselectServiceSchedulesSQL := selectServiceSchedulesSQL + query

	rows, err := ss.DBService.DB.QueryxContext(ctx, nselectServiceSchedulesSQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		serviceScheduleTmp := ovsstruct.ServiceSchedule{}
		err = rows.StructScan(&serviceScheduleTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		serviceSchedule, err := ss.getServiceScheduleStruct(ctx, &getRequest, serviceScheduleTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		serviceSchedules = append(serviceSchedules, serviceSchedule)

	}

	serviceScheduleResponse := ovsproto.GetServiceSchedulesResponse{}
	if len(serviceSchedules) != 0 {
		next := serviceSchedules[len(serviceSchedules)-1].ServiceScheduleD.Id
		next--
		nextc := common.EncodeCursor(next)
		serviceScheduleResponse = ovsproto.GetServiceSchedulesResponse{ServiceSchedules: serviceSchedules, NextCursor: nextc}
	} else {
		serviceScheduleResponse = ovsproto.GetServiceSchedulesResponse{ServiceSchedules: serviceSchedules, NextCursor: "0"}
	}
	return &serviceScheduleResponse, nil
}

// GetServiceSchedule - Get ServiceSchedule
func (ss *ServiceScheduleService) GetServiceSchedule(ctx context.Context, inReq *ovsproto.GetServiceScheduleRequest) (*ovsproto.GetServiceScheduleResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectServiceSchedulesSQL := selectServiceSchedulesSQL + ` where uuid4 = ?`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectServiceSchedulesSQL, uuid4byte)
	serviceScheduleTmp := ovsstruct.ServiceSchedule{}
	err = row.StructScan(&serviceScheduleTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	serviceSchedule, err := ss.getServiceScheduleStruct(ctx, &getRequest, serviceScheduleTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	serviceScheduleResponse := ovsproto.GetServiceScheduleResponse{}
	serviceScheduleResponse.ServiceSchedule = serviceSchedule
	return &serviceScheduleResponse, nil
}

// GetServiceScheduleByUniversalServiceReference - Get ServiceScheduleByUniversalServiceReference
func (ss *ServiceScheduleService) GetServiceScheduleByUniversalServiceReference(ctx context.Context, in *ovsproto.GetServiceScheduleByUniversalServiceReferenceRequest) (*ovsproto.GetServiceScheduleByUniversalServiceReferenceResponse, error) {
	nselectServiceSchedulesSQL := selectServiceSchedulesSQL + ` where universal_service_reference = ? and status_code = ?`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectServiceSchedulesSQL, in.UniversalServiceReference, "active")
	serviceScheduleTmp := ovsstruct.ServiceSchedule{}
	err := row.StructScan(&serviceScheduleTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	serviceSchedule, err := ss.getServiceScheduleStruct(ctx, &getRequest, serviceScheduleTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	serviceScheduleResponse := ovsproto.GetServiceScheduleByUniversalServiceReferenceResponse{}
	serviceScheduleResponse.ServiceSchedule = serviceSchedule
	return &serviceScheduleResponse, nil
}

// GetServiceScheduleByPk - Get ServiceSchedule By Primary key(Id)
func (ss *ServiceScheduleService) GetServiceScheduleByPk(ctx context.Context, inReq *ovsproto.GetServiceScheduleByPkRequest) (*ovsproto.GetServiceScheduleByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectServiceSchedulesSQL := selectServiceSchedulesSQL + ` where id = ?;`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectServiceSchedulesSQL, in.Id)
	serviceScheduleTmp := ovsstruct.ServiceSchedule{}
	err := row.StructScan(&serviceScheduleTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	serviceSchedule, err := ss.getServiceScheduleStruct(ctx, &getRequest, serviceScheduleTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	serviceScheduleResponse := ovsproto.GetServiceScheduleByPkResponse{}
	serviceScheduleResponse.ServiceSchedule = serviceSchedule
	return &serviceScheduleResponse, nil
}

// UpdateServiceScheduleByUniversalServiceReference - UpdateServiceScheduleByUniversalServiceReference
func (ss *ServiceScheduleService) UpdateServiceScheduleByUniversalServiceReference(ctx context.Context, in *ovsproto.UpdateServiceScheduleByUniversalServiceReferenceRequest) (*ovsproto.UpdateServiceScheduleByUniversalServiceReferenceResponse, error) {
	db := ss.DBService.DB
	tn := common.GetTimeDetails()
	stmt, err := db.PreparexContext(ctx, updateServiceScheduleSQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ss.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.CarrierServiceName,
			in.CarrierServiceCode,
			tn,
			in.UniversalServiceReference)
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

	return &ovsproto.UpdateServiceScheduleByUniversalServiceReferenceResponse{}, nil
}

// getServiceScheduleStruct - Get serviceSchedule
func (ss *ServiceScheduleService) getServiceScheduleStruct(ctx context.Context, in *commonproto.GetRequest, serviceScheduleTmp ovsstruct.ServiceSchedule) (*ovsproto.ServiceSchedule, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(serviceScheduleTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(serviceScheduleTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(serviceScheduleTmp.ServiceScheduleD.Uuid4)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	serviceScheduleTmp.ServiceScheduleD.IdS = uuid4Str
	serviceSchedule := ovsproto.ServiceSchedule{ServiceScheduleD: serviceScheduleTmp.ServiceScheduleD, CrUpdUser: serviceScheduleTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &serviceSchedule, nil
}
