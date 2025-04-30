package eblservices

import (
	"context"
	"net"
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eblstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ebl/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ConsignmentItemService - For accessing ConsignmentItem services
type ConsignmentItemService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eblproto.UnimplementedConsignmentItemServiceServer
}

// NewConsignmentItemService - Create ConsignmentItem service
func NewConsignmentItemService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *ConsignmentItemService {
	return &ConsignmentItemService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// InsertConsignmentItemSQL - Insert ConsignmentItemSQL Query
const InsertConsignmentItemSQL = `insert into consignment_items
	  ( 
  uuid4,
  description_of_goods,
  hs_code,
  shipping_instruction_id,
  weight,
  volume,
  weight_unit,
  volume_unit,
  shipment_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (:uuid4,
:description_of_goods,
:hs_code,
:shipping_instruction_id,
:weight,
:volume,
:weight_unit,
:volume_unit,
:shipment_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectConsignmentItemsSQL - select ConsignmentItemsSQL Query
const selectConsignmentItemsSQL = `select 
  id,
  uuid4,
  description_of_goods,
  hs_code,
  shipping_instruction_id,
  weight,
  volume,
  weight_unit,
  volume_unit,
  shipment_id,
  status_code,
  created_by_user_id,
  updated_by_user_id, 
  created_at,
  updated_at from consignment_items`

// StartEblServer - Start Ebl server
func StartEblServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
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
	consignmentItemService := NewConsignmentItemService(log, dbService, redisService, uc)
	shippingService := NewShippingService(log, dbService, redisService, uc)
	shippingInstructionService := NewShippingInstructionService(log, dbService, redisService, uc)
	transportDocumentService := NewTransportDocumentService(log, dbService, redisService, uc)
	surrenderRequestService := NewSurrenderRequestService(log, dbService, redisService, uc)
	surrenderRequestAnswerService := NewSurrenderRequestAnswerService(log, dbService, redisService, uc)
	issueRequestService := NewIssueRequestService(log, dbService, redisService, uc)
	issueRequestResponseService := NewIssueRequestResponseService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcEblServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	eblproto.RegisterConsignmentItemServiceServer(srv, consignmentItemService)
	eblproto.RegisterShippingServiceServer(srv, shippingService)
	eblproto.RegisterShippingInstructionServiceServer(srv, shippingInstructionService)
	eblproto.RegisterTransportDocumentServiceServer(srv, transportDocumentService)
	eblproto.RegisterSurrenderRequestServiceServer(srv, surrenderRequestService)
	eblproto.RegisterSurrenderRequestAnswerServiceServer(srv, surrenderRequestAnswerService)
	eblproto.RegisterIssueRequestServiceServer(srv, issueRequestService)
	eblproto.RegisterIssueRequestResponseServiceServer(srv, issueRequestResponseService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

// CreateConsignmentItem - CreateConsignmentItem
func (cis *ConsignmentItemService) CreateConsignmentItem(ctx context.Context, in *eblproto.CreateConsignmentItemRequest) (*eblproto.CreateConsignmentItemResponse, error) {
	consignmentItem, err := cis.ProcessConsignmentItemRequest(ctx, in)
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = cis.insertConsignmentItem(ctx, InsertConsignmentItemSQL, consignmentItem, InsertCargoItemSQL, InsertCargoLineItemSQL, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	consItemResponse := eblproto.CreateConsignmentItemResponse{}
	consItemResponse.ConsignmentItem = consignmentItem
	return &consItemResponse, nil
}

// CreateConsignmentItemsByShippingInstructionIdAndTOs - CreateConsignmentItemsByShippingInstructionIdAndTOs
func (cis *ConsignmentItemService) CreateConsignmentItemsByShippingInstructionIdAndTOs(ctx context.Context, inReq *eblproto.CreateConsignmentItemsByShippingInstructionIdAndTOsRequest) (*eblproto.CreateConsignmentItemsByShippingInstructionIdAndTOsResponse, error) {
	in := inReq.CreateConsignmentItemRequest
	consignmentItem, err := cis.ProcessConsignmentItemRequestFromShippingInstructionID(ctx, in)
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = cis.insertConsignmentItem(ctx, InsertConsignmentItemSQL, consignmentItem, InsertCargoItemSQL, InsertCargoLineItemSQL, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	consItemResponse := eblproto.CreateConsignmentItemsByShippingInstructionIdAndTOsResponse{}
	consItemResponse.ConsignmentItem = consignmentItem
	return &consItemResponse, nil
}

// ProcessConsignmentItemRequest - ProcessConsignmentItemRequest
func (cis *ConsignmentItemService) ProcessConsignmentItemRequest(ctx context.Context, in *eblproto.CreateConsignmentItemRequest) (*eblproto.ConsignmentItem, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, cis.UserServiceClient)
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	consignmentItemD := eblproto.ConsignmentItemD{}
	consignmentItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	consignmentItemD.DescriptionOfGoods = in.DescriptionOfGoods
	consignmentItemD.HsCode = in.HsCode
	consignmentItemD.ShippingInstructionId = in.ShippingInstructionId
	consignmentItemD.Weight = in.Weight
	consignmentItemD.Volume = in.Volume
	consignmentItemD.WeightUnit = in.WeightUnit
	consignmentItemD.VolumeUnit = in.VolumeUnit
	consignmentItemD.ShipmentId = in.ShipmentId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	consignmentItem := eblproto.ConsignmentItem{ConsignmentItemD: &consignmentItemD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	citems := []*eblproto.CargoItem{}
	// we will do for loop on references, UtilizedTransportEquipments, which is comes from client form
	for _, citem := range in.CargoItems {
		citem.UserId = in.UserId
		citem.UserEmail = in.UserEmail
		citem.RequestId = in.RequestId
		citem, err := cis.ProcessCargoItemRequest(ctx, citem)
		if err != nil {
			cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		citems = append(citems, citem)
	}
	consignmentItem.CargoItems = citems

	return &consignmentItem, nil
}

// insertConsignmentItem - Insert Consignment Item
func (cis *ConsignmentItemService) insertConsignmentItem(ctx context.Context, insertConsignmentItemSQL string, consignmentItem *eblproto.ConsignmentItem, insertCargoItemSQL string, insertCargoLineItemSQL string, userEmail string, requestID string) error {
	consignmentItemTmp, err := cis.CrConsignmentItemStruct(ctx, consignmentItem, userEmail, requestID)
	if err != nil {
		cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = cis.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertConsignmentItemSQL, consignmentItemTmp)
		if err != nil {
			cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		consignmentItem.ConsignmentItemD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(consignmentItem.ConsignmentItemD.Uuid4)
		if err != nil {
			cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		consignmentItem.ConsignmentItemD.IdS = uuid4Str
		for _, cargoItem := range consignmentItem.CargoItems {
			cargoItem.CargoItemD.ConsignmentItemId = consignmentItem.ConsignmentItemD.Id
			resp1, err := tx.NamedExecContext(ctx, insertCargoItemSQL, cargoItem)
			if err != nil {
				cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			uID1, err := resp1.LastInsertId()
			if err != nil {
				cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}

			for _, cargoLineItem := range cargoItem.CargoLineItems {
				cargoLineItem.CargoLineItemD.CargoItemId = uint32(uID1)
				_, err = tx.NamedExecContext(ctx, insertCargoLineItemSQL, cargoLineItem)
				if err != nil {
					cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
					return err
				}
			}

		}

		return nil
	})

	if err != nil {
		cis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrConsignmentItemStruct - process ConsignmentItem details
func (cis *ConsignmentItemService) CrConsignmentItemStruct(ctx context.Context, consignmentItem *eblproto.ConsignmentItem, userEmail string, requestID string) (*eblstruct.ConsignmentItem, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(consignmentItem.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(consignmentItem.CrUpdTime.UpdatedAt)
	consignmentItemTmp := eblstruct.ConsignmentItem{ConsignmentItemD: consignmentItem.ConsignmentItemD, CrUpdUser: consignmentItem.CrUpdUser, CrUpdTime: crUpdTime}

	return &consignmentItemTmp, nil
}

// FetchConsignmentItemsTOByShippingInstructionId - Get ConsignmentItemsTOByShippingInstructionID
func (cis *ConsignmentItemService) FetchConsignmentItemsTOByShippingInstructionId(ctx context.Context, in *eblproto.FetchConsignmentItemsTOByShippingInstructionIdRequest) (*eblproto.FetchConsignmentItemsTOByShippingInstructionIdResponse, error) {
	nselectConsignmentItemsSQL := selectConsignmentItemsSQL + ` where shipping_instruction_id = ?;`
	row := cis.DBService.DB.QueryRowxContext(ctx, nselectConsignmentItemsSQL, in.ShippingInstructionId)
	consignmentItemTmp := eblstruct.ConsignmentItem{}
	err := row.StructScan(&consignmentItemTmp)
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	consignmentItem, err := cis.getConsignmentItemStruct(ctx, &getRequest, consignmentItemTmp)
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	consItemResponse := eblproto.FetchConsignmentItemsTOByShippingInstructionIdResponse{}
	consItemResponse.ConsignmentItem = consignmentItem

	return &consItemResponse, nil
}

// getConsignmentItemStruct - Get consignmentItem
func (cis *ConsignmentItemService) getConsignmentItemStruct(ctx context.Context, in *commonproto.GetRequest, consignmentItemTmp eblstruct.ConsignmentItem) (*eblproto.ConsignmentItem, error) {
	uuid4Str, err := common.UUIDBytesToStr(consignmentItemTmp.ConsignmentItemD.Uuid4)
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(consignmentItemTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(consignmentItemTmp.CrUpdTime.UpdatedAt)

	consignmentItemTmp.ConsignmentItemD.IdS = uuid4Str

	consignmentItem := eblproto.ConsignmentItem{ConsignmentItemD: consignmentItemTmp.ConsignmentItemD, CrUpdUser: consignmentItemTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &consignmentItem, nil
}

// RemoveConsignmentItemsByShippingInstructionId - RemoveConsignmentItemsByShippingInstructionId
func (cis *ConsignmentItemService) RemoveConsignmentItemsByShippingInstructionId(ctx context.Context, in *eblproto.RemoveConsignmentItemsByShippingInstructionIdRequest) (*eblproto.RemoveConsignmentItemsByShippingInstructionIdResponse, error) {
	return &eblproto.RemoveConsignmentItemsByShippingInstructionIdResponse{}, nil
}

// ProcessConsignmentItemRequestFromShippingInstructionID - ProcessConsignmentItemRequestFromShippingInstructionID
func (cis *ConsignmentItemService) ProcessConsignmentItemRequestFromShippingInstructionID(ctx context.Context, in *eblproto.CreateConsignmentItemRequest) (*eblproto.ConsignmentItem, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, cis.UserServiceClient)
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	eblServ := &ShippingInstructionService{DBService: cis.DBService, RedisService: cis.RedisService, UserServiceClient: cis.UserServiceClient}
	inreq1 := eblproto.GetShippingInstructionByPkRequest{}
	req1 := commonproto.GetByIdRequest{}
	req1.Id = in.ShippingInstructionId
	req1.UserEmail = in.UserEmail
	req1.RequestId = in.RequestId
	inreq1.GetByIdRequest = &req1

	shippingInstructionResp, err := eblServ.GetShippingInstructionByPk(ctx, &inreq1)
	shippingInstruction := shippingInstructionResp.ShippingInstruction
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	consignmentItemD := eblproto.ConsignmentItemD{}
	consignmentItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	consignmentItemD.DescriptionOfGoods = in.DescriptionOfGoods
	consignmentItemD.HsCode = in.HsCode
	consignmentItemD.ShippingInstructionId = shippingInstruction.ShippingInstructionD.Id
	consignmentItemD.Weight = in.Weight
	consignmentItemD.Volume = in.Volume
	consignmentItemD.WeightUnit = in.WeightUnit
	consignmentItemD.VolumeUnit = in.VolumeUnit
	consignmentItemD.ShipmentId = in.ShipmentId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	consignmentItem := eblproto.ConsignmentItem{ConsignmentItemD: &consignmentItemD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	citems := []*eblproto.CargoItem{}
	// we will do for loop on references, UtilizedTransportEquipments, which is comes from client form
	for _, citem := range in.CargoItems {
		citem.UserId = in.UserId
		citem.UserEmail = in.UserEmail
		citem.RequestId = in.RequestId
		citem, err := cis.ProcessCargoItemRequest(ctx, citem)
		if err != nil {
			cis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		citems = append(citems, citem)
	}
	consignmentItem.CargoItems = citems

	return &consignmentItem, nil
}
