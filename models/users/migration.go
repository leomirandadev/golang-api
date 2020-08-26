package users

import "vcfConverter/settings"

func InitialMigration() {
	db := settings.ConnectDB()
	defer db.Close()

	db.AutoMigrate(&User{})
}
