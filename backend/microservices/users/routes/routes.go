package routes

import (
	mw "govueadmin/framework/middleware"
	"govueadmin/microservices/users/config"

	"govueadmin/framework/request"

	"govueadmin/microservices/users/controllers"
	"govueadmin/microservices/users/middlewares"

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
		u.POST("/auth/login", controllers.Login)
		u.POST("/auth/logout", controllers.Logout)
		u.POST("/register", controllers.Register)
		u.POST("/me", controllers.Me, mw.CheckAuth(), middlewares.Authenticate())
	}

	// r := request.Router(conf)

	// r.HandleFunc("/users/auth/login", controllers.Login).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/auth/logout", controllers.Logout).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/register", controllers.Register).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/me", middlewares.Authenticate(mw.CheckAuth(controllers.Me))).Methods(http.MethodPost, http.MethodOptions)

	return r
}
