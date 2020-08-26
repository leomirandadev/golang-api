package files

import (
	"vcfConverter/settings"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetAll() (bool, []File) {
	db := settings.ConnectDB()
	defer db.Close()

	var files []File
	db.Find(&files)

	return true, files
}

func Create(newFile File) (bool, File) {
	db := settings.ConnectDB()
	defer db.Close()

	db.Create(&newFile)

	return true, newFile
}

func Delete(id int64) (bool, File) {
	db := settings.ConnectDB()
	defer db.Close()

	var file File
	db.Where("ID = ?", id).Find(&file)
	db.Delete(&file)

	return true, file
}
