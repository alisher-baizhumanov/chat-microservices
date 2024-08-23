package logger

import (
	"context"
	"io"
	"log/slog"
)

var (
	globalLogger *slog.Logger
	writeCloser  io.WriteCloser
)

// IsDebugEnabled checks if the debug level logging is enabled.
func IsDebugEnabled() bool {
	if globalLogger == nil {
		return false
	}

	return globalLogger.Enabled(context.Background(), slog.LevelDebug)
}

// Debug logs a message at the debug level.
func Debug(msg string, attrs ...Field) {
	if globalLogger == nil {
		return
	}

	globalLogger.Debug(msg, convertFields(attrs)...)
}

// Info logs a message at the info level.
func Info(msg string, attrs ...Field) {
	if globalLogger == nil {
		return
	}

	globalLogger.Info(msg, convertFields(attrs)...)
}

// Warn logs a message at the warn level.
func Warn(msg string, attrs ...Field) {
	if globalLogger == nil {
		return
	}

	globalLogger.Warn(msg, convertFields(attrs)...)
}

// Error logs a message at the error level.
func Error(msg string, attrs ...Field) {
	if globalLogger == nil {
		return
	}

	globalLogger.Error(msg, convertFields(attrs)...)
}

// Close closes the write closer if it is initialized.
func Close() error {
	if writeCloser == nil {
		return nil
	}

	return writeCloser.Close()
}
