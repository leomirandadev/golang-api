package files

import "github.com/leomirandadev/golang-api/configs"

func InitialMigration() {
	db := configs.ConnectDB()
	defer db.Close()

	db.AutoMigrate(&File{})
	db.Model(&File{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
