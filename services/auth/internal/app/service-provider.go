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

func (s *serviceProvider) getDBClient() db.Client {
	return s.dbClient
}

func (s *serviceProvider) getUserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(
			s.getDBClient(),
		)
	}

	return s.userRepository
}

func (s *serviceProvider) getUserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.getUserRepository(),
		)
	}

	return s.userService
}

func (s *serviceProvider) getGRPCServerHandlers() *grpc.ServerHandlers {
	if s.gRPCServerHandlers == nil {
		s.gRPCServerHandlers = grpc.NewUserGRPCHandlers(
			s.getUserService(),
		)
	}

	return s.gRPCServerHandlers
}
