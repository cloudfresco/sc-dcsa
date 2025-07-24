package eblcontrollers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransportDocument(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	transportDocument := eblproto.CreateTransportDocumentRequest{}
	transportDocument.TransportDocumentReference = "0cc0bef0-a7c8-4c03"
	transportDocument.LocationId = uint32(8)
	transportDocument.IssueDate = "11/25/2020"
	transportDocument.ShippedOnboardDate = "12/24/2020"
	transportDocument.ReceivedForShipmentDate = "12/31/2020"
	transportDocument.NumberOfOriginals = uint32(12)
	transportDocument.CarrierId = uint32(3)
	transportDocument.ShippingInstructionId = uint32(8)
	transportDocument.DeclaredValueCurrency = "EUR"
	transportDocument.DeclaredValue = "1212"
	transportDocument.NumberOfRiderPages = int32(12)
	transportDocument.IssuingParty = "499918a2-d12d-4df6-840c-dd92357002df"
	transportDocument.CreatedDateTime = "11/28/2021"
	transportDocument.UpdatedDateTime = "12/01/2021"

	data, _ := json.Marshal(&transportDocument)

	req, err := http.NewRequest("POST", backendServerAddr+"/v2/transport-documents", bytes.NewBuffer(data))
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
