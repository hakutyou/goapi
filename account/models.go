package account

import (
	"encoding/hex"
	"github.com/hakutyou/goapi/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/sha3"
	"math/rand"
)

type password string

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`"x"`), nil
}

// func (p *password) UnmarshalJSON(data []byte) error {
// 	h := make([]byte, 64)
// 	c1 := sha3.NewShake256()
// 	if _, err := c1.Write(data); err != nil {
// 		return err
// 	}
// 	if _, err := c1.Read(h); err != nil {
// 		return err
// 	}
// 	*p = password(hex.EncodeToString(h))
// 	return nil
// }

func generateSalt(len int) (bytes []byte) {
	bytes = make([]byte, len)
	rand.Read(bytes)
	return
}

func (hashP *password) doHash(salt []byte) (err error) {
	h := make([]byte, 64)
	c1 := sha3.NewCShake256([]byte(""), salt)
	if _, err := c1.Write([]byte(*hashP)); err != nil {
		return err
	}
	if _, err := c1.Read(h); err != nil {
		return err
	}
	*hashP = password(hex.EncodeToString(h))
	return nil
}

func (u User) doValidate(p password) (ret bool) {
	ret = false

	db.First(&u, u.ID)
	if err := p.doHash(u.Salt); err != nil {
		return
	}

	if u.Password != p {
		return
	}
	ret = true
	return
}

type User struct {
	gorm.Model

	Name     string   `binding:"required" json:"name" gorm:"index;unique;not null;size:32"`
	Password password `binding:"required" json:"password" gorm:"size:255"`
	Salt     []byte   `json:"-" gorm:"size:64"`
	Status   bool     `json:"status"`
}

func Models(gormdb *gorm.DB) {
	gormdb.AutoMigrate(&User{})
}

func (u User) Login() (user User, ret bool) {
	ret = false
	db.Select("id").Where(User{Name: u.Name, Status: true}).First(&user)
	if user.ID <= 0 {
		ret = false
		return
	}
	ret = user.doValidate(u.Password)
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
