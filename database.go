package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

type redisConfig struct {
	Host string `yaml:"HOST"`
	Port string `yaml:"PORT"`
}

type databaseConfig struct {
	Engine   string `yaml:"ENGINE"`
	Host     string `yaml:"HOST"`
	Port     string `yaml:"PORT"`
	Username string `yaml:"USERNAME"`
	Password string `yaml:"PASSWORD"`
	Schema   string `yaml:"SCHEMA"`
}

func openRedis() {
	var (
		err error
		cfg redisConfig
	)

	if err := v.UnmarshalKey("REDIS", &cfg); err != nil {
		panic(err)
	}

	conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s",
		cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}
}

func closeRedis() {
	conn.Close()
}

func openDB() {
	var (
		err     error
		cfg     databaseConfig
		command string
	)

	if err := v.UnmarshalKey("DATABASE", &cfg); err != nil {
		panic(err)
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
	db, err = gorm.Open(cfg.Engine, command)
	if err != nil {
		panic(err)
	} else {
		db.SingularTable(true)
	}
}

func closeDB() {
	_ = db.Close()
}
