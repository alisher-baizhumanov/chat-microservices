package app

import (
	"context"
	"log/slog"

	gracefulshutdown "github.com/alisher-baizhumanov/chat-microservices/pkg/graceful-shutdown"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/chat-v1"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/config"
)

// App represents the application with its services and gRPC server.
type App struct {
	cfg         *config.Config
	server      *grpc.Server
	mongoClient any
}

// NewApp creates and initializes a new App instance.
func NewApp(_ context.Context, cfg *config.Config) (*App, error) {
	var mongoClient any

	services := newServiceProvider(mongoClient)

	gRPCHandlers := services.ServerHandlers()

	server, err := grpc.NewGRPCServer(cfg.GRPCServerPort, &desc.ChatServiceV1_ServiceDesc, gRPCHandlers)
	if err != nil {
		return nil, err
	}

	return &App{
		cfg:         cfg,
		server:      server,
		mongoClient: mongoClient,
	}, nil
}

// Run starts the gRPC server and waits for a termination signal to gracefully shut down the server.
func (a *App) Run() error {
	a.server.Start()
	defer a.server.Stop()
	slog.Info("Starting gRPC Server",
		slog.Int("port", a.cfg.GRPCServerPort),
	)

	stop := gracefulshutdown.WaitSignal()
	slog.Info("Stop application", slog.String("signal", stop.String()))

	return nil
}
