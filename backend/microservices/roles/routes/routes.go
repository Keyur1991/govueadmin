package routes

import (
	mw "govueadmin/framework/middleware"
	"govueadmin/microservices/roles/config"

	"govueadmin/framework/request"

	mwauth "govueadmin/microservices/common/middlewares/auth"
	"govueadmin/microservices/roles/controllers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	conf := &request.Config{
		Cors:      config.Cors(),
		RateLimit: config.RateLimit(),
	}

	r := request.Router(conf)

	u := r.Group("/roles")
	{
		u.Use(mw.CheckAuth())
		u.Use(mwauth.Authenticate())
		u.POST("/", controllers.CreateRoleHandler)
		u.PUT("/:id", controllers.UpdateRoleHandler)
		u.DELETE("/:id", controllers.DeleteRoleHandler)
		u.GET("/", controllers.GetRolesHandler)
		u.GET("/:id", controllers.ViewRoleHandler)
	}

	return r
}
