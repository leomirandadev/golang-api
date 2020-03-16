package settings

import "github.com/jinzhu/gorm"

func ConnectDB() *gorm.DB {

	var host string = "localhost"
	var user string = "root"
	var password string = "root"
	var dbname string = "golang_mysql"

	db, err := gorm.Open("mysql", user+":"+password+"@("+host+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
