package eblservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestConsignmentItemService_FetchConsignmentItemsTOByShippingInstructionId(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	consignmentItemService := NewConsignmentItemService(log, dbService, redisService, userServiceClient)
	consignmentItem, err := GetConsignmentItem(uint32(3), []byte{19, 22, 172, 59, 211, 211, 70, 34, 154, 54, 168, 98, 34, 123, 179, 133}, "1316ac3b-d3d3-4622-9a36-a862227bb385", "Leather Jackets", "411510", uint32(3), float64(4000), float64(0), "KGM", "", uint32(3), "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	consItemResponse := eblproto.FetchConsignmentItemsTOByShippingInstructionIdResponse{}
	consItemResponse.ConsignmentItem = consignmentItem

	form := eblproto.FetchConsignmentItemsTOByShippingInstructionIdRequest{}
	form.ShippingInstructionId = uint32(3)
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.FetchConsignmentItemsTOByShippingInstructionIdRequest
	}
	tests := []struct {
		cis     *ConsignmentItemService
		args    args
		want    *eblproto.FetchConsignmentItemsTOByShippingInstructionIdResponse
		wantErr bool
	}{
		{
			cis: consignmentItemService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &consItemResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		consItemResp, err := tt.cis.FetchConsignmentItemsTOByShippingInstructionId(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ConsignmentItemService.FetchConsignmentItemsTOByShippingInstructionId() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(consItemResp, tt.want) {
			t.Errorf("ConsignmentItemService.FetchConsignmentItemsTOByShippingInstructionId() = %v, want %v", consItemResp, tt.want)
		}
		consItemResult := consItemResp.ConsignmentItem
		assert.NotNil(t, consItemResult)
		assert.Equal(t, consItemResult.ConsignmentItemD.DescriptionOfGoods, "Leather Jackets", "they should be equal")
		assert.Equal(t, consItemResult.ConsignmentItemD.HsCode, "411510", "they should be equal")
		assert.Equal(t, consItemResult.ConsignmentItemD.WeightUnit, "KGM", "they should be equal")
	}
}

func GetConsignmentItem(id uint32, uuid4 []byte, idS string, descriptionOfGoods string, hsCode string, shippingInstructionId uint32, weight float64, volume float64, weightUnit string, volumeUnit string, shipmentId uint32, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*eblproto.ConsignmentItem, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	consignmentItemD := eblproto.ConsignmentItemD{}
	consignmentItemD.Id = id
	consignmentItemD.Uuid4 = uuid4
	consignmentItemD.IdS = idS
	consignmentItemD.DescriptionOfGoods = descriptionOfGoods
	consignmentItemD.HsCode = hsCode
	consignmentItemD.ShippingInstructionId = shippingInstructionId
	consignmentItemD.Weight = weight
	consignmentItemD.Volume = volume
	consignmentItemD.WeightUnit = weightUnit
	consignmentItemD.VolumeUnit = volumeUnit
	consignmentItemD.ShipmentId = shipmentId

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	consignmentItem := eblproto.ConsignmentItem{ConsignmentItemD: &consignmentItemD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &consignmentItem, nil
}
