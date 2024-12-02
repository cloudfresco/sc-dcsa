// https://github.com/dcsaorg/DCSA-BKG/blob/master/src/main/java/org/dcsa/bkg/controller/BKGController.java

package bkgcontrollers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/bkg/v2"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	bkgworkflows "github.com/cloudfresco/sc-dcsa/internal/workflows/bkgworkflows"

	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// BkgController - Create Bkg Controller
type BkgController struct {
	log               *zap.Logger
	UserServiceClient partyproto.UserServiceClient
	BkgServiceClient  bkgproto.BkgServiceClient
	wfHelper          common.WfHelper
	workflowClient    client.Client
}

// NewBkgController - Create Bkg Handler
func NewBkgController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, bkgServiceClient bkgproto.BkgServiceClient, wfHelper common.WfHelper, workflowClient client.Client) *BkgController {
	return &BkgController{
		log:               log,
		UserServiceClient: userServiceClient,
		BkgServiceClient:  bkgServiceClient,
		wfHelper:          wfHelper,
		workflowClient:    workflowClient,
	}
}

// ServeHTTP - parse url and call controller action
func (b *BkgController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := common.GetAuthData(r)
	cdata := partyproto.GetAuthUserDetailsRequest{}
	cdata.TokenString = data.TokenString
	cdata.Email = data.Email

	md := metadata.Pairs("authorization", "Bearer "+cdata.TokenString)
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	user, err := b.UserServiceClient.GetAuthUserDetails(ctx, &cdata)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}
	pathParts, queryString, err := common.ParseURL(r.URL.String())
	if err != nil {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}

	switch r.Method {
	case http.MethodGet:
		b.processGet(ctx, w, r, user, pathParts, queryString)
	case http.MethodPost:
		b.processPost(ctx, w, r, user, pathParts, data.TokenString)
	case http.MethodPut:
		b.processPut(ctx, w, r, user, pathParts, data.TokenString)
	case http.MethodPatch:
		b.processPatch(ctx, w, r, user, pathParts, data.TokenString)
	case http.MethodDelete:
		b.processDelete(ctx, w, r, user, pathParts, data.TokenString)
	default:
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processGet - Parse URL for all the GET paths and call the controller action
/*
	GET    "/v2/bookings/{carrierBookingRequestReference}"
*/

func (b *BkgController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
	if len(pathParts) == 3 {
		// GET    "/bookings/:id"
		b.GetBookingByCarrierBookingRequestReference(ctx, w, r, pathParts[2], user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPost - Parse URL for all the POST paths and call the controller action
/*
	POST "/v2/bookings"

*/
func (b *BkgController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
	if (len(pathParts) == 2) && (pathParts[1] == "bookings") {
		b.CreateBooking(ctx, w, r, user, tokenString)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPut - Parse URL for all the put paths and call the controller action
/*
PUT   "/v2/bookings/{carrierBookingRequestReference}" 3
*/

func (b *BkgController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
	if len(pathParts) == 3 {
		b.UpdateBooking(ctx, w, r, pathParts[2], user, tokenString)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPatch - Parse URL for all the patch paths and call the controller action
/*
 PATCH  "/v2/bookings/{carrierBookingRequestReference}"
*/

func (b *BkgController) processPatch(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
	if len(pathParts) == 3 {
		b.CancelBookingByCarrierBookingRequestReference(ctx, w, r, pathParts[2], user, tokenString)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processDelete - Parse URL for all the delete paths and call the controller action
/*
 DELETE   ""
*/

func (b *BkgController) processDelete(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, tokenString string) {
}

// CreateBooking - Create Booking Header
func (b *BkgController) CreateBooking(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, tokenString string) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        bkgworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := bkgproto.CreateBookingRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
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
	result := wHelper.StartWorkflow(workflowOptions, bkgworkflows.CreateBookingWorkflow, &form, tokenString, user, b.log)
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
func (b *BkgController) GetBookingByCarrierBookingRequestReference(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
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
}

// UpdateBooking - Update Booking
func (b *BkgController) UpdateBooking(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse, tokenString string) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        bkgworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
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
	result := wHelper.StartWorkflow(workflowOptions, bkgworkflows.UpdateBookingWorkflow, &form, tokenString, user, b.log)
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
func (b *BkgController) CancelBookingByCarrierBookingRequestReference(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse, tokenString string) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dcsa_" + uuid.New(),
		TaskList:                        bkgworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	form := bkgproto.CancelBookingByCarrierBookingReferenceRequest{CarrierBookingRequestReference: id, UserEmail: user.Email, RequestId: user.RequestId}
	wHelper := b.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, bkgworkflows.CancelBookingByCarrierBookingReferenceWorkflow, &form, tokenString, user, b.log)
	workflowClient := b.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var response bkgproto.CancelBookingByCarrierBookingReferenceResponse
	err := workflowRun.Get(ctx, &response)
	if err != nil {
		b.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, response)
}
