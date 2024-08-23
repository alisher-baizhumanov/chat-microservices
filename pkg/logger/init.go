package logger

import (
	"errors"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LogEnvironment represents the environment in which the logger is running.
type LogEnvironment string

const (
	// Development environment for development settings.
	Development LogEnvironment = "dev"

	// Production environment for production settings.
	Production = "prod"
)

// Constants for the logger writer settings.
const (
	maxSizeMegaBytes = 50
	maxAgeDays       = 14
	maxBackups       = 5
)

var (
	// ErrUnknownEnvironment is returned when an unknown environment is provided.
	ErrUnknownEnvironment = errors.New("unknown logger environment")
)

// Init initializes the global logger based on the provided environment.
func Init(env LogEnvironment) error {
	switch env {
	case Development:
		globalLogger = slog.New(tint.NewHandler(
			os.Stdout,
			&tint.Options{Level: slog.LevelDebug},
		))
	case Production:
		writeCloser = &lumberjack.Logger{
			Filename:   "logs/app.log",
			MaxSize:    maxSizeMegaBytes,
			MaxAge:     maxAgeDays,
			MaxBackups: maxBackups,
		}

		globalLogger = slog.New(slog.NewJSONHandler(
			writeCloser,
			&slog.HandlerOptions{Level: slog.LevelInfo},
		))
	default:
		return ErrUnknownEnvironment
	}

	slog.SetDefault(globalLogger)

	return nil
}
