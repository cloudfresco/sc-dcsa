package common

import (
	"context"

	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertAddressSQL = `insert into addresses
	  ( 
    uuid4,
    name1,
    street,
    street_number,
    floor1,
    postal_code,
    city,
    state_region,
    country_name)
  values (:uuid4,
    :name1,
    :street,
    :street_number,
    :floor1,
    :postal_code,
    :city,
    :state_region,
    :country_name);`

// CreateAddress - For Creating Address
func CreateAddress(ctx context.Context, in *commonproto.Address, userEmail string, requestID string) (*commonproto.Address, error) {
	var err error
	addr := commonproto.Address{}
	addr.Uuid4, err = GetUUIDBytes()
	if err != nil {
		log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return nil, err
	}
	addr.Name1 = in.Name1
	addr.Street = in.Street
	addr.StreetNumber = in.StreetNumber
	addr.Floor1 = in.Floor1
	addr.PostalCode = in.PostalCode
	addr.City = in.City
	addr.StateRegion = in.StateRegion
	addr.CountryName = in.CountryName

	return &addr, nil
}

// InsertAddress - For Inserting Address
func InsertAddress(ctx context.Context, tx *sqlx.Tx, addr *commonproto.Address, userEmail string, requestID string) (*commonproto.Address, error) {
	res, err := tx.NamedExecContext(ctx, insertAddressSQL, addr)
	if err != nil {
		log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return nil, err
	}

	uID, err := res.LastInsertId()
	if err != nil {
		log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return nil, err
	}
	addr.Id = uint32(uID)
	uuid4Str, err := UUIDBytesToStr(addr.Uuid4)
	if err != nil {
		log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return nil, err
	}
	addr.IdS = uuid4Str
	return addr, nil
}
