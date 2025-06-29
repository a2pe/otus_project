package notification

import (
	"context"
	//"fmt"
	"log"
	"time"

	pb "otus_project/internal/grpcpb"
)

type NotificationServer struct {
	pb.UnimplementedReminderServiceServer
}

func NewServer() *NotificationServer {
	return &NotificationServer{}
}

func (s *NotificationServer) AddReminder(ctx context.Context, req *pb.AddReminderRequest) (*pb.AddReminderResponse, error) {
	remindAt := req.RemindAt.AsTime()
	delay := time.Until(remindAt)
	if delay < 0 {
		delay = 0
	}

	log.Printf("[gRPC] Reminder scheduled: task %d at %s → %s", req.TaskId, remindAt.Format(time.RFC3339), req.Message)

	go func(taskID uint32, remindTime time.Time, msg string) {
		time.Sleep(delay)

		log.Printf("[Reminder] Task %d: %s", taskID, msg)

		if err := sendTelegramNotification(msg); err != nil {
			log.Printf("[ERROR] Telegram send failed: %v", err)
		} else {
			log.Println("[INFO] Telegram message sent successfully")
		}
	}(req.TaskId, remindAt, req.Message)

	return &pb.AddReminderResponse{Success: true}, nil
}
