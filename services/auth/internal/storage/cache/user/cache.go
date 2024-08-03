package user

import (
	"time"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/cache"
	cacheInterface "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
)

type userCache struct {
	cache cache.Cache
	ttl time.Duration
}

func NewCache(cache cache.Cache, ttl time.Duration) cacheInterface.UserCache {
	return &userCache{cache: cache, ttl: ttl}
}
