package users

import "github.com/leomirandadev/golang-api/configs"

func InitialMigration() {
	db := configs.ConnectDB()
	defer db.Close()

	db.AutoMigrate(&User{})
}
