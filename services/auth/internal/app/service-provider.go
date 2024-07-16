package app

import (
	"github.com/jackc/pgx/v5/pgxpool"

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
	_connectionPool *pgxpool.Pool
}

func newServiceProvider(connectionPool *pgxpool.Pool) serviceProvider {
	return serviceProvider{_connectionPool: connectionPool}
}

func (s *serviceProvider) connectionPool() *pgxpool.Pool {
	return s._connectionPool
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
