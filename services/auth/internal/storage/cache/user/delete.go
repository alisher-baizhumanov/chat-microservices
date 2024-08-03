package user

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

func (u *userCache) Delete(ctx context.Context, id int64) error {
	key := strconv.FormatInt(id, 10)
	if err := u.cache.Delete(ctx, key); err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrCache, err)
	}

	return nil
}
