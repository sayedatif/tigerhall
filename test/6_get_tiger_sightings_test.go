package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sayedatif/tigerhall/types"
	"github.com/stretchr/testify/assert"
)

func TestGetTigerSighting(t *testing.T) {
	req, err := http.NewRequest("GET", fmt.Sprintf(`/tigers/%d/sighting`, tigerId), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	var getTigerSightingResponse types.GetTigerSightingResponse
	err = json.NewDecoder(rr.Body).Decode(&getTigerSightingResponse)
	if err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	expected := "Fetched tiger sightings successfully"
	if getTigerSightingResponse.Message != expected {
		t.Errorf("expected message %q, got %q", expected, getTigerSightingResponse.Message)
	}

	minLength := 1
	if len(getTigerSightingResponse.Data) <= minLength {
		t.Errorf("expected length greater than %d, got %d", minLength, len(getTigerSightingResponse.Data))
	}
}
