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

func TestTransportDocumentService_GetTransportDocuments(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	transportDocumentService := NewTransportDocumentService(log, dbService, redisService, userServiceClient, currencyService)
	transportDocument1, err := GetTransportDocument(uint32(6), []byte{84, 74, 242, 38, 15, 235, 64, 110, 170, 113, 85, 41, 169, 108, 181, 91}, "544af226-0feb-406e-aa71-5529a96cb55b", "be038e58-5365", uint32(1), "2020-11-25T00:00:00Z", "2020-12-24T00:00:00Z", "2020-12-31T00:00:00Z", uint32(12), uint32(3), uint32(1), "EUR", int64(1212), "12.12", int32(12), "c49ea2d6-3806-46c8-8490-294affc71286", "2021-11-28T14:12:56Z", "2021-12-01T07:41:00Z")
	if err != nil {
		t.Error(err)
		return
	}

	transportDocument2, err := GetTransportDocument(uint32(5), []byte{207, 72, 173, 10, 154, 75, 72, 167, 183, 82, 194, 72, 251, 93, 136, 217}, "cf48ad0a-9a4b-48a7-b752-c248fb5d88d9", "c90a0ed6-ccc9-48e3", uint32(8), "2022-05-16T00:00:00Z", "2022-05-15T00:00:00Z", "2022-05-14T00:00:00Z", uint32(12), uint32(3), uint32(9), "EUR", int64(1212), "12.12", int32(12), "8e463a84-0a2d-47cd-9332-51e6cb36b635", "2021-11-28T14:12:56Z", "2021-12-01T07:41:00Z")
	if err != nil {
		t.Error(err)
		return
	}

	transportDocuments := []*eblproto.TransportDocument{}
	transportDocuments = append(transportDocuments, transportDocument1, transportDocument2)

	form := eblproto.GetTransportDocumentsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "NA=="
	transDocsResponse := eblproto.GetTransportDocumentsResponse{TransportDocuments: transportDocuments, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *eblproto.GetTransportDocumentsRequest
	}
	tests := []struct {
		tds     *TransportDocumentService
		args    args
		want    *eblproto.GetTransportDocumentsResponse
		wantErr bool
	}{
		{
			tds: transportDocumentService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &transDocsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transDocsResp, err := tt.tds.GetTransportDocuments(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("TransportDocumentService.GetTransportDocuments() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transDocsResp, tt.want) {
			t.Errorf("TransportDocumentService.GetTransportDocuments() = %v, want %v", transDocsResp, tt.want)
		}
		assert.NotNil(t, transDocsResp)
		transDocumentResult := transDocsResp.TransportDocuments[1]
		assert.Equal(t, transDocumentResult.TransportDocumentD.TransportDocumentReference, "c90a0ed6-ccc9-48e3", "they should be equal")
		assert.Equal(t, transDocumentResult.TransportDocumentD.DeclaredValueCurrency, "EUR", "they should be equal")
	}
}

func TestTransportDocumentService_FindTransportDocumentById(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	transportDocumentService := NewTransportDocumentService(log, dbService, redisService, userServiceClient, currencyService)
	transportDocument, err := GetTransportDocument(uint32(6), []byte{84, 74, 242, 38, 15, 235, 64, 110, 170, 113, 85, 41, 169, 108, 181, 91}, "544af226-0feb-406e-aa71-5529a96cb55b", "be038e58-5365", uint32(1), "2020-11-25T00:00:00Z", "2020-12-24T00:00:00Z", "2020-12-31T00:00:00Z", uint32(12), uint32(3), uint32(1), "EUR", int64(1212), "12.12", int32(12), "c49ea2d6-3806-46c8-8490-294affc71286", "2021-11-28T14:12:56Z", "2021-12-01T07:41:00Z")
	if err != nil {
		t.Error(err)
		return
	}

	tranDocResponse := eblproto.FindTransportDocumentByIdResponse{}
	tranDocResponse.TransportDocument = transportDocument

	form := eblproto.FindTransportDocumentByIdRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "544af226-0feb-406e-aa71-5529a96cb55b"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx context.Context
		in  *eblproto.FindTransportDocumentByIdRequest
	}
	tests := []struct {
		tds     *TransportDocumentService
		args    args
		want    *eblproto.FindTransportDocumentByIdResponse
		wantErr bool
	}{
		{
			tds: transportDocumentService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &tranDocResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transDocResp, err := tt.tds.FindTransportDocumentById(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("TransportDocumentService.FindTransportDocumentById() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transDocResp, tt.want) {
			t.Errorf("TransportDocumentService.FindTransportDocumentById() = %v, want %v", transDocResp, tt.want)
		}
		assert.NotNil(t, transDocResp)
		transDocumentResult := transDocResp.TransportDocument
		assert.Equal(t, transDocumentResult.TransportDocumentD.TransportDocumentReference, "be038e58-5365", "they should be equal")
		assert.Equal(t, transDocumentResult.TransportDocumentD.DeclaredValueCurrency, "EUR", "they should be equal")
	}
}

func TestTransportDocumentService_GetTransportDocumentByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	transportDocumentService := NewTransportDocumentService(log, dbService, redisService, userServiceClient, currencyService)
	transportDocument, err := GetTransportDocument(uint32(6), []byte{84, 74, 242, 38, 15, 235, 64, 110, 170, 113, 85, 41, 169, 108, 181, 91}, "544af226-0feb-406e-aa71-5529a96cb55b", "be038e58-5365", uint32(1), "2020-11-25T00:00:00Z", "2020-12-24T00:00:00Z", "2020-12-31T00:00:00Z", uint32(12), uint32(3), uint32(1), "EUR", int64(1212), "12.12", int32(12), "c49ea2d6-3806-46c8-8490-294affc71286", "2021-11-28T14:12:56Z", "2021-12-01T07:41:00Z")
	if err != nil {
		t.Error(err)
		return
	}

	tranDocResponse := eblproto.GetTransportDocumentByPkResponse{}
	tranDocResponse.TransportDocument = transportDocument

	form := eblproto.GetTransportDocumentByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(6)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx context.Context
		in  *eblproto.GetTransportDocumentByPkRequest
	}
	tests := []struct {
		tds     *TransportDocumentService
		args    args
		want    *eblproto.GetTransportDocumentByPkResponse
		wantErr bool
	}{
		{
			tds: transportDocumentService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &tranDocResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transDocResp, err := tt.tds.GetTransportDocumentByPk(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("TransportDocumentService.GetTransportDocumentByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transDocResp, tt.want) {
			t.Errorf("TransportDocumentService.GetTransportDocumentByPk() = %v, want %v", transDocResp, tt.want)
		}
		assert.NotNil(t, transDocResp)
		transDocumentResult := transDocResp.TransportDocument
		assert.Equal(t, transDocumentResult.TransportDocumentD.TransportDocumentReference, "be038e58-5365", "they should be equal")
		assert.Equal(t, transDocumentResult.TransportDocumentD.DeclaredValueCurrency, "EUR", "they should be equal")
	}
}

func TestTransportDocumentService_FindByTransportDocumentReference(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	transportDocumentService := NewTransportDocumentService(log, dbService, redisService, userServiceClient, currencyService)
	transportDocument, err := GetTransportDocument(uint32(6), []byte{84, 74, 242, 38, 15, 235, 64, 110, 170, 113, 85, 41, 169, 108, 181, 91}, "544af226-0feb-406e-aa71-5529a96cb55b", "be038e58-5365", uint32(1), "2020-11-25T00:00:00Z", "2020-12-24T00:00:00Z", "2020-12-31T00:00:00Z", uint32(12), uint32(3), uint32(1), "EUR", int64(1212), "12.12", int32(12), "c49ea2d6-3806-46c8-8490-294affc71286", "2021-11-28T14:12:56Z", "2021-12-01T07:41:00Z")
	if err != nil {
		t.Error(err)
		return
	}

	tranDocResponse := eblproto.FindByTransportDocumentReferenceResponse{}
	tranDocResponse.TransportDocument = transportDocument

	form := eblproto.FindByTransportDocumentReferenceRequest{}
	form.TransportDocumentReference = "be038e58-5365"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.FindByTransportDocumentReferenceRequest
	}
	tests := []struct {
		tds     *TransportDocumentService
		args    args
		want    *eblproto.FindByTransportDocumentReferenceResponse
		wantErr bool
	}{
		{
			tds: transportDocumentService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &tranDocResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transDocResp, err := tt.tds.FindByTransportDocumentReference(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("TransportDocumentService.FindByTransportDocumentReference() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transDocResp, tt.want) {
			t.Errorf("TransportDocumentService.FindByTransportDocumentReference() = %v, want %v", transDocResp, tt.want)
		}
		assert.NotNil(t, transDocResp)
		transDocumentResult := transDocResp.TransportDocument
		assert.Equal(t, transDocumentResult.TransportDocumentD.TransportDocumentReference, "be038e58-5365", "they should be equal")
		assert.Equal(t, transDocumentResult.TransportDocumentD.DeclaredValueCurrency, "EUR", "they should be equal")
	}
}

func GetTransportDocument(id uint32, uuid4 []byte, idS string, transportDocumentReference string, locationId uint32, issueDate string, shippedOnboardDate string, receivedForShipmentDate string, numberOfOriginals uint32, carrierId uint32, shippingInstructionId uint32, declaredValueCurrency string, declaredValue int64, declaredValueString string, numberOfRiderPages int32, issuingParty string, createdDateTime string, updatedDateTime string) (*eblproto.TransportDocument, error) {
	issueDate1, err := common.ConvertTimeToTimestamp(Layout, issueDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	shippedOnboardDate1, err := common.ConvertTimeToTimestamp(Layout, shippedOnboardDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	receivedForShipmentDate1, err := common.ConvertTimeToTimestamp(Layout, receivedForShipmentDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	createdDateTime1, err := common.ConvertTimeToTimestamp(Layout, createdDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedDateTime1, err := common.ConvertTimeToTimestamp(Layout, updatedDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	transportDocumentD := eblproto.TransportDocumentD{}
	transportDocumentD.Id = id
	transportDocumentD.Uuid4 = uuid4
	transportDocumentD.IdS = idS
	transportDocumentD.TransportDocumentReference = transportDocumentReference
	transportDocumentD.LocationId = locationId
	transportDocumentD.NumberOfOriginals = numberOfOriginals
	transportDocumentD.CarrierId = carrierId
	transportDocumentD.ShippingInstructionId = shippingInstructionId
	transportDocumentD.DeclaredValueCurrency = declaredValueCurrency
	transportDocumentD.DeclaredValue = declaredValue
	transportDocumentD.DeclaredValueString = declaredValueString
	transportDocumentD.NumberOfRiderPages = numberOfRiderPages
	transportDocumentD.IssuingParty = issuingParty

	transportDocumentT := eblproto.TransportDocumentT{}
	transportDocumentT.IssueDate = issueDate1
	transportDocumentT.ShippedOnboardDate = shippedOnboardDate1
	transportDocumentT.ReceivedForShipmentDate = receivedForShipmentDate1
	transportDocumentT.CreatedDateTime = createdDateTime1
	transportDocumentT.UpdatedDateTime = updatedDateTime1

	transportDocument := eblproto.TransportDocument{TransportDocumentD: &transportDocumentD, TransportDocumentT: &transportDocumentT}
	return &transportDocument, nil
}

func TestTransportDocumentService_CreateTransportDocument(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	transportDocumentService := NewTransportDocumentService(log, dbService, redisService, userServiceClient, currencyService)

	transportDocument := eblproto.CreateTransportDocumentRequest{}
	transportDocument.TransportDocumentReference = "0cc0bef0-a7c8-4c03"
	transportDocument.LocationId = uint32(8)
	transportDocument.IssueDate = "11/25/2020"
	transportDocument.ShippedOnboardDate = "12/24/2020"
	transportDocument.ReceivedForShipmentDate = "12/31/2020"
	transportDocument.NumberOfOriginals = uint32(12)
	transportDocument.CarrierId = uint32(3)
	transportDocument.ShippingInstructionId = uint32(8)
	transportDocument.DeclaredValueCurrency = "EUR"
	transportDocument.DeclaredValue = "1212"
	transportDocument.NumberOfRiderPages = int32(12)
	transportDocument.IssuingParty = "499918a2-d12d-4df6-840c-dd92357002df"
	transportDocument.CreatedDateTime = "11/28/2021"
	transportDocument.UpdatedDateTime = "12/01/2021"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateTransportDocumentRequest
	}
	tests := []struct {
		tds     *TransportDocumentService
		args    args
		wantErr bool
	}{
		{
			tds: transportDocumentService,
			args: args{
				ctx: ctx,
				in:  &transportDocument,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transDocResp, err := tt.tds.CreateTransportDocument(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("TransportDocumentService.CreateTransportDocument() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, transDocResp)
		transDocumentResult := transDocResp.TransportDocument
		assert.Equal(t, transDocumentResult.TransportDocumentD.TransportDocumentReference, "0cc0bef0-a7c8-4c03", "they should be equal")
		assert.Equal(t, transDocumentResult.TransportDocumentD.DeclaredValueCurrency, "EUR", "they should be equal")
		assert.Equal(t, transDocumentResult.TransportDocumentD.DeclaredValue, int64(121200), "they should be equal")
		assert.Equal(t, transDocumentResult.TransportDocumentD.NumberOfRiderPages, int32(12), "they should be equal")
	}
}
