package notification

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "otus_project/internal/grpcpb"
)

var (
	client pb.ReminderServiceClient
	once   sync.Once
)

func GetClient() pb.ReminderServiceClient {
	once.Do(func() {
		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect to notification service: %v", err)
		}
		client = pb.NewReminderServiceClient(conn)
	})
	return client
}

func ScheduleReminder(taskID uint32, remindAt string, message string) error {
	parsedTime, err := time.Parse(time.RFC3339, remindAt)
	if err != nil {
		return err
	}

	_, err = GetClient().AddReminder(context.Background(), &pb.AddReminderRequest{
		TaskId:   taskID,
		RemindAt: timestamppb.New(parsedTime),
		Message:  message,
	})
	return err
}
