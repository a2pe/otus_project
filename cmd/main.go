package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"otus_project/internal/app"
	"otus_project/internal/grpcpb"
	"otus_project/internal/notification"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 1. Запуск HTTP сервера
	httpApp, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}

	go func() {
		log.Println("Starting HTTP server...")
		if err := httpApp.Start(); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// 2. Запуск gRPC сервера
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen for gRPC: %v", err)
		}

		grpcServer := grpc.NewServer()
		grpcpb.RegisterReminderServiceServer(grpcServer, &notification.NotificationServer{})

		log.Println("Starting gRPC server on :50051...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	// 3. Ожидание завершения
	<-ctx.Done()
	log.Println("Shutting down both servers...")

	time.Sleep(2 * time.Second)
}
