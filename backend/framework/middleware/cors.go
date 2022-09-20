package middleware

import (
	"github.com/rs/cors"
	"govueadmin/framework/config"
	"net/http"
)

func Cors(conf *config.Cors) func(http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: conf.AllowedOrigins,
		AllowCredentials: conf.AllowCredentials,
		Debug: conf.Debug,
		AllowedHeaders: conf.AllowedHeaders,
		ExposedHeaders: conf.ExposedHeaders,
		AllowedMethods: conf.AllowedMethods,
	}).Handler
}