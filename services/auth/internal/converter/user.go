package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

func UserRegisterProtoToModel(user *desc.UserRegister) *model.UserRegister {
	return &model.UserRegister{
		Name:            user.Name,
		Email:           user.Email,
		Password:        []byte(user.Password),
		PasswordConfirm: []byte(user.PasswordConfirm),
	}
}

func UserModelToProto(user *model.User) *desc.UserInfo {
	var role desc.Role

	switch user.Role {
	case model.UserRole:
		role = desc.Role_USER
	case model.AdminRole:
		role = desc.Role_ADMIN
	}

	return &desc.UserInfo{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      role,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func UserOptionsProtoToModel(options *desc.UserUpdate) *model.UserUpdateOptions {
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
	}

	if options.Name != nil {
		name = &options.Name.Value
	}

	if options.Email != nil {
		email = &options.Email.Value
	}

	return &model.UserUpdateOptions{
		Role:  &role,
		Name:  name,
		Email: email,
	}
}
