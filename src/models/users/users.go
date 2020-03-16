package users

import (
	"vcfConverter/src/settings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	NickName string
	Name     string
	Email    string
	Password string
}

func InitialMigrationUser() {
	db := settings.ConnectDB()
	defer db.Close()

	db.AutoMigrate(&User{})
}

func GetAll() (bool, []User) {
	db := settings.ConnectDB()
	defer db.Close()

	var users []User
	db.Find(&users)

	return true, users
}

func GetById(id int64) (bool, []User) {
	db := settings.ConnectDB()
	defer db.Close()

	var user []User
	db.Where("ID = ?", id).Find(&user)

	return true, user
}

func Create(newUser User) (bool, User) {
	db := settings.ConnectDB()
	defer db.Close()

	db.Create(&newUser)

	return true, newUser
}

func Delete(id int64) (bool, User) {
	db := settings.ConnectDB()
	defer db.Close()

	var user User
	db.Where("ID = ?", id).Find(&user)
	db.Delete(&user)

	return true, user
}

func Update(id int64, userUpdate User) (bool, User) {
	db := settings.ConnectDB()
	defer db.Close()

	var user User
	db.Where("ID = ?", id).Find(&user)

	user.Email = ifExists(userUpdate.Email, userUpdate.Email, user.Email)
	user.Name = ifExists(userUpdate.Name, userUpdate.Name, user.Name)
	user.NickName = ifExists(userUpdate.NickName, userUpdate.NickName, user.NickName)

	db.Save(&user)
	return true, user
}

func ifExists(compare string, trueResponse string, falseResponse string) string {
	if len(compare) > 0 {
		return trueResponse
	} else {
		return falseResponse
	}
}
