package partyservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// LocationService - For accessing Location services
type LocationService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	partyproto.UnimplementedLocationServiceServer
}

// NewLocationService - Create Location service
func NewLocationService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *LocationService {
	return &LocationService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertLocationSQL - insert LocationSQL query
const insertLocationSQL = `insert into locations
	  ( 
  uuid4,
  location_name,
  latitude,
  longitude,
  un_location_code,
  address_id,
  facility_id
  )
  values (
:uuid4,
:location_name,
:latitude,
:longitude,
:un_location_code,
:address_id,
:facility_id);`

// selectLocationsSQL - select LocationsSQL query
const selectLocationsSQL = `select 
  id,
  uuid4,
  location_name,
  latitude,
  longitude,
  un_location_code,
  address_id,
  facility_id from locations`

// CreateLocation - CreateLocation
func (ls *LocationService) CreateLocation(ctx context.Context, in *partyproto.CreateLocationRequest) (*partyproto.CreateLocationResponse, error) {
	location, err := ls.ProcessLocationRequest(ctx, in)
	if err != nil {
		ls.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ls.insertLocation(ctx, insertLocationSQL, location, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ls.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	locationResponse := partyproto.CreateLocationResponse{}
	locationResponse.Location = location
	return &locationResponse, nil
}

// ProcessLocationRequest - Process LocationRequest
func (ls *LocationService) ProcessLocationRequest(ctx context.Context, in *partyproto.CreateLocationRequest) (*partyproto.Location, error) {
	var err error
	location := partyproto.Location{}
	location.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ls.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	location.LocationName = in.LocationName
	location.Latitude = in.Latitude
	location.Longitude = in.Longitude
	location.UnLocationCode = in.UnLocationCode
	location.AddressId = in.AddressId
	location.FacilityId = in.FacilityId
	return &location, nil
}

// insertLocation - Insert Document Party
func (ls *LocationService) insertLocation(ctx context.Context, insertLocationSQL string, location *partyproto.Location, userEmail string, requestID string) error {
	err := ls.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertLocationSQL, location)
		if err != nil {
			ls.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ls.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		location.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(location.Uuid4)
		if err != nil {
			ls.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		location.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ls.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// LoadLocations - Get Locations
func (ls *LocationService) LoadLocations(ctx context.Context, in *partyproto.LoadLocationsRequest) (*partyproto.LoadLocationsResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = ls.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	locations := []*partyproto.Location{}

	nselectLocationsSQL := selectLocationsSQL + query

	rows, err := ls.DBService.DB.QueryxContext(ctx, nselectLocationsSQL)
	if err != nil {
		ls.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		location := partyproto.Location{}
		err = rows.StructScan(&location)
		if err != nil {
			ls.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		locations = append(locations, &location)

	}

	locationsResponse := partyproto.LoadLocationsResponse{}
	if len(locations) != 0 {
		next := locations[len(locations)-1].Id
		next--
		nextc := common.EncodeCursor(next)
		locationsResponse = partyproto.LoadLocationsResponse{Locations: locations, NextCursor: nextc}
	} else {
		locationsResponse = partyproto.LoadLocationsResponse{Locations: locations, NextCursor: "0"}
	}
	return &locationsResponse, nil
}

// FetchLocationByID - FetchLocationByID
func (ls *LocationService) FetchLocationByID(ctx context.Context, inReq *partyproto.FetchLocationByIDRequest) (*partyproto.FetchLocationByIDResponse, error) {
	in := inReq.GetByIdRequest
	nselectLocationsSQL := selectLocationsSQL + ` where id = ?;`
	row := ls.DBService.DB.QueryRowxContext(ctx, nselectLocationsSQL, in.Id)
	location := partyproto.Location{}
	err := row.StructScan(&location)
	if err != nil {
		ls.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	locationResponse := partyproto.FetchLocationByIDResponse{}
	locationResponse.Location = &location
	return &locationResponse, nil
}
