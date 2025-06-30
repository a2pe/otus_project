package model

import "time"

type Tag struct {
	ID        uint      `json:"id" example:"1"`
	UserID    uint      `json:"user_id" example:"1"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" example:"2025-06-28T12:00:00Z"`
}

type TaskTag struct {
	TaskID uint
	TagID  uint
}

func (t *Tag) GetItem() uint {
	return t.ID
}

func (t *Tag) SetID(item uint) {
	t.ID = item
}

func (t *Tag) SetCreatedAt(date time.Time) {
	t.CreatedAt = date
}
