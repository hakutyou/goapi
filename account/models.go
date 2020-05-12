package account

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"index;unique;not null;size:255"`
	Age  uint   `binding:"required"`
}

var db *gorm.DB

func Models(gormdb *gorm.DB) {
	db = gormdb
	db.AutoMigrate(&User{})
}
