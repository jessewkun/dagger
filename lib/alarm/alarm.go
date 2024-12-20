package alarm

type Config struct {
	BarkIds []string `toml:"bark_ids" mapstructure:"bark_ids"`
}
