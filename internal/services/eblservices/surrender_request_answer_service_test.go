package eblservices

import (
	"context"
	"reflect"
	"testing"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
)

func TestSurrenderRequestAnswerService_CreateSurrenderRequestAnswer(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestAnswerService := NewSurrenderRequestAnswerService(log, dbService, redisService, userServiceClient)

	surrenderRequestAnswer := eblproto.CreateSurrenderRequestAnswerRequest{}
	surrenderRequestAnswer.SurrenderRequestReference = "Z12345"
	surrenderRequestAnswer.Action = "SURR"
	surrenderRequestAnswer.Comments = "comments"
	surrenderRequestAnswer.CreatedDateTime = "03/22/2024"
	surrenderRequestAnswer.SurrenderRequestId = uint32(1)
	surrenderRequestAnswer.UserId = "auth0|66fd06d0bfea78a82bb42459"
	surrenderRequestAnswer.UserEmail = "sprov300@gmail.com"
	surrenderRequestAnswer.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *eblproto.CreateSurrenderRequestAnswerRequest
	}
	tests := []struct {
		ss      *SurrenderRequestAnswerService
		args    args
		want    *eblproto.CreateSurrenderRequestAnswerResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestAnswerService,
			args: args{
				ctx: ctx,
				in:  &surrenderRequestAnswer,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		surrenderRequestAnswerResp, err := tt.ss.CreateSurrenderRequestAnswer(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestAnswerService.CreateSurrenderRequestAnswer() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, surrenderRequestAnswerResp)
		surrenderRequestAnswerResult := surrenderRequestAnswerResp.SurrenderRequestAnswer
		assert.Equal(t, surrenderRequestAnswerResult.SurrenderRequestAnswerD.SurrenderRequestReference, "Z12345", "they should be equal")
		assert.Equal(t, surrenderRequestAnswerResult.SurrenderRequestAnswerD.Action, "SURR", "they should be equal")
	}
}

func TestSurrenderRequestAnswerService_UpdateSurrenderRequestAnswer(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	surrenderRequestAnswerService := NewSurrenderRequestAnswerService(log, dbService, redisService, userServiceClient)

	surrenderRequestAnswer := eblproto.UpdateSurrenderRequestAnswerRequest{}
	surrenderRequestAnswer.SurrenderRequestReference = "B12345"
	surrenderRequestAnswer.Action = "UURR"
	surrenderRequestAnswer.Comments = "commentsare"
	surrenderRequestAnswer.Id = "f69275f8-2d6e-49ed-a6d0-100c95cc0958"
	surrenderRequestAnswer.UserId = "auth0|66fd06d0bfea78a82bb42459"
	surrenderRequestAnswer.UserEmail = "sprov300@gmail.com"
	surrenderRequestAnswer.RequestId = "bks1m1g91jau4nkks2f0"

	updateResponse := eblproto.UpdateSurrenderRequestAnswerResponse{}

	type args struct {
		ctx context.Context
		in  *eblproto.UpdateSurrenderRequestAnswerRequest
	}
	tests := []struct {
		ss      *SurrenderRequestAnswerService
		args    args
		want    *eblproto.UpdateSurrenderRequestAnswerResponse
		wantErr bool
	}{
		{
			ss: surrenderRequestAnswerService,
			args: args{
				ctx: ctx,
				in:  &surrenderRequestAnswer,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.ss.UpdateSurrenderRequestAnswer(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("SurrenderRequestAnswerService.UpdateSurrenderRequestAnswer() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SurrenderRequestAnswerService.UpdateSurrenderRequestAnswer() = %v, want %v", got, tt.want)
		}
	}
}
