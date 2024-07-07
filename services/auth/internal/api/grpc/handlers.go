package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/converter"
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
	id, err := h.userService.RegisterUser(
		ctx,
		converter.UserRegisterProtoToModel(createIn.UserRegister),
	)

	if err != nil {
		return nil, err
	}

	return &desc.CreateOut{Id: id}, nil
}

func (h *ServerHandlers) Get(ctx context.Context, getIn *desc.GetIn) (*desc.GetOut, error) {
	user, err := h.userService.GetById(ctx, getIn.Id)

	if err != nil {
		return nil, err
	}

	return &desc.GetOut{
			UserInfo: converter.UserModelToProto(user)},
		nil
}

func (h *ServerHandlers) Update(ctx context.Context, updateIn *desc.UpdateIn) (*emptypb.Empty, error) {
	err := h.userService.UpdateUserFields(
		ctx,
		updateIn.Id,
		converter.UserOptionsProtoToModel(updateIn.UserUpdate),
	)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *ServerHandlers) Delete(ctx context.Context, deleteIn *desc.DeleteIn) (*emptypb.Empty, error) {
	err := h.userService.DeleteById(ctx, deleteIn.Id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
