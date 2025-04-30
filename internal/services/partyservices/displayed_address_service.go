package partyservices

import (
	"context"

	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// insertDisplayedAddressSQL - insert DisplayedAddressSQL query
const insertDisplayedAddressSQL = `insert into displayed_addresses
	  ( 
  document_party_id,
  address_line_number,
  address_line_text
  )
  values (
  :document_party_id,
  :address_line_number,
  :address_line_text);`

// selectDisplayedAddressesSQL - select DisplayedAddressesSQL query
/*const selectDisplayedAddressesSQL = `select
  id,
  document_party_id,
  address_line_number,
  address_line_text from displayed_addresses`*/

// CreateDisplayedAddress - CreateDisplayedAddress
func (ps *PartyService) CreateDisplayedAddress(ctx context.Context, in *partyproto.CreateDisplayedAddressRequest) (*partyproto.CreateDisplayedAddressResponse, error) {
	displayedAddress, err := ps.ProcessDisplayedAddressRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertDisplayedAddress(ctx, insertDisplayedAddressSQL, displayedAddress, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	displayedAddressResponse := partyproto.CreateDisplayedAddressResponse{}
	displayedAddressResponse.DisplayedAddress = displayedAddress
	return &displayedAddressResponse, nil
}

// ProcessDisplayedAddressRequest - Process DisplayedAddressRequest
func (ps *PartyService) ProcessDisplayedAddressRequest(ctx context.Context, in *partyproto.CreateDisplayedAddressRequest) (*partyproto.DisplayedAddress, error) {
	displayedAddress := partyproto.DisplayedAddress{}

	displayedAddress.DocumentPartyId = in.DocumentPartyId
	displayedAddress.AddressLineNumber = in.AddressLineNumber
	displayedAddress.AddressLineText = in.AddressLineText

	return &displayedAddress, nil
}

// insertDisplayedAddress - Insert Document Party
func (ps *PartyService) insertDisplayedAddress(ctx context.Context, insertDisplayedAddressSQL string, displayedAddress *partyproto.DisplayedAddress, userEmail string, requestID string) error {
	err := ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertDisplayedAddressSQL, displayedAddress)
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
