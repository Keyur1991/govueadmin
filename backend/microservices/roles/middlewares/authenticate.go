package middlewares

import (
	"app/microservices/roles/client"
	"net/http"

	"app/microservices/common/feat"
	"app/microservices/roles/pb"
	"context"

	"github.com/Keyur1991/go-shreeva/response"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		pc := &client.AuthServiceClient{
			AuthClient: client.InitAuthServiceClient(),
		}

		res, err := pc.CheckAuthenticate(context.Background(), &pb.MeRequest{
			Token: feat.ExtractToken(c),
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
