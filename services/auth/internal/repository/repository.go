package repository

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, userRegister *model.UserRegister) (id int64, err error)
	Get(ctx context.Context, id int64) (user *model.User, err error)
	Update(ctx context.Context, id int64, userUpdate *model.UserUpdateOptions) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
