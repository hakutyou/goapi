package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func (cfg *databaseConfig) OpenDB() (err error) {
	var command string

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
	if cfg.DB, err = gorm.Open(cfg.Engine, command); err == nil {
		cfg.DB.SingularTable(true)
	}
	return
}
