package users

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/leomirandadev/golang-api/configs"
	"golang.org/x/crypto/bcrypt"
)

func GetAll() (bool, []User) {
	db := configs.ConnectDB()
	defer db.Close()

	var users []User
	db.Find(&users)

	users = removePasswordOnScope(users)

	return true, users
}

func GetById(id int64) (bool, []User) {
	db := configs.ConnectDB()
	defer db.Close()

	var user []User
	db.Where("ID = ?", id).Find(&user)

	user = removePasswordOnScope(user)

	return true, user
}

func GetByEmailPassword(email string, password string) (bool, []User) {
	db := configs.ConnectDB()
	defer db.Close()

	var user []User
	db.Where("email = ?", email).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(password))

	if err == nil {
		user = removePasswordOnScope(user)
		return true, user
	}

	return false, []User{}
}

func Create(newUser User) bool {
	db := configs.ConnectDB()
	defer db.Close()

	newUser.Password = hashPassword(newUser.Password)

	db.Create(&newUser)

	return true
}

func Delete(id int64) bool {
	db := configs.ConnectDB()
	defer db.Close()

	var user User
	db.Where("ID = ?", id).Find(&user)
	db.Delete(&user)

	return true
}

func Update(id int64, userUpdate User) bool {
	db := configs.ConnectDB()
	defer db.Close()

	var user User
	db.Where("ID = ?", id).Find(&user)

	user.Email = ifExists(userUpdate.Email, userUpdate.Email, user.Email)
	user.Name = ifExists(userUpdate.Name, userUpdate.Name, user.Name)
	user.NickName = ifExists(userUpdate.NickName, userUpdate.NickName, user.NickName)

	db.Save(&user)
	return true
}
