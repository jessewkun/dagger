package middleware

import (
	"context"
	"dagger/common"
	"dagger/lib/sys"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckLogin
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := 0
		if common.IsDebug() && c.Query("_debug") == "dagger" {
			userId = 1
		} else {
			token, err := c.Request.Cookie("dagger_token")
			if err != nil || len(token.Value) == 0 {
				c.JSON(http.StatusOK, sys.ErrorResp(c, sys.NewDaggerError(10001, errors.New("请重新登录"))))
				c.Abort()
				return
			}
			// TODO 实现自定义的 token 验证逻辑
		}
		c.Set(sys.CtxUserIDKey, userId)
		ctx := c.Request.Context()
		// 除了api接口层接受的是 gin.Context，其他地方都是 context.Context
		// 为了方便后续其他地方处理，比如后续代码逻辑获取 user_id 或者日志默认打印 user_id（config log transparent_parameter 配置中如果有），这里同步把 user_id 放到 context.Context 中
		ctx = context.WithValue(ctx, sys.CtxUserIDKey, userId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
