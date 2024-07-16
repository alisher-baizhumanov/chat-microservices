package converter

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	data "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository/user/model"
)

// UserCreateModelToData converts model.UserCreate to data.UserCreate.
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

// UserDataToModel converts data.User to model.User.
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

// UserUpdateOptionModelToData converts model.UserUpdateOptions to data.UserUpdateOptions.
func UserUpdateOptionModelToData(options *model.UserUpdateOptions) *data.UserUpdateOptions {
	if options == nil {
		return nil
	}

	var role *int8
	if options.Role != nil && *options.Role != model.NullRole {
		num := int8(*options.Role)
		role = &num
	}

	return &data.UserUpdateOptions{
		Role:  role,
		Name:  options.Name,
		Email: options.Email,
	}
}
