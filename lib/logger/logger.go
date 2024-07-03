package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logzap *zap.Logger

func Zap() *zap.Logger {
	return logzap
}

func InitLogger() {
	logzap = zap.New(initCore(), zap.AddCallerSkip(1), zap.AddCaller())
}

func initCore() zapcore.Core {
	logPath := "/dev/null"
	if closed := viper.GetBool("log.closed"); !closed {
		logPath = viper.GetString("log.path")
	}

	maxSize := viper.GetInt("log.max_size")
	if maxSize == 0 {
		maxSize = 5120
	}

	maxAge := viper.GetInt("log.max_age")
	if maxSize == 0 {
		maxAge = 7
	}

	maxBackup := viper.GetInt("log.max_backup")
	if maxBackup == 0 { // 如果未设置则默认保留5个
		maxBackup = 5
	}

	if maxBackup < 0 { // 如果设置为 -1 代表保留全部
		maxBackup = 0
	}

	opts := []zapcore.WriteSyncer{
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath, // ⽇志⽂件路径
			MaxSize:    maxSize, // 单位为MB,默认为100MB
			MaxAge:     maxAge,  // 文件最多保存多少天
			LocalTime:  true,    // 采用本地时间
			Compress:   false,   // 是否压缩日志
			MaxBackups: maxBackup,
		}),
	}

	syncWriter := zapcore.NewMultiWriteSyncer(opts...)

	// 自定义时间编码器
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller_line",
		LevelKey:       "level",
		MessageKey:     "msg",
		TimeKey:        "datetime",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,           // 自定义时间格式
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

	if tag != "" {
		fields = append(fields, zap.String("tag", tag))
	}

	hostname, _ := os.Hostname()
	fields = append(fields, zap.String("host", hostname))

	if customeParameter := viper.GetStringSlice("log.custome_parameter"); len(customeParameter) > 0 {
		for _, v := range customeParameter {
			if value := c.Value(v); value != nil {
				if valueInt, ok := value.(int); ok {
					fields = append(fields, zap.Int(v, valueInt))
				} else {
					fields = append(fields, zap.String(v, cast.ToString(value)))
				}
			}
		}
	}

	return fields
}

// Info log
func Info(c context.Context, tag string, template interface{}, args ...interface{}) {
	var msg string

	if tpl, flag := template.(string); flag {
		msg = fmt.Sprintf(tpl, args...)
	}

	if tpl, flag := template.(map[string]interface{}); flag {
		msg, _ = jsoniter.MarshalToString(tpl)
	}

	fields := formatField(c, tag)
	logzap.Info(msg, fields...)
}

// Error log
func Error(c context.Context, tag string, err error) {
	fields := formatField(c, tag)
	logzap.Error(err.Error(), fields...)
}

// ErrorWithMsg log
func ErrorWithMsg(c context.Context, tag string, template interface{}, args ...interface{}) {
	var msg string
	if tpl, flag := template.(string); flag {
		msg = fmt.Sprintf(tpl, args...)
	}

	if tpl, flag := template.(map[string]interface{}); flag {
		msg, _ = jsoniter.MarshalToString(tpl)
	}

	fields := formatField(c, tag)
	logzap.Error(msg, fields...)
}

// Debug log
func Debug(c context.Context, tag string, template interface{}, args ...interface{}) {
	var msg string
	if tpl, flag := template.(string); flag {
		msg = fmt.Sprintf(tpl, args...)
	}

	if tpl, flag := template.(map[string]interface{}); flag {
		msg, _ = jsoniter.MarshalToString(tpl)
	}

	fields := formatField(c, tag)
	logzap.Debug(msg, fields...)
}

// Warn log
func Warn(c context.Context, tag string, template interface{}, args ...interface{}) {
	var msg string
	if tpl, flag := template.(string); flag {
		msg = fmt.Sprintf(tpl, args...)
	}

	if tpl, flag := template.(map[string]interface{}); flag {
		msg, _ = jsoniter.MarshalToString(tpl)
	}

	fields := formatField(c, tag)
	logzap.Warn(msg, fields...)
}

// Panic log
func Panic(c context.Context, tag string, template interface{}, args ...interface{}) {
	var msg string
	if tpl, flag := template.(string); flag {
		msg = fmt.Sprintf(tpl, args...)
	}

	if tpl, flag := template.(map[string]interface{}); flag {
		msg, _ = jsoniter.MarshalToString(tpl)
	}

	fields := formatField(c, tag)
	logzap.Panic(msg, fields...)
}

// Fatal log
func Fatal(c context.Context, tag string, template interface{}, args ...interface{}) {
	var msg string
	if tpl, flag := template.(string); flag {
		msg = fmt.Sprintf(tpl, args...)
	}

	if tpl, flag := template.(map[string]interface{}); flag {
		msg, _ = jsoniter.MarshalToString(tpl)
	}

	fields := formatField(c, tag)
	logzap.Fatal(msg, fields...)
}
