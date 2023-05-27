package tests

import (
	"app/microservices/auth/routes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func init() {
	router = routes.Setup()
}

func request(method string, url string, data io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, data)
	setHeaders(req)

	return req
}

func response(req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	return w
}

func checkInternalServerError(w *httptest.ResponseRecorder, t *testing.T) {
	if w.Code == http.StatusInternalServerError {
		t.Fatalf("Internal server error: %s", w.Body)
	}
}

func setHeaders(r *http.Request) {
	r.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoiMjAyMi0xMS0xMVQxNDoyNjoyNS45OTg0MDEzMTJaIiwidG9rZW5faWQiOiI0ZTE0NTFkYS01ZTUxLTQ3YjQtYjJhYi04MzIyZjdkYTYwOGUiLCJ1c2VyX2lkIjoiNDhmZDc3YTktYzU2Ny00MzZkLTg3N2UtZTU1NzFjODk2NmE4In0.hVCSOsz-oR4I54tcADjUOGrNYlFC2G2at5GpSrLKKnQ")
	r.Header.Add("Content-Type", "application/json")
}

func TestLoginHandler(t *testing.T) {
	reqstr := `{"Email":"keyur.panchal@cygnetinfotech.com","Password":"Test@123"}`

	req := request("POST", "/auth/login", strings.NewReader(reqstr))

	w := response(req)

	checkInternalServerError(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMeHandler(t *testing.T) {
	req := request("POST", "/auth/me", nil)

	w := response(req)

	checkInternalServerError(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLogoutHandler(t *testing.T) {
	req := request("POST", "/auth/logout", nil)

	w := response(req)

	checkInternalServerError(w, t)

	assert.Equal(t, http.StatusOK, w.Code)
}
