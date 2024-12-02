package jitservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestServiceService_CreateService(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceService := NewServiceService(log, dbService, redisService, userServiceClient)

	service := jitproto.CreateServiceRequest{}
	service.CarrierId = uint32(5)
	service.CarrierServiceName = "A_carrier_service_name"
	service.CarrierServiceCode = "A_CSC"
	service.TradelaneId = ""
	service.UniversalServiceReference = "SR00001D"
	service.UserId = "auth0|66fd06d0bfea78a82bb42459"
	service.UserEmail = "sprov300@gmail.com"
	service.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *jitproto.CreateServiceRequest
	}
	tests := []struct {
		ss      *ServiceService
		args    args
		want    *jitproto.CreateServiceResponse
		wantErr bool
	}{
		{
			ss: serviceService,
			args: args{
				ctx: ctx,
				in:  &service,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		serviceResp, err := tt.ss.CreateService(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceService.CreateService() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		serviceResult := serviceResp.Service1
		assert.NotNil(t, serviceResult)
		assert.Equal(t, serviceResult.ServiceD.CarrierServiceName, "A_carrier_service_name", "they should be equal")
		assert.Equal(t, serviceResult.ServiceD.CarrierServiceCode, "A_CSC", "they should be equal")
		assert.Equal(t, serviceResult.ServiceD.UniversalServiceReference, "SR00001D", "they should be equal")
		assert.NotNil(t, serviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, serviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestServiceService_GetServices(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceService := NewServiceService(log, dbService, redisService, userServiceClient)

	service1, err := GetService(uint32(3), []byte{242, 106, 201, 13, 200, 154, 75, 255, 159, 211, 53, 193, 52, 163, 236, 49}, "f26ac90d-c89a-4bff-9fd3-35c134a3ec31", uint32(5), "B_carrier_service_name_1", "B_HLC", "", "SR00003H", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	service2, err := GetService(uint32(2), []byte{246, 80, 34, 241, 118, 231, 76, 242, 130, 135, 36, 28, 215, 174, 212, 222}, "f65022f1-76e7-4cf2-8287-241cd7aed4de", uint32(5), "B_carrier_service_name", "B_HLC", "", "SR00002B", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	services := []*jitproto.Service{}
	services = append(services, service1, service2)

	form := jitproto.GetServicesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MQ=="
	serviceResp := jitproto.GetServicesResponse{Services: services, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *jitproto.GetServicesRequest
	}
	tests := []struct {
		ss      *ServiceService
		args    args
		want    *jitproto.GetServicesResponse
		wantErr bool
	}{
		{
			ss: serviceService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &serviceResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		serviceResponse, err := tt.ss.GetServices(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceService.GetServices() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(serviceResponse, tt.want) {
			t.Errorf("ServiceService.GetServices() = %v, want %v", serviceResponse, tt.want)
		}
		serviceResult := serviceResponse.Services[0]
		assert.NotNil(t, serviceResult)
		assert.Equal(t, serviceResult.ServiceD.CarrierServiceName, "B_carrier_service_name_1", "they should be equal")
		assert.Equal(t, serviceResult.ServiceD.CarrierServiceCode, "B_HLC", "they should be equal")
		assert.Equal(t, serviceResult.ServiceD.UniversalServiceReference, "SR00003H", "they should be equal")
		assert.NotNil(t, serviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, serviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestServiceService_GetService(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceService := NewServiceService(log, dbService, redisService, userServiceClient)

	service1, err := GetService(uint32(2), []byte{246, 80, 34, 241, 118, 231, 76, 242, 130, 135, 36, 28, 215, 174, 212, 222}, "f65022f1-76e7-4cf2-8287-241cd7aed4de", uint32(5), "B_carrier_service_name", "B_HLC", "", "SR00002B", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	serviceResp := jitproto.GetServiceResponse{}
	serviceResp.Service1 = service1

	gform := commonproto.GetRequest{}
	gform.Id = "f65022f1-76e7-4cf2-8287-241cd7aed4de"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := jitproto.GetServiceRequest{}
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *jitproto.GetServiceRequest
	}
	tests := []struct {
		ss      *ServiceService
		args    args
		want    *jitproto.GetServiceResponse
		wantErr bool
	}{
		{
			ss: serviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &serviceResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		serviceResponse, err := tt.ss.GetService(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceService.GetService() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(serviceResponse, tt.want) {
			t.Errorf("ServiceService.GetService() = %v, want %v", serviceResponse, tt.want)
		}

		serviceResult := serviceResponse.Service1
		assert.NotNil(t, serviceResult)
		assert.Equal(t, serviceResult.ServiceD.CarrierServiceName, "B_carrier_service_name", "they should be equal")
		assert.Equal(t, serviceResult.ServiceD.CarrierServiceCode, "B_HLC", "they should be equal")
		assert.Equal(t, serviceResult.ServiceD.UniversalServiceReference, "SR00002B", "they should be equal")
		assert.NotNil(t, serviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, serviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestServiceService_GetServiceByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceService := NewServiceService(log, dbService, redisService, userServiceClient)
	service1, err := GetService(uint32(3), []byte{242, 106, 201, 13, 200, 154, 75, 255, 159, 211, 53, 193, 52, 163, 236, 49}, "f26ac90d-c89a-4bff-9fd3-35c134a3ec31", uint32(5), "B_carrier_service_name_1", "B_HLC", "", "SR00003H", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	serviceResp := jitproto.GetServiceByPkResponse{}
	serviceResp.Service1 = service1

	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(3)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := jitproto.GetServiceByPkRequest{}
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *jitproto.GetServiceByPkRequest
	}
	tests := []struct {
		ss      *ServiceService
		args    args
		want    *jitproto.GetServiceByPkResponse
		wantErr bool
	}{
		{
			ss: serviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &serviceResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		serviceResponse, err := tt.ss.GetServiceByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceService.GetServiceByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(serviceResponse, tt.want) {
			t.Errorf("ServiceService.GetServiceByPk() = %v, want %v", serviceResponse, tt.want)
		}
		serviceResult := serviceResponse.Service1
		assert.NotNil(t, serviceResult)
		assert.Equal(t, serviceResult.ServiceD.CarrierServiceName, "B_carrier_service_name_1", "they should be equal")
		assert.Equal(t, serviceResult.ServiceD.CarrierServiceCode, "B_HLC", "they should be equal")
		assert.Equal(t, serviceResult.ServiceD.UniversalServiceReference, "SR00003H", "they should be equal")
		assert.NotNil(t, serviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, serviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestServiceService_FindByCarrierServiceCode(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	serviceService := NewServiceService(log, dbService, redisService, userServiceClient)
	service1, err := GetService(uint32(1), []byte{3, 72, 34, 150, 239, 156, 17, 235, 154, 3, 2, 66, 172, 19, 25, 153}, "03482296-ef9c-11eb-9a03-0242ac131999", uint32(5), "A_carrier_service_name", "A_CSC", "", "SR00001D", "2020-03-07T12:12:12Z", "2020-04-07T12:12:12Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	serviceResp := jitproto.FindByCarrierServiceCodeResponse{}
	serviceResp.Service1 = service1

	form := jitproto.FindByCarrierServiceCodeRequest{}
	form.CarrierServiceCode = "A_CSC"
	form.UserId = "auth0|66fd06d0bfea78a82bb42459"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *jitproto.FindByCarrierServiceCodeRequest
	}
	tests := []struct {
		ss      *ServiceService
		args    args
		want    *jitproto.FindByCarrierServiceCodeResponse
		wantErr bool
	}{
		{
			ss: serviceService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &serviceResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		serviceResponse, err := tt.ss.FindByCarrierServiceCode(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ServiceService.FindByCarrierServiceCode() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(serviceResponse, tt.want) {
			t.Errorf("ServiceService.FindByCarrierServiceCode() = %v, want %v", serviceResponse, tt.want)
		}
		serviceResult := serviceResponse.Service1
		assert.NotNil(t, serviceResult)
		assert.Equal(t, serviceResult.ServiceD.CarrierServiceName, "A_carrier_service_name", "they should be equal")
		assert.Equal(t, serviceResult.ServiceD.CarrierServiceCode, "A_CSC", "they should be equal")
		assert.Equal(t, serviceResult.ServiceD.UniversalServiceReference, "SR00001D", "they should be equal")
		assert.NotNil(t, serviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, serviceResult.CrUpdTime.UpdatedAt)
	}
}

func GetService(id uint32, uuid4 []byte, idS string, carrierId uint32, carrierServiceName string, carrierServiceCode string, tradelaneId string, universalServiceReference string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*jitproto.Service, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	serviceD := new(jitproto.ServiceD)
	serviceD.Id = id
	serviceD.Uuid4 = uuid4
	serviceD.IdS = idS
	serviceD.CarrierId = carrierId
	serviceD.CarrierServiceName = carrierServiceName
	serviceD.CarrierServiceCode = carrierServiceCode
	serviceD.TradelaneId = tradelaneId
	serviceD.UniversalServiceReference = universalServiceReference

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	service := jitproto.Service{ServiceD: serviceD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &service, nil
}
