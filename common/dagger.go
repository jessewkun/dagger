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
	Mode  string                  `toml:"mode" mapstructure:"mode"`
	Port  string                  `toml:"port" mapstructure:"port"`
	Debug debug.Config            `toml:"debug" mapstructure:"debug"`
	Http  http.Config             `toml:"http" mapstructure:"http"`
	Log   logger.Config           `toml:"log" mapstructure:"log"`
	Mysql map[string]mysql.Config `toml:"mysql" mapstructure:"mysql"`
	Redis map[string]redis.Config `toml:"redis" mapstructure:"redis"`
}

var Cfg DaggerConfig
