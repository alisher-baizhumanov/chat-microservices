package app

import (
	"context"
	"log/slog"

	gracefulshutdown "github.com/alisher-baizhumanov/chat-microservices/pkg/graceful-shutdown"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/config"
	userRepository "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository/user"
)

// App represents the application with its services and gRPC server.
type App struct {
	cfg        *config.Config
	server     *grpc.Server
	repository *userRepository.Repository
}

// NewApp creates and initializes a new App instance.
func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	repository, err := userRepository.NewRepository(ctx, cfg.DSN)
	if err != nil {
		return nil, err
	}

	services := newServiceProvider(repository)

	gRPCHandlers := services.ServerHandlers()

	server, err := grpc.NewGRPCServer(cfg.GRPCServerPort, &desc.UserServiceV1_ServiceDesc, gRPCHandlers)
	if err != nil {
		return nil, err
	}

	return &App{
		cfg:        cfg,
		server:     server,
		repository: repository,
	}, nil
}

// Run starts the gRPC server and waits for a termination signal to gracefully shut down the server.
func (a *App) Run() error {
	defer a.repository.Stop()

	a.server.Start()
	defer a.server.Stop()
	slog.Info("Starting gRPC Server",
		slog.Int("port", a.cfg.GRPCServerPort),
	)

	stop := gracefulshutdown.WaitSignal()
	slog.Info("Stop application", slog.String("signal", stop.String()))

	return nil
}
