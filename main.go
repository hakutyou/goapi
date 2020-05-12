package main

import (
	"fmt"
	"github.com/hakutyou/goapi/account"
	"github.com/hakutyou/goapi/demo"
	"github.com/hakutyou/goapi/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	r *gin.Engine
)

func init() {
	// gorm 迁移
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect success")
		db.SingularTable(true)
	}
	defer db.Close()

	// gin
	r = gin.Default()
	// 中间件
	r.Use(middleware.LoggerMiddleware)

	// 数据库迁移
	account.Models(db)
	// 路由
	demo.Routes(r.Group("/go/demo"))
	account.Routes(r.Group("/go/account"))
}

func main() {
	r.Run(":8080")
}
