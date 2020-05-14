package account

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func SetDatabase(gormdb *gorm.DB) {
	db = gormdb
}
