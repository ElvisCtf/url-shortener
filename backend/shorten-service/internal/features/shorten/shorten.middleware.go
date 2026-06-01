package shorten

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// ipLimiter holds a per-IP token bucket limiter.
type ipLimiter struct {
	limiter *rate.Limiter
}

var (
	limiters sync.Map
)

// getLimiter returns the rate limiter for the given IP, creating one if needed.
// Each IP is allowed 5 requests/second with a burst of 10.
func getLimiter(ip string) *rate.Limiter {
	v, _ := limiters.LoadOrStore(ip, &ipLimiter{
		limiter: rate.NewLimiter(5, 10),
	})
	return v.(*ipLimiter).limiter
}

// RateLimitByIP is a Gin middleware that enforces per-IP rate limiting.
func RateLimitByIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !getLimiter(ip).Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		}
		c.Next()
	}
}
