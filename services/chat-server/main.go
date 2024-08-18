package main

import (
	"context"
	"log"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/app"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	application, err := app.NewApp(ctx)
	if err != nil {
		return err
	}

	return application.Run(ctx)
}
