package repository

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, userCreate *model.UserCreate) (id int64, err error)
	GetUser(ctx context.Context, id int64) (user *model.User, err error)
	UpdateUser(ctx context.Context, id int64, userUpdate *model.UserUpdateOptions) (err error)
	DeleteUser(ctx context.Context, id int64) (err error)
}
