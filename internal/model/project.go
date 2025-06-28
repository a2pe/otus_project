package model

import "time"

type Project struct {
	ID          uint
	UserID      uint
	Name        string
	Description string
	CreatedAt   time.Time
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
