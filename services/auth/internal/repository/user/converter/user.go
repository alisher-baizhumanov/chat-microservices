package converter

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	data "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository/user/model"
)

func UserCreateModelToData(user *model.UserCreate) *data.UserCreate {
	if user == nil {
		return nil
	}

	return &data.UserCreate{
		Name:           user.Name,
		Email:          user.Email,
		Role:           int8(user.Role),
		HashedPassword: user.HashedPassword,
		CreatedAt:      user.CreatedAt,
	}
}

func UserDataToModel(user *data.User) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      model.Role(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserUpdateOptionModelToData(options *model.UserUpdateOptions) *data.UserUpdateOptions {
	if options == nil {
		return nil
	}

	var role *int8
	if options.Role != nil {
		num := int8(*options.Role)
		role = &num
	}

	return &data.UserUpdateOptions{
		Role:  role,
		Name:  options.Name,
		Email: options.Email,
	}
}
