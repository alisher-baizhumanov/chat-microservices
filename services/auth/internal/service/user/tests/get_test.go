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

func TestGet(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx context.Context
		id  int64
	}

	type output struct {
		user model.User
		err  error
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = int64(1)
		name      = "name"
		email     = "example@gmail.com"
		role      = model.UserRole
		createdAt = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		err       = model.ErrNotFound

		userFromStorage = model.User{
			ID:        id,
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
			name: "success case get user",
			input: input{
				ctx: ctx,
				id:  id,
			},
			output: output{
				user: userFromStorage,
				err:  nil,
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetUserMock.Expect(ctx, id).Return(userFromStorage, nil)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, id).Return(model.User{}, err)
				mock.SetMock.Expect(ctx, userFromStorage).Return(nil)

				return mock
			},
			clock: clock.MockClock{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.userRepositoryMock(mc)
			cache := tt.userCacheMock(mc)
			service := userService.New(repository, cache, tt.clock)

			user, err := service.GetByID(tt.input.ctx, tt.input.id)

			require.Equal(t, tt.output.user, user)
			require.Equal(t, tt.output.err, err)
		})
	}
}
