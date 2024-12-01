package config

type Config struct {
	Address string `yaml:"address" json:"address" `
}

func (c *Config) SetDefault() {

}
