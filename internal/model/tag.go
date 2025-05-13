package model

type Tag struct {
	ID     uint
	UserID uint
	Name   string
}

type TaskTag struct {
	TaskID uint
	TagID  uint
}
