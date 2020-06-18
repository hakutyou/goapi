package Bang

import "github.com/jinzhu/gorm"

type Skill struct {
	gorm.Model

	TypeId   uint
	Name     string `gorm:"size:32"`
	MaxLevel uint
	MainDes  string `gorm:"size:64"`
}

func (Skill) TableName() string {
	return "skill"
}

type SkillDetail struct {
	gorm.Model

	Level    uint
	XiuWei   uint
	BangGong uint
	SuiYin   uint
	Des      string `gorm:"size:256"`
	Props    string `gorm:"size:32"`

	SkillID uint `gorm:"column:skill_id,index"`
}

func Models(gormdb *gorm.DB) {
	gormdb.AutoMigrate(&Skill{}, &SkillDetail{})
}
