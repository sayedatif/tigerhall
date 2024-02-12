package users

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sayedatif/tigerhall/server"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	router := server.NewRouter()

	w := httptest.NewRecorder()
	payload := []byte(`{"email": "test@test.com", "password": "secret"}`)
	req, _ := http.NewRequest("POST", "/users/login", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// r := gin.Default()
	// r.POST("/login", Login)

	// req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// rr := httptest.NewRecorder()

	// r.ServeHTTP(rr, req)

	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }

	// // expected := `{"message":"User logged in successfully"}`
	// bs, _ := json.Marshal(rr)
	// fmt.Println("rr", string(bs))
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}
