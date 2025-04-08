package service

import (
	"context"
	"dagger/app/user/dto"
	"dagger/common"
	"dagger/lib/logger"
	"dagger/lib/sys"
	"dagger/model/mysql"
	"dagger/utils"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type userService struct {
	ac *utils.AesCbc
}

func NewUserService() *userService {
	return &userService{
		ac: &utils.AesCbc{
			Key: "ygefPRL3DM85wTXUejJNrRuu",
			Iv:  "xiNFO9Zj8muFPey7",
		},
	}
}

// 邮箱密码登录或注册
func (s *userService) LoginOrRegister(ctx context.Context, req dto.ReqLoginHandler) (int, string, error) {
	user := mysql.User{}
	var err error
	if req.Action == "login" {
		user, err = mysql.NewUserModel().GetUserByWhere(ctx, map[string]string{"email": req.Email})
		if err != nil && err != gorm.ErrRecordNotFound {
			return 0, "", err
		}
		if err == gorm.ErrRecordNotFound {
			return 0, "", errors.New("请注册后再登录")
		}
		if utils.Md5X(req.Password+user.Salt) != user.Password {
			return 0, "", errors.New("密码错误")
		}
		if user.IsDestroy(ctx) {
			return 0, "", errors.New("该账号已注销")
		}
		if user.IsFreeze(ctx) {
			return 0, "", errors.New("该账号已冻结")
		}
	} else {
		user, err = mysql.NewUserModel().GetUserByWhere(ctx, map[string]string{"email": req.Email})
		if err != nil && err != gorm.ErrRecordNotFound {
			return 0, "", err
		}
		if user.Id > 0 {
			return 0, "", errors.New("该邮箱已使用，请更换邮箱注册")
		}
		user, err = mysql.NewUserModel().GetUserByWhere(ctx, map[string]string{"nickname": req.Nickname})
		if err != nil && err != gorm.ErrRecordNotFound {
			return 0, "", err
		}
		if user.Id > 0 {
			return 0, "", errors.New("该昵称已使用，请更换昵称注册")
		}
		salt := utils.RandomString(6)
		user = mysql.User{
			Nickname:       req.Nickname,
			Email:          req.Email,
			Salt:           salt,
			Password:       utils.Md5X(req.Password + salt),
			Status:         mysql.UserStatusNormal,
			LastActiveTime: "",
			BaseModel: mysql.BaseModel{
				CreateTime: utils.LocalTime(time.Now()),
				ModifyTime: utils.LocalTime(time.Now()),
			},
		}
		id, err := mysql.NewUserModel().Add(ctx, user)
		if err != nil {
			return 0, "", err
		}
		user.Id = id
	}

	token, err := s.EncodeToken(ctx, user.Id)
	return user.Id, token, err
}

// 获取用户信息
func (s *userService) GetUserInfo(ctx context.Context, userId int) (mysql.User, error) {
	return mysql.NewUserModel().GetUserByWhere(ctx, map[string]string{"id": cast.ToString(userId)})
}

// 更新用户最后活跃时间
func (s *userService) UpdateUserLastActiveTime(ctx context.Context, userId int) error {
	user := map[string]interface{}{}
	user["id"] = userId
	user["last_active_time"] = utils.Now()
	_, err := mysql.NewUserModel().Update(ctx, user)
	return err
}

// 解密token
func (s *userService) DecodeToken(c *gin.Context) error {
	userId := 0
	if common.IsDebug() && c.Query("_debug") == "dagger" {
		userId = 1
	} else {
		// 从 Authorization header 获取 token
		// 获取从 cookie 获取，取决于项目的具体实现方式
		// tokenStr, err := c.Request.Cookie("dagger_token")
		// if err != nil || len(tokenStr.Value) == 0 {
		// 	return errors.New("请重新登录")
		// }
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			return errors.New("请重新登录")
		}

		// 如果 token 带有 Bearer 前缀，去掉前缀
		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		tokenStr, err := s.ac.Decode(tokenStr)
		if err != nil {
			return errors.New("登录态解密失败，请重新登录")
		}
		var token dto.Token
		err = json.Unmarshal([]byte(tokenStr), &token)
		if err != nil {
			return errors.New("登录态解析失败，请重新登录")
		}
		if condition := token.Expire < utils.TimeDifference(utils.Now(), token.LoginTime); condition {
			return errors.New("登录已失效，请重新登录")
		}
		userId = token.UserId
	}
	c.Set(sys.CtxUserIDKey, userId)
	// 除了api接口层接受的是 gin.Context，其他地方都是 context.Context
	// 为了方便后续其他地方处理，比如后续代码逻辑获取 user_id 或者日志默认打印 user_id（config log transparent_parameter 配置中如果有），这里同步把 user_id 放到 context.Context 中
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, sys.CtxUserIDKey, userId)
	c.Request = c.Request.WithContext(ctx)
	return nil
}

