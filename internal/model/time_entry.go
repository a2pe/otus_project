package model

import "time"

type TimeEntry struct {
	ID        uint
	UserID    uint
	TaskID    uint
	StartTime time.Time
	StopTime  time.Time
	Note      string
	CreatedAt time.Time
}

func (t *TimeEntry) GetItem() uint {
	return t.ID
}

func (t *TimeEntry) SetID(item uint) {
	t.TaskID = item
}

func (t *TimeEntry) SetCreatedAt(date time.Time) {
	t.CreatedAt = date
}
