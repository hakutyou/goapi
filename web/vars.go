package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	v     *viper.Viper
	r     *gin.Engine
	db    *gorm.DB
	conn  redis.Conn
	sugar *zap.SugaredLogger
)
