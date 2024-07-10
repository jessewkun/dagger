package middleware

import (
	"bytes"
	"dagger/lib/logger"
	"dagger/lib/sys"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

// DaggerRecovery 中间件自定义 panic 恢复
func DaggerRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				trace := PanicTrace(2)
				logger.ErrorWithField(c.Request.Context(), TAGNAME, "PANIC", map[string]interface{}{
					"recover": r,
					"panic":   string(trace),
				})
				if sys.IsDebug() {
					fmt.Printf("recover: %+v\n", r)
					fmt.Printf("panic: %+v\n", string(trace))
				}
				c.JSON(http.StatusOK, sys.SystemErrorResp(c))
				c.Abort()
			}
		}()
		c.Next()
	}
}

func PanicTrace(kb int) []byte {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return stack
}
