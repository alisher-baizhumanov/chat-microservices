package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// RegisterUser registers a new user in the system with the given registration details.
// It converts the UserRegister model to a UserCreate model, sets the default user role,
// and assigns the current time as the creation time.
func (s *service) RegisterUser(ctx context.Context, userRegister model.UserRegister) (int64, error) {
	createdAt := time.Now()

	hash, err := s.hasher.Hash(userRegister.Password)
	if err != nil {
		logger.Warn(model.ErrInvalidToken.Error(), logger.String("error", err.Error()))
		return 0, fmt.Errorf("%w, message: %w", model.ErrPasswordHashing, err)
	}

	userCreate := model.UserCreate{
		Name:           userRegister.Name,
		Email:          userRegister.Email,
		Role:           model.UserRole,
		HashedPassword: hash,
		CreatedAt:      createdAt,
	}

	id, err := s.userRepository.CreateUser(ctx, userCreate)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNonUniqueEmail):
			logger.Info("email is not unique", logger.String("email", userCreate.Email))
		case errors.Is(err, model.ErrNonUniqueUsername):
			logger.Info("username is not unique", logger.String("username", userCreate.Name))
		default:
			logger.Warn("database error", logger.String("error", err.Error()))
		}

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
		logger.Warn("Register user: not created user in cache",
			logger.String("error", err.Error()),
			logger.Int64("id", id),
		)
	}

	logger.Info("user registered", logger.Int64("id", id))
	return id, nil
}
