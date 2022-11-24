package routes

import (
	"govueadmin/microservices/users/config"

	"govueadmin/framework/request"

	"govueadmin/microservices/users/controllers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	conf := &request.Config{
		Cors:      config.Cors(),
		RateLimit: config.RateLimit(),
	}

	r := request.Router(conf)

	u := r.Group("/users")
	{
		u.POST("/register", controllers.RegisterHandler)

	}

	// r := request.Router(conf)

	// r.HandleFunc("/users/auth/login", controllers.Login).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/auth/logout", controllers.Logout).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/register", controllers.Register).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/me", middlewares.Authenticate(mw.CheckAuth(controllers.Me))).Methods(http.MethodPost, http.MethodOptions)

	return r
}
