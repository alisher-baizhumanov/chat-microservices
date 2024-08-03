package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/clock"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
	cache "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	cacheMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/mocks"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	repositoryMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/mocks"
)

func TestRegister(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx  context.Context
		user model.UserRegister
	}

	type output struct {
		id  int64
		err error
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

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

	tests := []struct {
		name               string
		input              input
		output             output
		userRepositoryMock func(mc *minimock.Controller) repository.UserRepository
		userCacheMock      func(mc *minimock.Controller) cache.UserCache
		clock              clock.Clock
	}{
		{
			name: "success case create user",
			input: input{
				ctx:  ctx,
				user: userRegister,
			},
			output: output{
				id:  1,
				err: nil,
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Expect(ctx, userDB).Return(1, nil)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.SetMock.Expect(ctx, userCache).Return(nil)

				return mock
			},
			clock: clock.MockClock{CurrentTime: createdAt},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.userRepositoryMock(mc)
			cache := tt.userCacheMock(mc)
			service := userService.New(repository, cache, tt.clock)

			id, err := service.RegisterUser(tt.input.ctx, tt.input.user)

			require.Equal(t, tt.output.id, id)
			require.Equal(t, tt.output.err, err)
		})
	}
}
