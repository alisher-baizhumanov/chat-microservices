package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/api/grpc/converter"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
)

// UserHandlers implements the gRPC server for user-related operations.
type UserHandlers struct {
	desc.UnimplementedUserServiceV1Server

	userService service.UserService
}

// NewUserHandlers creates and returns a new UserHandlers instance.
func NewUserHandlers(userService service.UserService) *UserHandlers {
	return &UserHandlers{userService: userService}
}

// Create registers a new user based on the provided CreateIn message.
func (h *UserHandlers) Create(ctx context.Context, createIn *desc.CreateIn) (*desc.CreateOut, error) {
	id, err := h.userService.RegisterUser(
		ctx,
		converter.UserRegisterProtoToModel(createIn.GetUserRegister()),
	)
	if err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return &desc.CreateOut{Id: id}, nil
}

// Get retrieves a user's information based on the provided GetIn message.
func (h *UserHandlers) Get(ctx context.Context, getIn *desc.GetIn) (*desc.GetOut, error) {
	user, err := h.userService.GetByID(ctx, getIn.GetId())
	if err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return &desc.GetOut{
		UserInfo: converter.UserModelToProto(user),
	}, nil
}

// Update modifies a user's information based on the provided UpdateIn message.
func (h *UserHandlers) Update(ctx context.Context, updateIn *desc.UpdateIn) (*emptypb.Empty, error) {
	err := h.userService.UpdateUserFields(
		ctx,
		updateIn.GetId(),
		converter.UserOptionsProtoToModel(updateIn.GetUserUpdate()),
	)
	if err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return nil, nil
}

// Delete removes a user based on the provided DeleteIn message.
func (h *UserHandlers) Delete(ctx context.Context, deleteIn *desc.DeleteIn) (*emptypb.Empty, error) {
	err := h.userService.DeleteByID(ctx, deleteIn.GetId())
	if err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return nil, nil
}