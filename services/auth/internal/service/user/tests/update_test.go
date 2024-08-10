package user_test

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
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

		expID = int64(1)
		name  = "name"
		email = "example@gmail.com"
		role  = model.UserRole
		err   = error(nil)

		expOptions = model.UserUpdateOptions{
			Name:  &name,
			Email: &email,
			Role:  &role,
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
			name: "success case update user",
			input: input{
				ctx:     ctx,
				id:      expID,
				options: expOptions,
			},
			output: output{
				err: err,
			},
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.UpdateUserMock.Inspect(func(_ context.Context, actualID int64, actualOptions model.UserUpdateOptions, _ time.Time) {
					require.Equal(t, expID, actualID)
					require.Equal(t, expOptions, actualOptions)
				}).Return(nil)

				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, expID).Return(nil)

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

			actualErr := service.UpdateUserFields(
				testCase.input.ctx,
				testCase.input.id,
				testCase.input.options,
			)

			require.Equal(t, testCase.output.err, actualErr)
		})
	}
}
