package logger

// Config 日志配置
type Config struct {
	// ⽇志⽂件路径，绝对路径
	Path string `toml:"path" mapstructure:"path"`

	// 是否关闭日志
	closed bool `toml:"closed" mapstructure:"closed"`

	// 单位为MB,默认为100MB
	MaxSize int `toml:"max_size" mapstructure:"max_size"`

	// 文件最多保存多少天,默认7天
	MaxAge int `toml:"max_age" mapstructure:"max_age"`

	// 保留多少个备份,默认10
	MaxBackup int `toml:"max_backup" mapstructure:"max_backup"`

	// 透传参数，继承上下文中的参数
	TransparentParameter []string `toml:"transparent_parameter" mapstructure:"transparent_parameter"`
}
