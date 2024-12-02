package eventcoreservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestTransportCallService_CreateTransportCall(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	transportCallService := NewTransportCallService(log, dbService, redisService, userServiceClient)

	transportCall := eventcoreproto.CreateTransportCallRequest{}
	transportCall.TransportCallReference = "TC-REF-08_02-B"
	transportCall.TransportCallSequenceNumber = uint32(1)
	transportCall.FacilityId = uint32(834)
	transportCall.FacilityTypeCode = "POTE"
	transportCall.OtherFacility = ""
	transportCall.LocationId = uint32(0)
	transportCall.ModeOfTransportCode = "Maritime transport"
	transportCall.VesselId = uint32(1)
	transportCall.ImportVoyageId = uint32(10)
	transportCall.ExportVoyageId = uint32(10)
	transportCall.PortCallStatusCode = ""
	transportCall.PortVisitReference = ""
	transportCall.UserId = "auth0|66fd06d0bfea78a82bb42459"
	transportCall.UserEmail = "sprov300@gmail.com"
	transportCall.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eventcoreproto.CreateTransportCallRequest
	}
	tests := []struct {
		tcs     *TransportCallService
		args    args
		want    *eventcoreproto.CreateTransportCallResponse
		wantErr bool
	}{
		{
			tcs: transportCallService,
			args: args{
				ctx: ctx,
				in:  &transportCall,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		transportCallResp, err := tt.tcs.CreateTransportCall(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("TransportCallService.CreateTransportCall() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		transportCallResult := transportCallResp.TransportCall
		assert.NotNil(t, transportCallResult)
		assert.Equal(t, transportCallResult.TransportCallD.TransportCallReference, "TC-REF-08_02-B", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.TransportCallSequenceNumber, uint32(1), "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.FacilityTypeCode, "POTE", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.ModeOfTransportCode, "Maritime transport", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.VesselId, uint32(1), "they should be equal")
		assert.NotNil(t, transportCallResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, transportCallResult.CrUpdTime.UpdatedAt)

	}
}

func TestTransportCallService_GetTransportCalls(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	transportCallService := NewTransportCallService(log, dbService, redisService, userServiceClient)

	transportCall1, err := GetTransportCall(uint32(10), []byte{183, 133, 49, 122, 35, 64, 77, 183, 143, 179, 200, 223, 177, 237, 250, 96}, "b785317a-2340-4db7-8fb3-c8dfb1edfa60", "TC-REF-08_03-B", uint32(2), uint32(732), "POTE", "", uint32(0), "Maritime transport", uint32(3), uint32(1), uint32(1), "", "", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	transportCall2, err := GetTransportCall(uint32(9), []byte{127, 45, 131, 60, 44, 127, 79, 197, 167, 26, 229, 16, 136, 29, 166, 74}, "7f2d833c-2c7f-4fc5-a71a-e510881da64a", "TC-REF-08_03-A", uint32(1), uint32(834), "BRTH", "", uint32(0), "Maritime transport", uint32(3), uint32(1), uint32(1), "", "", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	transportCalls := []*eventcoreproto.TransportCall{}
	transportCalls = append(transportCalls, transportCall1, transportCall2)

	form := eventcoreproto.GetTransportCallsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "OA=="
	transportCallResponse := eventcoreproto.GetTransportCallsResponse{TransportCalls: transportCalls, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *eventcoreproto.GetTransportCallsRequest
	}
	tests := []struct {
		tcs     *TransportCallService
		args    args
		want    *eventcoreproto.GetTransportCallsResponse
		wantErr bool
	}{
		{
			tcs: transportCallService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &transportCallResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		transportCallResp, err := tt.tcs.GetTransportCalls(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("TransportCallService.GetTransportCalls() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transportCallResp, tt.want) {
			t.Errorf("TransportCallService.GetTransportCalls() = %v, want %v", transportCallResp, tt.want)
		}

		transportCallResult := transportCallResp.TransportCalls[0]
		assert.NotNil(t, transportCallResult)
		assert.Equal(t, transportCallResult.TransportCallD.TransportCallReference, "TC-REF-08_03-B", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.TransportCallSequenceNumber, uint32(2), "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.FacilityTypeCode, "POTE", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.ModeOfTransportCode, "Maritime transport", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.VesselId, uint32(3), "they should be equal")
		assert.NotNil(t, transportCallResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, transportCallResult.CrUpdTime.UpdatedAt)
	}
}

func TestTransportCallService_FindTransportCall(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	transportCallService := NewTransportCallService(log, dbService, redisService, userServiceClient)

	transportCall, err := GetTransportCall(uint32(9), []byte{127, 45, 131, 60, 44, 127, 79, 197, 167, 26, 229, 16, 136, 29, 166, 74}, "7f2d833c-2c7f-4fc5-a71a-e510881da64a", "TC-REF-08_03-A", uint32(1), uint32(834), "BRTH", "", uint32(0), "Maritime transport", uint32(3), uint32(1), uint32(1), "", "", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	transportCallResponse := eventcoreproto.FindTransportCallResponse{}
	transportCallResponse.TransportCall = transportCall

	gform := commonproto.GetRequest{}
	gform.Id = "7f2d833c-2c7f-4fc5-a71a-e510881da64a"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := eventcoreproto.FindTransportCallRequest{}
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *eventcoreproto.FindTransportCallRequest
	}
	tests := []struct {
		tcs     *TransportCallService
		args    args
		want    *eventcoreproto.FindTransportCallResponse
		wantErr bool
	}{
		{
			tcs: transportCallService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &transportCallResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		transportCallResp, err := tt.tcs.FindTransportCall(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("TransportCallService.FindTransportCall() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transportCallResp, tt.want) {
			t.Errorf("TransportCallService.FindTransportCall() = %v, want %v", transportCallResp, tt.want)
		}

		transportCallResult := transportCallResp.TransportCall
		assert.NotNil(t, transportCallResult)
		assert.Equal(t, transportCallResult.TransportCallD.TransportCallReference, "TC-REF-08_03-A", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.TransportCallSequenceNumber, uint32(1), "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.FacilityTypeCode, "BRTH", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.ModeOfTransportCode, "Maritime transport", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.VesselId, uint32(3), "they should be equal")
		assert.NotNil(t, transportCallResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, transportCallResult.CrUpdTime.UpdatedAt)
	}
}

func TestTransportCallService_GetTransportCallByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	transportCallService := NewTransportCallService(log, dbService, redisService, userServiceClient)

	transportCall, err := GetTransportCall(uint32(9), []byte{127, 45, 131, 60, 44, 127, 79, 197, 167, 26, 229, 16, 136, 29, 166, 74}, "7f2d833c-2c7f-4fc5-a71a-e510881da64a", "TC-REF-08_03-A", uint32(1), uint32(834), "BRTH", "", uint32(0), "Maritime transport", uint32(3), uint32(1), uint32(1), "", "", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	transportCallResponse := eventcoreproto.GetTransportCallByPkResponse{}
	transportCallResponse.TransportCall = transportCall

	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(9)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := eventcoreproto.GetTransportCallByPkRequest{}
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *eventcoreproto.GetTransportCallByPkRequest
	}
	tests := []struct {
		tcs     *TransportCallService
		args    args
		want    *eventcoreproto.GetTransportCallByPkResponse
		wantErr bool
	}{
		{
			tcs: transportCallService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &transportCallResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		transportCallResp, err := tt.tcs.GetTransportCallByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("TransportCallService.GetTransportCallByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(transportCallResp, tt.want) {
			t.Errorf("TransportCallService.GetTransportCallByPk() = %v, want %v", transportCallResp, tt.want)
		}
		transportCallResult := transportCallResp.TransportCall
		assert.NotNil(t, transportCallResult)
		assert.Equal(t, transportCallResult.TransportCallD.TransportCallReference, "TC-REF-08_03-A", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.TransportCallSequenceNumber, uint32(1), "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.FacilityTypeCode, "BRTH", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.ModeOfTransportCode, "Maritime transport", "they should be equal")
		assert.Equal(t, transportCallResult.TransportCallD.VesselId, uint32(3), "they should be equal")
		assert.NotNil(t, transportCallResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, transportCallResult.CrUpdTime.UpdatedAt)
	}
}

func GetTransportCall(id uint32, uuid4 []byte, idS string, transportCallReference string, transportCallSequenceNumber uint32, facilityId uint32, facilityTypeCode string, otherFacility string, locationId uint32, modeOfTransportCode string, vesselId uint32, importVoyageId uint32, exportVoyageId uint32, portCallStatusCode string, portVisitReference string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*eventcoreproto.TransportCall, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	transportCallD := new(eventcoreproto.TransportCallD)
	transportCallD.Id = id
	transportCallD.Uuid4 = uuid4
	transportCallD.IdS = idS
	transportCallD.TransportCallReference = transportCallReference
	transportCallD.TransportCallSequenceNumber = transportCallSequenceNumber
	transportCallD.FacilityId = facilityId
	transportCallD.FacilityTypeCode = facilityTypeCode
	transportCallD.OtherFacility = otherFacility
	transportCallD.LocationId = locationId
	transportCallD.ModeOfTransportCode = modeOfTransportCode
	transportCallD.VesselId = vesselId
	transportCallD.ImportVoyageId = importVoyageId
	transportCallD.ExportVoyageId = exportVoyageId
	transportCallD.PortCallStatusCode = portCallStatusCode
	transportCallD.PortVisitReference = portVisitReference

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	transportCall := eventcoreproto.TransportCall{TransportCallD: transportCallD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &transportCall, nil
}
