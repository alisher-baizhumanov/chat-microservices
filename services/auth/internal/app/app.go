package app

import (
	"log/slog"

	gracefulshutdown "github.com/alisher-baizhumanov/chat-microservices/pkg/graceful-shutdown"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
)

// App represents the application with its services and gRPC server.
type App struct {
	services serviceProvider
	server   *grpc.Server
}

// NewApp creates and initializes a new App instance.
func NewApp() (*App, error) {
	app := &App{}

	gRPCHandlers := app.services.ServerHandlers()

	cfg := app.services.Config()
	server, err := grpc.NewGRPCServer(cfg.GRPCServerPort, &desc.UserServiceV1_ServiceDesc, gRPCHandlers)
	if err != nil {
		return nil, err
	}

	app.server = server

	return app, nil
}

// Run starts the gRPC server and waits for a termination signal to gracefully shut down the server.
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
