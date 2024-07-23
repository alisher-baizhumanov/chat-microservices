package app

import (
	"context"
	"log/slog"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo"
	mg "github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo/mg"
	gracefulshutdown "github.com/alisher-baizhumanov/chat-microservices/pkg/graceful-shutdown"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/chat-v1"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/config"
)

// App represents the application with its services and gRPC server.
type App struct {
	cfg    *config.Config
	server *grpc.Server
	client mongo.Client
}

// NewApp creates and initializes a new App instance.
func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	client, err := mg.NewClient(ctx, cfg.DSN, cfg.Database)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx); err != nil {
		return nil, err
	}

	services := newServiceProvider(client)

	gRPCHandlers := services.ServerHandlers()

	server, err := grpc.NewGRPCServer(cfg.GRPCServerPort, &desc.ChatServiceV1_ServiceDesc, gRPCHandlers)
	if err != nil {
		return nil, err
	}

	return &App{
		cfg:    cfg,
		server: server,
		client: client,
	}, nil
}

// Run starts the gRPC server and waits for a termination signal to gracefully shut down the server.
func (a *App) Run(ctx context.Context) {
	defer func() {
		if err := a.client.Close(ctx); err != nil {
			slog.Warn("error to close connection to DB",
				slog.Any("error", err),
			)
		}
	}()

	a.server.Start()
	defer a.server.Stop()
	slog.Info("Starting gRPC Server",
		slog.Int("port", a.cfg.GRPCServerPort),
	)

	stop := gracefulshutdown.WaitSignal()
	slog.Info("Stop application", slog.String("signal", stop.String()))
}
