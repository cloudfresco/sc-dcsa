package partyservices

import (
	"context"

	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// insertPartyFunctionSQL - insert PartyFunctionSQL query
const insertPartyFunctionSQL = `insert into party_functions
	  ( 
  party_function_code,
  party_function_name,
  party_function_description
  )
  values (
  :party_function_code,
  :party_function_name,
  :party_function_description);`

// selectPartyFunctionsSQL - select PartyFunctionsSQL query
/*const selectPartyFunctionsSQL = `select
  id,
  party_function_code,
  party_function_name,
  party_function_description from party_functions`*/

// CreatePartyFunction - CreatePartyFunction
func (ps *PartyService) CreatePartyFunction(ctx context.Context, in *partyproto.CreatePartyFunctionRequest) (*partyproto.CreatePartyFunctionResponse, error) {
	partyFunction, err := ps.ProcessPartyFunctionRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertPartyFunction(ctx, insertPartyFunctionSQL, partyFunction, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	partyFunctionResponse := partyproto.CreatePartyFunctionResponse{}
	partyFunctionResponse.PartyFunction = partyFunction
	return &partyFunctionResponse, nil
}

// ProcessPartyFunctionRequest - Process PartyFunctionRequest
func (ps *PartyService) ProcessPartyFunctionRequest(ctx context.Context, in *partyproto.CreatePartyFunctionRequest) (*partyproto.PartyFunction, error) {
	partyFunction := partyproto.PartyFunction{}

	partyFunction.PartyFunctionCode = in.PartyFunctionCode
	partyFunction.PartyFunctionName = in.PartyFunctionName
	partyFunction.PartyFunctionDescription = in.PartyFunctionDescription

	return &partyFunction, nil
}

// insertPartyFunction - Insert Document Party
func (ps *PartyService) insertPartyFunction(ctx context.Context, insertPartyFunctionSQL string, partyFunction *partyproto.PartyFunction, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertPartyFunctionSQL, partyFunction)
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
