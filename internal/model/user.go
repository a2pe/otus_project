package model

import "time"

type User struct {
	ID        uint
	email     string
	Password  string
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
