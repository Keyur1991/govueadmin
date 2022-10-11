package middleware

import (
	"fmt"
	"govueadmin/framework/cookie"
	"govueadmin/framework/jwt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Check if request is authenticated or not
// This method first will check if request is from
// web or the request has Authentication Cookie available then
// it will validate Authentication cookie token
// Otherwise it will look for Authorization header
// and validate that header.
// Else return invalid json response
func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authToken string

		// Extract auth-token cookie from request
		authCookie, err := c.Cookie("auth-token")

		if err == nil {
			// get auth token from cookie
			authToken, err = GetAuthTokenFromCookie(authCookie)

			if err != nil {
				// return Internal server error response
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": http.StatusText(http.StatusInternalServerError),
				})
			}
		}

		if authToken == "" {
			// Extract authorization header from request
			authToken = c.GetHeader("Authorization")
		}

		status := false

		if authToken != "" {
			// validate authentication token
			status, err = validateJwtToken(authToken)
		}

		if !status {
			// return Unauthorized response
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("%s", err),
			})
		}

		c.Next()
	}
}

// Get auth token string from cookie
func GetAuthTokenFromCookie(encodedToken string) (string, error) {
	var token string

	err := cookie.DecodeCookieString(encodedToken, "auth-token", &token, os.Getenv("SECRET_KEY_COOKIE"))

	return token, err
}

// Validate authorization header
func validateJwtToken(token string) (bool, error) {
	// get the original token
	jwtToken, err := jwt.GetOriginalToken(&token)

	if err != nil {
		return false, fmt.Errorf("%s", err)
	} else if !jwtToken.Valid { // check if valid authorization token
		return false, fmt.Errorf("%s", "Invalid authorization token.")
	} else {
		// get claims from the token
		claims := jwt.GetJWTClaims(jwtToken)

		// calculate the expiration time of the token
		expTime, _ := time.Parse(time.RFC3339Nano, claims["expires"].(string))

		// calculate difference of expiration time with current time
		diff := time.Now().Sub(expTime)

		if diff > 0 {
			return false, fmt.Errorf("%s", "Token expired")
		}
	}

	return true, nil
}
