package debug

import "context"

// Config debug config
type Config struct {
	Module []string `toml:"module"` // debug模块, 可选值 mysql,http, + 自定义业务模块
	Way    string   `toml:"way"`    // debug输出方式, log, console
}

type DebugFunc func(c context.Context, format string, v ...interface{})
