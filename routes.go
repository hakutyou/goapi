package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/account"
	"github.com/hakutyou/goapi/demo"
	"github.com/hakutyou/goapi/middleware"
)

func Route() {
	// 路由
	demo.Routes(r.Group("/go/demo"))
	account.Routes(r.Group("/go/account"))
}

func MiddleWare() {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware)
}

func Migrations() {
	account.Models(db)
}
