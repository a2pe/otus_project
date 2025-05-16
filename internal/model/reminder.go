package model

import "time"

type Reminder struct {
	ID        uint
	UserID    uint
	TaskID    uint
	RemindAt  time.Time
	IsSent    bool
	CreatedAt time.Time
}

func (r Reminder) GetItem() uint {
	return r.ID
}
