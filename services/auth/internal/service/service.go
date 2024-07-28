package service

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// UserService defines the methods for managing users at the service layer.
type UserService interface {
	RegisterUser(ctx context.Context, userRegister *model.UserRegister) (int64, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	UpdateUserFields(ctx context.Context, id int64, userUpdate *model.UserUpdateOptions) error
	DeleteByID(ctx context.Context, id int64) error
}
