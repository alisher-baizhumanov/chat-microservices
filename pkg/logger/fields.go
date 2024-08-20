package logger

import (
	"log/slog"
	"time"
)

// Field is a type alias for slog.Attr, representing a single log attribute.
type Field slog.Attr

// String creates a new Field with a string value.
func String(key, value string) Field {
	return Field(slog.String(key, value))
}

// Int creates a new Field with an int value.
func Int(key string, value int) Field {
	return Field(slog.Int(key, value))
}

// Int64 creates a new Field with an int64 value.
func Int64(key string, value int64) Field {
	return Field(slog.Int64(key, value))
}

// Duration creates a new Field with a time.Duration value.
func Duration(key string, value time.Duration) Field {
	return Field(slog.Duration(key, value))
}

// Any creates a new Field with a value of any type.
func Any(key string, value any) Field {
	return Field(slog.Any(key, value))
}

// convertFields converts a slice of Field to a slice of any type.
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
