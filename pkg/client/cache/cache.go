package cache

import (
	"context"
	"time"
)

type Client interface {
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
	Cache() Cache
}

type Cache interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string, dest any) error
	Delete(ctx context.Context, key string) error
}
