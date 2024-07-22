package app

import (
	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/api/grpc"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository"
	userRepository "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository/user"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
)

type serviceProvider struct {
	_gRPCServer     *grpc.ServerHandlers
	_userService    service.UserService
	_userRepository repository.UserRepository
	_dbClient       db.Client
}

func newServiceProvider(dbClient db.Client) serviceProvider {
	return serviceProvider{_dbClient: dbClient}
}

func (s *serviceProvider) connectionPool() db.Client {
	return s._dbClient
}

func (s *serviceProvider) userRepository() repository.UserRepository {
	if s._userRepository == nil {
		s._userRepository = userRepository.NewRepository(
			s.connectionPool(),
		)
	}

	return s._userRepository
}

func (s *serviceProvider) userService() service.UserService {
	if s._userService == nil {
		s._userService = userService.NewService(
			s.userRepository(),
		)
	}

	return s._userService
}

func (s *serviceProvider) serverHandlers() *grpc.ServerHandlers {
	if s._gRPCServer == nil {
		s._gRPCServer = grpc.NewUserGRPCHandlers(
			s.userService(),
		)
	}

	return s._gRPCServer
}
