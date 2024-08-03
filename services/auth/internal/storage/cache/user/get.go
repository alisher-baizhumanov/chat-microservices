package user

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/user/converter"
	cacheData "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/user/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

func (u *userCache) Get(ctx context.Context, id int64) (model.User, error) {
	var user cacheData.User

	key := strconv.FormatInt(id, 10)
	if err := u.cache.Get(ctx, key, &user); err != nil {
		if errors.Is(err, redis.Nil) {
			return model.User{}, model.ErrNotFound
		}

		return model.User{}, fmt.Errorf("%w, message: %w", model.ErrCache, err)
	}

	userModel, err := converter.UserCacheDataToModel(id, user)
	if err != nil {
		return model.User{}, fmt.Errorf("%w, message: %w", model.ErrTimeConverting, err)
	}

	return userModel, nil
}
