// @title Productivity Tracker API
// @version 1.0
// @description API for tracking users, projects, tasks and more.
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"context"
	"log"
	"otus_project/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newApp, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = newApp.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting server")

}
