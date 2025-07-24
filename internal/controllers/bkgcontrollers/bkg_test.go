package bkgcontrollers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	"github.com/cloudfresco/sc-dcsa/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	bkg := bkgproto.CreateBookingRequest{}
	bkg.CarrierBookingRequestReference = "ef223019-ff16-4870-be69-9dbaaaae9b11"
	bkg.DocumentStatus = "Received"
	bkg.ReceiptTypeAtOrigin = "CY"
	bkg.DeliveryTypeAtDestination = "CY"
	bkg.CargoMovementTypeAtOrigin = "FCL"
	bkg.CargoMovementTypeAtDestination = "FCL"
	bkg.ServiceContractReference = ""
	bkg.PaymentTermCode = "PRE"
	bkg.IsPartialLoadAllowed = false
	bkg.IsExportDeclarationRequired = true
	bkg.ExportDeclarationReference = "export_declaration_reference"
	bkg.IsImportLicenseRequired = true
	bkg.ImportLicenseReference = "import_license_reference"
	bkg.SubmissionDateTime = "02/16/2023"
	bkg.IsAmsAciFilingRequired = false
	bkg.IsDestinationFilingRequired = false
	bkg.ContractQuotationReference = "contractRef"
	bkg.ExpectedDepartureDate = "12/10/2023"
	bkg.ExpectedArrivalAtPlaceOfDeliveryStartDate = "11/09/2023"
	bkg.ExpectedArrivalAtPlaceOfDeliveryEndDate = "11/15/2023"
	bkg.TransportDocumentTypeCode = "BOL"
	bkg.TransportDocumentReference = ""
	bkg.BookingChannelReference = ""
	bkg.IncoTerms = ""
	bkg.CommunicationChannelCode = ""
	bkg.IsEquipmentSubstitutionAllowed = false
	bkg.VesselName = "Rum Runner"
	bkg.VesselImoNumber = "9321483"
	bkg.ExportVoyageNumber = "export-voyage-number"
	bkg.PreCarriageModeOfTransportCode = ""
	bkg.DeclaredValueCurrency = "EUR"
	bkg.DeclaredValue = "10000"
	bkg.InvoicePayableAt = ""

	commodity := bkgproto.CreateCommodityRequest{}
	commodity.CommodityType = "Mobile phones"
	commodity.HsCode = "720711"
	commodity.CargoGrossWeight = 12000.00
	commodity.CargoGrossWeightUnit = "KGM"
	commodity.CargoGrossVolume = float64(0)
	commodity.CargoGrossVolumeUnit = ""
	commodity.NumberOfPackages = uint32(0)
	commodity.ExportLicenseIssueDate = "02/22/2023"
	commodity.ExportLicenseExpiryDate = "02/22/2023"

	commodities := []*bkgproto.CreateCommodityRequest{}
	commodities = append(commodities, &commodity)
	bkg.Commodities = commodities

	reference := bkgproto.CreateReferenceRequest{}
	reference.ReferenceTypeCode = "FF"
	reference.ReferenceValue = "test"

	references := []*bkgproto.CreateReferenceRequest{}
	references = append(references, &reference)
	bkg.References = references

	valueAddedServiceRequest := bkgproto.CreateValueAddedServiceRequest{}
	valueAddedServiceRequest.ValueAddedServiceCode = "CDECL"

	valueAddedServiceRequests := []*bkgproto.CreateValueAddedServiceRequest{}
	valueAddedServiceRequests = append(valueAddedServiceRequests, &valueAddedServiceRequest)
	bkg.ValueAddedServiceRequests = valueAddedServiceRequests

	requestedEquipment := bkgproto.CreateRequestedEquipmentRequest{}
	requestedEquipment.RequestedEquipmentSizetype = "22GP"
	requestedEquipment.RequestedEquipmentUnits = 3

	requestedEquipments := []*bkgproto.CreateRequestedEquipmentRequest{}
	requestedEquipments = append(requestedEquipments, &requestedEquipment)
	bkg.RequestedEquipments = requestedEquipments

	shipmentLocation := bkgproto.CreateShipmentLocationRequest{}
	shipmentLocation.ShipmentLocationTypeCode = "FCD"
	shipmentLocation.DisplayedName = "Singapore"
	shipmentLocation.EventDateTime = "02/22/2023"

	shipmentLocations := []*bkgproto.CreateShipmentLocationRequest{}
	shipmentLocations = append(shipmentLocations, &shipmentLocation)
	bkg.ShipmentLocations = shipmentLocations

	data, _ := json.Marshal(&bkg)

	req, err := http.NewRequest("POST", backendServerAddr+"/v2/bookings", bytes.NewBuffer(data))
	if err != nil {
		t.Error(err)
		return
	}

	req = common.SetEmailToken(req, tokenString, email)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	// Check the status code is what we expect.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
		return
	}

	assert.NotNil(t, w.Body.String())
}

func TestUpdateBooking(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	form := bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceRequest{}
	form.DocumentStatus = "PENU"
	form.ReceiptTypeAtOrigin = "CY"
	form.DeliveryTypeAtDestination = "CFS"
	form.CargoMovementTypeAtOrigin = "FCL"
	form.CargoMovementTypeAtDestination = "LCL"
	form.ServiceContractReference = "SERVICE_CONTRACT_REFERENCE_01"
	form.PaymentTermCode = "PRE"

	data, _ := json.Marshal(&form)

	req, err := http.NewRequest("PUT", backendServerAddr+"/v2/bookings/CARRIER_BOOKING_REQUEST_REFERENCE_01", bytes.NewBuffer(data))
	if err != nil {
		t.Error(err)
		return
	}

	req = common.SetEmailToken(req, tokenString, email)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	// Check the status code is what we expect.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
		return
	}

	expected := string(`"Updated Successfully"` + "\n")
	assert.Equal(t, w.Body.String(), expected, "they should be equal")
}
