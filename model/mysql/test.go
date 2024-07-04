package mysql

import (
	"context"
	"dagger/lib/logger"
	"dagger/utils"

	"gorm.io/gorm/clause"
)

type Test struct {
	Id         int             `json:"id" orm:"id"`
	Name       string          `json:"name" orm:"name"`
	Email      string          `json:"email" orm:"email"`
	CreateTime utils.LocalTime `json:"create_time" orm:"create_time"`
	ModifyTime utils.LocalTime `json:"modify_time" orm:"modify_time"`
}

const TAGTEST = "model test"

func (table *Test) TableName() string {
	return "test"
}

func (w Test) AddTest(ctx context.Context, t Test) (id int, _err error) {
	_err = mainDb().Table(w.TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "modify_time"}),
	}).Create(&t).Error
	if _err != nil {
		logger.ErrorWithMsg(ctx, TAGTEST, "AddTest error, test: %+v, err: %+v", t, _err)
		return 0, _err
	}
	id = t.Id
	return
}

func (w Test) ListByCondition(ctx context.Context, req Test) (resp []*Test, _err error) {
	db := mainDb().Table(w.TableName())
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
		logger.ErrorWithMsg(ctx, TAGTEST, "ListByCondition err, req: %+v, err: %+v", req, _err)
		return
	}
	return
}

func (w Test) CountByCondition(ctx context.Context, req Test) (count int64, _err error) {
	db := mainDb().Table(w.TableName())
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
		logger.ErrorWithMsg(ctx, TAGTEST, "ListByCondition err, req: %+v, err: %+v", req, _err)
		return
	}
	return
}
