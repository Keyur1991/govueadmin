package middlewares

import (
	"errors"
	"govueadmin/framework/response"
	"govueadmin/microservices/users/controllers"
	"govueadmin/microservices/users/models"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func Authenticate(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := controllers.GetJwtClaims(r)

		if claims == nil {
			response.Json(&w, http.StatusUnauthorized, map[string]string{
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		authToken := &models.AuthToken{
			UserId:  claims["user_id"].(string),
			TokenId: claims["token_id"].(string),
		}

		var value interface{}

		rec := authToken.Find()

		err := rec.Decode(value)

		if errors.Is(err, mongo.ErrNoDocuments) {
			response.Json(&w, http.StatusUnauthorized, map[string]string{
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		h.ServeHTTP(w, r)
	})
}
