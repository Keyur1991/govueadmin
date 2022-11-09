package middlewares

import (
	"errors"
	"fmt"
	"govueadmin/microservices/users/controllers"
	"govueadmin/microservices/users/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	//"go.mongodb.org/mongo-driver/mongo"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := controllers.GetJwtClaims(c)
		fmt.Println(claims)
		if claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
			return
		}

		authToken := &models.AuthToken{}

		err := authToken.Find(claims["user_id"].(string), claims["token_id"].(string))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
			return
		}

		c.Next()
	}
}
