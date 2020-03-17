package files

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type File struct {
	gorm.Model
	Name string
	Type string
}
