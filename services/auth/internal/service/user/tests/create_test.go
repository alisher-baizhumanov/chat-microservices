package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	cacheMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/mocks"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	repositoryMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/mocks"
)

var (
	ctx = context.Background()

	name      = "name"
	email     = "example@gmail.com"
	role      = model.UserRole
	password  = []byte("secret_password")
	createdAt = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	userRegister = model.UserRegister{
		Name:            name,
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
	}

	expUserDB = model.UserCreate{
		Name:           name,
		Email:          email,
		Role:           role,
		CreatedAt:      createdAt,
		HashedPassword: password,
	}

	expUserCache = model.User{
		ID:        1,
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
)

func createUserRepositoryMock(mc *minimock.Controller, t *testing.T) repository.UserRepository {
	mock := repositoryMocks.NewUserRepositoryMock(mc)
	mock.CreateUserMock.Inspect(func(_ context.Context, actualUserDB model.UserCreate) {
		require.Equal(t, expUserDB.Name, actualUserDB.Name)
		require.Equal(t, expUserDB.Email, actualUserDB.Email)
		require.Equal(t, expUserDB.Role, actualUserDB.Role)
		require.Equal(t, expUserDB.HashedPassword, actualUserDB.HashedPassword)
	}).Return(1, nil)

	return mock
}

func createUserCacheMock(mc *minimock.Controller, t *testing.T) cache.UserCache {
	mock := cacheMocks.NewUserCacheMock(mc)
	mock.SetMock.Inspect(func(_ context.Context, actualUserCache model.User) {
		require.Equal(t, expUserCache.ID, actualUserCache.ID)
		require.Equal(t, expUserCache.Name, actualUserCache.Name)
		require.Equal(t, expUserCache.Email, actualUserCache.Email)
		require.Equal(t, expUserCache.Role, actualUserCache.Role)
	}).Return(nil)

	return mock
}

func TestRegister(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)

	cases := []struct {
		name               string
		mc                 *minimock.Controller
		input              model.UserRegister
		expectedID         int64
		expectedErr        error
		userRepositoryMock func(mc *minimock.Controller) repository.UserRepository
		userCacheMock      func(mc *minimock.Controller) cache.UserCache
	}{
		{
			name:        "success case create user",
			input:       userRegister,
			expectedID:  1,
			expectedErr: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return createUserRepositoryMock(mc, t)
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				return createUserCacheMock(mc, t)
			},
		},
		{
			name: "error case user already exists",
			input: model.UserRegister{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
			},
			expectedID:  0,
			expectedErr: model.ErrNonUniqueUsername,
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
			input: model.UserRegister{
				Name:            name,
				Email:           "email",
				Password:        password,
				PasswordConfirm: password,
			},
			expectedID:  0,
			expectedErr: model.ErrNonUniqueEmail,
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
			input: model.UserRegister{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
			},
			expectedID:  0,
			expectedErr: model.ErrDatabase,
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
			input: model.UserRegister{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
			},
			expectedID:  1,
			expectedErr: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return createUserRepositoryMock(mc, t)
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

			id, err := service.RegisterUser(ctx, test.input)

			require.Equal(t, test.expectedID, id)
			require.Equal(t, test.expectedErr, err)
		})
	}
}
