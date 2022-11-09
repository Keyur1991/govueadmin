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
		u.Use(mw.CheckAuth())
		u.Use(middlewares.Authenticate())
		//u.POST("/me", mw.CheckAuth(), middlewares.Authenticate(), controllers.Me)
		u.POST("/me", controllers.Me)

		ur := u.Group("/roles")
		ur.POST("/", controllers.CreateRole)
		ur.PUT("/:id", controllers.UpdateRole)
		ur.DELETE("/:id", controllers.DeleteRole)
		ur.GET("/", controllers.GetRoles)
		ur.GET("/:id", controllers.ViewRole)
	}

	// r := request.Router(conf)

	// r.HandleFunc("/users/auth/login", controllers.Login).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/auth/logout", controllers.Logout).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/register", controllers.Register).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/me", middlewares.Authenticate(mw.CheckAuth(controllers.Me))).Methods(http.MethodPost, http.MethodOptions)

	return r
}