// 解密token，不返回错误，用于不强制登录的场景
func (s *userService) DecodeTokenWithNoError(c *gin.Context) {
	userId := 0
	if common.IsDebug() && c.Query("_debug") == "dagger" {
		userId = 1
	} else {
		// 从 Authorization header 获取 token
		// 获取从 cookie 获取，取决于项目的具体实现方式
		// tokenStr, err := c.Request.Cookie("dagger_token")
		// if err != nil || len(tokenStr.Value) == 0 {
		// 	return errors.New("请重新登录")
		// }
		headertoken := c.GetHeader("Authorization")
		// 如果 token 带有 Bearer 前缀，去掉前缀
		if len(headertoken) > 7 && headertoken[:7] == "Bearer " {
			if tokenStr, err := s.ac.Decode(headertoken[7:]); err == nil {
				var token dto.Token
				if err = json.Unmarshal([]byte(tokenStr), &token); err == nil {
					if condition := token.Expire >= utils.TimeDifference(utils.Now(), token.LoginTime); condition {
						userId = token.UserId
					}
				}
			}
		}
	}
	c.Set(sys.CtxUserIDKey, userId)
	// 除了api接口层接受的是 gin.Context，其他地方都是 context.Context
	// 为了方便后续其他地方处理，比如后续代码逻辑获取 user_id 或者日志默认打印 user_id（config log transparent_parameter 配置中如果有），这里同步把 user_id 放到 context.Context 中
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, sys.CtxUserIDKey, userId)
	c.Request = c.Request.WithContext(ctx)
}

// 生成token
func (s *userService) EncodeToken(ctx context.Context, userId int) (string, error) {
	tokens := dto.Token{
		UserId:    userId,
		LoginTime: utils.Now(),
		Expire:    30 * 24 * 3600,
	}
	tokenByte, _ := json.Marshal(tokens)
	token, err := s.ac.Encode(string(tokenByte))
	if err != nil {
		logger.Error(ctx, "EncodeToken", err)
		return "", errors.New("token生成失败")
	}
	return token, nil
}

// 修改用户信息
func (s *userService) ModifyUserInfo(c *gin.Context, req dto.ReqModifyUserInfoHandler) error {
	modifyInfo := map[string]interface{}{}
	// 检查邮箱是否已使用
	if req.Email != "" {
		if !utils.IsEmail(req.Email) {
			return errors.New("邮箱格式错误")
		}
		user, err := mysql.NewUserModel().GetUserByWhere(c.Request.Context(), map[string]string{"email": req.Email})
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if user.Id > 0 && user.Id != c.MustGet("user_id").(int) {
			return errors.New("该邮箱已使用，请更换邮箱注册")
		}
		modifyInfo["email"] = req.Email
	}
	// 检查昵称是否已使用
	if req.Nickname != "" {
		if len(req.Nickname) < 2 || len(req.Nickname) > 10 {
			return errors.New("昵称长度必须在2到10个字符之间")
		}
		user, err := mysql.NewUserModel().GetUserByWhere(c.Request.Context(), map[string]string{"nickname": req.Nickname})
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if user.Id > 0 && user.Id != c.MustGet("user_id").(int) {
			return errors.New("该昵称已使用，请更换昵称注册")
		}
		modifyInfo["nickname"] = req.Nickname
	}

	// 检查手机号是否已使用
	if req.Phone != "" {
		if !utils.IsChinesePhoneNumber(req.Phone) {
			return errors.New("手机号格式错误")
		}
		user, err := mysql.NewUserModel().GetUserByWhere(c.Request.Context(), map[string]string{"phone": req.Phone})
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if user.Id > 0 && user.Id != c.MustGet("user_id").(int) {
			return errors.New("该手机号已使用，请更换手机号注册")
		}
		modifyInfo["phone"] = req.Phone
	}

	if len(modifyInfo) < 1 {
		return nil
	}

	modifyInfo["modify_time"] = utils.Now()
	modifyInfo["id"] = c.MustGet("user_id").(int)
	_, err := mysql.NewUserModel().Update(c.Request.Context(), modifyInfo)
	return err
}

