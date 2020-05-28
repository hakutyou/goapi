package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/spf13/viper"
)

func LoadConfigure() error {
	v = viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".config.yaml")
	v.SetConfigType("yaml")

	return v.ReadInConfig()
}

func openRedis() (err error) {
	var (
		cfg redisConfig
	)
	if err = v.UnmarshalKey("REDIS", &cfg); err != nil {
		return err
	}

	redis = &asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Index,
	}

	err = nil
	return
}
