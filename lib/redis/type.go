package redis

import "time"

// Config redis config
type Config struct {
	// redis addrs ip:port
	Addrs []string `json:"addrs" mapstructure:"addrs"`

	// redis password
	Password string `json:"password" mapstructure:"password"`

	// redis db, default 0
	Db int `json:"db" mapstructure:"db"`

	// 是否记录日志, default false,开启后会记录每次redis操作的耗时和命令
	IsLog bool `toml:"is_log" mapstructure:"is_log"`

	// 连接池大小，default 500
	PoolSize int `json:"pool_size" mapstructure:"pool_size"`

	// 空闲连接超时时间，default 1s
	IdleTimeout time.Duration `json:"idle_timeout" mapstructure:"idle_timeout"`

	// 空闲连接检查频率，default 10s
	IdleCheckFrequency time.Duration `json:"idle_check_frequency" mapstructure:"idle_check_frequency"`

	// 最小空闲连接数，default 3
	MinIdleConns int `json:"min_idle_conns" mapstructure:"min_idle_conns"`

	// 最大重试次数，default 3
	MaxRetries int `json:"max_retries" mapstructure:"max_retries"`

	// 连接超时时间，default 2s
	DialTimeout time.Duration `json:"dial_timeout" mapstructure:"dial_timeout"`
}
