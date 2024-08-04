package app

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/cache"
	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	clock "github.com/alisher-baizhumanov/chat-microservices/pkg/clock"
	grpcLibrary "github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	httpLibrary "github.com/alisher-baizhumanov/chat-microservices/pkg/http-gateway"
	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
	grpcHandler "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/api/grpc"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/config"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
	cacheInterface "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	userCache "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/user"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	userRepository "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/user"
)

type serviceProvider struct {
	gRPCHandlers   *grpcHandler.ServerHandlers
	userService    service.UserService
	userRepository repository.UserRepository
	dbClient       db.Client
	cacheClient    cache.Client
	userCache      cacheInterface.UserCache
	cfg            *config.Config
}

func newServiceProvider(dbClient db.Client, cacheClient cache.Client, cfg *config.Config) serviceProvider {
	return serviceProvider{
		dbClient:    dbClient,
		cacheClient: cacheClient,
		cfg:         cfg,
	}
}

func (s *serviceProvider) getConfig() *config.Config {
	return s.cfg
}

func (s *serviceProvider) getDBClient() db.Client {
	return s.dbClient
}

func (s *serviceProvider) getCacheClient() cache.Client {
	return s.cacheClient
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
			s.getConfig().CacheTTL,
		)
	}

	return s.userCache
}

func (s *serviceProvider) getUserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.New(
			s.getUserRepository(),
			s.getUserCache(),
			&clock.RealClock{},
		)
	}

	return s.userService
}

func (s *serviceProvider) getGRPCHandlers() *grpcHandler.ServerHandlers {
	if s.gRPCHandlers == nil {
		s.gRPCHandlers = grpcHandler.NewUserGRPCHandlers(
			s.getUserService(),
		)
	}

	return s.gRPCHandlers
}

func (s *serviceProvider) getGRPCServer() (*grpcLibrary.Server, error) {
	return grpcLibrary.NewGRPCServer(
		s.getConfig().GRPCServerPort,
		&desc.UserServiceV1_ServiceDesc,
		s.getGRPCHandlers(),
	)
}

func (s *serviceProvider) getHTTPserver(ctx context.Context) (*httpLibrary.Server, error) {
	mux := runtime.NewServeMux()

	if err := desc.RegisterUserServiceV1HandlerFromEndpoint(
		ctx,
		mux,
		s.getConfig().GRPCAddress(),
		s.getConfig().GRPCDialOptions(),
	); err != nil {
		return nil, err
	}

	return httpLibrary.NewHTTPServer(
		s.getConfig().HTTPServerPort,
		mux,
	), nil
}
