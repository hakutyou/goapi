package account

import (
	"encoding/hex"
	"github.com/hakutyou/goapi/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/sha3"
)

type password string

func (p *password) UnmarshalJSON(data []byte) error {
	h := make([]byte, 64)
	c1 := sha3.NewShake256()
	if _, err := c1.Write(data); err != nil {
		return err
	}
	if _, err := c1.Read(h); err != nil {
		return err
	}
	*p = password(hex.EncodeToString(h))
	return nil
}

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`"x"`), nil
}

type User struct {
	gorm.Model

	Name     string   `binding:"required" json:"name" gorm:"index;unique;not null;size:32"`
	Password password `binding:"required" json:"password" gorm:"size:255"`
	Status   bool     `json:"status"`
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
