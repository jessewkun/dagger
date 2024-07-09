package common

import (
	"dagger/lib/debug"
	"dagger/lib/http"
	"dagger/lib/logger"
	"dagger/lib/mysql"
	"dagger/lib/redis"
)

// 项目通用配置
type DaggerConfig struct {
	Mode  string                  `toml:"mode" mapstructure:"mode"`   // 运行模式, debug 开发, release 生产, test 测试
	Port  string                  `toml:"port" mapstructure:"port"`   // 服务端口, 默认 :8000
	Debug debug.Config            `toml:"debug" mapstructure:"debug"` // 调试配置
	Http  http.Config             `toml:"http" mapstructure:"http"`   // http 配置
	Log   logger.Config           `toml:"log" mapstructure:"log"`     // 日志配置
	Mysql map[string]mysql.Config `toml:"mysql" mapstructure:"mysql"` // mysql 配置
	Redis map[string]redis.Config `toml:"redis" mapstructure:"redis"` // redis 配置
}

// 项目配置
var Cfg DaggerConfig
