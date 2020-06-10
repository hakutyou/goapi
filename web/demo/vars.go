package demo

import (
	"github.com/garyburd/redigo/redis"
)

var (
	conn redis.Conn
)

func SetRedis(c redis.Conn) {
	conn = c
}
