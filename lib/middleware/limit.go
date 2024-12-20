package middleware

import (
	"dagger/lib/sys"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// 限制每秒允许1次请求，最多积累3个请求
// var limiter = rate.NewLimiter(1, 3)
// r.POST("/test", middleware.RateLimiter(rate.NewLimiter(1, 3)))

func RateLimiter(l *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !l.Allow() {
			c.JSON(http.StatusOK, sys.RateLimiterErrorResp(c))
			c.Abort()
			return
		}
		c.Next()
	}
}
