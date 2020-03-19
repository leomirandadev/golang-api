package files

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type File struct {
	ID        int `gorm:"primary_key"`
	Name      string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
