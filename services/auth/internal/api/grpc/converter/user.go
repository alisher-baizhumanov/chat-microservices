package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// UserRegisterProtoToModel converts a UserRegister protobuf message to a model.UserRegister struct.
func UserRegisterProtoToModel(user *desc.UserRegister) model.UserRegister {
	return model.UserRegister{
		Name:            user.GetName(),
		Email:           user.GetEmail(),
		Password:        []byte(user.GetPassword()),
		PasswordConfirm: []byte(user.GetPasswordConfirm()),
	}
}

// UserModelToProto converts a model.User struct to a UserInfo protobuf message.
func UserModelToProto(user model.User) *desc.UserInfo {
	var role desc.Role

	switch user.Role {
	case model.UserRole:
		role = desc.Role_USER
	case model.AdminRole:
		role = desc.Role_ADMIN
	default:
		role = desc.Role_NULL
	}

	return &desc.UserInfo{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      role,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

// UserOptionsProtoToModel converts a UserUpdate protobuf message to a model.UserUpdateOptions struct.
func UserOptionsProtoToModel(options *desc.UserUpdate) model.UserUpdateOptions {
	var (
		role  model.Role
		name  *string
		email *string
	)

	switch options.Role {
	case desc.Role_USER:
		role = model.UserRole
	case desc.Role_ADMIN:
		role = model.AdminRole
	default:
		role = model.NullRole
	}

	if options.GetName() != nil {
		name = &options.GetName().Value
	}

	if options.GetEmail() != nil {
		email = &options.GetEmail().Value
	}

	return model.UserUpdateOptions{
		Role:  &role,
		Name:  name,
		Email: email,
	}
}
