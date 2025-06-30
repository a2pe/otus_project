package model

import "time"

type Task struct {
	ID          uint      `json:"id" bson:"id" example:"1" description:"Уникальный идентификатор задачи"`
	ProjectID   uint      `json:"project_id" bson:"project-id" example:"42" description:"ID проекта, к которому относится задача"`
	Title       string    `json:"title" bson:"title" example:"Написать репозиторий" description:"Название задачи" validate:"required"`
	Description string    `json:"description" bson:"description" example:"Создать функции CRUD для трекера"`
	Status      string    `json:"status" bson:"status" example:"in_progress" description:"Статус задачи" validate:"required"`
	DueDate     time.Time `json:"due_date" bson:"due-date" example:"2025-07-01T15:04:05Z" description:"Срок исполнения" validate:"required"`
	EstimateHrs float64   `json:"estimate_hrs" bson:"estimate-hrs" example:"3.5" description:"Оценка времени выполнения в часах"`
	CreatedAt   time.Time `json:"created_at" bson:"created-at" example:"2025-06-28T12:00:00Z" description:"Дата создания"`
}

func (t *Task) GetItem() uint {
	return t.ID
}

func (t *Task) SetID(item uint) {
	t.ID = item
}

func (t *Task) SetCreatedAt(date time.Time) {
	t.CreatedAt = date
}
