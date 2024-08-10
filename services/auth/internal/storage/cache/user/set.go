package user

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/user/converter"
)

func (u *userCache) Set(ctx context.Context, userConverted model.User) error {
	user := converter.UserModelToCacheData(userConverted)

	key := strconv.FormatInt(userConverted.ID, 10)
	if err := u.cache.Set(
		ctx,
		key,
		user,
		u.ttl,
	); err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrCache, err)
	}

	return nil
}
