package partyservices

import (
	"context"

	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// insertPartyIdentifyingCodeSQL - insert PartyIdentifyingCodeSQL query
const insertPartyIdentifyingCodeSQL = `insert into party_identifying_codes
	  ( 
  dcsa_responsible_agency_code,
  party_id,
  code_list_name,
  party_code
  )
  values (
  :dcsa_responsible_agency_code,
  :party_id,
  :code_list_name,
  :party_code);`

// selectPartyIdentifyingCodesSQL - select PartyIdentifyingCodesSQL query
/*const selectPartyIdentifyingCodesSQL = `select
  id,
  dcsa_responsible_agency_code,
  party_id,
  code_list_name,
  party_code from party_identifying_codes`*/

// CreatePartyIdentifyingCode - CreatePartyIdentifyingCode
func (ps *PartyService) CreatePartyIdentifyingCode(ctx context.Context, in *partyproto.CreatePartyIdentifyingCodeRequest) (*partyproto.CreatePartyIdentifyingCodeResponse, error) {
	partyIdentifyingCode, err := ps.ProcessPartyIdentifyingCodeRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertPartyIdentifyingCode(ctx, insertPartyIdentifyingCodeSQL, partyIdentifyingCode, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	partyIdentifyingCodeResponse := partyproto.CreatePartyIdentifyingCodeResponse{}
	partyIdentifyingCodeResponse.PartyIdentifyingCode = partyIdentifyingCode
	return &partyIdentifyingCodeResponse, nil
}

// ProcessPartyIdentifyingCodeRequest - Process PartyIdentifyingCodeRequest
func (ps *PartyService) ProcessPartyIdentifyingCodeRequest(ctx context.Context, in *partyproto.CreatePartyIdentifyingCodeRequest) (*partyproto.PartyIdentifyingCode, error) {
	partyIdentifyingCode := partyproto.PartyIdentifyingCode{}

	partyIdentifyingCode.DcsaResponsibleAgencyCode = in.DcsaResponsibleAgencyCode
	partyIdentifyingCode.PartyId = in.PartyId
	partyIdentifyingCode.CodeListName = in.CodeListName
	partyIdentifyingCode.PartyCode = in.PartyCode

	return &partyIdentifyingCode, nil
}

// insertPartyIdentifyingCode - Insert Document Party
func (ps *PartyService) insertPartyIdentifyingCode(ctx context.Context, insertPartyIdentifyingCodeSQL string, partyIdentifyingCode *partyproto.PartyIdentifyingCode, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertPartyIdentifyingCodeSQL, partyIdentifyingCode)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
