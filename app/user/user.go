package user

import (
	"context"
	"dagger/app/user/dto"
	"dagger/common"
	xhttp "dagger/lib/http"
	"dagger/lib/logger"
	"dagger/lib/sys"
	"dagger/service"
	"dagger/utils"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// AuthHandler 邮箱密码登录或注册
func AuthHandler(c *gin.Context) {
	var req = dto.ReqLoginHandler{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	req.ConfirmPassword = strings.TrimSpace(req.ConfirmPassword)
	if req.Action != "login" && req.Action != "register" {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusOK, sys.ErrorResp(c, errors.New("请输入邮箱或密码")))
		return
	}
	if len(req.Password) < 8 {
		c.JSON(http.StatusOK, sys.ErrorResp(c, errors.New("密码不少于8个字符,必须包含大小写和特殊字符")))
		return
	}
	if req.Action == "register" {
		if req.Password != req.ConfirmPassword {
			c.JSON(http.StatusOK, sys.ErrorResp(c, errors.New("两次密码不一致")))
			return
		}

		// if len(req.Nickname) < 3 {
		// 	c.JSON(http.StatusOK, sys.ErrorResp(c, errors.New("昵称不能少于3个字符")))
		// 	return
		// }
	}
	userService := service.NewUserService()
	userId, token, err := userService.LoginOrRegister(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}

	userInfo, err := userService.GetUserInfo(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}

	c.JSON(http.StatusOK, sys.SuccResp(c, gin.H{
		"token": token,
		"user_info": gin.H{
			"id":          userInfo.Id,
			"nickname":    userInfo.Nickname,
			"email":       userInfo.Email,
			"status":      userInfo.Status,
			"phone":       utils.MaskPhoneNumber(userInfo.Phone),
			"create_time": userInfo.CreateTime,
			"is_admin":    userInfo.IsAdmin,
		},
	}))
}

// AuthCodeHandler 发送验证码
func AuthCodeHandler(c *gin.Context) {
	var req = dto.ReqAuthCodeHandler{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}

	req.Phone = strings.TrimSpace(req.Phone)

	if !service.NewUserService().CanSendAuthCode(req.Phone) {
		c.JSON(http.StatusOK, sys.ErrorResp(c, errors.New("请求过于频繁，请稍后再试")))
		return
	}
	if !service.NewUserService().CanSendAuthCode(c.ClientIP()) {
		c.JSON(http.StatusOK, sys.ErrorResp(c, errors.New("请求过于频繁，请稍后再试")))
		return
	}

	if err := service.NewUserService().SendAuthCode(c, req); err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}

	c.JSON(http.StatusOK, sys.SuccResp(c, nil))
}

// AuthPhoneHandler 手机号验证码登录或注册
func AuthPhoneHandler(c *gin.Context) {
	var req = dto.ReqAuthPhoneHandler{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	req.Phone = strings.TrimSpace(req.Phone)
	req.Code = strings.TrimSpace(req.Code)
	if err = service.NewUserService().VerifyAuthPhoneCode(c, req); err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}

	token, err := service.NewUserService().LoginOrRegisterByPhone(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}

	setCookie(c, token, 24*365*time.Hour)
	c.JSON(http.StatusOK, sys.SuccResp(c, nil))
}

// CheckLoginHandler
// 用于前端检查登录态，检查的操作在中间件中，这里直接返回用户信息
func CheckLoginHandler(c *gin.Context) {
	userInfo, err := service.NewUserService().GetUserInfo(c.Request.Context(), c.MustGet("user_id").(int))
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}

	ctx := sys.CopyCtx(c.Request.Context())
	go func(ctx context.Context) {
		err := service.NewUserService().UpdateUserLastActiveTime(ctx, userInfo.Id)
		if err != nil {
			logger.Error(ctx, "UpdateUserLastActiveTime error", err)
		}
	}(ctx)

	c.JSON(http.StatusOK, sys.SuccResp(c, gin.H{
		"user_info": gin.H{
			"id":          userInfo.Id,
			"nickname":    userInfo.Nickname,
			"email":       userInfo.Email,
			"status":      userInfo.Status,
			"phone":       utils.MaskPhoneNumber(userInfo.Phone),
			"create_time": userInfo.CreateTime,
			"is_admin":    userInfo.IsAdmin,
		},
	}))
}

func LogoutHandler(c *gin.Context) {
	setCookie(c, "", -1)
	c.JSON(http.StatusOK, sys.SuccResp(c, nil))
}

func setCookie(c *gin.Context, token string, maxAge time.Duration) {
	secure := false
	if common.IsRelease() {
		secure = true
	}
	xhttp.SetCookie(c.Writer, "dagger_token", token, maxAge, "/", common.Cfg.Domain, secure, true)
}

// ModifyHandler 修改用户信息
func ModifyHandler(c *gin.Context) {
	var req = dto.ReqModifyUserInfoHandler{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	userService := service.NewUserService()
	err = userService.ModifyUserInfo(c, req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}
	userInfo, err := userService.GetUserInfo(c.Request.Context(), c.MustGet("user_id").(int))
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}
	c.JSON(http.StatusOK, sys.SuccResp(c, gin.H{
		"user_info": gin.H{
			"id":          userInfo.Id,
			"nickname":    userInfo.Nickname,
			"email":       userInfo.Email,
			"status":      userInfo.Status,
			"phone":       utils.MaskPhoneNumber(userInfo.Phone),
			"create_time": userInfo.CreateTime,
			"is_admin":    userInfo.IsAdmin,
		},
	}))
}

// ModifyPasswordHandler 修改用户密码
func ModifyPasswordHandler(c *gin.Context) {
	var req = dto.ReqModifyUserPasswordHandler{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	userService := service.NewUserService()
	err = userService.ModifyPassword(c, req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}
	c.JSON(http.StatusOK, sys.SuccResp(c, nil))
}

func ResetPasswordHandler(c *gin.Context) {
	var req = dto.ReqResetUserPasswordHandler{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	userService := service.NewUserService()
	err = userService.ResetPassword(c, req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}
	c.JSON(http.StatusOK, sys.SuccResp(c, nil))
}

func SetPasswordHandler(c *gin.Context) {
	var req = dto.ReqResetPasswordSetNewHandler{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ParamErrorResp(c))
		return
	}
	userService := service.NewUserService()
	err = userService.SetPassword(c, req)
	if err != nil {
		c.JSON(http.StatusOK, sys.ErrorResp(c, err))
		return
	}
	c.JSON(http.StatusOK, sys.SuccResp(c, nil))
}
