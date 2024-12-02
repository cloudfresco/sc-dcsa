package jitservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestTimestampService_CreateTimestamp(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	timestampService := NewTimestampService(log, dbService, redisService, userServiceClient)
	tstamp := jitproto.CreateTimestampRequest{}
	tstamp.EventTypeCode = "ARRI"
	tstamp.EventClassifierCode = "ACT"
	tstamp.DelayReasonCode = "ANA"
	tstamp.ChangeRemark = "Authorities not available"
	tstamp.EventDateTime = "11/09/2023"
	tstamp.UserId = "auth0|66fd06d0bfea78a82bb42459"
	tstamp.UserEmail = "sprov300@gmail.com"
	tstamp.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *jitproto.CreateTimestampRequest
	}
	tests := []struct {
		ts      *TimestampService
		args    args
		want    *jitproto.CreateTimestampResponse
		wantErr bool
	}{
		{
			ts: timestampService,
			args: args{
				ctx: ctx,
				in:  &tstamp,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tstampResp, err := tt.ts.CreateTimestamp(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("TimestampService.CreateTimestamp() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		tstampResult := tstampResp.Timestamp1
		assert.NotNil(t, tstampResult)
		assert.Equal(t, tstampResult.TimestampD.EventTypeCode, "ARRI", "they should be equal")
		assert.Equal(t, tstampResult.TimestampD.EventClassifierCode, "ACT", "they should be equal")
		assert.Equal(t, tstampResult.TimestampD.DelayReasonCode, "ANA", "they should be equal")
		assert.Equal(t, tstampResult.TimestampD.ChangeRemark, "Authorities not available", "they should be equal")
	}
}

func TestTimestampService_GetTimestamps(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	timestampService := NewTimestampService(log, dbService, redisService, userServiceClient)
	tstamp1, err := GetTimestamp(uint32(2), []byte{68, 221, 35, 196, 84, 120, 67, 86, 163, 40, 66, 22, 91, 149, 209, 41}, "TTA", "PLN", "WEA", "Bad weather", "2020-03-07T12:12:12Z")
	if err != nil {
		t.Error(err)
		return
	}

	tstamp2, err := GetTimestamp(uint32(1), []byte{79, 106, 3, 103, 46, 50, 70, 238, 188, 79, 162, 243, 173, 192, 63, 102}, "ARRI", "ACT", "ANA", "Authorities not available", "2020-03-07T12:12:12Z")
	if err != nil {
		t.Error(err)
		return
	}

	tstamps := []*jitproto.Timestamp{}
	tstamps = append(tstamps, tstamp1, tstamp2)

	form := jitproto.GetTimestampsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	tstampResponse := jitproto.GetTimestampsResponse{Timestamps: tstamps, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *jitproto.GetTimestampsRequest
	}
	tests := []struct {
		ts      *TimestampService
		args    args
		want    *jitproto.GetTimestampsResponse
		wantErr bool
	}{
		{
			ts: timestampService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &tstampResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tstampResp, err := tt.ts.GetTimestamps(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("TimestampService.GetTimestamps() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(tstampResp, tt.want) {
			t.Errorf("TimestampService.GetTimestamps() = %v, want %v", tstampResp, tt.want)
		}
		tstampResult := tstampResp.Timestamps[1]
		assert.NotNil(t, tstampResult)
		assert.Equal(t, tstampResult.TimestampD.EventTypeCode, "ARRI", "they should be equal")
		assert.Equal(t, tstampResult.TimestampD.EventClassifierCode, "ACT", "they should be equal")
		assert.Equal(t, tstampResult.TimestampD.DelayReasonCode, "ANA", "they should be equal")
		assert.Equal(t, tstampResult.TimestampD.ChangeRemark, "Authorities not available", "they should be equal")

	}
}

func GetTimestamp(id uint32, uuid4 []byte, eventTypeCode string, eventClassifierCode string, delayReasonCode string, changeRemark string, eventDateTime string) (*jitproto.Timestamp, error) {
	eventDateTime1, err := common.ConvertTimeToTimestamp(Layout, eventDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	timestampD := new(jitproto.TimestampD)
	timestampD.Id = id
	timestampD.Uuid4 = uuid4
	timestampD.EventTypeCode = eventTypeCode
	timestampD.EventClassifierCode = eventClassifierCode
	timestampD.DelayReasonCode = delayReasonCode
	timestampD.ChangeRemark = changeRemark

	timestampT := new(jitproto.TimestampT)
	timestampT.EventDateTime = eventDateTime1

	tstamp := jitproto.Timestamp{TimestampD: timestampD, TimestampT: timestampT}

	return &tstamp, nil
}
