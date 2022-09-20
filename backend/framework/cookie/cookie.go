package cookie

import (
	"github.com/gorilla/securecookie"
)

func EncodeCookieString(val *string, key string, secretKey string) (string, error) {
	s := securecookie.New([]byte(secretKey), nil)
	return s.Encode(key, *val)
}

func DecodeCookieString(val string, key string, receiver *string, secretKey string) error {
	s := securecookie.New([]byte(secretKey), nil)
	return s.Decode(key, val, receiver)
}
