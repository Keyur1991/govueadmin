package jwt

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	JWT "github.com/dgrijalva/jwt-go"
)

func GetOriginalToken(authToken *string) (*JWT.Token, error) {
	return JWT.Parse(*authToken, func(token *JWT.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid authorization token.")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}

func GenerateJWTToken(claims *JWT.MapClaims, sk string) (string, error) {
	signingKey := []byte(sk)

	token := JWT.NewWithClaims(JWT.SigningMethodHS256, *claims)

	return token.SignedString(signingKey)
}

func GetJWTClaims(token *JWT.Token) map[string]interface{} {
	claims, _ := (*token).Claims.(JWT.MapClaims)
	return claims
}
