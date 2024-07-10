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

func (w Demo) Add(ctx context.Context, t Demo) (id int, _err error) {
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

func (w Demo) GetDemoList(ctx context.Context, id int, pagesize int) (demos []Demo, err error) {
	err = mainDb(ctx).Table(w.TableName()).Where("id > ?", id).Limit(pagesize).Find(&demos).Error
	if err != nil && err.Error() != "record not found" {
		logger.ErrorWithMsg(ctx, TAGDEMO, "GetDemoList error, id: %d, err: %+v", id, err)
		return
	}
	return
}
