package service

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// UserService defines the methods for managing users at the service layer.
type UserService interface {
	RegisterUser(ctx context.Context, userRegister model.UserRegister) (int64, error)
	GetByID(ctx context.Context, id int64) (model.User, error)
	UpdateUserFields(ctx context.Context, id int64, userUpdate model.UserUpdateOptions) error
	DeleteByID(ctx context.Context, id int64) error
}

// AuthService defines the methods for managing authentication at the service layer.
type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	CheckAccess(ctx context.Context, path, accessToken string) error
}
