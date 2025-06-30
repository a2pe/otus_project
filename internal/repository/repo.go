package repository

import (
	"context"
	"errors"
	"fmt"
	logging "otus_project/internal/audit"
	"otus_project/internal/model"
	"otus_project/internal/model/common"
	"otus_project/internal/storage/mongo"
	"sync"
	"sync/atomic"
	"time"
)

var (
	userCounter        uint64
	projectCounter     uint64
	taskCounter        uint64
	reminderCounter    uint64
	tagCounter         uint64
	timeEntriesCounter uint64

	UsersMu       sync.RWMutex
	ProjectsMu    sync.RWMutex
	TasksMu       sync.RWMutex
	RemindersMu   sync.RWMutex
	TagsMu        sync.RWMutex
	TimeEntriesMu sync.RWMutex

	Users       []*model.User
	Projects    []*model.Project
	Reminders   []*model.Reminder
	Tags        []*model.Tag
	TimeEntries []*model.TimeEntry
	Tasks       []*model.Task

	storage *mongo.MongoStorage

	redisLogger *logging.RedisLogger
)

func Init(ctx context.Context) error {
	var err error

	redisLogger = logging.NewRedisLogger("redis:6379", 24*time.Hour)

	storage, err = mongo.NewMongoStorage(ctx, "mongodb://mongo:27017", "tracker_db")
	if err != nil {
		return fmt.Errorf("failed to initialize mongo storage: %w", err)
	}
	fmt.Println("MongoDB connected successfully")
	return nil
}

func getType(item common.Item) string {
	switch item.(type) {
	case *model.User:
		return "user"
	case *model.Task:
		return "task"
	case *model.Project:
		return "project"
	case *model.Reminder:
		return "reminder"
	case *model.Tag:
		return "tag"
	case *model.TimeEntry:
		return "time_entry"
	default:
		return "unknown"
	}
}

func SaveItem(item common.Item, itemType string) error {
	switch v := item.(type) {
	case *model.User:
		prepareNewItem(v, &userCounter)
	case *model.Project:
		prepareNewItem(v, &projectCounter)
	case *model.Task:
		prepareNewItem(v, &taskCounter)
	case *model.Reminder:
		prepareNewItem(v, &reminderCounter)
	case *model.Tag:
		prepareNewItem(v, &tagCounter)
	case *model.TimeEntry:
		prepareNewItem(v, &timeEntriesCounter)
	default:
		return errors.New("unsupported type in SaveItem")
	}
	err := storage.SaveItem(context.Background(), item, itemType)
	if err == nil {
		_ = redisLogger.LogChange(context.Background(), "create", getType(item), item.GetItem(), item)
	}
	return err
}

func GetByID(itemType string, id int) (common.Item, bool) {
	raw, err := storage.GetByID(context.Background(), itemType, uint(id))
	if err != nil || raw == nil {
		return nil, false
	}

	item, ok := raw.(common.Item)
	if !ok {
		return nil, false
	}

	return item, true
}

func prepareNewItem[T common.Item](item T, counter *uint64) {
	item.SetID(uint(atomic.AddUint64(counter, 1)))
	item.SetCreatedAt(time.Now())
}

func GetAllItems(itemType string) (any, error) {
	return storage.GetAllItems(context.Background(), itemType)
}

func UpdateItem(itemType string, updated common.Item) bool {
	err := storage.UpdateByID(context.Background(), itemType, updated.GetItem(), updated)
	if err != nil {
		return false
	} else {
		_ = redisLogger.LogChange(context.Background(), "update", itemType, updated.GetItem(), updated)
		return true
	}
}

func DeleteItem(itemType string, id int) bool {
	_, err := storage.DeleteByID(context.Background(), itemType, uint(id))
	if err != nil {
		return false
	} else {
		_ = redisLogger.LogChange(context.Background(), "delete", itemType, uint(id), map[string]any{"deleted": true})
		return true
	}
}
