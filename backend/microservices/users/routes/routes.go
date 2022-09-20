package routes

import (
	mw "govueadmin/framework/middleware"
	"govueadmin/framework/request"
	"govueadmin/microservices/users/config"
	"govueadmin/microservices/users/controllers"
	"govueadmin/microservices/users/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	conf := &request.Config{
		Cors:      config.Cors(),
		RateLimit: config.RateLimit(),
	}

	r := request.Router(conf)

	r.HandleFunc("/users/auth/login", controllers.Login).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/users/auth/logout", controllers.Logout).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/users/register", controllers.Register).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/users/me", middlewares.Authenticate(mw.CheckAuth(controllers.Me))).Methods(http.MethodPost, http.MethodOptions)

	return r
}
