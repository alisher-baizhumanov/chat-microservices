package config

import (
	"log/slog"
	"time"

	"github.com/caarlos0/env/v11"
)

// Config holds configuration settings for the application.
type Config struct {
	GRPCServerPort int           `env:"GRPC_SERVER_PORT,required"`
	DSN            string        `env:"DB_DSN,required"`
	CacheDSN       string        `env:"CACHE_DSN,required"`
	CacheTTL       time.Duration `env:"CACHE_TTL,required"`
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
