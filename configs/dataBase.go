package configs

import (
	"os"

	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {

	var host string = os.Getenv("DB_HOST_GOLANG")
	var user string = os.Getenv("DB_USER_GOLANG")
	var password string = os.Getenv("DB_PASSWORD_GOLANG")
	var dbname string = os.Getenv("DB_NAME_GOLANG")

	db, err := gorm.Open("mysql", user+":"+password+"@("+host+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
