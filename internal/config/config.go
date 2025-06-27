package config

type Config struct {
	Host   string
	Port   int
	Secret string
}

var Cfg Config

func NewConfig() *Config {
	Cfg = Config{
		Host:   "localhost",
		Port:   8089,
		Secret: "secret",
	}
	return &Cfg
}
