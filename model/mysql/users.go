package mysql

import (
	"context"
	"dagger/app/user/dto"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Nickname       string `json:"nickname"`         // 昵称
	Phone          string `json:"phone"`            // 手机号
	Email          string `json:"email"`            // 邮箱
	Salt           string `json:"salt"`             // 盐
	Password       string `json:"password"`         // 密码
	Status         int    `json:"status"`           // 0:正常 10:注销 20:冻结
	IsAdmin        int    `json:"is_admin"`         // 是否是管理员, 0:否 1:是
	LastActiveTime string `json:"last_active_time"` // 最后活跃时间
}

const TAGUSER = "model user"

const UserStatusNormal = 0
const UserStatusDestroy = 10
const UserStatusFreeze = 20

func (table *User) TableName() string {
	return "users"
}

func NewUserModel() *User {
	return &User{}
}

func (w *User) IsDestroy(ctx context.Context) bool {
	return w.Status == UserStatusDestroy
}

func (w *User) IsFreeze(ctx context.Context) bool {
	return w.Status == UserStatusFreeze
}

func (w *User) Add(ctx context.Context, d User) (id int, _err error) {
	_err = mainDb(ctx).Table(w.TableName()).Create(&d).Error
	if _err != nil {
		return 0, _err
	}
	id = d.Id
	return
}

// 更新用户信息
// 注意参数为 map[string]interface{}， 不要使用 User 结构体，避免 0 值更新失败
func (w *User) Update(ctx context.Context, d map[string]interface{}) (rows int, _err error) {
	res := mainDb(ctx).Table(w.TableName()).Model(w).Where("id", d["id"]).Updates(d)
	if res.Error != nil {
		// mysql lib 支持在 error 的时候自动打印日志，所以这里不需要再打印日志
		return 0, res.Error
	}
	return int(res.RowsAffected), nil
}

func (w *User) GetUserByWhere(ctx context.Context, where map[string]string) (user User, err error) {
	db := mainDb(ctx).Table(w.TableName())
	if condition, ok := where["email"]; ok {
		db = db.Where("email", condition)
	}
	if condition, ok := where["nickname"]; ok {
		db = db.Where("nickname", condition)
	}
	if condition, ok := where["phone"]; ok {
		db = db.Where("phone", condition)
	}
	if condition, ok := where["id"]; ok {
		db = db.Where("id", cast.ToInt(condition))
	}
	err = db.First(&user).Error
	if err != nil {
		// mysql lib 支持在 error 的时候自动打印日志，所以这里不需要再打印日志
		return
	}
	return
}

func (w *User) GetUserListByWhere(ctx context.Context, req dto.ReqUserListHandler) (articles []*User, err error) {
	db := mainDb(ctx).Table(w.TableName())
	if len(req.Ids) > 0 {
		db = db.Where("id IN ?", req.Ids)
	}
	if req.UID > 0 {
		db = db.Where("uid = ?", req.UID)
	}
	if req.Phone != "" {
		db = db.Where("phone = ?", req.Phone)
	}
	if req.Email != "" {
		db = db.Where("email = ?", req.Email)
	}
	if req.Nickname != "" {
		db = db.Where("nickname = ?", req.Nickname)
	}
	if req.Status >= 0 {
		db = db.Where("status = ?", req.Status)
	}
	if req.StartDate != "" {
		db = db.Where("create_time >= ?", req.StartDate)
	}
	if req.EndDate != "" {
		db = db.Where("create_time <= ?", req.EndDate)
	}
	if req.ActiveStartDate != "" {
		db = db.Where("last_active_time >= ?", req.ActiveStartDate)
	}
	if req.ActiveEndDate != "" {
		db = db.Where("last_active_time <= ?", req.ActiveEndDate)
	}
	err = db.Order("id DESC").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&articles).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*User{}, nil
		}
		// mysql lib 支持在 error 的时候自动打印日志，所以这里不需要再打印日志
		return
	}
	return
}

func (w *User) GetUserCountByWhere(ctx context.Context, req dto.ReqUserListHandler) (count int64, err error) {
	db := mainDb(ctx).Table(w.TableName())
	if len(req.Ids) > 0 {
		db = db.Where("id IN ?", req.Ids)
	}
	if req.UID > 0 {
		db = db.Where("uid = ?", req.UID)
	}
	if req.Phone != "" {
		db = db.Where("phone = ?", req.Phone)
	}
	if req.Email != "" {
		db = db.Where("email = ?", req.Email)
	}
	if req.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+req.Nickname+"%")
	}
	if req.Status >= 0 {
		db = db.Where("status = ?", req.Status)
	}
	if req.StartDate != "" {
		db = db.Where("create_time >= ?", req.StartDate)
	}
	if req.EndDate != "" {
		db = db.Where("create_time <= ?", req.EndDate)
	}
	if req.ActiveStartDate != "" {
		db = db.Where("last_active_time >= ?", req.ActiveStartDate)
	}
	if req.ActiveEndDate != "" {
		db = db.Where("last_active_time <= ?", req.ActiveEndDate)
	}
	err = db.Count(&count).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		// mysql lib 支持在 error 的时候自动打印日志，所以这里不需要再打印日志
		return
	}
	return
}
