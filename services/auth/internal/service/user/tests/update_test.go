package user_test

import (
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	cacheMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/mocks"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	cases := []struct {
		name               string
		id                 int64
		options            model.UserUpdateOptions
		expErr             error
		userRepositoryMock func(mc *minimock.Controller) repository.UserRepository
		userCacheMock      func(mc *minimock.Controller) cache.UserCache
	}{
		{
			name:    "success case update user",
			id:      id,
			options: updateOptions,
			expErr:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return createUserRepositoryUpdateMock(mc, t, id, nil)
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, int64(1)).Return(nil)

				return mock
			},
		},
		{
			name:    "error case invalid user ID",
			id:      0,
			options: updateOptions,
			expErr:  model.ErrInvalidID,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return createUserRepositoryUpdateMock(mc, t, 0, model.ErrInvalidID)
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, int64(0)).Return(model.ErrInvalidID)

				return mock
			},
		},
		{
			name:    "error case repository error",
			id:      id,
			options: updateOptions,
			expErr:  model.ErrDatabase,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return createUserRepositoryUpdateMock(mc, t, id, model.ErrDatabase)
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)

				return mock
			},
		},
		{
			name:    "error case cache error",
			id:      id,
			options: updateOptions,
			expErr:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return createUserRepositoryUpdateMock(mc, t, id, nil)
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(model.ErrCache)

				return mock
			},
		},
	}

	for _, oneCase := range cases {
		test := oneCase

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := test.userRepositoryMock(mc)
			cacheMock := test.userCacheMock(mc)
			service := userService.New(repositoryMock, cacheMock)

			actualErr := service.UpdateUserFields(
				ctx,
				test.id,
				test.options,
			)

			require.Equal(t, test.expErr, actualErr)
		})
	}
}
