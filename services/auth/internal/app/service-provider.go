package app

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/api/grpc"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/config"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository"
	userRepository "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository/user"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
)

type serviceProvider struct {
	gRPCServer     *grpc.ServerHandlers
	userService    service.UserService
	userRepository repository.UserRepository
	cfg            *config.Config
}

func (s *serviceProvider) Config() *config.Config {
	return s.cfg
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository()
	}

	return s.userRepository
}

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
		)
	}

	return s.userService
}

func (s *serviceProvider) ServerHandlers() *grpc.ServerHandlers {
	if s.gRPCServer == nil {
		s.gRPCServer = grpc.NewUserGRPCHandlers(
			s.UserService(),
		)
	}

	return s.gRPCServer
}
