package mysql

import (
	"context"
	"dagger/lib/debug"
	mysqllib "dagger/lib/mysql"
	"dagger/utils"

	"gorm.io/gorm"
)

var debuglog = debug.InitDebug("mysql")

type BaseModel struct {
	Id         int             `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreateTime utils.LocalTime `json:"create_time" gorm:"default:CURRENT_TIMESTAMP"`                             // 创建时间
	ModifyTime utils.LocalTime `json:"modify_time" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"` // 修改时间
}

func mainDb(ctx context.Context) *gorm.DB {
	d := mysqllib.GetConn("main").WithContext(ctx)
	if debug.IsDebug("mysql") {
		return d.Debug()
	}
	return d
}
