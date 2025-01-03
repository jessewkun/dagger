package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 配置 CORS
func Cros(allowedOrigins map[string]bool) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type, Authorization, Content-Length,Keep-Alive,credentials,Cache-Control,X-Requested-With,If-Modified-Since,Cache-Control,Pragma,Last-Modified,Accept,Accept-Encoding,Accept-Language,Connection,Host,Referer,User-Agent,Origin,Sec-Ch-Ua,Sec-Ch-Ua-Mobile,Sec-Ch-Ua-Platform,Sec-Fetch-Dest,Sec-Fetch-Mode,Sec-Fetch-Site"},
		AllowCredentials: true,
		AllowOriginWithContextFunc: func(c *gin.Context, origin string) bool {
			return allowedOrigins[origin]
		},
	})
}
