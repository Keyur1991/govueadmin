package request

import (
	mw "govueadmin/framework/middleware"

	"github.com/gorilla/mux"
)

func Router(config *Config) *mux.Router {
	r := mux.NewRouter()

	r.Use(mw.Cors(config.Cors))
	r.Use(mw.RateLimit(config.RateLimit))
	r.Use(mw.Logger)

	return r
}
