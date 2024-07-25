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
	gRPCServerHandlers *grpc.ServerHandlers
	userService        service.UserService
	userRepository     repository.UserRepository
	dbClient           db.Client
}

func newServiceProvider(dbClient db.Client) serviceProvider {
	return serviceProvider{dbClient: dbClient}
}

func (s *serviceProvider) DBClient() db.Client {
	return s.dbClient
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(
			s.DBClient(),
		)
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

func (s *serviceProvider) GRPCServerHandlers() *grpc.ServerHandlers {
	if s.gRPCServerHandlers == nil {
		s.gRPCServerHandlers = grpc.NewUserGRPCHandlers(
			s.UserService(),
		)
	}

	return s.gRPCServerHandlers
}
