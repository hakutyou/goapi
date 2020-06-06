package services

import (
	"github.com/garyburd/redigo/redis"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
	"go.uber.org/zap"
)

var (
	conn             redis.Conn
	sugar            *zap.SugaredLogger
	tencentSmsClient *sms.Client
)

func SetRedis(c redis.Conn) {
	conn = c
}

func SetLogger(sugarLogger *zap.SugaredLogger) {
	sugar = sugarLogger
}
