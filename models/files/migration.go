package files

import "vcfConverter/settings"

func InitialMigration() {
	db := settings.ConnectDB()
	defer db.Close()

	db.AutoMigrate(&File{})
	db.Model(&File{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
