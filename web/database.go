package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/hibiken/asynq"
	"github.com/jinzhu/gorm"
)

type redisConfig struct {
	Index    int    `yaml:"Index"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
}

type redisAsynqConfig struct {
	Index    int    `yaml:"TaskIndex"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
}

type databaseConfig struct {
	Engine   string `yaml:"Engine"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Schema   string `yaml:"Schema"`
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

func openDB() (err error) {
	var (
		cfg     databaseConfig
		command string
	)

	if err = v.UnmarshalKey("DATABASE", &cfg); err != nil {
		return
	}
	if cfg.Engine == "mysql" {
		command = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	} else if cfg.Engine == "postgres" {
		command = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Schema)
	} else {
		cfg.Engine = "sqlite3"
		command = "./gorm.db"
	}
	if db, err = gorm.Open(cfg.Engine, command); err == nil {
		db.SingularTable(true)
	}
	return
}

func closeDB() error {
	return db.Close()
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
