package config

import "log"

const port = 55051

// Config holds configuration settings for the application.
type Config struct {
	GRPCServerPort int
}

// MustLoad attempts to load the application configuration.
// It returns a pointer to Config. If loading fails, it logs a fatal error.
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
