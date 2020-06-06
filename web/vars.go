package main

import (
	"github.com/hakutyou/goapi/web/services"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	v          *viper.Viper
	r          *gin.Engine
	db         *gorm.DB
	conn       redis.Conn
	client     *asynq.Client
	sugar      *zap.SugaredLogger
	tencentSms *services.TencentSms
)
