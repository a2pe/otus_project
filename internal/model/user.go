package model

import "time"

type User struct {
	ID        uint      `json:"id" bson:"id" example:"1"`
	email     string    `example:"user@example.com" bson:"email" validate:"required"`
	password  string    `example:"123456" bson:"password" validate:"required"`
	Name      string    `json:"name" bson:"name" validate:"required,min=2,max=100" example:"Alice"`
	Timezone  string    `json:"timezone" bson:"timezone" example:"America/Los_Angeles"`
	CreatedAt time.Time `json:"created_at" bson:"created-at" example:"2024-01-01T15:04:05Z07:00"`
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}
func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) GetItem() uint {
	return u.ID
}

func (u *User) SetID(item uint) {
	u.ID = item
}

func (u *User) SetCreatedAt(date time.Time) {
	u.CreatedAt = date
}
