package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/hibiken/asynq"
)

type redisConfig struct {
	Index    int    `yaml:"Index"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"port"`
	Password string `yaml:"Password"`
}

type redisAsynqConfig struct {
	Index    int    `yaml:"TaskIndex"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"port"`
	Password string `yaml:"Password"`
}

func openRedis() (err error) {
	var (
		cfg redisConfig
	)

	if err = v.UnmarshalKey("REDIS", &cfg); err != nil {
		return
	}
	conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s",
		cfg.Host, cfg.Port),
		redis.DialPassword(cfg.Password), redis.DialDatabase(cfg.Index))
	return
}

func closeRedis() error {
	return conn.Close()
}

func initAsynq() (err error) {
	var cfg redisAsynqConfig

	if err = v.UnmarshalKey("REDIS", &cfg); err != nil {
		return
	}
	client = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Index,
	})
	return
}
