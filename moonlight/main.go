package main

import (
	"fmt"
	"github.com/hakutyou/goapi/moonlight/Bang"
	"github.com/hakutyou/goapi/moonlight/Rank"
	"github.com/hakutyou/goapi/moonlight/database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/smallnest/rpcx/server"
)

func init() {
	var err error

	if err = LoadConfigure(); err != nil {
		panic(err)
	}
	// 连接数据库
	if err = database.DBCfg.OpenDB(); err != nil {
		panic(err)
	}
	// 迁移
	Bang.Models(database.DBCfg.DB)
	Rank.Models(database.DBCfg.DB)
}

func main() {
	var (
		err error
		s   *server.Server
	)

	s = server.NewServer()
	if err = s.RegisterName("Bang", new(Bang.Bang), ""); err != nil {
		panic(err)
	}
	if err = s.RegisterName("Rank", new(Rank.Rank), ""); err != nil {
		panic(err)
	}
	if err = s.Serve("tcp",
		fmt.Sprintf("localhost:%d", cfg.Port)); err != nil {
		panic(err)
	}
}
