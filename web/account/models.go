package account

import (
	"encoding/hex"
	"math/rand"

	"github.com/hakutyou/goapi/web/utils"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/sha3"
)

type password string

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`"x"`), nil
}

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

type User struct {
	gorm.Model

	Name     string   `binding:"required" form:"name" json:"name" gorm:"index;unique;not null;size:32" example:"hakutyou"`
	Password password `binding:"required" form:"password" json:"password" gorm:"size:255" example:"myPassword"`
	Salt     []byte   `form:"-" json:"-" gorm:"size:64"`
	Status   bool     `form:"-" json:"-"`
}

func Models(gormdb *gorm.DB) {
	gormdb.AutoMigrate(&User{})
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

func (u User) login() (user User, ret bool) {
	ret = false
	db.Select("id").Where(User{Name: u.Name, Status: true}).First(&user)
	if user.ID <= 0 {
		ret = false
		return
	}
	ret = user.doValidate(u.Password)
	return
}

func getUserInfo(userID uint) (user User, ret bool) {
	ret = false
	db.First(&user, userID)
	if user.Name != "" {
		ret = true
	}
	return
}

func (u User) generateToken() (string, error) {
	return utils.GenerateToken(u.ID)
}
