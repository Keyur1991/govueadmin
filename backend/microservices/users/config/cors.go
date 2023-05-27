package config

import (
	"os"
	"net/http"
	fwconfig "github.com/Keyur1991/go-shreeva/config"
)

func Cors() *fwconfig.Cors {
	return &fwconfig.Cors{
		AllowedOrigins: []string{os.Getenv("FRONTEND_APP_URL")},
		AllowCredentials: true,
		Debug: false,
		AllowedHeaders: []string{"Content-Type", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposedHeaders: []string{"X-CSRF-Token", "Authorization"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
	}
}

