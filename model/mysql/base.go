package mysql

import (
	"dagger/lib/debug"
	mysqllib "dagger/lib/mysql"

	"gorm.io/gorm"
)

var debuglog = debug.InitDebug("mysql")

func mainDb() *gorm.DB {
	d := mysqllib.GetConn("main")
	if debug.IsDebug("mysql") {
		return d.Debug()
	}
	return d
}
