package ovsservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestServiceScheduleService_CreateServiceSchedule(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceScheduleService := NewServiceScheduleService(log, dbService, redisService, userServiceClient)

	serviceSchedule := ovsproto.CreateServiceScheduleRequest{}
	serviceSchedule.CarrierServiceName = "B_carrier_service_name_1"
	serviceSchedule.CarrierServiceCode = "B_HLC"
	serviceSchedule.UniversalServiceReference = "SR00003H"
	serviceSchedule.UserId = "auth0|66fd06d0bfea78a82bb42459"
	serviceSchedule.UserEmail = "sprov300@gmail.com"
	serviceSchedule.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *ovsproto.CreateServiceScheduleRequest
	}
	tests := []struct {
		ss      *ServiceScheduleService
		args    args
		want    *ovsproto.CreateServiceScheduleResponse
		wantErr bool
	}{
		{
			ss: serviceScheduleService,
			args: args{
				ctx: ctx,
				in:  &serviceSchedule,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		serviceScheduleResp, err := tt.ss.CreateServiceSchedule(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceScheduleService.CreateServiceSchedule() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		serviceScheduleResult := serviceScheduleResp.ServiceSchedule
		assert.NotNil(t, serviceScheduleResult)
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.CarrierServiceName, "B_carrier_service_name_1", "they should be equal")
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.CarrierServiceCode, "B_HLC", "they should be equal")
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.UniversalServiceReference, "SR00003H", "they should be equal")
		assert.NotNil(t, serviceScheduleResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, serviceScheduleResult.CrUpdTime.UpdatedAt)
	}
}

func TestServiceScheduleService_GetServiceSchedules(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceScheduleService := NewServiceScheduleService(log, dbService, redisService, userServiceClient)
	serviceSchedule1, err := GetServiceSchedule(uint32(2), []byte{128, 19, 139, 40, 180, 29, 66, 185, 178, 95, 116, 113, 246, 134, 201, 27}, "80138b28-b41d-42b9-b25f-7471f686c91b", "B_carrier_service_name_1", "B_HLC", "SR00003H", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	serviceSchedule2, err := GetServiceSchedule(uint32(1), []byte{215, 99, 245, 80, 31, 119, 72, 225, 175, 83, 78, 3, 141, 96, 74, 211}, "d763f550-1f77-48e1-af53-4e038d604ad3", "B_carrier_service_name", "B_HLC", "SR00002B", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	serviceSchedules := []*ovsproto.ServiceSchedule{}
	serviceSchedules = append(serviceSchedules, serviceSchedule1, serviceSchedule2)

	form := ovsproto.GetServiceSchedulesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	serviceScheduleResp := ovsproto.GetServiceSchedulesResponse{ServiceSchedules: serviceSchedules, NextCursor: nextc}
	type args struct {
		ctx context.Context
		in  *ovsproto.GetServiceSchedulesRequest
	}
	tests := []struct {
		ss      *ServiceScheduleService
		args    args
		want    *ovsproto.GetServiceSchedulesResponse
		wantErr bool
	}{
		{
			ss: serviceScheduleService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &serviceScheduleResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		serviceScheduleResponse, err := tt.ss.GetServiceSchedules(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceScheduleService.GetServiceSchedules() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(serviceScheduleResponse, tt.want) {
			t.Errorf("ServiceScheduleService.GetServiceSchedules() = %v, want %v", serviceScheduleResponse, tt.want)
		}

		serviceScheduleResult := serviceScheduleResponse.ServiceSchedules[0]
		assert.NotNil(t, serviceScheduleResult)
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.CarrierServiceName, "B_carrier_service_name_1", "they should be equal")
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.CarrierServiceCode, "B_HLC", "they should be equal")
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.UniversalServiceReference, "SR00003H", "they should be equal")
		assert.NotNil(t, serviceScheduleResult.CrUpdTime.CreatedAt)
	}
}

func TestServiceScheduleService_GetServiceSchedule(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceScheduleService := NewServiceScheduleService(log, dbService, redisService, userServiceClient)

	serviceSchedule1, err := GetServiceSchedule(uint32(1), []byte{215, 99, 245, 80, 31, 119, 72, 225, 175, 83, 78, 3, 141, 96, 74, 211}, "d763f550-1f77-48e1-af53-4e038d604ad3", "B_carrier_service_name", "B_HLC", "SR00002B", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	serviceScheduleResp := ovsproto.GetServiceScheduleResponse{}
	serviceScheduleResp.ServiceSchedule = serviceSchedule1

	gform := commonproto.GetRequest{}
	gform.Id = "d763f550-1f77-48e1-af53-4e038d604ad3"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := ovsproto.GetServiceScheduleRequest{}
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *ovsproto.GetServiceScheduleRequest
	}
	tests := []struct {
		ss      *ServiceScheduleService
		args    args
		want    *ovsproto.GetServiceScheduleResponse
		wantErr bool
	}{
		{
			ss: serviceScheduleService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &serviceScheduleResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		serviceScheduleResponse, err := tt.ss.GetServiceSchedule(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceScheduleService.GetServiceSchedule() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(serviceScheduleResponse, tt.want) {
			t.Errorf("ServiceScheduleService.GetServiceSchedule() = %v, want %v", serviceScheduleResponse, tt.want)
		}
		serviceScheduleResult := serviceScheduleResponse.ServiceSchedule
		assert.NotNil(t, serviceScheduleResult)
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.CarrierServiceName, "B_carrier_service_name", "they should be equal")
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.CarrierServiceCode, "B_HLC", "they should be equal")
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.UniversalServiceReference, "SR00002B", "they should be equal")
		assert.NotNil(t, serviceScheduleResult.CrUpdTime.CreatedAt)
	}
}

