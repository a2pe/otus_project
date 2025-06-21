package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"otus_project/internal/repository"
	"otus_project/internal/service"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		<-sigChan
		log.Println("Program interrupted: shutting down")
		cancel()
	}()

	repository.StartSliceLogger(ctx)
	service.GenerateData(ctx)

	log.Println("Program shut down successfully")

}
