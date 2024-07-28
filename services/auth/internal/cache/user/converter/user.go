package converter

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	cacheData "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/cache/user/model"
)
func UserModelToCacheData(user model.User) cacheData.User {
	return cacheData.User{
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserCacheDataToModel(id int64, user cacheData.User) model.User {
	return model.User{
		ID: id,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}