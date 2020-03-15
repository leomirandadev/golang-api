package users

import (
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

func connectDataBase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost)/golang_mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func InitialMigrationUser() {
	db := connectDataBase()
	defer db.Close()

	db.AutoMigrate(&User{})
}

func GetAll() (bool, []User) {
	db := connectDataBase()
	defer db.Close()

	var users []User
	db.Find(&users)
	return true, users
}

func Create(newUser User) (bool, User) {
	db := connectDataBase()
	defer db.Close()

	db.Create(&newUser)
	return true, newUser
}

// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	name := vars["name"]

// 	var user User
// 	db.Where("name = ?", name).Find(&user)
// 	db.Delete(&user)

// 	fmt.Fprintf(w, "Successfully Deleted User")
// }

// func updateUser(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	name := vars["name"]
// 	email := vars["email"]

// 	var user User
// 	db.Where("name = ?", name).Find(&user)

// 	user.Email = email

// 	db.Save(&user)
// 	fmt.Fprintf(w, "Successfully Updated User")
// }
