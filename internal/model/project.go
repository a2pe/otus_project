package model

import "time"

type Project struct {
	ID          uint      `json:"id" bson:"id" example:"42"`
	UserID      uint      `json:"user_id" bson:"user-id" example:"1"`
	Name        string    `json:"name" bson:"name" example:"My Project" validate:"required"`
	Description string    `json:"description,omitempty" bson:"description" example:"A sample project"`
	CreatedAt   time.Time `json:"created_at" bson:"created-at" example:"2025-06-28T19:53:32.953897+04:00"`
}

func (p *Project) GetItem() uint {
	return p.ID
}

func (p *Project) SetID(item uint) {
	p.ID = item
}

func (p *Project) SetCreatedAt(date time.Time) {
	p.CreatedAt = date
}
