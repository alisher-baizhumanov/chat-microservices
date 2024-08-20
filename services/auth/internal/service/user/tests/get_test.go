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
	repositoryMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/mocks"
)

func TestGetByID(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	cases := []struct {
		name               string
		id                 int64
		expUser            model.User
		err                error
		userRepositoryMock func(mc *minimock.Controller) repository.UserRepository
		userCacheMock      func(mc *minimock.Controller) cache.UserCache
	}{
		{
			name: "success case get user",
			id:   id,

			expUser: expUserInfo,
			err:     nil,

			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetUserMock.Expect(ctx, id).Return(expUserInfo, nil)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, id).Return(model.User{}, model.ErrNotFound)
				mock.SetMock.Expect(ctx, expUserInfo).Return(nil)

				return mock
			},
		},
		{
			name: "error case user not found",
			id:   id,

			expUser: model.User{},
			err:     model.ErrNotFound,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetUserMock.Expect(ctx, id).Return(model.User{}, model.ErrNotFound)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, id).Return(model.User{}, model.ErrNotFound)

				return mock
			},
		},
		{
			name:    "error case invalid user ID",
			id:      0,
			expUser: model.User{},
			err:     model.ErrInvalidID,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetUserMock.Expect(ctx, int64(0)).Return(model.User{}, model.ErrInvalidID)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, int64(0)).Return(model.User{}, model.ErrInvalidID)

				return mock
			},
		},
		{
			name:    "error case repository error",
			id:      id,
			expUser: model.User{},
			err:     model.ErrDatabase,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetUserMock.Expect(ctx, id).Return(model.User{}, model.ErrDatabase)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, id).Return(model.User{}, model.ErrNotFound)

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
			service := userService.New(repositoryMock, cacheMock, nil)

			user, err := service.GetByID(ctx, test.id)

			require.Equal(t, test.expUser, user)
			require.Equal(t, test.err, err)
		})
	}
}
