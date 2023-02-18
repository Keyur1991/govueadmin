package middlewares

import (
	"govueadmin/microservices/roles/client"
	"net/http"

	"context"
	"govueadmin/framework/response"
	"govueadmin/framework/utils"
	"govueadmin/microservices/roles/pb"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		pc := &client.AuthServiceClient{
			AuthClient: client.InitAuthServiceClient(),
		}

		res, err := pc.CheckAuthenticate(context.Background(), &pb.MeRequest{
			Token: utils.ExtractToken(c),
		})

		if err != nil {
			response.InternalServerError(c, err)
			return
		}

		if res.GetStatus() == http.StatusOK {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
			return
		}
	}
}
