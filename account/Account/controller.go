package Account

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/hakutyou/goapi/account/database"
	"github.com/jinzhu/gorm"
)

type UserToken struct {
	Token string
}

func (Account) CreateAccount(_ context.Context,
	user User, reply *User) (err error) {
	var tmpUser = new(User)

	user.Status = true
	if user.Salt, err = generateSalt(32); err != nil {
		// 随机数生成失败
		return errors.New("服务器繁忙")
	}
	if err = user.Password.doHash(user.Salt); err != nil {
		// Hash 计算失败
		return errors.New("服务器繁忙")
	}
	// 检测用户是否存在
	if err = database.DBCfg.DB.Where(User{
		Name: user.Name,
	}).First(&tmpUser).Error; err != gorm.ErrRecordNotFound {
		return errors.New("用户名已存在")
	}
	err = nil
	if dbc := database.DBCfg.DB.Create(&user); dbc.Error != nil {
		// driverErr := dbc.Error.Error()
		// print(driverErr)
		return errors.New("服务器繁忙")
	}
	reply = &user
	return
}

func (Account) LoginAccount(_ context.Context,
	user User, reply *UserToken) (err error) {
	var ret bool

	ret = user.login()
	if ret == false {
		return errors.New("用户名或密码错误")
	}
	// 生成 jwt Token
	if reply.Token, err = user.generateToken(); err != nil {
		return errors.New("服务器繁忙")
	}
	return
}

func (Account) ParseToken(_ context.Context,
	token UserToken, user *User) (err error) {
	var tokenClaims *jwt.Token

	if tokenClaims, _ = jwt.ParseWithClaims(token.Token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return JwtCfg.JwtSecret, nil
		}); tokenClaims == nil {
		return errors.New("服务器繁忙")
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		// 获取用户信息
		database.DBCfg.DB.First(&user, claims.UserID)
		if user.Name == "" || user.Status == false {
			return errors.New("用户不存在")
		}
	}
	return
}
