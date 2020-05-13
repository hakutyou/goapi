package account

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"index;unique;not null;size:255" json:"name" binding:"required"`
	Age  uint   `binding:"required" json:"age"`
}

var db *gorm.DB

func Models(gormdb *gorm.DB) {
	gormdb.AutoMigrate(&User{})
}

func SetDatabase(gormdb *gorm.DB) {
	db = gormdb
}