// 修改密码
func (s *userService) ModifyPassword(c *gin.Context, req dto.ReqModifyUserPasswordHandler) error {
	user, err := mysql.NewUserModel().GetUserByWhere(c.Request.Context(), map[string]string{"id": cast.ToString(c.MustGet("user_id").(int))})
	if err != nil {
		return err
	}
	if utils.Md5X(req.OldPassword+user.Salt) != user.Password {
		return errors.New("旧密码错误")
	}
	if req.Password != req.ConfirmPassword {
		return errors.New("两次输入的密码不一致")
	}
	salt := utils.RandomString(6)
	modifyInfo := map[string]interface{}{}
	modifyInfo["password"] = utils.Md5X(req.Password + salt)
	modifyInfo["salt"] = salt
	modifyInfo["modify_time"] = utils.Now()
	modifyInfo["id"] = c.MustGet("user_id").(int)
	_, err = mysql.NewUserModel().Update(c.Request.Context(), modifyInfo)
	return err
}

// 重置密码，发送邮件
func (s *userService) ResetPassword(c *gin.Context, req dto.ReqResetUserPasswordHandler) error {
	if !utils.IsEmail(req.Email) {
		return errors.New("邮箱格式错误")
	}
	_, err := mysql.NewUserModel().GetUserByWhere(c.Request.Context(), map[string]string{"email": req.Email})
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return errors.New("该邮箱尚未注册")
	}

	// 生成重置密码token，包含email和过期时间
	resetToken := dto.ResetPasswordToken{
		Email:     req.Email,
		ResetTime: utils.Now(),
		Expire:    3600, // 1小时后过期
	}
	tokenBytes, err := json.Marshal(resetToken)
	if err != nil {
		logger.Error(c.Request.Context(), "生成重置密码token失败", err)
		return errors.New("生成重置密码token失败")
	}

	// 加密token
	encryptedToken, err := s.ac.Encode(string(tokenBytes))
	if err != nil {
		logger.Error(c.Request.Context(), "加密token失败", err)
		return errors.New("加密token失败")
	}

	// 构建重置密码链接
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", common.Cfg.Domain, encryptedToken)

	// 发送重置密码邮件
	mailContent := fmt.Sprintf(`
		<p>您好，</p>
		<p>请点击以下链接重置密码：</p>
		<p><a href="%s">重置密码</a></p>
		<p>链接有效期为1小时，请尽快完成密码重置。</p>
		<p>如果这不是您的操作，请忽略此邮件。</p>
	`, resetLink)

	err = utils.SendMail(req.Email, "重置密码", mailContent)
	if err != nil {
		logger.Error(c.Request.Context(), "发送重置密码邮件失败", err)
		return errors.New("发送重置密码邮件失败")
	}

	return nil
}

// 设置密码
func (s *userService) SetPassword(c *gin.Context, req dto.ReqResetPasswordSetNewHandler) error {
	if req.Token == "" {
		return errors.New("重置密码token不能为空")
	}

	// 解密token
	tokenStr, err := s.ac.Decode(req.Token)
	if err != nil {
		logger.Error(c.Request.Context(), "解密重置密码token失败", err)
		return errors.New("重置密码token无效")
	}

	// 解析token
	var resetToken dto.ResetPasswordToken
	err = json.Unmarshal([]byte(tokenStr), &resetToken)
	if err != nil {
		logger.Error(c.Request.Context(), "解析重置密码token失败", err)
		return errors.New("重置密码token无效")
	}

	// 验证token是否过期
	if condition := resetToken.Expire < utils.TimeDifference(utils.Now(), resetToken.ResetTime); condition {
		return errors.New("重置密码链接已过期，请重新发起重置密码请求")
	}

	// 验证密码
	if req.Password != req.ConfirmPassword {
		return errors.New("两次输入的密码不一致")
	}

	// 获取用户信息
	user, err := mysql.NewUserModel().GetUserByWhere(c.Request.Context(), map[string]string{"email": resetToken.Email})
	if err != nil {
		return err
	}

	// 生成新的盐和密码
	salt := utils.RandomString(6)
	modifyInfo := map[string]interface{}{
		"salt":        salt,
		"password":    utils.Md5X(req.Password + salt),
		"modify_time": utils.Now(),
		"id":          user.Id,
	}

	// 更新密码
	_, err = mysql.NewUserModel().Update(c.Request.Context(), modifyInfo)
	if err != nil {
		logger.Error(c.Request.Context(), "更新密码失败", err)
		return errors.New("设置新密码失败")
	}

	return nil
}

