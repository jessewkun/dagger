package middleware

import (
	"dagger/lib/constant"
	"dagger/lib/logger"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// IOLog
// 返回结果前记录接口返回数据
func IOLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		bodyByte := []byte{}
		if c.Request.Method == http.MethodPost {
			bodyByte, _ = io.ReadAll(c.Request.Body)
		}
		logger.InfoWithField(c.Request.Context(), "iolog", "", map[string]interface{}{
			"duration":        time.Since(t),
			"request_uri":     c.Request.RequestURI,
			"method":          c.Request.Method,
			"domain":          c.Request.Host,
			"remote_ip":       c.ClientIP(),
			"user_agent":      c.Request.UserAgent(),
			"status":          c.Writer.Status(),
			"response":        c.GetString(constant.CTX_DAGGER_OUTPUT),
			"response_length": c.Writer.Size(),
			"body":            string(bodyByte),
		})
	}
}