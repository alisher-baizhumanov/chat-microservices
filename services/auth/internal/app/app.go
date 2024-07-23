package app

import (
	"context"
	"log/slog"

	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	pg "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres/pg"
	gracefulshutdown "github.com/alisher-baizhumanov/chat-microservices/pkg/graceful-shutdown"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/config"
)

// App represents the application with its services and gRPC server.
type App struct {
	cfg      *config.Config
	server   *grpc.Server
	dbClient db.Client
}

// NewApp creates and initializes a new App instance.
func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	dbClient, err := pg.New(ctx, cfg.DSN)
	if err != nil {
		return nil, err
	}

	if err = dbClient.DB().Ping(ctx); err != nil {
		return nil, err
	}

	services := newServiceProvider(dbClient)

	gRPCHandlers := services.serverHandlers()

	server, err := grpc.NewGRPCServer(cfg.GRPCServerPort, &desc.UserServiceV1_ServiceDesc, gRPCHandlers)
	if err != nil {
		return nil, err
	}

	return &App{
		cfg:      cfg,
		server:   server,
		dbClient: dbClient,
	}, nil
}

// Run starts the gRPC server and waits for a termination signal to gracefully shut down the server.
func (a *App) Run() {
	defer a.dbClient.DB().Close()

	a.server.Start()
	defer a.server.Stop()
	slog.Info("Starting gRPC Server",
		slog.Int("port", a.cfg.GRPCServerPort),
	)

	stop := gracefulshutdown.WaitSignal()
	slog.Info("Stop application", slog.String("signal", stop.String()))
}
