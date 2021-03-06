package main

import (
	"fmt"
	"github.com/hakutyou/goapi/account/Account"
	"github.com/hakutyou/goapi/account/database"
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
	Account.Models(database.DBCfg.DB)
}

func main() {
	var (
		err error
		s   *server.Server
	)

	s = server.NewServer()
	if err = s.RegisterName("Account", new(Account.Account), ""); err != nil {
		panic(err)
	}
	if err = s.Serve("tcp",
		fmt.Sprintf("localhost:%d", cfg.Port)); err != nil {
		panic(err)
	}
}
