package converter

import (
	"time"

	cacheData "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/cache/user/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

const timeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

func UserModelToCacheData(user model.User) cacheData.User {
	return cacheData.User{
		Name:      user.Name,
		Email:     user.Email,
		Role:      int8(user.Role),
		CreatedAt: user.CreatedAt.Format(timeFormat),
		UpdatedAt: user.UpdatedAt.Format(timeFormat),
	}
}

func UserCacheDataToModel(id int64, user cacheData.User) (model.User, error) {
	var (
		createdAt, updatedAt time.Time
		err                  error
	)

	if user.CreatedAt != "" {
		createdAt, err = time.Parse(timeFormat, user.CreatedAt)
		if err != nil {
			return model.User{}, err
		}
	}

	if user.UpdatedAt != "" {
		updatedAt, err = time.Parse(timeFormat, user.UpdatedAt)
		if err != nil {
			return model.User{}, err
		}
	}

	return model.User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      model.Role(user.Role),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
