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
	Mode  string                  `toml:"mode"`
	Port  string                  `toml:"port"`
	Debug debug.Config            `toml:"debug"`
	Http  http.Config             `toml:"http"`
	Log   logger.Config           `toml:"log"`
	Mysql map[string]mysql.Config `toml:"mysql"`
	Redis map[string]redis.Config `toml:"redis"`
}

var Cfg DaggerConfig
