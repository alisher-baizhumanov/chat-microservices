package main

import (
	"log"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/app"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
