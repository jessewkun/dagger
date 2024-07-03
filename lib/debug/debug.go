package debug

import (
	"context"
	"fmt"

	"dagger/lib/logger"

	"github.com/spf13/viper"
)

const TAGNAME = "DAGGER_DEBUG"

// InitDebug 初始化debug
func InitDebug(flag string) DebugFunc {
	enable := IsDebug(flag)
	return func(c context.Context, format string, v ...interface{}) {
		if enable {
			hookPrint(c, format, v)
		}
	}
}

// IsDebug 是否开启debug
func IsDebug(flag string) bool {
	enable := false
	for _, part := range viper.GetStringSlice("debug.module") {
		if part == flag {
			enable = true
			break
		}
	}
	return enable
}

// hookPrint 输出debug信息
func hookPrint(c context.Context, format string, v ...interface{}) {
	if viper.GetString("debug.way") == "log" {
		logger.Debug(c, TAGNAME, format, v)
	} else {
		fmt.Sprintln(">>> "+format, v)
	}
}
