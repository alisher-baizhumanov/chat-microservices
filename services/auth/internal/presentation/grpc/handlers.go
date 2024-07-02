package grpc

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
)

type userServer struct {
	desc.UnimplementedUserServiceV1Server
}

func (s *userServer) Create(ctx context.Context, createIn *desc.CreateIn) (*desc.CreateOut, error) {
	user := createIn.GetUserRegister()

	slog.InfoContext(ctx, "created user",
		slog.String("name", user.Name),
		slog.String("email", user.Email),
		slog.String("password", user.Password),
		slog.String("password_confirm", user.PasswordConfirm),
	)

	return &desc.CreateOut{Id: gofakeit.Int64()}, nil
}

func (s *userServer) Get(ctx context.Context, getIn *desc.GetIn) (*desc.GetOut, error) {
	slog.InfoContext(ctx, "get user", slog.Int64("id", getIn.Id))

	return &desc.GetOut{UserInfo: &desc.UserInfo{
		Id:        getIn.Id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}}, nil
}
