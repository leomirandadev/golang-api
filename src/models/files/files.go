package files

import (
	"vcfConverter/src/settings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type File struct {
	gorm.Model
	Name string
	Type string
}

func InitialMigration() {
	db := settings.ConnectDB()
	defer db.Close()

	db.AutoMigrate(&File{})
}

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
