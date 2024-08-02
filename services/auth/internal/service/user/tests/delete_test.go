package user_test

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/clock"
	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
	cache "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	cacheMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/mocks"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	repositoryMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/mocks"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type input struct {
		ctx context.Context
		id  int64
	}

	type output struct {
		err error
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id  = int64(1)
		err = error(nil)
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
			name: "success case delete user",
			input: input{
				ctx: ctx,
				id:  1,
			},
			output: output{
				err: nil,
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.DeleteUserMock.Expect(ctx, id).Return(err)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(t)
				mock.DeleteMock.Expect(ctx, id).Return(err)

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

			err := service.DeleteByID(tt.input.ctx, tt.input.id)

			require.Equal(t, tt.output.err, err)
		})
	}
}
