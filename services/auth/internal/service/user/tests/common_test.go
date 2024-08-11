package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	cacheMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/mocks"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	repositoryMocks "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/mocks"
)

var (
	ctx = context.Background()

	id        = int64(1)
	name      = "name"
	email     = "example@gmail.com"
	role      = model.UserRole
	password  = []byte("secret_password")
	createdAt = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	updateOptions = model.UserUpdateOptions{
		Name:  &name,
		Email: &email,
		Role:  &role,
	}

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

	expUserInfo = model.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
)

func createUserRepositoryCreateMock(mc *minimock.Controller, t *testing.T) repository.UserRepository {
	mock := repositoryMocks.NewUserRepositoryMock(mc)
	mock.CreateUserMock.Inspect(func(_ context.Context, actualUserDB model.UserCreate) {
		require.Equal(t, expUserDB.Name, actualUserDB.Name)
		require.Equal(t, expUserDB.Email, actualUserDB.Email)
		require.Equal(t, expUserDB.Role, actualUserDB.Role)
		require.Equal(t, expUserDB.HashedPassword, actualUserDB.HashedPassword)
	}).Return(id, nil)

	return mock
}

func createUserCacheMock(mc *minimock.Controller, t *testing.T) cache.UserCache {
	mock := cacheMocks.NewUserCacheMock(mc)
	mock.SetMock.Inspect(func(_ context.Context, actualUserCache model.User) {
		require.Equal(t, expUserInfo.ID, actualUserCache.ID)
		require.Equal(t, expUserInfo.Name, actualUserCache.Name)
		require.Equal(t, expUserInfo.Email, actualUserCache.Email)
		require.Equal(t, expUserInfo.Role, actualUserCache.Role)
	}).Return(nil)

	return mock
}

func createUserRepositoryUpdateMock(mc *minimock.Controller, t *testing.T, expID int64, returnErr error) repository.UserRepository {
	mock := repositoryMocks.NewUserRepositoryMock(mc)
	mock.UpdateUserMock.Inspect(func(_ context.Context, id int64, actualUserUpdate model.UserUpdateOptions, _ time.Time) {
		require.Equal(t, expID, id)
		require.Equal(t, updateOptions, actualUserUpdate)
	}).Return(returnErr)

	return mock
}
