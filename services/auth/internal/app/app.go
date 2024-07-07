package app

import (
	"log/slog"

	gracefulshutdown "github.com/alisher-baizhumanov/chat-microservices/pkg/graceful-shutdown"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/api/grpc"
)

type App struct {
	services serviceProvider
	server   *grpc.Server
}

func NewApp() (*App, error) {
	var err error

	app := &App{}

	gRPCHandlers := app.services.ServerHandlers()

	cfg := app.services.Config()
	app.server, err = grpc.NewGRPCServer(cfg.GRPCServerPort, gRPCHandlers)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	a.server.Start()
	defer a.server.Stop()
	slog.Info("Starting gRPC Server",
		slog.Int("port", a.services.Config().GRPCServerPort),
	)

	stop := gracefulshutdown.WaitSignal()
	slog.Info("Stop application", slog.String("signal", stop.String()))

	return nil
}
