package app

import (
	"time"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/cache"
	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/api/grpc"
	cacheInterface "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/cache"
	userCache "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/cache/user"
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
	cacheClient        cache.Client
	userCache          cacheInterface.UserCache
	ttl                time.Duration
}

func newServiceProvider(dbClient db.Client, cacheClient cache.Client, ttl time.Duration) serviceProvider {
	return serviceProvider{
		dbClient:    dbClient,
		cacheClient: cacheClient,
		ttl:         ttl,
	}
}

func (s *serviceProvider) getDBClient() db.Client {
	return s.dbClient
}

func (s *serviceProvider) getCacheClient() cache.Client {
	return s.cacheClient
}

func (s *serviceProvider) getTTL() time.Duration {
	return s.ttl
}

func (s *serviceProvider) getUserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(
			s.getDBClient(),
		)
	}

	return s.userRepository
}

func (s *serviceProvider) getUserCache() cacheInterface.UserCache {
	if s.userCache == nil {
		s.userCache = userCache.NewCache(
			s.getCacheClient().Cache(),
			s.getTTL(),
		)
	}

	return s.userCache
}

func (s *serviceProvider) getUserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.getUserRepository(),
			s.getUserCache(),
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
