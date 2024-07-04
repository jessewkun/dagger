package middleware

import (
	"context"
	"dagger/lib/constant"
	"dagger/lib/logger"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TAGNAME = "middleware"

// 检查登录态
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// 这里模拟一个用户id，业务自己来实现真实的登录态检查
		user_id := 123
		c.Set("user_id", user_id)
		// 除了api接口层接受的是 gin.Context，其他地方都是 context.Context
		// 为了方便后续其他地方处理，比如后续代码逻辑获取 user_id 或者日志默认打印 user_id（config log transparent_parameter 配置中如果有），这里同步把 user_id 放到 context.Context 中
		ctx = context.WithValue(ctx, "user_id", user_id)
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

// Traceid
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("trace_id")
		if len(traceId) < 1 {
			traceId = uuid.New().String()
		}
		c.Set("trace_id", traceId)
		ctx := c.Request.Context()
		// 除了api接口层接受的是 gin.Context，其他地方都是 context.Context
		// 为了方便后续其他地方处理，比如后续代码逻辑获取 trace_id 或者日志默认打印 trace_id（config log transparent_parameter 配置中如果有），这里同步把 trace_id 放到 context.Context 中
		ctx = context.WithValue(ctx, "trace_id", traceId)
		c.Request = c.Request.WithContext(ctx)
		host, _ := os.Hostname()
		c.Header("server", host)
		c.Next()
	}
}

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
		logger.InfoWithField(c.Request.Context(), TAGNAME, "iolog", map[string]interface{}{
			"duration":        time.Since(t),
			"request_uri":     c.Request.RequestURI,
			"method":          c.Request.Method,
			"domain":          c.Request.Host,
			"remote_ip":       c.ClientIP(),
			"user_agent":      c.Request.UserAgent(),
			"status":          c.Writer.Status(),
			"response":        c.GetString(constant.DAGGER_OUTPUT),
			"response_length": c.Writer.Size(),
			"body":            string(bodyByte),
		})
	}
}
