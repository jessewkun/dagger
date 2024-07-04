package redis

// Config redis config
type Config struct {
	Nodes    []string `json:"nodes" mapstructure:"nodes"`       // redis nodes ip:port
	Password string   `json:"password" mapstructure:"password"` // redis password
	Db       int      `json:"db" mapstructure:"db"`             // redis db
}
