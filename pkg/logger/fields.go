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

func Any(key string, value any) Field {
	return Field(slog.Any(key, value))
}

func convertFields(fields []Field) []any {
	if len(fields) == 0 {
		return nil
	}

	attrs := make([]any, len(fields))
	for i := range fields {
		attrs[i] = slog.Attr(fields[i])
	}

	return attrs
}