type authCodeRecord struct {
	lastSent time.Time // 上次发送时间
	expireAt time.Time // 过期时间
	code     string    // 验证码
}

var (
	authCodeCache sync.Map
	cacheDuration = 5 * time.Minute
	cleanInterval = 20 * time.Minute
)

// 暂时不用redis，这里用内存缓存做一下安全处理，避免刷短信的同时也避免内存占用过大
func init() {
	go cleanExpiredRecords()
}

func cleanExpiredRecords() {
	for {
		time.Sleep(cleanInterval)
		now := time.Now()
		authCodeCache.Range(func(key, value interface{}) bool {
			record := value.(authCodeRecord)
			if now.After(record.expireAt) {
				authCodeCache.Delete(key)
			}
			return true
		})
	}
}

// 是否可以发送验证码
func (s *userService) CanSendAuthCode(key string) bool {
	now := time.Now()
	if record, ok := authCodeCache.Load(key); ok {
		r := record.(authCodeRecord)
		if now.Before(r.lastSent.Add(cacheDuration)) {
			return false
		}
	}
	return true
}

// 验证验证码
func (s *userService) VerifyAuthPhoneCode(c *gin.Context, req dto.ReqAuthPhoneHandler) error {

	if sys.IsDebug() && req.Code == "888888" {
		return nil
	}
	// 从缓存中获取验证码记录
	record, ok := authCodeCache.Load(req.Phone)
	if !ok {
		return errors.New("验证码不存在或已过期")
	}

	authRecord := record.(authCodeRecord)
	now := time.Now()

	// 检查验证码是否过期
	if now.After(authRecord.expireAt) {
		authCodeCache.Delete(req.Phone)
		return errors.New("验证码已过期")
	}

	if authRecord.code != req.Code {
		return errors.New("验证码不正确")
	}

	authCodeCache.Delete(req.Phone)
	authCodeCache.Delete(c.ClientIP())
	return nil
}

func (s *userService) SendAuthCode(c *gin.Context, req dto.ReqAuthCodeHandler) error {

	code := utils.RandomCode(6)
	// 业务逻辑，暂时不用
	// err := utils.SendSmsCode(c.Request.Context(), req.Phone, code)
	// if err != nil {
	// 	logger.ErrorWithField(c.Request.Context(), "SendAuthCode", err.Error(), map[string]interface{}{
	// 		"phone": req.Phone,
	// 	})
	// 	return errors.New("发送验证码失败")
	// }

	now := time.Now()
	// 手机号的缓存时间设置用来限制单个手机号发送频率
	authCodeCache.Store(req.Phone, authCodeRecord{
		lastSent: now,
		expireAt: now.Add(5 * time.Minute),
		code:     code,
	})
	// ip 的缓存时间设置用来限制ip发送频率，防止刷短信
	authCodeCache.Store(c.ClientIP(), authCodeRecord{
		lastSent: now,
		expireAt: now.Add(1 * time.Minute),
	})

	return nil
}

func (s *userService) LoginOrRegisterByPhone(ctx context.Context, req dto.ReqAuthPhoneHandler) (string, error) {
	user, err := mysql.NewUserModel().GetUserByWhere(ctx, map[string]string{"phone": req.Phone})
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}
	if user.Id < 1 {
		user = mysql.User{
			Nickname:       "",
			Email:          "",
			Phone:          req.Phone,
			Salt:           "",
			Password:       "",
			Status:         mysql.UserStatusNormal,
			LastActiveTime: "",
			BaseModel: mysql.BaseModel{
				CreateTime: utils.LocalTime(time.Now()),
				ModifyTime: utils.LocalTime(time.Now()),
			},
		}
		id, err := mysql.NewUserModel().Add(ctx, user)
		if err != nil {
			return "", err
		}
		user.Id = id
	}

	token, err2 := s.EncodeToken(ctx, user.Id)
	if err2 != nil {
		logger.Error(ctx, "LoginOrRegisterByPhone", err2)
		return "", errors.New("token生成失败")
	}
	return token, nil
}
