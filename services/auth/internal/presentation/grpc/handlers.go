package grpc

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
)

type userServer struct {
	desc.UnimplementedUserServiceV1Server
}

func (s *userServer) Create(ctx context.Context, createIn *desc.CreateIn) (*desc.CreateOut, error) {
	user := createIn.GetUserRegister()

	slog.InfoContext(ctx, "created user",
		slog.String("name", user.GetName()),
		slog.String("email", user.GetEmail()),
		slog.String("password", user.GetPassword()),
		slog.String("password_confirm", user.GetPasswordConfirm()),
	)

	return &desc.CreateOut{Id: gofakeit.Int64()}, nil
}

func (s *userServer) Get(ctx context.Context, getIn *desc.GetIn) (*desc.GetOut, error) {
	slog.InfoContext(ctx, "get user", slog.Int64("id", getIn.GetId()))

	return &desc.GetOut{UserInfo: &desc.UserInfo{
		Id:        getIn.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}}, nil
}

func (s *userServer) Update(ctx context.Context, updateIn *desc.UpdateIn) (*emptypb.Empty, error) {
	updateArgs := updateIn.GetUserUpdate()

	slog.InfoContext(ctx, "update user",
		slog.With("id", updateIn.GetId()),
		slog.With("role", updateArgs.GetRole()),
		slog.With("name", updateArgs.GetName()),
		slog.With("email", updateArgs.GetEmail()),
	)

	return nil, nil
}

func (s *userServer) Delete(ctx context.Context, deleteIn *desc.DeleteIn) (*emptypb.Empty, error) {
	slog.InfoContext(ctx, "delete user",
		slog.Int64("id", deleteIn.GetId()),
	)

	return nil, nil
}
