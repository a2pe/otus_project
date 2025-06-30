package notification

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"otus_project/internal/grpcpb"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	grpcpb.RegisterReminderServiceServer(server, NewServer())

	log.Println("Notification gRPC server started on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
