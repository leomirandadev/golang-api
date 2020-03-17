package users

import "vcfConverter/src/settings"

func InitialMigration() {
	db := settings.ConnectDB()
	defer db.Close()

	db.AutoMigrate(&User{})
}
