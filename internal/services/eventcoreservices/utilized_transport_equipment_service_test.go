package eventcoreservices

import (
	"context"
	"testing"

	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
)

func TestUtilizedTransportEquipmentService_CreateUtilizedTransportEquipment(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	utilizedTransportEquipmentService := NewUtilizedTransportEquipmentService(log, dbService, redisService, userServiceClient)

	utilizedTransportEquipment := eventcoreproto.CreateUtilizedTransportEquipmentRequest{}
	utilizedTransportEquipment.EquipmentReference = "BMOU2149612"
	utilizedTransportEquipment.CargoGrossWeight = float64(3000)
	utilizedTransportEquipment.CargoGrossWeightUnit = "KGM"
	utilizedTransportEquipment.IsShipperOwned = false
	utilizedTransportEquipment.UserId = "auth0|66fd06d0bfea78a82bb42459"
	utilizedTransportEquipment.UserEmail = "sprov300@gmail.com"
	utilizedTransportEquipment.RequestId = "bks1m1g91jau4nkks2f0"

	equipment := eventcoreproto.CreateEquipmentRequest{}
	equipment.EquipmentReference = "BMOU2149612"
	equipment.IsoEquipmentCode = "22G1"
	equipment.TareWeight = float64(2000)
	equipment.WeightUnit = "KGM"
	equipment.UserId = "auth0|66fd06d0bfea78a82bb42459"
	equipment.UserEmail = "sprov300@gmail.com"
	equipment.RequestId = "bks1m1g91jau4nkks2f0"

	utilizedTransportEquipment.Equipment = &equipment

	type args struct {
		ctx context.Context
		in  *eventcoreproto.CreateUtilizedTransportEquipmentRequest
	}
	tests := []struct {
		uts     *UtilizedTransportEquipmentService
		args    args
		want    *eventcoreproto.CreateUtilizedTransportEquipmentResponse
		wantErr bool
	}{
		{
			uts: utilizedTransportEquipmentService,
			args: args{
				ctx: ctx,
				in:  &utilizedTransportEquipment,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		utilizedTransportEquipmentResp, err := tt.uts.CreateUtilizedTransportEquipment(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("UtilizedTransportEquipmentService.CreateUtilizedTransportEquipment() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		utilizedTransportEquipmentResult := utilizedTransportEquipmentResp.UtilizedTransportEquipment
		assert.NotNil(t, utilizedTransportEquipmentResult)
		assert.Equal(t, utilizedTransportEquipmentResult.UtilizedTransportEquipmentD.EquipmentReference, "BMOU2149612", "they should be equal")
		assert.Equal(t, utilizedTransportEquipmentResult.UtilizedTransportEquipmentD.CargoGrossWeight, float64(3000), "they should be equal")
		assert.Equal(t, utilizedTransportEquipmentResult.UtilizedTransportEquipmentD.CargoGrossWeightUnit, "KGM", "they should be equal")
		assert.NotNil(t, utilizedTransportEquipmentResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, utilizedTransportEquipmentResult.CrUpdTime.UpdatedAt)
	}
}
