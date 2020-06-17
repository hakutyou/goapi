package database

import "github.com/jinzhu/gorm"

var (
	DBCfg databaseConfig
)

type databaseConfig struct {
	Engine   string `yaml:"Engine"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Schema   string `yaml:"Schema"`
	DB       *gorm.DB
}
