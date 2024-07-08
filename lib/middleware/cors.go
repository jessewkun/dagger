package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 配置 CORS
func Cros() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"http://example.com"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET"},
	})
}
