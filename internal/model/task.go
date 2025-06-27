package model

import "time"

type Task struct {
	ID          uint
	ProjectID   uint
	Title       string
	Description string
	Status      string
	DueDate     time.Time
	EstimateHrs float64
	CreatedAt   time.Time
}

func (t *Task) GetItem() uint {
	return t.ID
}

func (t *Task) SetID(item uint) {
	t.ID = item
}
