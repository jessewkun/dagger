package redis

// Config redis config
type Config struct {
	Nodes    []string `json:"nodes"`    // redis nodes ip:port
	Password string   `json:"password"` // redis password
	Db       int      `json:"db"`       // redis db
}
