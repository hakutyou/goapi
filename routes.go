package main

import (
	"github.com/hakutyou/goapi/account"
	"github.com/hakutyou/goapi/demo"
	"github.com/hakutyou/goapi/middleware"

	"github.com/gin-gonic/gin"
)

func Route() {
	// 路由
	demo.Routes(r.Group("/go/demo"))
	account.Routes(r.Group("/go/account"))
}

func MiddleWare() {
	// r.Use(gin.Logger())
	r.Use(middleware.LoggerMiddleware)
	r.Use(gin.Recovery())
}

func Migrations() {
	account.Models(db)
}
