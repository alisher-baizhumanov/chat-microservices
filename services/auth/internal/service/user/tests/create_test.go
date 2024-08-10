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

func TestRegister(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	cases := []struct {
		name               string
		mc                 *minimock.Controller
		userRegister       model.UserRegister
		expID              int64
		expErr             error
		userRepositoryMock func(mc *minimock.Controller) repository.UserRepository
		userCacheMock      func(mc *minimock.Controller) cache.UserCache
	}{
		{
			name:         "success case create user",
			userRegister: userRegister,
			expID:        1,
			expErr:       nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return createUserRepositoryCreateMock(mc, t)
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				return createUserCacheMock(mc, t)
			},
		},
		{
			name: "error case user already exists",
			userRegister: model.UserRegister{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
			},
			expID:  0,
			expErr: model.ErrNonUniqueUsername,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Return(0, model.ErrNonUniqueUsername)
				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				return cacheMocks.NewUserCacheMock(mc)
			},
		},
		{
			name: "error case email already exists",
			userRegister: model.UserRegister{
				Name:            name,
				Email:           "email",
				Password:        password,
				PasswordConfirm: password,
			},
			expID:  0,
			expErr: model.ErrNonUniqueEmail,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Return(0, model.ErrNonUniqueEmail)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				return cacheMocks.NewUserCacheMock(mc)
			},
		},
		{
			name: "error case repository error",
			userRegister: model.UserRegister{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
			},
			expID:  0,
			expErr: model.ErrDatabase,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Return(0, model.ErrDatabase)
				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				return cacheMocks.NewUserCacheMock(mc)
			},
		},
		{
			name: "error case cache error",
			userRegister: model.UserRegister{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
			},
			expID:  1,
			expErr: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return createUserRepositoryCreateMock(mc, t)
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.SetMock.Return(model.ErrCache)

				return mock
			},
		},
	}

	for _, oneTest := range cases {
		test := oneTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := test.userRepositoryMock(mc)
			cacheMock := test.userCacheMock(mc)
			service := userService.New(repositoryMock, cacheMock)

			actualID, actualErr := service.RegisterUser(ctx, test.userRegister)

			require.Equal(t, test.expID, actualID)
			require.Equal(t, test.expErr, actualErr)
		})
	}
}
