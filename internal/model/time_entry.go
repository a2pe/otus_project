package model

import "time"

type TimeEntry struct {
	ID        uint
	UserID    uint
	TaskID    uint
	StartTime time.Time
	EndTime   time.Time
	Note      string
	CreatedAt time.Time
}
