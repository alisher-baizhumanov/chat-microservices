package logger

import (
	"log/slog"
	"time"
)

type Field slog.Attr

func String(key, value string) Field {
	return Field(slog.String(key, value))
}

func Int(key string, value int) Field {
	return Field(slog.Int(key, value))
}

func Int64(key string, value int64) Field {
	return Field(slog.Int64(key, value))
}

func Duration(key string, value time.Duration) Field {
	return Field(slog.Duration(key, value))
}
