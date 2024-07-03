package logger

// Config 日志配置
type Config struct {
	Path                 string   `toml:"path"`                  // ⽇志⽂件路径
	closed               bool     `toml:"closed"`                // 是否关闭日志
	MaxSize              int      `toml:"max_size"`              // 单位为MB,默认为100MB
	MaxAge               int      `toml:"max_age"`               // 文件最多保存多少天
	MaxBackup            int      `toml:"max_backup"`            // 保留多少个备份
	TransparentParameter []string `toml:"transparent_parameter"` // 透传参数，继承上下文中的参数
}