func TestServiceScheduleService_GetServiceScheduleByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceScheduleService := NewServiceScheduleService(log, dbService, redisService, userServiceClient)

	serviceSchedule1, err := GetServiceSchedule(uint32(1), []byte{215, 99, 245, 80, 31, 119, 72, 225, 175, 83, 78, 3, 141, 96, 74, 211}, "d763f550-1f77-48e1-af53-4e038d604ad3", "B_carrier_service_name", "B_HLC", "SR00002B", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	serviceScheduleResp := ovsproto.GetServiceScheduleByPkResponse{}
	serviceScheduleResp.ServiceSchedule = serviceSchedule1

	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(1)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := ovsproto.GetServiceScheduleByPkRequest{}
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *ovsproto.GetServiceScheduleByPkRequest
	}
	tests := []struct {
		ss      *ServiceScheduleService
		args    args
		want    *ovsproto.GetServiceScheduleByPkResponse
		wantErr bool
	}{
		{
			ss: serviceScheduleService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &serviceScheduleResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		serviceScheduleResponse, err := tt.ss.GetServiceScheduleByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceScheduleService.GetServiceScheduleByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(serviceScheduleResponse, tt.want) {
			t.Errorf("ServiceScheduleService.GetServiceScheduleByPk() = %v, want %v", serviceScheduleResponse, tt.want)
		}
		serviceScheduleResult := serviceScheduleResponse.ServiceSchedule
		assert.NotNil(t, serviceScheduleResult)
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.CarrierServiceName, "B_carrier_service_name", "they should be equal")
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.CarrierServiceCode, "B_HLC", "they should be equal")
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.UniversalServiceReference, "SR00002B", "they should be equal")
		assert.NotNil(t, serviceScheduleResult.CrUpdTime.CreatedAt)
	}
}

