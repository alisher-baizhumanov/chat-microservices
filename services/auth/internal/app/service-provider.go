package app

import (
	"context"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/jwt"
	token_manager "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/jwt/token-manager"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/cache"
	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	grpcLibrary "github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	httpLibrary "github.com/alisher-baizhumanov/chat-microservices/pkg/http-gateway"
	descAuth "github.com/alisher-baizhumanov/chat-microservices/protos/generated/auth-v1"
	descUser "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
	grpcHandler "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/api/grpc"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/config"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
	authService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/auth"
	userService "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service/user"
	cacheInterface "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	userCache "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache/user"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	userRepository "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/user"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/hasher"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/hasher/argon2id"
)

type serviceProvider struct {
	cfg *config.Config

	userHandlers *grpcHandler.UserHandlers
	authHandlers *grpcHandler.AuthHandlers

	userService service.UserService
	authService service.AuthService

	tokenManager   jwt.TokenManager
	passwordHasher hasher.PasswordHasher

	userRepository repository.UserRepository
	userCache      cacheInterface.UserCache

	dbClient    db.Client
	cacheClient cache.Client
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

func (s *serviceProvider) getPasswordHasher() hasher.PasswordHasher {
	if s.passwordHasher == nil {
		s.passwordHasher = argon2id.New(hasher.DefaultOptions)
	}

	return s.passwordHasher
}

func (s *serviceProvider) getTokenManager() jwt.TokenManager {
	if s.tokenManager == nil {
		s.tokenManager = token_manager.New(
			s.getConfig().AccessSecretKey,
			s.getConfig().RefreshSecretKey,
			s.getConfig().AccessTokenTTL,
			s.getConfig().RefreshTokenTTL,
		)
	}

	return s.tokenManager
}

func (s *serviceProvider) getAuthService() service.AuthService {
	if s.authService == nil {
		s.authService = authService.New(
			s.getUserRepository(),
			s.getPasswordHasher(),
			s.getTokenManager(),
		)
	}

	return s.authService
}

func (s *serviceProvider) getUserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.New(
			s.getUserRepository(),
			s.getUserCache(),
			s.getPasswordHasher(),
		)
	}

	return s.userService
}

func (s *serviceProvider) getUserHandlers() *grpcHandler.UserHandlers {
	if s.userHandlers == nil {
		s.userHandlers = grpcHandler.NewUserHandlers(
			s.getUserService(),
		)
	}

	return s.userHandlers
}

func (s *serviceProvider) getAuthHandlers() *grpcHandler.AuthHandlers {
	if s.authHandlers == nil {
		s.authHandlers = grpcHandler.NewAuthHandlers(
			s.authService,
		)
	}

	return s.authHandlers
}

func (s *serviceProvider) getGRPCServer() (*grpcLibrary.Server, error) {
	return grpcLibrary.NewGRPCServer(
		s.getConfig().GRPCServerPort,
		[]grpcLibrary.Service{
			{
				ServiceDesc: &descUser.UserServiceV1_ServiceDesc,
				Handler:     s.getUserHandlers(),
			},
			{
				ServiceDesc: &descAuth.AuthServiceV1_ServiceDesc,
				Handler:     s.getAuthHandlers(),
			},
		},
	)
}

func (s *serviceProvider) getHTTPServer(ctx context.Context) (*httpLibrary.Server, error) {
	mux := runtime.NewServeMux()

	if err := descUser.RegisterUserServiceV1HandlerFromEndpoint(
		ctx,
		mux,
		s.getConfig().GRPCAddress(),
		s.getConfig().GRPCDialOptions(),
	); err != nil {
		return nil, err
	}

	if err := descAuth.RegisterAuthServiceV1HandlerFromEndpoint(
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
