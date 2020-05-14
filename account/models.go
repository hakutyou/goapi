package account

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

var jwtSecret = []byte("123123123123")

type User struct {
	gorm.Model

	Name     string `binding:"required" json:"name" gorm:"index;unique;not null;size:32"`
	Password string `binding:"required" gorm:"size:255"`
	Status   bool
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func Models(gormdb *gorm.DB) {
	gormdb.AutoMigrate(&User{})
}

func (u User) CheckAuth() bool {
	var user User
	db.Select("id").Where(User{Name: u.Name, Password: u.Password}).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func (u User) GenerateToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		u.Name,
		u.Password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "sirat",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}
