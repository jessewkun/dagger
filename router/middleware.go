package router

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 检查签名
func CheckSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "user_id", 123)
		ctx = context.WithValue(ctx, "trace_id", "123")
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// 配置 CORS
func Cros() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"http://example.com"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET"},
	})
}

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("trace_id")
		if len(traceId) < 1 {
			traceId = uuid.New().String()
		}
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "trace_id", traceId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}

}
