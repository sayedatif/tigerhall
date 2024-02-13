package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sayedatif/tigerhall/types"
	"github.com/stretchr/testify/assert"
)

func TestGetTiger(t *testing.T) {
	req, err := http.NewRequest("GET", "/tigers", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	var getTigerResponse types.GetTigerResponse
	err = json.NewDecoder(rr.Body).Decode(&getTigerResponse)
	if err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	expected := "Fetched tigers successfully"
	if getTigerResponse.Message != expected {
		t.Errorf("expected message %q, got %q", expected, getTigerResponse.Message)
	}

	minLength := 1
	if len(getTigerResponse.Data) <= minLength {
		t.Errorf("expected length greater than %d, got %d", minLength, len(getTigerResponse.Data))
	}
}
