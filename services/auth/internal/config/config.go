package config

import (
	"log/slog"
	"net"
	"strconv"
	"time"

	"github.com/caarlos0/env/v11"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Config holds configuration settings for the application.
type Config struct {
	GRPCServerHost string        `env:"GRPC_SERVER_HOST,required"`
	GRPCServerPort int           `env:"GRPC_SERVER_PORT,required"`
	HTTPServerPort int           `env:"HTTP_SERVER_PORT,required"`
	DSN            string        `env:"DB_DSN,required"`
	CacheDSN       string        `env:"CACHE_DSN,required"`
	CacheTTL       time.Duration `env:"CACHE_TTL,required"`
}

// GRPCAddress returns the gRPC server address in the format "host:port".
// This is a helper method that constructs the address string from the
// GRPCServerHost and GRPCServerPort configuration values.
func (c Config) GRPCAddress() string {
	return net.JoinHostPort(c.GRPCServerHost, strconv.Itoa(c.GRPCServerPort))
}

// GRPCDialOptions returns the gRPC dial options to be used when connecting to the gRPC server.
// This includes using insecure credentials, as the gRPC server is assumed to be running locally.
func (c Config) GRPCDialOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
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
