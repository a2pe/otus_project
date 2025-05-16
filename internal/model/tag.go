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

func (t Tag) GetItem() uint {
	return t.ID
}
