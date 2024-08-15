package app

import (
	"context"
	"log/slog"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/cache"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/cache/redis"
	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	pg "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres/pg"
	gracefulshutdown "github.com/alisher-baizhumanov/chat-microservices/pkg/graceful-shutdown"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	http "github.com/alisher-baizhumanov/chat-microservices/pkg/http-gateway"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/config"
)

// App represents the application with its services and gRPC server.
type App struct {
	cfg         *config.Config
	grpcServer  *grpc.Server
	httpServer  *http.Server
	dbClient    db.Client
	cacheClient cache.Client
}

// NewApp creates and initializes a new App instance.
func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	dbClient, err := pg.New(ctx, cfg.DSN)
	if err != nil {
		return nil, err
	}

	cacheClient := redis.NewClient(cfg.CacheDSN)

	services := newServiceProvider(dbClient, cacheClient, cfg)

	grpcServer, err := services.getGRPCServer()
	if err != nil {
		return nil, err
	}

	httpServer, err := services.getHTTPServer(ctx)
	if err != nil {
		return nil, err
	}

	return &App{
		cfg:         cfg,
		grpcServer:  grpcServer,
		dbClient:    dbClient,
		cacheClient: cacheClient,
		httpServer:  httpServer,
	}, nil
}

// Run starts the gRPC server and waits for a termination signal to gracefully shut down the server.
func (a *App) Run(ctx context.Context) (err error) {
	if err = a.dbClient.DB().Ping(ctx); err != nil {
		return err
	}
	defer a.dbClient.DB().Close()

	if err = a.cacheClient.Ping(ctx); err != nil {
		return err
	}
	defer func() {
		if errClose := a.cacheClient.Close(ctx); errClose != nil {
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
