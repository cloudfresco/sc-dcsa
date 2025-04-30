package eventcoreservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eventcorestruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/eventcore/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// VoyageService - For accessing Voyage services
type VoyageService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eventcoreproto.UnimplementedVoyageServiceServer
}

// NewVoyageService - Create Voyage service
func NewVoyageService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *VoyageService {
	return &VoyageService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertVoyageSQL - insert VoyageSQL query
const insertVoyageSQL = `insert into voyages
	  (
uuid4,
carrier_voyage_number,
universal_voyage_reference,
service_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:carrier_voyage_number,
:universal_voyage_reference,
:service_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectVoyagesSQL - select VoyagesSQL query
const selectVoyagesSQL = `select 
  id,
  uuid4,
  carrier_voyage_number,
  universal_voyage_reference,
  service_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from voyages`

// CreateVoyage - Create Voyage
func (vs *VoyageService) CreateVoyage(ctx context.Context, in *eventcoreproto.CreateVoyageRequest) (*eventcoreproto.CreateVoyageResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, vs.UserServiceClient)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	voyageD := eventcoreproto.VoyageD{}
	voyageD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	voyageD.CarrierVoyageNumber = in.CarrierVoyageNumber
	voyageD.UniversalVoyageReference = in.UniversalVoyageReference
	voyageD.ServiceId = in.ServiceId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	voyage := eventcoreproto.Voyage{VoyageD: &voyageD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = vs.insertVoyage(ctx, insertVoyageSQL, &voyage, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	voyageResponse := eventcoreproto.CreateVoyageResponse{}
	voyageResponse.Voyage = &voyage
	return &voyageResponse, nil
}

// insertVoyage - Insert Voyage
func (vs *VoyageService) insertVoyage(ctx context.Context, insertVoyageSQL string, voyage *eventcoreproto.Voyage, userEmail string, requestID string) error {
	voyageTmp, err := vs.crVoyageStruct(ctx, voyage, userEmail, requestID)
	if err != nil {
		vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = vs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertVoyageSQL, voyageTmp)
		if err != nil {
			vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		voyage.VoyageD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(voyage.VoyageD.Uuid4)
		if err != nil {
			vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		voyage.VoyageD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crVoyageStruct - process Voyage details
func (vs *VoyageService) crVoyageStruct(ctx context.Context, voyage *eventcoreproto.Voyage, userEmail string, requestID string) (*eventcorestruct.Voyage, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(voyage.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(voyage.CrUpdTime.UpdatedAt)
	voyageTmp := eventcorestruct.Voyage{VoyageD: voyage.VoyageD, CrUpdUser: voyage.CrUpdUser, CrUpdTime: crUpdTime}

	return &voyageTmp, nil
}

// GetVoyages - Get  Voyages
func (vs *VoyageService) GetVoyages(ctx context.Context, in *eventcoreproto.GetVoyagesRequest) (*eventcoreproto.GetVoyagesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = vs.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	voyages := []*eventcoreproto.Voyage{}

	nselectVoyagesSQL := selectVoyagesSQL + query

	rows, err := vs.DBService.DB.QueryxContext(ctx, nselectVoyagesSQL)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		voyageTmp := eventcorestruct.Voyage{}
		err = rows.StructScan(&voyageTmp)
		if err != nil {
			vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		voyage, err := vs.getVoyageStruct(ctx, &getRequest, voyageTmp)
		if err != nil {
			vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		voyages = append(voyages, voyage)

	}

	voyageResponse := eventcoreproto.GetVoyagesResponse{}
	if len(voyages) != 0 {
		next := voyages[len(voyages)-1].VoyageD.Id
		next--
		nextc := common.EncodeCursor(next)
		voyageResponse = eventcoreproto.GetVoyagesResponse{Voyages: voyages, NextCursor: nextc}
	} else {
		voyageResponse = eventcoreproto.GetVoyagesResponse{Voyages: voyages, NextCursor: "0"}
	}
	return &voyageResponse, nil
}

// FindByCarrierVoyageNumberAndServiceId - Find ByCarrierVoyageNumberAndServiceID
func (vs *VoyageService) FindByCarrierVoyageNumberAndServiceId(ctx context.Context, in *eventcoreproto.FindByCarrierVoyageNumberAndServiceIdRequest) (*eventcoreproto.FindByCarrierVoyageNumberAndServiceIdResponse, error) {
	nselectVoyagesSQL := selectVoyagesSQL + ` where carrier_voyage_number = ? and service_id = ?;`
	row := vs.DBService.DB.QueryRowxContext(ctx, nselectVoyagesSQL, in.CarrierVoyageNumber, in.ServiceId)
	voyageTmp := eventcorestruct.Voyage{}
	err := row.StructScan(&voyageTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	voyage, err := vs.getVoyageStruct(ctx, &getRequest, voyageTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	voyageResponse := eventcoreproto.FindByCarrierVoyageNumberAndServiceIdResponse{}
	voyageResponse.Voyage = voyage
	return &voyageResponse, nil
}

// FindByCarrierVoyageNumber - Find ByCarrierVoyageNumber
func (vs *VoyageService) FindByCarrierVoyageNumber(ctx context.Context, in *eventcoreproto.FindByCarrierVoyageNumberRequest) (*eventcoreproto.FindByCarrierVoyageNumberResponse, error) {
	nselectVoyagesSQL := selectVoyagesSQL + ` where carrier_voyage_number = ?;`
	row := vs.DBService.DB.QueryRowxContext(ctx, nselectVoyagesSQL, in.CarrierVoyageNumber)
	voyageTmp := eventcorestruct.Voyage{}
	err := row.StructScan(&voyageTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	voyage, err := vs.getVoyageStruct(ctx, &getRequest, voyageTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	voyageResponse := eventcoreproto.FindByCarrierVoyageNumberResponse{}
	voyageResponse.Voyage = voyage
	return &voyageResponse, nil
}

// FindCarrierVoyageNumbersByTransportDocumentRef - FindCarrierVoyageNumbersByTransportDocumentRef
func (vs *VoyageService) FindCarrierVoyageNumbersByTransportDocumentRef(ctx context.Context, in *eventcoreproto.FindCarrierVoyageNumbersByTransportDocumentRefRequest) (*eventcoreproto.FindCarrierVoyageNumbersByTransportDocumentRefResponse, error) {
	nselectVoyagesSQL := selectVoyagesSQL + `v inner join transport_call tc 
		 ON tc.import_voyage_id = v.id OR tc.export_voyage_id = v.id 
		inner join transport t 
		 ON t.load_transport_call_id = tc.id 
		inner join shipment_transport st 
		 ON st.transport_id = t.id 
		inner join utilized_transport_equipment ute
		 ON ute.shipment_id = st.shipment_id
		inner join cargo_item ci 
		 ON ci.utilized_transport_equipment_id = ute.id 
		inner join transport_document td 
		 ON td.shipping_instruction_id = ci.shipping_instruction_id 
		WHERE td.transport_document_reference = ? );`

	row := vs.DBService.DB.QueryRowxContext(ctx, nselectVoyagesSQL, in.TransportDocumentRef)
	voyageTmp := eventcorestruct.Voyage{}
	err := row.StructScan(&voyageTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	voyage, err := vs.getVoyageStruct(ctx, &getRequest, voyageTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	voyageResponse := eventcoreproto.FindCarrierVoyageNumbersByTransportDocumentRefResponse{}
	voyageResponse.Voyage = voyage
	return &voyageResponse, nil
}

// FindCarrierVoyageNumbersByCarrierBookingRef - FindCarrierVoyageNumbersByCarrierBookingRef
func (vs *VoyageService) FindCarrierVoyageNumbersByCarrierBookingRef(ctx context.Context, in *eventcoreproto.FindCarrierVoyageNumbersByCarrierBookingRefRequest) (*eventcoreproto.FindCarrierVoyageNumbersByCarrierBookingRefResponse, error) {
	nselectVoyagesSQL := selectVoyagesSQL + `v inner join transport_call tc
		ON tc.import_voyage_id = v.id OR tc.export_voyage_id = v.id
		inner join transport t
		ON t.load_transport_call_id = tc.id
		inner join shipment_transport st
		ON st.transport_id = t.id
		inner join shipment s
		ON s.id = st.shipment_id
		WHERE s.carrier_booking_reference = ? );`

	row := vs.DBService.DB.QueryRowxContext(ctx, nselectVoyagesSQL, in.CarrierBookingRef)
	voyageTmp := eventcorestruct.Voyage{}
	err := row.StructScan(&voyageTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	voyage, err := vs.getVoyageStruct(ctx, &getRequest, voyageTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	voyageResponse := eventcoreproto.FindCarrierVoyageNumbersByCarrierBookingRefResponse{}
	voyageResponse.Voyage = voyage
	return &voyageResponse, nil
}

// FindCarrierVoyageNumbersByShippingInstructionId - FindCarrierVoyageNumbersByShippingInstructionId
func (vs *VoyageService) FindCarrierVoyageNumbersByShippingInstructionId(ctx context.Context, in *eventcoreproto.FindCarrierVoyageNumbersByShippingInstructionIdRequest) (*eventcoreproto.FindCarrierVoyageNumbersByShippingInstructionIdResponse, error) {
	nselectVoyagesSQL := selectVoyagesSQL + `v inner join transport_call tc
		ON tc.import_voyage_id = v.id OR tc.export_voyage_id = v.id
		inner join transport t
		ON t.load_transport_call_id = tc.id
		inner join shipment_transport st
		ON st.transport_id = t.id
		inner join utilized_transport_equipment ute
		ON ute.shipment_id = st.shipment_id
		inner join cargo_item ci
		ON ci.utilized_transport_equipment_id = ute.id
		LEFT JOIN reference r
		ON r.shipment_id = st.shipment_id
		WHERE (ci.shipping_instruction_id = ? OR reference.shipping_instruction_id = ? ));`

	row := vs.DBService.DB.QueryRowxContext(ctx, nselectVoyagesSQL, in.ShippingInstructionId)
	voyageTmp := eventcorestruct.Voyage{}
	err := row.StructScan(&voyageTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	voyage, err := vs.getVoyageStruct(ctx, &getRequest, voyageTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	voyageResponse := eventcoreproto.FindCarrierVoyageNumbersByShippingInstructionIdResponse{}
	voyageResponse.Voyage = voyage
	return &voyageResponse, nil
}

// getVoyageStruct - Get Voyage struct
func (vs *VoyageService) getVoyageStruct(ctx context.Context, in *commonproto.GetRequest, voyageTmp eventcorestruct.Voyage) (*eventcoreproto.Voyage, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(voyageTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(voyageTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(voyageTmp.VoyageD.Uuid4)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	voyageTmp.VoyageD.IdS = uuid4Str
	voyage := eventcoreproto.Voyage{VoyageD: voyageTmp.VoyageD, CrUpdUser: voyageTmp.CrUpdUser, CrUpdTime: crUpdTime}
	return &voyage, nil
}
