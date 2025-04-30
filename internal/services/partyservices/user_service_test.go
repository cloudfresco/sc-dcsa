package partyservices

import (
	"context"
	"reflect"
	"testing"

	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUsers(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	userService := NewUserService(log, dbService, redisService, mailerService, jwtOpt, userOpt, serverOpt)

	user1, err := GetUser("auth0|66fd06d0bfea78a82bb42459", "sprov300@gmail.com", "https://s.gravatar.com/avatar/52ab1cc37bb42deb67ea939fd68ff7d4?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsp.png", "sprov300@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}

	user2, err := GetUser("auth0|66fcdfb6d20dcb68e3fcbc3b", "sprov200@gmail.com", "https://s.gravatar.com/avatar/06004bcbe9705b0ba5d7c4923fef0061?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsp.png", "sprov200@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}

	users := []*partyproto.User{}
	users = append(users, user1)
	users = append(users, user2)

	usersResponse := partyproto.GetUsersResponse{}
	usersResponse.Users = users

	form := partyproto.GetUsersRequest{}
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *partyproto.GetUsersRequest
	}
	tests := []struct {
		u       *UserService
		args    args
		want    *partyproto.GetUsersResponse
		wantErr bool
	}{
		{
			u: userService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &usersResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		usersResp, err := tt.u.GetUsers(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("UserService.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		assert.NotNil(t, usersResp)
		userResult := usersResp.Users[0]
		assert.Equal(t, userResult.Name, "sprov300@gmail.com", "they should be equal")
		assert.Equal(t, userResult.Picture, "https://s.gravatar.com/avatar/52ab1cc37bb42deb67ea939fd68ff7d4?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsp.png", "they should be equal")
		assert.Equal(t, userResult.Email, "sprov300@gmail.com", "they should be equal")
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	userService := NewUserService(log, dbService, redisService, mailerService, jwtOpt, userOpt, serverOpt)

	user, err := GetUser("auth0|66fd06d0bfea78a82bb42459", "sprov300@gmail.com", "https://s.gravatar.com/avatar/52ab1cc37bb42deb67ea939fd68ff7d4?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsp.png", "sprov300@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}

	userResponse := partyproto.GetUserByEmailResponse{}
	userResponse.User = user

	form := partyproto.GetUserByEmailRequest{}
	form.Email = "sprov300@gmail.com"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *partyproto.GetUserByEmailRequest
	}
	tests := []struct {
		u       *UserService
		args    args
		want    *partyproto.GetUserByEmailResponse
		wantErr bool
	}{
		{
			u: userService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &userResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		userResp, err := tt.u.GetUserByEmail(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("UserService.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(userResp, tt.want) {
			t.Errorf("UserService.GetUserByEmail() = %v, want %v", userResp, tt.want)
		}
		assert.NotNil(t, userResp)
		userResult := userResp.User
		assert.Equal(t, userResult.Name, "sprov300@gmail.com", "they should be equal")
		assert.Equal(t, userResult.Picture, "https://s.gravatar.com/avatar/52ab1cc37bb42deb67ea939fd68ff7d4?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsp.png", "they should be equal")
		assert.Equal(t, userResult.Email, "sprov300@gmail.com", "they should be equal")
	}
}

func TestUserService_GetUser(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	userService := NewUserService(log, dbService, redisService, mailerService, jwtOpt, userOpt, serverOpt)

	user, err := GetUser("auth0|66fd06d0bfea78a82bb42459", "sprov300@gmail.com", "https://s.gravatar.com/avatar/52ab1cc37bb42deb67ea939fd68ff7d4?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsp.png", "sprov300@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}

	userResponse := partyproto.GetUserResponse{}
	userResponse.User = user

	form := partyproto.GetUserRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "auth0|66fd06d0bfea78a82bb42459"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx context.Context
		in  *partyproto.GetUserRequest
	}
	tests := []struct {
		u       *UserService
		args    args
		want    *partyproto.GetUserResponse
		wantErr bool
	}{
		{
			u: userService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &userResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		userResp, err := tt.u.GetUser(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("UserService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(userResp, tt.want) {
			t.Errorf("UserService.GetUser() = %v, want %v", userResp, tt.want)
		}
		assert.NotNil(t, userResp)
		userResult := userResp.User
		assert.Equal(t, userResult.Name, "sprov300@gmail.com", "they should be equal")
		assert.Equal(t, userResult.Picture, "https://s.gravatar.com/avatar/52ab1cc37bb42deb67ea939fd68ff7d4?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsp.png", "they should be equal")
		assert.Equal(t, userResult.Email, "sprov300@gmail.com", "they should be equal")
	}
}

func GetUser(id string, email string, picture string, name string) (*partyproto.User, error) {
	user := new(partyproto.User)
	user.Id = id
	user.Email = email
	user.Picture = picture
	user.Name = name
	return user, nil
}
