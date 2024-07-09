package mysql

import (
	"context"
	dlog "dagger/lib/logger"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewMysqlLogger 创建一个mysql日志记录器
func NewMysqlLogger(slowThreshold time.Duration, level logger.LogLevel, ignore bool) *MysqlLogger {
	return &MysqlLogger{
		SlowThreshold:             slowThreshold,
		LogLevel:                  level,
		IgnoreRecordNotFoundError: ignore,
	}
}

var _ logger.Interface = (*MysqlLogger)(nil)

func (ml *MysqlLogger) LogMode(lev logger.LogLevel) logger.Interface {
	return &MysqlLogger{}
}

func (ml *MysqlLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	dlog.Info(ctx, TAGNAME, msg, args)
}

func (ml *MysqlLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	dlog.Warn(ctx, TAGNAME, msg, args)
}

func (ml *MysqlLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	dlog.ErrorWithMsg(ctx, TAGNAME, msg, args)
}

func (ml *MysqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil && (err != gorm.ErrRecordNotFound || !ml.IgnoreRecordNotFoundError) {
		dlog.ErrorWithMsg(ctx, "SQL_ERROR", "sql=%v, rows=%v, elapsed=%v, err=%v", sql, rows, elapsed, err)
		return
	}

	if ml.SlowThreshold != 0 && elapsed > ml.SlowThreshold {
		dlog.Warn(ctx, "SLOW_QUERY", "sql=%v, rows=%v, elapsed=%v, slowthreshold=%v", sql, rows, elapsed, ml.SlowThreshold)
	} else {
		dlog.Info(ctx, "MySQL_Query", "sql=%v, rows=%v, elapsed=%v", sql, rows, elapsed)
	}
}
