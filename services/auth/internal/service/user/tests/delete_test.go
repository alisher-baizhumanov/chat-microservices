package user_test

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	cacheMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/mocks"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	repositoryMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/mocks"
)

func TestDeleteByID(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	cases := []struct {
		name               string
		id                 int64
		expErr             error
		userRepositoryMock func(mc *minimock.Controller) repository.UserRepository
		userCacheMock      func(mc *minimock.Controller) cache.UserCache
	}{
		{
			name:   "success case delete user",
			id:     1,
			expErr: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.DeleteUserMock.Expect(ctx, id).Return(nil)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)

				return mock
			},
		},
		{
			name:   "error case invalid user ID",
			id:     0,
			expErr: model.ErrInvalidID,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.DeleteUserMock.Expect(ctx, int64(0)).Return(model.ErrInvalidID)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, int64(0)).Return(model.ErrInvalidID)

				return mock
			},
		},
		{
			name:   "error case repository error",
			id:     1,
			expErr: model.ErrDatabase,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.DeleteUserMock.Expect(ctx, id).Return(model.ErrDatabase)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)

				return mock
			},
		},
		{
			name:   "error case cache error",
			id:     1,
			expErr: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.DeleteUserMock.Expect(ctx, id).Return(nil)

				return mock
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

			actualErr := service.DeleteByID(ctx, test.id)

			require.Equal(t, test.expErr, actualErr)
		})
	}
}
