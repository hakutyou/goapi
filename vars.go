package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	r     *gin.Engine
	db    *gorm.DB
	conn  redis.Conn
	sugar *zap.SugaredLogger
)
