package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/sayedatif/tigerhall/types"
	"github.com/stretchr/testify/assert"
)

var tigerSightingId int64

func TestCreateTigerSighting(t *testing.T) {
	formData := url.Values{}
	formData.Add("lat", "19.23")
	formData.Add("long", "72.85")

	req, err := http.NewRequest("POST", fmt.Sprintf(`/tigers/%d/sighting`, tigerId), strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, authToken))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	var createTigerSightingResponse types.CreateTigerSightingResponse
	err = json.NewDecoder(rr.Body).Decode(&createTigerSightingResponse)
	if err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	expected := "Created tiger sighting successfully"
	if createTigerSightingResponse.Message != expected {
		t.Errorf("expected message %q, got %q", expected, createTigerSightingResponse.Message)
	}
	tigerSightingId = createTigerSightingResponse.Data.TigerSightingId
}
