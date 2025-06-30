package model

import "time"

type Task struct {
	ID          uint      `json:"id" example:"1" description:"Уникальный идентификатор задачи"`
	ProjectID   uint      `json:"project_id" example:"42" description:"ID проекта, к которому относится задача"`
	Title       string    `json:"title" example:"Написать репозиторий" description:"Название задачи"`
	Description string    `json:"description" example:"Создать функции CRUD для трекера"`
	Status      string    `json:"status" example:"in_progress" description:"Статус задачи"`
	DueDate     time.Time `json:"due_date" example:"2025-07-01T15:04:05Z" description:"Срок исполнения"`
	EstimateHrs float64   `json:"estimate_hrs" example:"3.5" description:"Оценка времени выполнения в часах"`
	CreatedAt   time.Time `json:"created_at" example:"2025-06-28T12:00:00Z" description:"Дата создания"`
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
