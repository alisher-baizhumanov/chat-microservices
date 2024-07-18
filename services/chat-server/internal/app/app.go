package app

import (
	"context"
	"log/slog"

	gracefulshutdown "github.com/alisher-baizhumanov/chat-microservices/pkg/graceful-shutdown"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/chat-v1"
	mongoLibrary "go.mongodb.org/mongo-driver/mongo"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/config"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/mongo"
)

// App represents the application with its services and gRPC server.
type App struct {
	cfg         *config.Config
	server      *grpc.Server
	mongoClient *mongoLibrary.Client
}

// NewApp creates and initializes a new App instance.
func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	mongoClient, err := mongo.NewConnectionPool(ctx, cfg.DSN)
	if err != nil {
		return nil, err
	}

	services := newServiceProvider(mongoClient.Database(cfg.Database))

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
func (a *App) Run(ctx context.Context) error {
	defer func() {
		if err := mongo.CloseConnectionPool(ctx, a.mongoClient); err != nil {
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

	return nil
}
