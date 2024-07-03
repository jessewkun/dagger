package debug

import (
	"context"
	"fmt"
	"strings"

	"dagger/lib/logger"

	"github.com/spf13/viper"
)

type Debug func(c context.Context, format string, v ...interface{})

const TAGNAME = "DAGGER_DEBUG"

// hookPrint 输出debug信息
var hookPrint = func(c context.Context, input string) {
	if viper.GetString("debug.way") == "log" {
		logger.Debug(c, TAGNAME, input)
	} else {
		fmt.Println(input)
	}
}

// defaultDebugSwitch 默认的debug hook
var defaultDebugSwitch = func() string {
	return viper.GetString("debug.module")
}

// InitDebug 初始化debug
func InitDebug(flag string) Debug {
	enable := IsDebug(flag)
	return func(c context.Context, format string, v ...interface{}) {
		if enable {
			hookPrint(c, fmt.Sprintf(">>>"+format, v...))
		}
	}
}

// IsDebug 是否开启debug
func IsDebug(flag string) bool {
	enable := false
	switcher := viper.GetString("debug.module")
	parts := strings.Split(switcher, ",")
	for _, part := range parts {
		if part == flag {
			enable = true
			break
		}
	}
	return enable
}
