package config

import (
	"log/slog"

	"github.com/caarlos0/env/v11"
)

// Config holds configuration settings for the application.
type Config struct {
	GRPCServerPort int    `env:"GRPC_SERVER_PORT,required"`
	DSN            string `env:"DB_DSN,required"`
}

// Load attempts to load the application configuration.
func Load() (*Config, error) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
