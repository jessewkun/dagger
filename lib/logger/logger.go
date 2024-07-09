package logger

import (
	"context"
	"dagger/utils"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logzap *zap.Logger
var logcfg Config

func Zap() *zap.Logger {
	return logzap
}

func InitLogger(cfg Config) {
	logcfg = cfg
	logzap = zap.New(initCore(), zap.AddCallerSkip(1), zap.AddCaller())
}

func initCore() zapcore.Core {
	opts := []zapcore.WriteSyncer{
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   logcfg.Path,      // ⽇志⽂件路径
			MaxSize:    logcfg.MaxSize,   // 单位为MB,默认为100MB
			MaxAge:     logcfg.MaxAge,    // 文件最多保存多少天
			LocalTime:  true,             // 采用本地时间
			Compress:   false,            // 是否压缩日志
			MaxBackups: logcfg.MaxBackup, // 保留旧文件的最大个数
		}),
	}

	syncWriter := zapcore.NewMultiWriteSyncer(opts...)

	encoderConf := zapcore.EncoderConfig{
		CallerKey:     "caller_line",
		LevelKey:      "level",
		MessageKey:    "msg",
		TimeKey:       "datetime",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}, // 自定义时间格式
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 日志级别改为大写
		EncodeCaller:   zapcore.FullCallerEncoder,   // 全路径编码器
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	return zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf),
		syncWriter, zap.NewAtomicLevelAt(zapcore.DebugLevel))
}

// formatField 格式化字段
func formatField(c context.Context, tag string) []zapcore.Field {
	fields := make([]zapcore.Field, 0)
	fields = append(fields, zap.String("tag", tag))

	hostname, _ := os.Hostname()
	fields = append(fields, zap.String("host", hostname))
	ip, _ := utils.GetLocalIP()
	fields = append(fields, zap.String("ip", ip))

	if len(logcfg.TransparentParameter) > 0 {
		for _, v := range logcfg.TransparentParameter {
			if value := c.Value(v); value != nil {
				fields = append(fields, zap.Any(v, value))
			}
		}
	}

	return fields
}

// Info log
func Info(c context.Context, tag string, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	fields := formatField(c, tag)
	logzap.Info(msg, fields...)
}

// InfoWithField log
func InfoWithField(c context.Context, tag string, msg string, field map[string]interface{}) {
	fields := formatField(c, tag)
	for k, v := range field {
		fields = append(fields, zap.Any(k, v))
	}
	logzap.Info(msg, fields...)
}

// Error log
func Error(c context.Context, tag string, err error) {
	fields := formatField(c, tag)
	logzap.Error(err.Error(), fields...)
}

// ErrorWithMsg log
func ErrorWithMsg(c context.Context, tag string, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	fields := formatField(c, tag)
	logzap.Error(msg, fields...)
}

// Debug log
func Debug(c context.Context, tag string, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	fields := formatField(c, tag)
	logzap.Debug(msg, fields...)
}

// Warn log
func Warn(c context.Context, tag string, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	fields := formatField(c, tag)
	logzap.Warn(msg, fields...)
}

// Panic log
func Panic(c context.Context, tag string, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	fields := formatField(c, tag)
	logzap.Panic(msg, fields...)
}

// Fatal log
func Fatal(c context.Context, tag string, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	fields := formatField(c, tag)
	logzap.Fatal(msg, fields...)
}
