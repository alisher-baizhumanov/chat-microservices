package cache

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

type UserCache interface {
	Get(ctx context.Context, id int64) (model.User, error)
	Set(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id int64) error
}