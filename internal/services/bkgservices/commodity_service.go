// https://github.com/dcsaorg/DCSA-Edocumentation/blob/9bd5082561a0d1857439ea49b97e20716bbe2c92/edocumentation-service/src/main/java/org/dcsa/edocumentation/service/ReferenceService.java
package bkgservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	bkgstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/bkg/v2"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// insertCommoditySQL - insert CommoditySQL query
const insertCommoditySQL = `insert into commodities
	  (
uuid4,
booking_id,
commodity_type,
hs_code,
cargo_gross_weight,
cargo_gross_weight_unit,
cargo_gross_volume,
cargo_gross_volume_unit,
number_of_packages,
export_license_issue_date,
export_license_expiry_date,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (:uuid4,
:booking_id,
:commodity_type,
:hs_code,
:cargo_gross_weight,
:cargo_gross_weight_unit,
:cargo_gross_volume,
:cargo_gross_volume_unit,
:number_of_packages,
:export_license_issue_date,
:export_license_expiry_date,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectCommoditiesSQL - select CommoditiesSQL query
/*const selectCommoditiesSQL = `select
  id,
  uuid4,
  booking_id,
  commodity_type,
  hs_code,
  cargo_gross_weight,
  cargo_gross_weight_unit,
  cargo_gross_volume,
  cargo_gross_volume_unit,
  number_of_packages,
  export_license_issue_date,
  export_license_expiry_date,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from commodities`*/

// CreateCommodity - CreateCommodity
func (bs *BkgService) CreateCommodity(ctx context.Context, in *bkgproto.CreateCommodityRequest) (*bkgproto.CreateCommodityResponse, error) {
	commodity, err := bs.ProcessCommodityRequest(ctx, in)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = bs.insertCommodity(ctx, insertCommoditySQL, commodity, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	commodityResponse := bkgproto.CreateCommodityResponse{}
	commodityResponse.Commodity = commodity
	return &commodityResponse, nil
}

// ProcessCommodityRequest - ProcessCommodityRequest
func (bs *BkgService) ProcessCommodityRequest(ctx context.Context, in *bkgproto.CreateCommodityRequest) (*bkgproto.Commodity, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, bs.UserServiceClient)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	exportLicenseIssueDate, err := time.Parse(common.Layout, in.ExportLicenseIssueDate)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	exportLicenseExpiryDate, err := time.Parse(common.Layout, in.ExportLicenseExpiryDate)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	commodityD := bkgproto.CommodityD{}
	commodityD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	commodityD.BookingId = in.BookingId
	commodityD.CommodityType = in.CommodityType
	commodityD.HsCode = in.HsCode
	commodityD.CargoGrossWeight = in.CargoGrossWeight
	commodityD.CargoGrossWeightUnit = in.CargoGrossWeightUnit
	commodityD.CargoGrossVolume = in.CargoGrossVolume
	commodityD.CargoGrossVolumeUnit = in.CargoGrossVolumeUnit
	commodityD.NumberOfPackages = in.NumberOfPackages

	commodityT := bkgproto.CommodityT{}
	commodityT.ExportLicenseIssueDate = common.TimeToTimestamp(exportLicenseIssueDate.UTC().Truncate(time.Second))
	commodityT.ExportLicenseExpiryDate = common.TimeToTimestamp(exportLicenseExpiryDate.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	commodity := bkgproto.Commodity{CommodityD: &commodityD, CommodityT: &commodityT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &commodity, nil
}

// insertCommodity - Insert Commodity
func (bs *BkgService) insertCommodity(ctx context.Context, insertCommoditySQL string, commodity *bkgproto.Commodity, userEmail string, requestID string) error {
	commodityTmp, err := bs.CrCommodityStruct(ctx, commodity, userEmail, requestID)
	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = bs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertCommoditySQL, commodityTmp)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		commodity.CommodityD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(commodity.CommodityD.Uuid4)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		commodity.CommodityD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrCommodityStruct - process Commodity details
func (bs *BkgService) CrCommodityStruct(ctx context.Context, commodity *bkgproto.Commodity, userEmail string, requestID string) (*bkgstruct.Commodity, error) {
	commodityT := new(bkgstruct.CommodityT)
	commodityT.ExportLicenseIssueDate = common.TimestampToTime(commodity.CommodityT.ExportLicenseIssueDate)
	commodityT.ExportLicenseExpiryDate = common.TimestampToTime(commodity.CommodityT.ExportLicenseExpiryDate)

	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(commodity.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(commodity.CrUpdTime.UpdatedAt)

	commodityTmp := bkgstruct.Commodity{CommodityD: commodity.CommodityD, CommodityT: commodityT, CrUpdUser: commodity.CrUpdUser, CrUpdTime: crUpdTime}
	return &commodityTmp, nil
}
