package users

import (
	"vcfConverter/src/settings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	NickName string
	Name     string
	Email    string
	Password string
}

func InitialMigration() {
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

func GetByEmailPassword(email string, password string) (bool, []User) {
	db := settings.ConnectDB()
	defer db.Close()

	var user []User
	db.Where("email = ?", email).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(password))

	if err == nil {
		return true, user
	}

	return false, []User{}
}

func Create(newUser User) (bool, User) {
	db := settings.ConnectDB()
	defer db.Close()

	newUser.Password = hashPassword(newUser.Password)

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

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
