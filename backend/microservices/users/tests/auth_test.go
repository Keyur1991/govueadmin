package tests

import (
	"encoding/json"
	"fmt"
	"govueadmin/microservices/users/controllers"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestParseLoginRequest(t *testing.T) {
	lr := controllers.LoginRequest{
		Email:    "abc@test.com",
		Password: "1234",
	}

	lrs, err := json.Marshal(lr)

	if err != nil {
		t.Fatalf("Expected nil but got error %s", err)
	}

	r := strings.NewReader(string(lrs))

	ior := io.NopCloser(r)

	err = controllers.ParseLoginRequest(ior)

	if err != nil {
		t.Fatalf("Expected nil but got error %s", err)
	}
}

func TestSetJwtClaims(t *testing.T) {
	fields := []string{
		"12asdaskdhiahsn123asdfa",
		"asdf;asdhlfnasdnfkasnfdkasdfasd",
	}

	res := controllers.SetJwtClaims(fields...)

	if reflect.ValueOf(*res).Kind() != reflect.Map {
		t.Fatalf("Expected returned type Map but got different.")
	}
}

func TestLogin(t *testing.T) {
	endpoint := fmt.Sprintf("%s/users/auth/login", os.Getenv("API_URL"))

	r := strings.NewReader(`{
		"Email" : "keyur.panchal.test1@yopmail.com",
		"Password" : "Test@123"
	}`)

	req := httptest.NewRequest(http.MethodPost, endpoint, r)

	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()

	controllers.Login(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected %d but got %d", http.StatusOK, w.Code)
	}
}

func cookieAndToken() (*http.Cookie, string) {
	cookie := &http.Cookie{
		Name:     "auth-token",
		Value:    "MTY2MzY2MzY3NnxfZ0VGREFELUFRQmxlVXBvWWtkamFVOXBTa2xWZWtreFRtbEpjMGx1VWpWalEwazJTV3R3V0ZaRFNqa3VaWGxLYkdWSVFuQmpiVlo2U1dwdmFVMXFRWGxOYVRCM1QxTXdlVTFHVVhkUFZHOHdUbnB2TVU1cE5EVlBWRWswVGtSVk1FMXFXbUZKYVhkcFpFYzVjbHBYTldaaFYxRnBUMmxLYTAxdFdtdFBWRTV0VG5rd2VFNVhWbXhNVkZKcFQwZFJkRmxxVVhoTmFURnJXVlJWTkU1VVZUTk5iVmt4VFhwTmFVeERTakZqTWxaNVdESnNhMGxxYjJsT2FrMTRXV3BKZUZsdFdtMWFSR2Q2VFdwck5VNXRSWGxPYWxVeFdXMUpkMGx1TUM1Q04wTXlaREJQUkRScGJsSXlhM054YkU1cExXd3hNbEl5TnpWYU16ZFpSRzh5ZURCT2FrcHBhamhOfO7RvrqK4yoLzGkdgaSrD1AkxWmdqREZ5QjEk3ti2u08",
		Path:     "/",
		HttpOnly: true,
	}

	authToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoiMjAyMi0wOS0yMFQwOTo0Nzo1Ni45OTI4NDU0MjZaIiwidG9rZW5faWQiOiJkMmZkOTNmNy0xNWVlLTRiOGQtYjQxMi1kYTU4NTU3MmY1MzMiLCJ1c2VyX2lkIjoiNjMxYjIxYmZmZDgzMjk5NmEyNjU1YmIwIn0.B7C2d0OD4inR2ksqlNi-l12R275Z37YDo2x0NjJij8M"

	return cookie, authToken
}

func TestGetJwtClaims(t *testing.T) {
	cookie, token := cookieAndToken()

	recorder := httptest.NewRecorder()

	http.SetCookie(recorder, cookie)

	req := &http.Request{
		Header: http.Header{
			"Cookie":        recorder.HeaderMap["Set-Cookie"],
			"Authorization": []string{token},
		},
	}

	jwtMapClaims := controllers.GetJwtClaims(req)

	if reflect.ValueOf(jwtMapClaims).Kind() != reflect.Map {
		t.Fatalf("Expected returned type map but got different")
	}
}

func TestMe(t *testing.T) {
	endpoint := fmt.Sprintf("%s/users/me", os.Getenv("API_URL"))

	cookie, token := cookieAndToken()

	req := httptest.NewRequest(http.MethodPost, endpoint, nil)

	req.AddCookie(cookie)

	req.Header.Add("Authorization", token)

	w := httptest.NewRecorder()

	controllers.Me(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected code %d but got %d", http.StatusOK, w.Code)
	}

}
