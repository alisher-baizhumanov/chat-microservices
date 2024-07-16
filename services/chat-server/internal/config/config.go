package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/google/uuid"
)

// Config holds configuration settings for the application.
type Config struct {
	GRPCServerPort int    `env:"GRPC_SERVER_PORT,required"`
	DSN            string `env:"DB_DSN,required"`
}

// Load attempts to load the application configuration.
func Load() (*Config, error) {
	uuid.EnableRandPool()

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
