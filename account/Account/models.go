package Account

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hakutyou/goapi/account/database"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model

	Name     string   `gorm:"index;unique;not null;size:32"`
	Password password `gorm:"size:255"`
	Salt     []byte   `gorm:"size:64"`
	Status   bool
}

func Models(gormdb *gorm.DB) {
	gormdb.AutoMigrate(&User{})
}

func (u *User) login() (ret bool) {
	ret = false
	database.DBCfg.DB.Select("id").Where(User{Name: u.Name, Status: true}).First(&u)
	if u.ID <= 0 {
		return
	}
	ret = u.doValidate(u.Password)
	return
}

func (u User) doValidate(p password) bool {
	database.DBCfg.DB.First(&u, u.ID)
	if err := p.doHash(u.Salt); err != nil {
		// Hash 错误
		return false
	}
	if u.Password != p {
		// 用户名密码错误
		return false
	}
	return true
}

func (u User) generateToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			// 平台标识
			Issuer: "hakutyou",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(JwtCfg.JwtSecret)
}
