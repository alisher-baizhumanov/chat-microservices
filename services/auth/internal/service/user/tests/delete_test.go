package user_test

import (
	"context"
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

	type input struct {
		ctx context.Context
		id  int64
	}

	type output struct {
		err error
	}

	var (
		mc = minimock.NewController(t)

		id  = int64(1)
		err = error(nil)
	)

	cases := []struct {
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
		{
			name: "error case invalid user ID",
			input: input{
				ctx: ctx,
				id:  0,
			},
			output: output{
				err: model.ErrInvalidID,
			},
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
			name: "error case repository error",
			input: input{
				ctx: ctx,
				id:  1,
			},
			output: output{
				err: model.ErrDatabase,
			},
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
			name: "error case cache error",
			input: input{
				ctx: ctx,
				id:  1,
			},
			output: output{
				err: nil,
			},
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

			err := service.DeleteByID(test.input.ctx, test.input.id)

			require.Equal(t, test.output.err, err)
		})
	}
}
