package users

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	NickName string
	Name     string
	Email    string
	Password string
}
