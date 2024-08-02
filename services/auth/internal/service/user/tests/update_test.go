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

func TestUpdate(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx     context.Context
		id      int64
		options model.UserUpdateOptions
	}

	type output struct {
		err error
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = int64(1)
		name      = "name"
		email     = "example@gmail.com"
		role      = model.UserRole
		updatedAt = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		err       = error(nil)

		options = model.UserUpdateOptions{
			Name:  &name,
			Email: &email,
			Role:  &role,
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
			name: "success case update user",
			input: input{
				ctx:     ctx,
				id:      id,
				options: options,
			},
			output: output{
				err: err,
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.UpdateUserMock.Expect(ctx, id, options, updatedAt).Return(err)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(t)
				mock.DeleteMock.Expect(ctx, id).Return(nil)

				return mock
			},
			clock: clock.MockClock{CurrentTime: updatedAt},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.userRepositoryMock(mc)
			cache := tt.userCacheMock(mc)
			service := userService.New(repository, cache, tt.clock)

			err := service.UpdateUserFields(
				tt.input.ctx,
				tt.input.id,
				tt.input.options,
			)

			require.Equal(t, tt.output.err, err)
		})
	}
}
