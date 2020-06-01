package demo

import (
	"github.com/hakutyou/goapi/web/services"

	"github.com/garyburd/redigo/redis"
	"github.com/hibiken/asynq"
)

var (
	conn     redis.Conn
	baiduOcr services.BaiduApi
	client   *asynq.Client
)

func SetRedis(c redis.Conn) {
	conn = c
}

func SetBaiduOcr(b services.BaiduApi) {
	baiduOcr = b
}

func SetAsynq(c *asynq.Client) {
	client = c
}
