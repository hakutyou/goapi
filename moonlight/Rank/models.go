package Rank

import "github.com/jinzhu/gorm"

type AccountRank struct {
	gorm.Model

	OrigName string `gorm:"index;unique;not null;size:32"`
	Name     string `gorm:"size:32"`
	OldName1 string `gorm:"size:32"`
	OldName2 string `gorm:"size:32"`

	BangPai  string `gorm:"size:32"`
	LianMeng string `gorm:"size:32"`
        Sex      string `gorm:"size:32"`
	// Sex      string `gorm:"type:enum('male','female','little_female')"`
	Class    string `gorm:"size:32"`
	// Class    string `gorm:"type:enum('zw','tb','sw','gb','tm','wd','sl','tx','sd','yh')"`
	GongLi   uint
}

func Models(gormdb *gorm.DB) {
	gormdb.AutoMigrate(&AccountRank{})
}
