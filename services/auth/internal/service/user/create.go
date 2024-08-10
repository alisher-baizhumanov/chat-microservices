package user

import (
	"context"
	"log/slog"
	"time"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// RegisterUser registers a new user in the system with the given registration details.
// It converts the UserRegister model to a UserCreate model, sets the default user role,
// and assigns the current time as the creation time.
func (s *service) RegisterUser(ctx context.Context, userRegister model.UserRegister) (int64, error) {
	createdAt := time.Now()

	userCreate := model.UserCreate{
		Name:           userRegister.Name,
		Email:          userRegister.Email,
		Role:           model.UserRole,
		HashedPassword: userRegister.Password,
		CreatedAt:      createdAt,
	}

	id, err := s.userRepository.CreateUser(ctx, userCreate)
	if err != nil {
		return 0, err
	}

	user := model.User{
		ID:        id,
		Name:      userCreate.Name,
		Email:     userCreate.Email,
		Role:      userCreate.Role,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	if err = s.userCache.Set(ctx, user); err != nil {
		slog.ErrorContext(ctx, "not created user in cache",
			slog.String("error", err.Error()),
			slog.Int64("id", id),
		)
	}

	return id, nil
}
