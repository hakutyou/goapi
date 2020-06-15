package main

import (
	"github.com/hibiken/asynq"
	"github.com/spf13/viper"
)

var (
	v     *viper.Viper
	redis *asynq.RedisClientOpt
)

type redisConfig struct {
	Index    int    `yaml:"TaskIndex"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"port"`
	Password string `yaml:"Password"`
}
