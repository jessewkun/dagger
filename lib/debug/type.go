package debug

import "context"

// Config debug config
type Config struct {
	// debug模块, 可选值 mysql,http, + 自定义业务模块
	Module []string `toml:"module" mapstructure:"module"`

	// debug方式, 可选值 log, console
	Mode string `toml:"mode" mapstructure:"mode"`
}

type DebugFunc func(c context.Context, format string, v ...interface{})
