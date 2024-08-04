package app

import (
	"context"
	"log/slog"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo"
	mg "github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo/mg"
	gracefulshutdown "github.com/alisher-baizhumanov/chat-microservices/pkg/graceful-shutdown"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	http "github.com/alisher-baizhumanov/chat-microservices/pkg/http-gateway"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/config"
)

// App represents the application with its services and gRPC server.
type App struct {
	cfg        *config.Config
	grpcServer *grpc.Server
	httpServer *http.Server
	client     mongo.Client
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

	services := newServiceProvider(client, cfg)

	grpcServer, err := services.getGRPCServer()
	if err != nil {
		return nil, err
	}

	httpServer, err := services.getHTTPserver(ctx)
	if err != nil {
		return nil, err
	}

	return &App{
		cfg:        cfg,
		grpcServer: grpcServer,
		httpServer: httpServer,
		client:     client,
	}, nil
}

// Run starts the gRPC server and waits for a termination signal to gracefully shut down the server.
func (a *App) Run(ctx context.Context) (err error) {
	if err = a.client.Ping(ctx); err != nil {
		return err
	}
	defer func() {
		if errClose := a.client.Close(ctx); errClose != nil {
			err = errClose
		}
	}()

	a.grpcServer.Start()
	defer a.grpcServer.Stop()
	slog.Info("Starting gRPC Server",
		slog.Int("port", a.cfg.GRPCServerPort),
	)

	a.httpServer.Start()
	defer func() {
		if errClose := a.httpServer.Stop(ctx); errClose != nil {
			err = errClose
		}
	}()
	slog.Info("Starting HTTP Server",
		slog.Int("port", a.cfg.HTTPServerPort),
	)

	stop := gracefulshutdown.WaitSignal()
	slog.Info("Stop application", slog.String("signal", stop.String()))

	return nil
}
