package model

import "time"

type Project struct {
	ID          uint
	UserID      uint
	Name        string
	Description string
	CreatedAt   time.Time
}
