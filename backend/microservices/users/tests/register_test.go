package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"govueadmin/microservices/users/controllers"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type Regbody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func register(entry interface{}) *httptest.ResponseRecorder {
	endpoint := fmt.Sprintf("%s/users/register", os.Getenv("API_URL"))

	body, _ := json.Marshal(&entry)
	req := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	controllers.Register(w, req)

	return w
}

func TestRegister(t *testing.T) {

	for _, entry := range entries("Registration") {

		w := register(entry)

		if w.Code != http.StatusOK {
			t.Fatalf("code: got %v, want %d", w.Code, http.StatusOK)
		}
	}
}

func TestRegisterValidate(t *testing.T) {
	for _, entry := range entries("Registration") {

		w := register(entry)

		if w.Code != http.StatusUnprocessableEntity {
			t.Fatalf("code: got %v, want %d", w.Code, http.StatusUnprocessableEntity)
		}
	}
}
