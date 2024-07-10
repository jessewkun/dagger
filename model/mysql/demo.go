package mysql

import (
	"context"
	"dagger/utils"
)

type Demo struct {
	Id         int             `json:"id" gorm:"primaryKey"`                                                     // 主键 id
	Name       string          `json:"name" gorm:"size:64;not null"`                                             // 名称
	Email      string          `json:"email" gorm:"size:64;not null"`                                            // 邮箱
	CreateTime utils.LocalTime `json:"create_time" gorm:"default:CURRENT_TIMESTAMP"`                             // 创建时间
	ModifyTime utils.LocalTime `json:"modify_time" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"` // 修改时间
}

const TAGDEMO = "model demo"

func (table *Demo) TableName() string {
	return "demo"
}

func NewDemoModel() *Demo {
	return &Demo{}
}

func (w *Demo) Add(ctx context.Context, d Demo) (id int, _err error) {
	_err = mainDb(ctx).Table(w.TableName()).Create(&d).Error
	if _err != nil {
		return 0, _err
	}
	id = d.Id
	return
}

func (w *Demo) Update(ctx context.Context, d Demo) (rows int, _err error) {
	res := mainDb(ctx).Table(w.TableName()).Model(w).Where("id", d.Id).UpdateColumns(d)
	if res.Error != nil {
		// mysql lib 支持在 error 的时候自动打印日志，所以这里不需要再打印日志
		return 0, res.Error
	}
	return int(res.RowsAffected), nil
}

func (w *Demo) Delete(ctx context.Context, id int) (err error) {
	err = mainDb(ctx).Table(w.TableName()).Where("id", id).Delete(&Demo{}).Error
	if err != nil {
		// mysql lib 支持在 error 的时候自动打印日志，所以这里不需要再打印日志
		return
	}
	return
}

func (w *Demo) GetDemoById(ctx context.Context, id int) (demo Demo, err error) {
	err = mainDb(ctx).Table(w.TableName()).Where("id", id).First(&demo).Error
	if err != nil && err.Error() != "record not found" {
		// mysql lib 支持在 error 的时候自动打印日志，所以这里不需要再打印日志
		return
	}
	return
}

func (w *Demo) GetDemoList(ctx context.Context, id int, pagesize int) (demos []Demo, err error) {
	err = mainDb(ctx).Table(w.TableName()).Where("id > ?", id).Limit(pagesize).Find(&demos).Error
	if err != nil && err.Error() != "record not found" {
		// mysql lib 支持在 error 的时候自动打印日志，所以这里不需要再打印日志
		return
	}
	return
}
