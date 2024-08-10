package user_test

import (
	"context"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	cacheMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/mocks"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	repositoryMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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
