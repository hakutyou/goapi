package services

import (
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

var (
	conn  redis.Conn
	sugar *zap.SugaredLogger
)

func SetRedis(c redis.Conn) {
	conn = c
}

func SetLogger(sugarLogger *zap.SugaredLogger) {
	sugar = sugarLogger
}
