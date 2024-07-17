package main

import (
	"context"
	"log"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/app"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	application, err := app.NewApp(ctx, cfg)
	if err != nil {
		return err
	}

	return application.Run(ctx)
}
