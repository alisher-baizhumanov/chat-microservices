package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	cacheInterface "github.com/alisher-baizhumanov/chat-microservices/pkg/client/cache"
)

type clientRedis struct {
	client *redis.Client
}

func NewClient(addr string) (cacheInterface.Client) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &clientRedis{
		client: client,
	}
}

func (c *clientRedis) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

func (c *clientRedis) Close(_ context.Context) error {
	return c.client.Close()
}

func (c *clientRedis) Cache() cacheInterface.Cache {
	return &cacheRedis{client: c.client}
}
