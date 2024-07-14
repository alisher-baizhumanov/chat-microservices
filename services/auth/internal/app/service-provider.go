package app

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/api/grpc"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
)

type serviceProvider struct {
	_gRPCServer     *grpc.ServerHandlers
	_userService    service.UserService
	_userRepository repository.UserRepository
}

func newServiceProvider(repo repository.UserRepository) serviceProvider {
	return serviceProvider{_userRepository: repo}
}

func (s *serviceProvider) userRepository() repository.UserRepository {
	return s._userRepository
}

func (s *serviceProvider) UserService() service.UserService {
	if s._userService == nil {
		s._userService = userService.NewService(
			s.userRepository(),
		)
	}

	return s._userService
}

func (s *serviceProvider) ServerHandlers() *grpc.ServerHandlers {
	if s._gRPCServer == nil {
		s._gRPCServer = grpc.NewUserGRPCHandlers(
			s.UserService(),
		)
	}

	return s._gRPCServer
}
