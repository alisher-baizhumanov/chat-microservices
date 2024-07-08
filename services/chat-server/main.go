package main

import (
	"log"
	"log/slog"

	gracefulshutdown "github.com/alisher-baizhumanov/chat-microservices/pkg/graceful-shutdown"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/presentation/grpc"
)

const port = 55052

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	server, err := grpc.NewGRPCServer(port)
	if err != nil {
		return err
	}

	server.Start()
	defer server.Stop()
	slog.Info("Starting gRPC Server", slog.Int64("port", port))

	stop := gracefulshutdown.WaitSignal()
	slog.Info("Stop application", slog.String("signal", stop.String()))

	return nil
}
