package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheRedis struct {
	client *redis.Client
}

func (c *cacheRedis) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	if err := c.client.HSet(ctx, key, value).Err(); err != nil {
		return err
	}

	return c.client.Expire(ctx, key, ttl).Err()
}

func (c *cacheRedis) Get(ctx context.Context, key string, dest any) error {
	mapStringCmd := c.client.HGetAll(ctx, key)
	if mapStringCmd.Err() != nil {
		return mapStringCmd.Err()
	}

	if len(mapStringCmd.Val()) == 0 {
		return redis.Nil
	}

	return mapStringCmd.Scan(dest)
}

func (c *cacheRedis) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}