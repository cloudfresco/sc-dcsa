package jitservices

import (
	"context"
	"net"
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	jitstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/jit/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ServiceService - For accessing Service services
type ServiceService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	jitproto.UnimplementedServiceServiceServer
}

// NewServiceService - Create Service service
func NewServiceService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *ServiceService {
	return &ServiceService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertServiceSQL - insert ServiceSQL Query
const insertServiceSQL = `insert into services
	  ( 
  uuid4,
  carrier_id,
  carrier_service_name,
  carrier_service_code,
  tradelane_id,
  universal_service_reference,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:carrier_id,
:carrier_service_name,
:carrier_service_code,
:tradelane_id,
:universal_service_reference,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectServiceSQL - select ServiceSQL Query
const selectServicesSQL = `select 
  id,
  uuid4,
  carrier_id,
  carrier_service_name,
  carrier_service_code,
  tradelane_id,
  universal_service_reference,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from services`

// StartJitServer - Start Jit server
func StartJitServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
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
	serviceService := NewServiceService(log, dbService, redisService, uc)
	timestampService := NewTimestampService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcJitServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	jitproto.RegisterServiceServiceServer(srv, serviceService)
	jitproto.RegisterTimestampServiceServer(srv, timestampService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

// CreateService - Create  Service
func (ss *ServiceService) CreateService(ctx context.Context, in *jitproto.CreateServiceRequest) (*jitproto.CreateServiceResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ss.UserServiceClient)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	serviceD := jitproto.ServiceD{}
	serviceD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	serviceD.CarrierId = in.CarrierId
	serviceD.CarrierServiceName = in.CarrierServiceName
	serviceD.CarrierServiceCode = in.CarrierServiceCode
	serviceD.TradelaneId = in.TradelaneId
	serviceD.UniversalServiceReference = in.UniversalServiceReference

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	service := jitproto.Service{ServiceD: &serviceD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ss.insertService(ctx, insertServiceSQL, &service, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	serviceResponse := jitproto.CreateServiceResponse{}
	serviceResponse.Service1 = &service
	return &serviceResponse, nil
}

// insertService - Insert Service
func (ss *ServiceService) insertService(ctx context.Context, insertServiceSQL string, service *jitproto.Service, userEmail string, requestID string) error {
	serviceTmp, err := ss.crServiceStruct(ctx, service, userEmail, requestID)
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertServiceSQL, serviceTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		service.ServiceD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(service.ServiceD.Uuid4)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		service.ServiceD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crServiceStruct - process Servicedetails
func (ss *ServiceService) crServiceStruct(ctx context.Context, service *jitproto.Service, userEmail string, requestID string) (*jitstruct.Service, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(service.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(service.CrUpdTime.UpdatedAt)
	serviceTmp := jitstruct.Service{ServiceD: service.ServiceD, CrUpdUser: service.CrUpdUser, CrUpdTime: crUpdTime}

	return &serviceTmp, nil
}

// GetServices - Get  Services
func (ss *ServiceService) GetServices(ctx context.Context, in *jitproto.GetServicesRequest) (*jitproto.GetServicesResponse, error) {
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

	services := []*jitproto.Service{}

	nselectServicesSQL := selectServicesSQL + query

	rows, err := ss.DBService.DB.QueryxContext(ctx, nselectServicesSQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		serviceTmp := jitstruct.Service{}
		err = rows.StructScan(&serviceTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		service, err := ss.getServiceStruct(ctx, &getRequest, serviceTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		services = append(services, service)

	}

	serviceResponse := jitproto.GetServicesResponse{}
	if len(services) != 0 {
		next := services[len(services)-1].ServiceD.Id
		next--
		nextc := common.EncodeCursor(next)
		serviceResponse = jitproto.GetServicesResponse{Services: services, NextCursor: nextc}
	} else {
		serviceResponse = jitproto.GetServicesResponse{Services: services, NextCursor: "0"}
	}
	return &serviceResponse, nil
}

// GetService - Get Service
func (ss *ServiceService) GetService(ctx context.Context, inReq *jitproto.GetServiceRequest) (*jitproto.GetServiceResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectServicesSQL := selectServicesSQL + ` where uuid4 = ?`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectServicesSQL, uuid4byte)
	serviceTmp := jitstruct.Service{}
	err = row.StructScan(&serviceTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	service, err := ss.getServiceStruct(ctx, &getRequest, serviceTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	serviceResponse := jitproto.GetServiceResponse{}
	serviceResponse.Service1 = service
	return &serviceResponse, nil
}

// GetServiceByPk - Get Service By Primary key(Id)
func (ss *ServiceService) GetServiceByPk(ctx context.Context, inReq *jitproto.GetServiceByPkRequest) (*jitproto.GetServiceByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectServicesSQL := selectServicesSQL + ` where id = ?;`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectServicesSQL, in.Id)
	serviceTmp := jitstruct.Service{}
	err := row.StructScan(&serviceTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	service, err := ss.getServiceStruct(ctx, &getRequest, serviceTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	serviceResponse := jitproto.GetServiceByPkResponse{}
	serviceResponse.Service1 = service
	return &serviceResponse, nil
}

// FindByCarrierServiceCode -FindByCarrierServiceCode
func (ss *ServiceService) FindByCarrierServiceCode(ctx context.Context, in *jitproto.FindByCarrierServiceCodeRequest) (*jitproto.FindByCarrierServiceCodeResponse, error) {
	nselectServicesSQL := selectServicesSQL + ` where carrier_service_code = ?;`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectServicesSQL, in.CarrierServiceCode)
	serviceTmp := jitstruct.Service{}
	err := row.StructScan(&serviceTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	service, err := ss.getServiceStruct(ctx, &getRequest, serviceTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	serviceResponse := jitproto.FindByCarrierServiceCodeResponse{}
	serviceResponse.Service1 = service
	return &serviceResponse, nil
}

// getServiceStruct - Get service
func (ss *ServiceService) getServiceStruct(ctx context.Context, in *commonproto.GetRequest, serviceTmp jitstruct.Service) (*jitproto.Service, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(serviceTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(serviceTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(serviceTmp.ServiceD.Uuid4)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	serviceTmp.ServiceD.IdS = uuid4Str

	service := jitproto.Service{ServiceD: serviceTmp.ServiceD, CrUpdUser: serviceTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &service, nil
}
