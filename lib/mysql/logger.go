package mysql

import (
	"context"
	dlog "dagger/lib/logger"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// newMysqlLogger 创建一个mysql日志记录器
func newMysqlLogger(slowThreshold time.Duration, level logger.LogLevel, ignore bool) *mysqlLogger {
	return &mysqlLogger{
		SlowThreshold:             slowThreshold,
		LogLevel:                  level,
		IgnoreRecordNotFoundError: ignore,
	}
}

var _ logger.Interface = (*mysqlLogger)(nil)

func (ml *mysqlLogger) LogMode(lev logger.LogLevel) logger.Interface {
	return &mysqlLogger{}
}

func (ml *mysqlLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	dlog.Info(ctx, TAGNAME, msg, args)
}

func (ml *mysqlLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	dlog.Warn(ctx, TAGNAME, msg, args)
}

func (ml *mysqlLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	dlog.ErrorWithMsg(ctx, TAGNAME, msg, args)
}

func (ml *mysqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil && (!errors.Is(err, gorm.ErrRecordNotFound) || !ml.IgnoreRecordNotFoundError) {
		dlog.ErrorWithMsg(ctx, "SQL_ERROR", "sql=%v, rows=%v, elapsed=%v, err=%v", sql, rows, elapsed, err)
		return
	}

	if ml.SlowThreshold != 0 && elapsed > ml.SlowThreshold {
		dlog.Warn(ctx, "SLOW_QUERY", "sql=%v, rows=%v, elapsed=%v, slowthreshold=%v", sql, rows, elapsed, ml.SlowThreshold)
	} else {
		dlog.Info(ctx, "MySQL_Query", "sql=%v, rows=%v, elapsed=%v", sql, rows, elapsed)
	}
}
