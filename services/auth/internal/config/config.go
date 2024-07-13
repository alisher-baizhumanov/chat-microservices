package config

const port = 55051

// Config holds configuration settings for the application.
type Config struct {
	GRPCServerPort int
}

// Load attempts to load the application configuration.
func Load() (*Config, error) {
	return &Config{GRPCServerPort: port}, nil
}
