// https://github.com/dcsaorg/DCSA-BKG/blob/master/src/main/java/org/dcsa/bkg/controller/BKGController.java

package bkgcontrollers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	bkgworkflows "github.com/cloudfresco/sc-dcsa/internal/workflows/bkgworkflows"

	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)

// BkgController - Create Bkg Controller
type BkgController struct {
	log               *zap.Logger
	UserServiceClient partyproto.UserServiceClient
	BkgServiceClient  bkgproto.BkgServiceClient
	wfHelper          common.WfHelper
	workflowClient    client.Client
	ServerOpt         *config.ServerOptions
}

// NewBkgController - Create Bkg Handler
func NewBkgController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, bkgServiceClient bkgproto.BkgServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *BkgController {
	return &BkgController{
		log:               log,
		UserServiceClient: userServiceClient,
		BkgServiceClient:  bkgServiceClient,
		wfHelper:          wfHelper,
		workflowClient:    workflowClient,
		ServerOpt:         serverOpt,
	}
}

// CreateBooking - Create Booking Header
func (b *BkgController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"bkg:cud"}, b.ServerOpt.Auth0Audience, b.ServerOpt.Auth0Domain, b.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        bkgworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := bkgproto.CreateBookingRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		b.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	v, err := protovalidate.New()
	if err != nil {
		b.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	if err = v.Validate(&form); err != nil {
		b.log.Error("Validation failed", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	wHelper := b.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, bkgworkflows.CreateBookingWorkflow, &form, token, user, b.log)
	workflowClient := b.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var booking bkgproto.CreateBookingResponse
	err = workflowRun.Get(ctx, &booking)
	if err != nil {
		b.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1310", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, booking)
}

// GetBookingByCarrierBookingRequestReference - GetBookingByCarrierBookingRequestReference Booking
func (b *BkgController) GetBookingByCarrierBookingRequestReference(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"bkg:read"}, b.ServerOpt.Auth0Audience, b.ServerOpt.Auth0Domain, b.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("carrierBookingRequestReference")

	booking, err := b.BkgServiceClient.GetBookingByCarrierBookingRequestReference(ctx, &bkgproto.GetBookingByCarrierBookingRequestReferenceRequest{CarrierBookingRequestReference: id, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		b.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, booking)
}

// UpdateBooking - Update Booking
func (b *BkgController) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"bkg:cud"}, b.ServerOpt.Auth0Audience, b.ServerOpt.Auth0Domain, b.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("carrierBookingRequestReference")

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        bkgworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		b.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	form.CarrierBookingRequestReference = id
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	v, err := protovalidate.New()
	if err != nil {
		b.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	if err = v.Validate(&form); err != nil {
		b.log.Error("Validation failed", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	wHelper := b.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, bkgworkflows.UpdateBookingWorkflow, &form, token, user, b.log)
	workflowClient := b.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var response string
	err = workflowRun.Get(ctx, &response)
	if err != nil {
		b.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, response)
}

// CancelBookingByCarrierBookingRequestReference - CancelBookingByCarrierBookingReference
func (b *BkgController) CancelBookingByCarrierBookingRequestReference(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"bkg:cud"}, b.ServerOpt.Auth0Audience, b.ServerOpt.Auth0Domain, b.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("carrierBookingRequestReference")

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        bkgworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	form := bkgproto.CancelBookingByCarrierBookingReferenceRequest{CarrierBookingRequestReference: id, UserEmail: user.Email, RequestId: user.RequestId}
	wHelper := b.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, bkgworkflows.CancelBookingByCarrierBookingReferenceWorkflow, &form, token, user, b.log)
	workflowClient := b.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var response bkgproto.CancelBookingByCarrierBookingReferenceResponse
	err = workflowRun.Get(ctx, &response)
	if err != nil {
		b.log.Error("Error", zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, response)
}
