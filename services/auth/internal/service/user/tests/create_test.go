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

	testCases := []struct {
		name               string
		input              input
		output             output
		userRepositoryMock func(mc *minimock.Controller) repository.UserRepository
		userCacheMock      func(mc *minimock.Controller) cache.UserCache
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
				mock.CreateUserMock.Inspect(func(_ context.Context, actualUserDB model.UserCreate) {
					require.Equal(t, expUserDB.Name, actualUserDB.Name)
					require.Equal(t, expUserDB.Email, actualUserDB.Email)
					require.Equal(t, expUserDB.Role, actualUserDB.Role)
					require.Equal(t, expUserDB.HashedPassword, actualUserDB.HashedPassword)
				}).Return(1, nil)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.SetMock.Inspect(func(_ context.Context, actualUserCache model.User) {
					require.Equal(t, expUserCache.ID, actualUserCache.ID)
					require.Equal(t, expUserCache.Name, actualUserCache.Name)
					require.Equal(t, expUserCache.Email, actualUserCache.Email)
					require.Equal(t, expUserCache.Role, actualUserCache.Role)
				}).Return(nil)

				return mock
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := testCase.userRepositoryMock(mc)
			cacheMock := testCase.userCacheMock(mc)
			service := userService.New(repositoryMock, cacheMock)

			id, err := service.RegisterUser(testCase.input.ctx, testCase.input.user)

			require.Equal(t, testCase.output.id, id)
			require.Equal(t, testCase.output.err, err)
		})
	}
}
