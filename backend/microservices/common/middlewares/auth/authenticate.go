package auth

import (
	"context"
	"fmt"
	"govueadmin/framework/response"
	cauth "govueadmin/microservices/common/client/auth"
	"govueadmin/microservices/common/feat"
	pbauth "govueadmin/microservices/common/pb/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		pc := cauth.InitAuthServiceClient()

		res, err := pc.AuthClient.Me(context.Background(), &pbauth.AuthRequest{
			Token: feat.ExtractToken(c),
		})
		fmt.Println(err)
		fmt.Println(res)
		if err != nil {
			response.InternalServerError(c, err)
			return
		}

		if res.GetStatus() == http.StatusOK {
			c.Next()
		} else {
			response.AbortWithStatusJSON(c)
			return
		}
	}
}
