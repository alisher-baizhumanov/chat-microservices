package config

import "log"

const port = 55051

type Config struct {
	GRPCServerPort int
}

func MustLoad() *Config {
	config, err := load()
	if err != nil {
		log.Fatal("failed to load config file", err)
	}

	return config
}

func load() (*Config, error) {
	return &Config{GRPCServerPort: port}, nil
}
