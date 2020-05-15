package account

import (
	"github.com/hakutyou/goapi/utils"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name     string `binding:"required" json:"name" gorm:"index;unique;not null;size:32"`
	Password string `binding:"required" gorm:"size:255"`
	Status   bool
}

func Models(gormdb *gorm.DB) {
	gormdb.AutoMigrate(&User{})
}

func (u User) Login() (user User, ret bool) {
	ret = false
	db.Select("id").Where(User{Name: u.Name, Password: u.Password, Status: true}).First(&user)
	if user.ID > 0 {
		ret = true
	}
	return
}

func GetUserInfo(userID uint) (user User, ret bool) {
	ret = false
	db.First(&user, userID)
	if user.Name != "" {
		ret = true
	}
	return
}

func (u User) GenerateToken() (string, error) {
	return utils.GenerateToken(u.ID)
}
