package settings

import "github.com/jinzhu/gorm"

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost)/golang_mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
