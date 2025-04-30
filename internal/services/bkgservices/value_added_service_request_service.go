package bkgservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// insertValueAddedServiceRequestSQL - insert ValueAddedServiceRequestSQL query
const insertValueAddedServiceRequestSQL = `insert into value_added_service_requests
	  (
uuid4,
booking_id,
value_added_service_code)
  values (
:uuid4,
:booking_id,
:value_added_service_code);`

// selectValueAddedServiceRequestsSQL - select ValueAddedServiceRequestsSQL query
/*const selectValueAddedServiceRequestsSQL = `select
  id,
  uuid4,
  booking_id,
  value_added_service_code from value_added_service_requests`*/

// CreateValueAddedService - CreateValueAddedService
func (bs *BkgService) CreateValueAddedService(ctx context.Context, in *bkgproto.CreateValueAddedServiceRequest) (*bkgproto.CreateValueAddedServiceResponse, error) {
	valueAddedServiceRequest, err := bs.ProcessValueAddedServiceRequest(ctx, in)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = bs.insertValueAddedServiceRequest(ctx, insertValueAddedServiceRequestSQL, valueAddedServiceRequest, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	valueAddedServiceResponse := bkgproto.CreateValueAddedServiceResponse{}
	valueAddedServiceResponse.ValueAddedServiceRequest = valueAddedServiceRequest
	return &valueAddedServiceResponse, nil
}

// ProcessValueAddedServiceRequest - ProcessValueAddedServiceRequest
func (bs *BkgService) ProcessValueAddedServiceRequest(ctx context.Context, in *bkgproto.CreateValueAddedServiceRequest) (*bkgproto.ValueAddedServiceRequest, error) {
	var err error
	valueAddedServiceRequest := bkgproto.ValueAddedServiceRequest{}
	valueAddedServiceRequest.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	valueAddedServiceRequest.BookingId = in.BookingId
	valueAddedServiceRequest.ValueAddedServiceCode = in.ValueAddedServiceCode

	return &valueAddedServiceRequest, nil
}

// insertValueAddedService - Insert ValueAddedService
func (bs *BkgService) insertValueAddedServiceRequest(ctx context.Context, insertValueAddedServiceRequestSQL string, valueAddedServiceRequest *bkgproto.ValueAddedServiceRequest, userEmail string, requestID string) error {
	err := bs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertValueAddedServiceRequestSQL, valueAddedServiceRequest)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		valueAddedServiceRequest.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(valueAddedServiceRequest.Uuid4)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		valueAddedServiceRequest.IdS = uuid4Str
		return nil
	})
	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	return nil
}
