package mysql

import (
	"context"
	"dagger/lib/debug"
	mysqllib "dagger/lib/mysql"

	"gorm.io/gorm"
)

var debuglog = debug.InitDebug("mysql")

func mainDb(ctx context.Context) *gorm.DB {
	d := mysqllib.GetConn("main").WithContext(ctx)
	if debug.IsDebug("mysql") {
		return d.Debug()
	}
	return d
}