func TestServiceScheduleService_GetServiceScheduleByUniversalServiceReference(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceScheduleService := NewServiceScheduleService(log, dbService, redisService, userServiceClient)

	serviceSchedule, err := GetServiceSchedule(uint32(1), []byte{215, 99, 245, 80, 31, 119, 72, 225, 175, 83, 78, 3, 141, 96, 74, 211}, "d763f550-1f77-48e1-af53-4e038d604ad3", "B_carrier_service_name", "B_HLC", "SR00002B", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	serviceScheduleResp := ovsproto.GetServiceScheduleByUniversalServiceReferenceResponse{}
	serviceScheduleResp.ServiceSchedule = serviceSchedule

	form := ovsproto.GetServiceScheduleByUniversalServiceReferenceRequest{}
	form.UniversalServiceReference = "SR00002B"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx   context.Context
		inReq *ovsproto.GetServiceScheduleByUniversalServiceReferenceRequest
	}
	tests := []struct {
		ss      *ServiceScheduleService
		args    args
		want    *ovsproto.GetServiceScheduleByUniversalServiceReferenceResponse
		wantErr bool
	}{
		{
			ss: serviceScheduleService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &serviceScheduleResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		serviceScheduleResponse, err := tt.ss.GetServiceScheduleByUniversalServiceReference(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceScheduleService.GetServiceScheduleByUniversalServiceReference() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(serviceScheduleResponse, tt.want) {
			t.Errorf("ServiceScheduleService.GetServiceScheduleByUniversalServiceReference() = %v, want %v", serviceScheduleResponse, tt.want)
		}
		serviceScheduleResult := serviceScheduleResponse.ServiceSchedule
		assert.NotNil(t, serviceScheduleResult)
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.CarrierServiceName, "B_carrier_service_name", "they should be equal")
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.CarrierServiceCode, "B_HLC", "they should be equal")
		assert.Equal(t, serviceScheduleResult.ServiceScheduleD.UniversalServiceReference, "SR00002B", "they should be equal")
		assert.NotNil(t, serviceScheduleResult.CrUpdTime.CreatedAt)
	}
}

func TestServiceScheduleService_UpdateServiceScheduleByUniversalServiceReference(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceScheduleService := NewServiceScheduleService(log, dbService, redisService, userServiceClient)

	form := ovsproto.UpdateServiceScheduleByUniversalServiceReferenceRequest{}
	form.CarrierServiceName = "C_carrier_service_name"
	form.CarrierServiceCode = "C_HLC"
	form.UniversalServiceReference = "SR00002B"

	updateResponse := ovsproto.UpdateServiceScheduleByUniversalServiceReferenceResponse{}

	type args struct {
		ctx context.Context
		in  *ovsproto.UpdateServiceScheduleByUniversalServiceReferenceRequest
	}
	tests := []struct {
		ss      *ServiceScheduleService
		args    args
		want    *ovsproto.UpdateServiceScheduleByUniversalServiceReferenceResponse
		wantErr bool
	}{
		{
			ss: serviceScheduleService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.ss.UpdateServiceScheduleByUniversalServiceReference(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceScheduleService.UpdateServiceScheduleByUniversalServiceReference() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ServiceScheduleService.UpdateServiceScheduleByUniversalServiceReference() = %v, want %v", got, tt.want)
		}
	}
}

func GetServiceSchedule(id uint32, uuid4 []byte, idS string, carrierServiceName string, carrierServiceCode string, universalServiceReference string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*ovsproto.ServiceSchedule, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	serviceScheduleD := new(ovsproto.ServiceScheduleD)
	serviceScheduleD.Id = id
	serviceScheduleD.Uuid4 = uuid4
	serviceScheduleD.IdS = idS
	serviceScheduleD.CarrierServiceName = carrierServiceName
	serviceScheduleD.CarrierServiceCode = carrierServiceCode
	serviceScheduleD.UniversalServiceReference = universalServiceReference

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	serviceSchedule := ovsproto.ServiceSchedule{ServiceScheduleD: serviceScheduleD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &serviceSchedule, nil
}
