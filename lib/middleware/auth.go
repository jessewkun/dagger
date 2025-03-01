package middleware

import (
	"dagger/lib/sys"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 自定义登录态检查签名
//
// 下边的 CheckLogin 中间件仅仅是一个容器，具体的登录态检查逻辑需要在业务中实现该签名
type CheckLoginFunc func(c *gin.Context) error

// CheckLogin
func CheckLogin(fun CheckLoginFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fun(c); err != nil {
			c.JSON(http.StatusOK, sys.ErrorResp(c, sys.NewDaggerError(10001, err)))
			c.Abort()
			return
		}
		c.Next()
	}
}
