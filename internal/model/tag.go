package model

import "time"

type Tag struct {
	ID        uint
	UserID    uint
	Name      string
	CreatedAt time.Time
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
