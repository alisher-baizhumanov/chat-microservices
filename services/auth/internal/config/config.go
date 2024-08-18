package config

import (
	"net"
	"strconv"
	"time"

	"github.com/caarlos0/env/v11"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Config holds configuration settings for the application.
type Config struct {
	Env string `env:"ENVIRONMENT,required"`

	GRPCServerHost string `env:"GRPC_SERVER_HOST,required"`
	GRPCServerPort int    `env:"GRPC_SERVER_PORT,required"`
	HTTPServerPort int    `env:"HTTP_SERVER_PORT,required"`

	DatabaseDSN string        `env:"DB_DSN,required"`
	CacheDSN    string        `env:"CACHE_DSN,required"`
	CacheTTL    time.Duration `env:"CACHE_TTL,required"`

	AccessTokenTTL   time.Duration `env:"ACCESS_TOKEN_TTL,required"`
	RefreshTokenTTL  time.Duration `env:"REFRESH_TOKEN_TTL,required"`
	AccessSecretKey  string        `env:"ACCESS_SECRET_KEY,required"`
	RefreshSecretKey string        `env:"REFRESH_SECRET_KEY,required"`
}

// GRPCAddress returns the gRPC server address in the format "host:port".
// This is a helper method that constructs the address string from the
// GRPCServerHost and GRPCServerPort configuration values.
func (c Config) GRPCAddress() string {
	return net.JoinHostPort(c.GRPCServerHost, strconv.Itoa(c.GRPCServerPort))
}

// GRPCClientDialOptions returns the gRPC dial options to be used when connecting to the gRPC server.
// This includes using insecure credentials, as the gRPC server is assumed to be running locally.
func (c Config) GRPCClientDialOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
}

// Load attempts to load the application configuration.
func Load() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
