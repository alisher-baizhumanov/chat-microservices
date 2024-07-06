package service

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, userRegister *model.UserRegister) (id int64, err error)
	GetById(ctx context.Context, id int64) (user *model.User, err error)
	UpdateUserFields(ctx context.Context, id int64, userUpdate *model.UserUpdateOptions) (err error)
	DeleteById(ctx context.Context, id int64) (err error)
}
