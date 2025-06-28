package main

import (
	"context"
	"log"
	"otus_project/internal/app"
	"otus_project/internal/repository"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := repository.Init(ctx); err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}

	newApp, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}

	log.Println("Starting server...")

	if err := newApp.Start(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
