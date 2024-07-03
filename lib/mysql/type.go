package mysql

// Config mysql config
type Config struct {
	Driver      string   `toml:"driver"`        // 驱动
	Dsn         []string `toml:"dsn"`           // 数据源
	MaxConn     int      `toml:"max_conn"`      // 最大连接数
	MaxIdleConn int      `toml:"max_idle_conn"` // 最大空闲连接数
	ConnMaxLife int      `toml:"conn_max_life"` // 连接最长持续时间， 默认1小时，单位秒
	IsLog       bool     `toml:"is_log"`        // 是否记录日志  日志级别为info
}
