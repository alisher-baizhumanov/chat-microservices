package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	cache "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	cacheMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/mocks"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	repositoryMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/mocks"
)

func TestRegister(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx  context.Context
		user model.UserRegister
	}

	var (
		ctx = context.Background()
		_  = minimock.NewController(t)

		name      = "name"
		email     = "example@gmail.com"
		role      = model.UserRole
		password  = []byte("secret_password")
		createdAt = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

		userRegister = model.UserRegister{
			Name:            "name",
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
		}

		userDB = model.UserCreate{
			Name:           name,
			Email:          email,
			Role:           role,
			CreatedAt:      createdAt,
			HashedPassword: password,
		}

		userCache = model.User{
			ID:        1,
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		}
	)

	_ = []struct {
		name               string
		ctx                context.Context
		user               model.UserRegister
		id                 int64
		err                error
		userRepositoryMock func(mc *minimock.Controller) repository.UserRepository
		userCacheMock      func(mc *minimock.Controller) cache.UserCache
	}{
		{
			name: "success case",
			ctx:  ctx,
			user: userRegister,
			id:   1,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Expect(ctx, userDB).Return(1, nil)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(t)
				mock.SetMock.Expect(ctx, userCache).Return(nil)

				return mock
			},
		},
	}
}
