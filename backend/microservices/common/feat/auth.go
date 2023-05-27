package feat

import (
	"github.com/Keyur1991/go-shreeva/cookie"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

func ExtractToken(c *gin.Context) string {
	var token string
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		// Extract auth-token cookie from request
		authCookie, err := c.Cookie("auth-token")

		if err == nil {
			mu.Lock()
			cookie.DecodeCookieString(authCookie, "auth-token", &token, os.Getenv("SECRET_KEY_COOKIE"))
			mu.Unlock()
		}

		wg.Done()
	}()

	go func() {
		mu.Lock()
		token = c.GetHeader("Authorization")
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()

	return token
}
