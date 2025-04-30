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

func TestShippingInstructionService_GetShippingInstructions(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	shippingInstructionService := NewShippingInstructionService(log, dbService, redisService, userServiceClient)
	shippingInstruction1, err := GetShippingInstruction(uint32(10), []byte{193, 68, 198, 223, 68, 14, 64, 101, 132, 48, 244, 107, 159, 166, 126, 101}, "c144c6df-440e-4065-8430-f46b9fa67e65", "c144c6dff46b9fa67e65", "RECE", true, uint32(2), uint32(4), true, true, true, false, uint32(0), "", "", "", "", "", "", "2021-12-24T00:00:00Z", "2021-12-31T00:00:00Z")
	if err != nil {
		t.Error(err)
		return
	}
	shippingInstruction2, err := GetShippingInstruction(uint32(9), []byte{44, 51, 127, 204, 40, 20, 66, 179, 167, 82, 241, 132, 126, 116, 203, 167}, "2c337fcc-2814-42b3-a752-f1847e74cba7", "SI_REF_10", "DRFT", true, uint32(2), uint32(4), true, true, true, false, uint32(0), "", "", "", "", "", "", "2021-12-24T00:00:00Z", "2021-12-31T00:00:00Z")
	if err != nil {
		t.Error(err)
		return
	}
	shippingInstructions := []*eblproto.ShippingInstruction{}
	shippingInstructions = append(shippingInstructions, shippingInstruction1, shippingInstruction2)

	form := eblproto.GetShippingInstructionsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "OA=="
	shipInstsResponse := eblproto.GetShippingInstructionsResponse{ShippingInstructions: shippingInstructions, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *eblproto.GetShippingInstructionsRequest
	}
	tests := []struct {
		sis     *ShippingInstructionService
		args    args
		want    *eblproto.GetShippingInstructionsResponse
		wantErr bool
	}{
		{
			sis: shippingInstructionService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &shipInstsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		shipInstsResp, err := tt.sis.GetShippingInstructions(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ShippingInstructionService.GetShippingInstructions() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(shipInstsResp, tt.want) {
			t.Errorf("ShippingInstructionService.GetShippingInstructions() = %v, want %v", shipInstsResp, tt.want)
		}
		assert.NotNil(t, shipInstsResp)
		shippingInstructionResult := shipInstsResp.ShippingInstructions[1]
		assert.Equal(t, shippingInstructionResult.ShippingInstructionD.ShippingInstructionReference, "SI_REF_10", "they should be equal")
		assert.Equal(t, shippingInstructionResult.ShippingInstructionD.DocumentStatus, "DRFT", "they should be equal")
		assert.True(t, shippingInstructionResult.ShippingInstructionD.IsShippedOnboardType, "Its true")
		assert.True(t, shippingInstructionResult.ShippingInstructionD.IsToOrder, "Its true")
	}
}

func TestShippingInstructionService_FindById(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	shippingInstructionService := NewShippingInstructionService(log, dbService, redisService, userServiceClient)
	shippingInstruction, err := GetShippingInstruction(uint32(10), []byte{193, 68, 198, 223, 68, 14, 64, 101, 132, 48, 244, 107, 159, 166, 126, 101}, "c144c6df-440e-4065-8430-f46b9fa67e65", "c144c6dff46b9fa67e65", "RECE", true, uint32(2), uint32(4), true, true, true, false, uint32(0), "", "", "", "", "", "", "2021-12-24T00:00:00Z", "2021-12-31T00:00:00Z")
	if err != nil {
		t.Error(err)
		return
	}
	shipInstResponse := eblproto.FindByIdResponse{}
	shipInstResponse.ShippingInstruction = shippingInstruction

	form := eblproto.FindByIdRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "c144c6df-440e-4065-8430-f46b9fa67e65"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx context.Context
		in  *eblproto.FindByIdRequest
	}
	tests := []struct {
		sis     *ShippingInstructionService
		args    args
		want    *eblproto.FindByIdResponse
		wantErr bool
	}{
		{
			sis: shippingInstructionService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &shipInstResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		shipInstResp, err := tt.sis.FindById(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ShippingInstructionService.FindById() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(shipInstResp, tt.want) {
			t.Errorf("ShippingInstructionService.FindById() = %v, want %v", shipInstResp, tt.want)
		}
		assert.NotNil(t, shipInstResp)
		shippingInstructionResult := shipInstResp.ShippingInstruction
		assert.Equal(t, shippingInstructionResult.ShippingInstructionD.ShippingInstructionReference, "c144c6dff46b9fa67e65", "they should be equal")
		assert.Equal(t, shippingInstructionResult.ShippingInstructionD.DocumentStatus, "RECE", "they should be equal")
		assert.True(t, shippingInstructionResult.ShippingInstructionD.IsShippedOnboardType, "Its true")
		assert.True(t, shippingInstructionResult.ShippingInstructionD.IsToOrder, "Its true")
	}
}

func TestShippingInstructionService_GetShippingInstructionByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	shippingInstructionService := NewShippingInstructionService(log, dbService, redisService, userServiceClient)
	shippingInstruction, err := GetShippingInstruction(uint32(10), []byte{193, 68, 198, 223, 68, 14, 64, 101, 132, 48, 244, 107, 159, 166, 126, 101}, "c144c6df-440e-4065-8430-f46b9fa67e65", "c144c6dff46b9fa67e65", "RECE", true, uint32(2), uint32(4), true, true, true, false, uint32(0), "", "", "", "", "", "", "2021-12-24T00:00:00Z", "2021-12-31T00:00:00Z")
	if err != nil {
		t.Error(err)
		return
	}
	shipInstResponse := eblproto.GetShippingInstructionByPkResponse{}
	shipInstResponse.ShippingInstruction = shippingInstruction

	form := eblproto.GetShippingInstructionByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(10)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx context.Context
		in  *eblproto.GetShippingInstructionByPkRequest
	}
	tests := []struct {
		sis     *ShippingInstructionService
		args    args
		want    *eblproto.GetShippingInstructionByPkResponse
		wantErr bool
	}{
		{
			sis: shippingInstructionService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &shipInstResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		shipInstResp, err := tt.sis.GetShippingInstructionByPk(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ShippingInstructionService.GetShippingInstructionByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(shipInstResp, tt.want) {
			t.Errorf("ShippingInstructionService.GetShippingInstructionByPk() = %v, want %v", shipInstResp, tt.want)
		}
		assert.NotNil(t, shipInstResp)
		shippingInstructionResult := shipInstResp.ShippingInstruction
		assert.Equal(t, shippingInstructionResult.ShippingInstructionD.ShippingInstructionReference, "c144c6dff46b9fa67e65", "they should be equal")
		assert.Equal(t, shippingInstructionResult.ShippingInstructionD.DocumentStatus, "RECE", "they should be equal")
		assert.True(t, shippingInstructionResult.ShippingInstructionD.IsShippedOnboardType, "Its true")
		assert.True(t, shippingInstructionResult.ShippingInstructionD.IsToOrder, "Its true")
	}
}

func TestShippingInstructionService_FindByReference(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	shippingInstructionService := NewShippingInstructionService(log, dbService, redisService, userServiceClient)
	shippingInstruction, err := GetShippingInstruction(uint32(10), []byte{193, 68, 198, 223, 68, 14, 64, 101, 132, 48, 244, 107, 159, 166, 126, 101}, "c144c6df-440e-4065-8430-f46b9fa67e65", "c144c6dff46b9fa67e65", "RECE", true, uint32(2), uint32(4), true, true, true, false, uint32(0), "", "", "", "", "", "", "2021-12-24T00:00:00Z", "2021-12-31T00:00:00Z")
	if err != nil {
		t.Error(err)
		return
	}

	shipInstResponse := eblproto.FindByReferenceResponse{}
	shipInstResponse.ShippingInstruction = shippingInstruction

	form := eblproto.FindByReferenceRequest{}
	form.ShippingInstructionReference = "c144c6dff46b9fa67e65"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.FindByReferenceRequest
	}
	tests := []struct {
		sis     *ShippingInstructionService
		args    args
		want    *eblproto.FindByReferenceResponse
		wantErr bool
	}{
		{
			sis: shippingInstructionService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &shipInstResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		shipInstResp, err := tt.sis.FindByReference(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ShippingInstructionService.FindByReference() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(shipInstResp, tt.want) {
			t.Errorf("ShippingInstructionService.FindByReference() = %v, want %v", shipInstResp, tt.want)
		}
		assert.NotNil(t, shipInstResp)
		shippingInstructionResult := shipInstResp.ShippingInstruction
		assert.Equal(t, shippingInstructionResult.ShippingInstructionD.ShippingInstructionReference, "c144c6dff46b9fa67e65", "they should be equal")
		assert.Equal(t, shippingInstructionResult.ShippingInstructionD.DocumentStatus, "RECE", "they should be equal")
		assert.True(t, shippingInstructionResult.ShippingInstructionD.IsShippedOnboardType, "Its true")
		assert.True(t, shippingInstructionResult.ShippingInstructionD.IsToOrder, "Its true")
	}
}

func TestShippingInstructionService_UpdateShippingInstructionByShippingInstructionReference(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}
	ctx := LoginUser()

	shippingInstructionService := NewShippingInstructionService(log, dbService, redisService, userServiceClient)

	form := eblproto.UpdateShippingInstructionByShippingInstructionReferenceRequest{}
	form.ShippingInstructionReference = "SI_REF_4"
	form.DocumentStatus = "APPR"
	form.TransportDocumentTypeCode = ""
	form.DisplayedNameForPlaceOfReceipt = ""

	updateResponse := eblproto.UpdateShippingInstructionByShippingInstructionReferenceResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateShippingInstructionByShippingInstructionReferenceRequest
	}
	tests := []struct {
		sis     *ShippingInstructionService
		args    args
		want    *eblproto.UpdateShippingInstructionByShippingInstructionReferenceResponse
		wantErr bool
	}{
		{
			sis: shippingInstructionService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.sis.UpdateShippingInstructionByShippingInstructionReference(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ShippingInstructionService.UpdateShippingInstructionByShippingInstructionReference() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ShippingInstructionService.UpdateShippingInstructionByShippingInstructionReference() = %v, want %v", got, tt.want)
		}
	}
}

func GetShippingInstruction(id uint32, uuid4 []byte, idS string, shippingInstructionReference string, documentStatus string, isShippedOnboardType bool, numberOfCopies uint32, numberOfOriginals uint32, isElectronic bool, isToOrder bool, areChargesDisplayedOnOriginals bool, areChargesDisplayedOnCopies bool, locationId uint32, transportDocumentTypeCode string, displayedNameForPlaceOfReceipt string, displayedNameForPortOfLoad string, displayedNameForPortOfDischarge string, displayedNameForPlaceOfDelivery string, amendToTransportDocument string, createdDateTime string, updatedDateTime string) (*eblproto.ShippingInstruction, error) {
	createdDateTime1, err := common.ConvertTimeToTimestamp(Layout, createdDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedDateTime1, err := common.ConvertTimeToTimestamp(Layout, updatedDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	shippingInstructionD := eblproto.ShippingInstructionD{}
	shippingInstructionD.Id = id
	shippingInstructionD.Uuid4 = uuid4
	shippingInstructionD.IdS = idS
	shippingInstructionD.ShippingInstructionReference = shippingInstructionReference
	shippingInstructionD.DocumentStatus = documentStatus
	shippingInstructionD.IsShippedOnboardType = isShippedOnboardType
	shippingInstructionD.NumberOfCopies = numberOfCopies
	shippingInstructionD.NumberOfOriginals = numberOfOriginals
	shippingInstructionD.IsElectronic = isElectronic
	shippingInstructionD.IsToOrder = isToOrder
	shippingInstructionD.AreChargesDisplayedOnOriginals = areChargesDisplayedOnOriginals
	shippingInstructionD.AreChargesDisplayedOnCopies = areChargesDisplayedOnCopies
	shippingInstructionD.LocationId = locationId
	shippingInstructionD.TransportDocumentTypeCode = transportDocumentTypeCode
	shippingInstructionD.DisplayedNameForPlaceOfReceipt = displayedNameForPlaceOfReceipt
	shippingInstructionD.DisplayedNameForPortOfLoad = displayedNameForPortOfLoad
	shippingInstructionD.DisplayedNameForPortOfDischarge = displayedNameForPortOfDischarge
	shippingInstructionD.DisplayedNameForPlaceOfDelivery = displayedNameForPlaceOfDelivery
	shippingInstructionD.AmendToTransportDocument = amendToTransportDocument

	shippingInstructionT := eblproto.ShippingInstructionT{}
	shippingInstructionT.CreatedDateTime = createdDateTime1
	shippingInstructionT.UpdatedDateTime = updatedDateTime1

	shippingInstruction := eblproto.ShippingInstruction{ShippingInstructionD: &shippingInstructionD, ShippingInstructionT: &shippingInstructionT}

	return &shippingInstruction, nil
}
