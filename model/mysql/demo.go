package mysql

import (
	"context"
	"dagger/lib/logger"
	"dagger/utils"

	"gorm.io/gorm/clause"
)

type Demo struct {
	Id         int             `json:"id" orm:"id"`                   // 主键 id
	Name       string          `json:"name" orm:"name"`               // 名称
	Email      string          `json:"email" orm:"email"`             // 邮箱
	CreateTime utils.LocalTime `json:"create_time" orm:"create_time"` // 创建时间
	ModifyTime utils.LocalTime `json:"modify_time" orm:"modify_time"` // 修改时间
}

const TAGDEMO = "model demo"

func (table *Demo) TableName() string {
	return "demo"
}

func (w Demo) AddTest(ctx context.Context, t Demo) (id int, _err error) {
	_err = mainDb(ctx).Table(w.TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "modify_time"}),
	}).Create(&t).Error
	if _err != nil {
		logger.ErrorWithMsg(ctx, TAGDEMO, "AddTest error, test: %+v, err: %+v", t, _err)
		return 0, _err
	}
	id = t.Id
	return
}

func (w Demo) GetDemoById(ctx context.Context, id int) (demo Demo, err error) {
	err = mainDb(ctx).Table(w.TableName()).Where("id", id).First(&demo).Error
	if err != nil && err.Error() != "record not found" {
		logger.ErrorWithMsg(ctx, TAGDEMO, "GetDemoById error, id: %d, err: %+v", id, err)
		return
	}
	return
}

func (w Demo) ListByCondition(ctx context.Context, req Demo) (resp []*Demo, _err error) {
	db := mainDb(ctx).Table(w.TableName())
	if req.Name != "" {
		db = db.Where("name = ?", req.Name)
	}
	if req.Email != "" {
		db = db.Where("email = ?", req.Email)
	}
	// if len(req.StartTime) > 0 {
	// 	db = db.Where("create_time >= ?", req.StartTime)
	// }
	// if len(req.EndTime) > 0 {
	// 	db = db.Where("create_time <= ?", req.EndTime)
	// }

	// _err = mainDb().Raw(sql).Scan(&resp).Error
	_err = db.Find(&resp).Error
	// _err = db.Count(&count).Error
	if _err == nil || _err.Error() == "record not found" {
		_err = nil
	}

	if _err != nil {
		logger.ErrorWithMsg(ctx, TAGDEMO, "ListByCondition err, req: %+v, err: %+v", req, _err)
		return
	}
	return
}

func (w Demo) CountByCondition(ctx context.Context, req Demo) (count int64, _err error) {
	db := mainDb(ctx).Table(w.TableName())
	if req.Name != "" {
		db = db.Where("name = ?", req.Name)
	}
	if req.Email != "" {
		db = db.Where("email = ?", req.Email)
	}
	// if len(req.StartTime) > 0 {
	// 	db = db.Where("create_time >= ?", req.StartTime)
	// }
	// if len(req.EndTime) > 0 {
	// 	db = db.Where("create_time <= ?", req.EndTime)
	// }

	_err = db.Count(&count).Error
	if _err == nil || _err.Error() == "record not found" {
		_err = nil
	}

	if _err != nil {
		logger.ErrorWithMsg(ctx, TAGDEMO, "ListByCondition err, req: %+v, err: %+v", req, _err)
		return
	}
	return
}
