package model

import "time"

type TimeEntry struct {
	ID        uint      `json:"id" example:"1" description:"Уникальный идентификатор записи времени"`
	UserID    uint      `json:"user_id" example:"10" description:"ID пользователя, связанного с этой записью"`
	TaskID    uint      `json:"task_id" example:"42" description:"ID задачи, к которой относится запись времени"`
	StartTime time.Time `json:"start_time" example:"2025-06-28T09:00:00Z" description:"Время начала работы"`
	StopTime  time.Time `json:"stop_time" example:"2025-06-28T12:30:00Z" description:"Время завершения работы"`
	Note      string    `json:"note" example:"Работа над бэкендом" description:"Дополнительная заметка о работе"`
	CreatedAt time.Time `json:"created_at" example:"2025-06-28T12:35:00Z" description:"Дата создания записи"`
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
