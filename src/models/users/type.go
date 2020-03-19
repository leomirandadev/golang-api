package users

import "time"

type User struct {
	ID        int `gorm:"primary_key"`
	NickName  string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
