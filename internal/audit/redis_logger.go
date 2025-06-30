package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisLogger struct {
	client *redis.Client
	ttl    time.Duration
}

type LogEntry struct {
	Action    string      `json:"action"`
	ItemType  string      `json:"item_type"`
	ItemID    uint        `json:"item_id"`
	Timestamp time.Time   `json:"timestamp"`
	Payload   interface{} `json:"payload"`
}

func NewRedisLogger(addr string, ttl time.Duration) *RedisLogger {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisLogger{client: rdb, ttl: ttl}
}

func (r *RedisLogger) LogChange(ctx context.Context, action string, itemType string, itemID uint, payload interface{}) error {
	entry := LogEntry{
		Action:    action,
		ItemType:  itemType,
		ItemID:    itemID,
		Timestamp: time.Now(),
		Payload:   payload,
	}
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("log:%s:%d:%d", itemType, itemID, time.Now().UnixNano())
	return r.client.Set(ctx, key, data, r.ttl).Err()
}
