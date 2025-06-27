package model

import "time"

type User struct {
	ID        uint
	email     string
	password  string
	Name      string
	Timezone  string
	CreatedAt time.Time
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
