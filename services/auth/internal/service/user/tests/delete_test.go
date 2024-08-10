package user_test

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
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

	testCases := []struct {
		name               string
		input              input
		output             output
		userRepositoryMock func(mc *minimock.Controller) repository.UserRepository
		userCacheMock      func(mc *minimock.Controller) cache.UserCache
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
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(err)

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

			err := service.DeleteByID(testCase.input.ctx, testCase.input.id)

			require.Equal(t, testCase.output.err, err)
		})
	}
}
