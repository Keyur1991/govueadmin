package routes

import (
	mw "github.com/Keyur1991/go-shreeva/middleware"
	"app/microservices/roles/config"

	"github.com/Keyur1991/go-shreeva/request"

	mwauth "app/microservices/common/middlewares/auth"
	"app/microservices/roles/controllers"

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
