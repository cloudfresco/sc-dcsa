package partycontrollers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// PartyController - Create Party Controller
type PartyController struct {
	log                *zap.Logger
	PartyServiceClient partyproto.PartyServiceClient
	UserServiceClient  partyproto.UserServiceClient
}

// NewPartyController - Create Party Handler
func NewPartyController(log *zap.Logger, partyServiceClient partyproto.PartyServiceClient, userServiceClient partyproto.UserServiceClient) *PartyController {
	return &PartyController{
		log:                log,
		PartyServiceClient: partyServiceClient,
		UserServiceClient:  userServiceClient,
	}
}

// ServeHTTP - parse url and call controller action
func (pp *PartyController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := common.GetAuthData(r)
	cdata := partyproto.GetAuthUserDetailsRequest{}
	cdata.TokenString = data.TokenString
	cdata.Email = data.Email

	md := metadata.Pairs("authorization", "Bearer "+cdata.TokenString)
	ctx := metadata.NewOutgoingContext(r.Context(), md)
	user, err := pp.UserServiceClient.GetAuthUserDetails(ctx, &cdata)
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
		pp.processGet(ctx, w, r, user, pathParts, queryString)
	case http.MethodPost:
		pp.processPost(ctx, w, r, user, pathParts)
	case http.MethodPut:
		pp.processPut(ctx, w, r, user, pathParts)
	case http.MethodDelete:
		pp.processDelete(ctx, w, r, user, pathParts)
	default:
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processGet - Parse URL for all the GET paths and call the controller action
/*
 GET  "/v1/parties/"
 GET  "/v1/parties/{id}"
*/

func (pp *PartyController) processGet(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string, queryString url.Values) {
	if (len(pathParts) == 2) && (pathParts[1] == "parties") {
		limit := queryString.Get("limit")
		cursor := queryString.Get("cursor")
		pp.GetParties(ctx, w, r, limit, cursor, user)
	} else if len(pathParts) == 3 {
		if pathParts[1] == "parties" {
			pp.GetParty(ctx, w, r, pathParts[2], user)
		} else {
			common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
			return
		}
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPost - Parse URL for all the POST paths and call the controller action
/*
 POST  "/v1/parties/create/"
 POST  "/v1/parties/{id}/partyContactDetailcreate/"
*/
func (pp *PartyController) processPost(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
	if (len(pathParts) == 3) && (pathParts[1] == "parties") {
		if pathParts[2] == "create" {
			pp.CreateParty(ctx, w, r, user)
		} else {
			common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
			return
		}
	} else if (len(pathParts) == 4) && (pathParts[1] == "parties") {
		if pathParts[3] == "partyContactDetailcreate" {
			pp.CreatePartyContactDetail(ctx, w, r, pathParts[2], user)
		} else {
			common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
			return
		}
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processPut - Parse URL for all the put paths and call the controller action
/*
 PUT  "/v1/parties/{id}"
*/

func (pp *PartyController) processPut(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
	if (len(pathParts) == 3) && (pathParts[1] == "parties") {
		pp.UpdateParty(ctx, w, r, pathParts[2], user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// processDelete - Parse URL for all the delete paths and call the controller action
/*
 DELETE  "/v1/parties/{id}"
*/

func (pp *PartyController) processDelete(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse, pathParts []string) {
	if (len(pathParts) == 3) && (pathParts[1] == "parties") {
		pp.DeleteParty(ctx, w, r, pathParts[2], user)
	} else {
		common.RenderErrorJSON(w, "1000", "Invalid Request", 400, user.RequestId)
		return
	}
}

// GetParties - used to view all Parties
func (pp *PartyController) GetParties(ctx context.Context, w http.ResponseWriter, r *http.Request, limit string, cursor string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		parties, err := pp.PartyServiceClient.GetParties(ctx, &partyproto.GetPartiesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
		if err != nil {
			pp.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4016", err.Error(), 402, user.RequestId)
			return
		}
		common.RenderJSON(w, parties)
	}
}

// GetParty - used to view Party
func (pp *PartyController) GetParty(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		party, err := pp.PartyServiceClient.GetParty(ctx, &partyproto.GetPartyRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
		if err != nil {
			pp.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4016", err.Error(), 402, user.RequestId)
			return
		}

		common.RenderJSON(w, party)
	}
}

// CreateParty - used to Create Party
func (pp *PartyController) CreateParty(ctx context.Context, w http.ResponseWriter, r *http.Request, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		form := partyproto.CreatePartyRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&form)
		if err != nil {
			pp.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
			return
		}
		form.UserId = user.UserId
		form.UserEmail = user.Email
		form.RequestId = user.RequestId

		party, err := pp.PartyServiceClient.CreateParty(ctx, &form)
		if err != nil {
			pp.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4014", err.Error(), 402, user.RequestId)
			return
		}
		common.RenderJSON(w, party)
	}
}

// UpdateParty - Update Party
func (pp *PartyController) UpdateParty(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		form := partyproto.UpdatePartyRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&form)
		if err != nil {
			pp.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
			return
		}
		form.Id = id
		form.UserId = user.UserId
		form.UserEmail = user.Email
		form.RequestId = user.RequestId
		_, err = pp.PartyServiceClient.UpdateParty(ctx, &form)
		if err != nil {
			pp.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4016", err.Error(), 402, user.RequestId)
			return
		}
		common.RenderJSON(w, "Updated Successfully")
	}
}

// DeleteParty - delete Party
func (pp *PartyController) DeleteParty(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		_, err := pp.PartyServiceClient.DeleteParty(ctx, &partyproto.DeletePartyRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
		if err != nil {
			pp.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4016", err.Error(), 402, user.RequestId)
			return
		}
		common.RenderJSON(w, "Deleted Successfully")
	}
}

// CreatePartyContactDetail - used to Create Party Contact
func (pp *PartyController) CreatePartyContactDetail(ctx context.Context, w http.ResponseWriter, r *http.Request, id string, user *partyproto.GetAuthUserDetailsResponse) {
	select {
	case <-ctx.Done():
		common.RenderErrorJSON(w, "1002", "Client closed connection", 402, user.RequestId)
		return
	default:
		form := partyproto.CreatePartyContactDetailRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&form)
		if err != nil {
			pp.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4004", err.Error(), 402, user.RequestId)
			return
		}
		pcid, err := strconv.ParseUint((id), 10, 0)
		if err != nil {
			pp.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4004", err.Error(), 402, user.RequestId)
			return
		}
		form.UserId = user.UserId
		form.PartyId = uint32(pcid)
		form.UserEmail = user.Email
		form.RequestId = user.RequestId
		partyContactDetail, err := pp.PartyServiceClient.CreatePartyContactDetail(ctx, &form)
		if err != nil {
			pp.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
			common.RenderErrorJSON(w, "4016", err.Error(), 402, user.RequestId)
			return
		}
		common.RenderJSON(w, partyContactDetail)
	}
}
