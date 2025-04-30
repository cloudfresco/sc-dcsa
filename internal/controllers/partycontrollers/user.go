package partycontrollers

import (
	"encoding/json"
	"net/http"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)

// UserController - used for
type UserController struct {
	log               *zap.Logger
	UserServiceClient partyproto.UserServiceClient
	wfHelper          common.WfHelper
	workflowClient    client.Client
	ServerOpt         *config.ServerOptions
}

// NewUserController - Used to create a users handler
func NewUserController(log *zap.Logger, s partyproto.UserServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *UserController {
	return &UserController{
		log:               log,
		UserServiceClient: s,
		wfHelper:          wfHelper,
		workflowClient:    workflowClient,
		ServerOpt:         serverOpt,
	}
}

// GetUsers - Get Users
func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"users:read"}, uc.ServerOpt.Auth0Audience, uc.ServerOpt.Auth0Domain, uc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	users, err := uc.UserServiceClient.GetUsers(ctx, &partyproto.GetUsersRequest{UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		uc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, users)
}

// GetUser - Get User Details
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"users:read"}, uc.ServerOpt.Auth0Audience, uc.ServerOpt.Auth0Domain, uc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	usr, err := uc.UserServiceClient.GetUser(ctx, &partyproto.GetUserRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		uc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1303", err.Error(), 400, user.RequestId)
		return
	}

	common.RenderJSON(w, usr)
}

// ChangePassword - Changes Password
func (uc *UserController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"users:cud"}, uc.ServerOpt.Auth0Audience, uc.ServerOpt.Auth0Domain, uc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	form := partyproto.ChangePasswordRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		uc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1306", err.Error(), 402, user.RequestId)
		return
	}
	form.UserEmail = user.Email
	form.RequestId = user.RequestId
	_, err = uc.UserServiceClient.ChangePassword(ctx, &form)
	if err != nil {
		uc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1307", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, "We've just sent you an email to reset your password.")
}

// GetUserByEmail - Get User By email
func (uc *UserController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"users:read"}, uc.ServerOpt.Auth0Audience, uc.ServerOpt.Auth0Domain, uc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	form := partyproto.GetUserByEmailRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		uc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1308", err.Error(), 402, user.RequestId)
		return
	}
	form.UserEmail = user.Email
	form.RequestId = user.RequestId
	usr, err := uc.UserServiceClient.GetUserByEmail(ctx, &partyproto.GetUserByEmailRequest{Email: form.Email, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		uc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1309", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, usr)
}

// UpdateUser - Update User
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"users:cud"}, uc.ServerOpt.Auth0Audience, uc.ServerOpt.Auth0Domain, uc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	form := partyproto.UpdateUserRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		uc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1310", err.Error(), 402, user.RequestId)
		return
	}
	form.Id = id
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId
	_, err = uc.UserServiceClient.UpdateUser(ctx, &form)
	if err != nil {
		uc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Int("msgnum", 1311), zap.Error(err))
		common.RenderErrorJSON(w, "1311", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, "Updated Successfully")
}

// DeleteUser - delete user
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"users:cud"}, uc.ServerOpt.Auth0Audience, uc.ServerOpt.Auth0Domain, uc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	_, err = uc.UserServiceClient.DeleteUser(ctx, &partyproto.DeleteUserRequest{UserId: id, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		uc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Int("msgnum", 1312), zap.Error(err))
		common.RenderErrorJSON(w, "1312", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, "Deleted Successfully")
}
