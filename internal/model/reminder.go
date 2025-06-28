package model

import "time"

type Reminder struct {
	ID        uint      `json:"id" example:"1"`
	UserID    uint      `json:"user_id" example:"1"`
	TaskID    uint      `json:"task_id" example:"1"`
	RemindAt  time.Time `json:"remind_at" example:"2025-06-28T12:00:00Z"`
	IsSent    bool      `json:"is_sent" example:"true"`
	CreatedAt time.Time `json:"created_at" example:"2025-06-28T12:00:00Z"`
}

func (r *Reminder) GetItem() uint {
	return r.ID
}

func (r *Reminder) SetID(item uint) {
	r.ID = item
}

func (r *Reminder) SetCreatedAt(date time.Time) {
	r.CreatedAt = date
}
