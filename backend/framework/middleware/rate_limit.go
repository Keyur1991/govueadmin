package middleware

import (
	"net"
	"net/http"
	"govueadmin/framework/config"
	"golang.org/x/time/rate"
	"time"
	"sync"
	"govueadmin/framework/response"
	"github.com/gorilla/mux"
	"fmt"
)

type visitor struct {
	limiter *rate.Limiter
	lastSeen time.Time
}

var visitors = make(map[string]*visitor)

var mu sync.Mutex

var tokenBurst int
func init() {
    go cleanupVisitors()
}

func getVisitor(ip string, conf *config.RateLimit) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	tokenBurst = conf.RequestsPerMinute
	v, exists := visitors[ip]

	if !exists {
		limiter := rate.NewLimiter(rate.Every(1*time.Minute), tokenBurst)
		visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	fmt.Println(v)
	return v.limiter
}

func cleanupVisitors() {
	for {
		time.Sleep(2 * time.Second)
			
		mu.Lock()
		for ip, v := range visitors {
			fmt.Println("Time since", time.Since(v.lastSeen).Seconds())
		
			if time.Since(v.lastSeen) > 1 * time.Minute {
				delete(visitors, ip)
			} 
		}
		mu.Unlock()
	}
}

func RateLimit(conf *config.RateLimit) mux.MiddlewareFunc {
	return func (h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				response.Json(&w, http.StatusInternalServerError, "Internal server error")
				return
			}
	
			limiter := getVisitor(ip, conf)
	
			if limiter.Allow() == false {
				response.Json(&w, http.StatusTooManyRequests, http.StatusText(429))
				return
			}
			
			h.ServeHTTP(w, r)
		})
	}
}