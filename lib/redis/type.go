package redis

import "time"

// Config redis config
type Config struct {
	Addrs              []string      `json:"addrs" mapstructure:"addrs"`                               // redis addrs ip:port
	Password           string        `json:"password" mapstructure:"password"`                         // redis password
	Db                 int           `json:"db" mapstructure:"db"`                                     // redis db
	IsLog              bool          `toml:"is_log" mapstructure:"is_log"`                             // 是否记录日志
	PoolSize           int           `json:"pool_size" mapstructure:"pool_size"`                       // 连接池大小
	IdleTimeout        time.Duration `json:"idle_timeout" mapstructure:"idle_timeout"`                 // 空闲连接超时时间
	IdleCheckFrequency time.Duration `json:"idle_check_frequency" mapstructure:"idle_check_frequency"` // 空闲连接检查频率
	MinIdleConns       int           `json:"min_idle_conns" mapstructure:"min_idle_conns"`             // 最小空闲连接数
	MaxRetries         int           `json:"max_retries" mapstructure:"max_retries"`                   // 最大重试次数
	DialTimeout        time.Duration `json:"dial_timeout" mapstructure:"dial_timeout"`                 // 连接超时时间
}
