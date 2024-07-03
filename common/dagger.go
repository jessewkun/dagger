package common

import (
	"dagger/lib/mysql"
)

// 项目通用配置
type DaggerConfig struct {
	Mode string       `toml:"mode"`
	Port string       `toml:"port"`
	Db   mysql.Config `toml:"db"`
}

var DaggerCfg DaggerConfig
