package grpc

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
)

type ServerHandlers struct {
	desc.UnimplementedUserServiceV1Server
	userService service.UserService
}

func NewUserGRPCHandlers(userService service.UserService) *ServerHandlers {
	return &ServerHandlers{userService: userService}
}

func (h *ServerHandlers) Create(ctx context.Context, createIn *desc.CreateIn) (*desc.CreateOut, error) {
	user := createIn.UserRegister

	slog.InfoContext(ctx, "created user",
		slog.String("name", user.Name),
		slog.String("email", user.Email),
		slog.String("password", user.Password),
		slog.String("password_confirm", user.PasswordConfirm),
	)

	return &desc.CreateOut{Id: gofakeit.Int64()}, nil
}

func (h *ServerHandlers) Get(ctx context.Context, getIn *desc.GetIn) (*desc.GetOut, error) {
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

func (h *ServerHandlers) Update(ctx context.Context, updateIn *desc.UpdateIn) (*emptypb.Empty, error) {
	updateArgs := updateIn.UserUpdate

	log := slog.With("id", updateIn.Id)

	if updateArgs.Role != desc.Role_NULL {
		log = log.With("role", updateArgs.Role)
	}

	if updateArgs.Name != nil {
		log = log.With("name", updateArgs.Name.Value)
	}

	if updateArgs.Email != nil {
		log = log.With("email", updateArgs.Email.Value)
	}

	log.InfoContext(ctx, "update user")
	return nil, nil
}

func (h *ServerHandlers) Delete(ctx context.Context, deleteIn *desc.DeleteIn) (*emptypb.Empty, error) {
	slog.InfoContext(ctx, "delete user", slog.Int64("id", deleteIn.Id))

	return nil, nil
}
