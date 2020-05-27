package demo

import (
	"github.com/garyburd/redigo/redis"
	"github.com/hakutyou/goapi/services"
)

var (
	conn     redis.Conn
	baiduOcr services.BaiduApi
)

func SetRedis(c redis.Conn) {
	conn = c
}

func SetBaiduOcr(b services.BaiduApi) {
	baiduOcr = b
}
