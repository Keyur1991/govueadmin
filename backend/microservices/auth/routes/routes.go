package routes

import (
	mw "github.com/Keyur1991/go-shreeva/middleware"
	"app/microservices/auth/config"

	"github.com/Keyur1991/go-shreeva/request"

	"app/microservices/auth/controllers"
	"app/microservices/auth/middlewares"

	"github.com/gin-gonic/gin"
)

type GrpcServer struct{}

func Setup() *gin.Engine {
	conf := &request.Config{
		Cors:      config.Cors(),
		RateLimit: config.RateLimit(),
	}

	r := request.Router(conf)

	// create grpc service client
	svc := &ServiceClient{
		Client: InitServiceClient(),
	}
	u := r.Group("/auth")
	{
		u.POST("/login", controllers.LoginHandler)
		u.POST("/logout", controllers.LogoutHandler)
		u.Use(mw.CheckAuth())
		u.Use(middlewares.Authenticate())
		//u.POST("/me", mw.CheckAuth(), middlewares.Authenticate(), controllers.Me)
		u.POST("/me", svc.MeHandler)
	}

	// r := request.Router(conf)

	// r.HandleFunc("/users/auth/login", controllers.Login).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/auth/logout", controllers.Logout).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/register", controllers.Register).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/users/me", middlewares.Authenticate(mw.CheckAuth(controllers.Me))).Methods(http.MethodPost, http.MethodOptions)

	return r
}

func (svc *ServiceClient) MeHandler(c *gin.Context) {
	// invokes grpc service client method
	controllers.MeHandler(c, svc.Client)
}
