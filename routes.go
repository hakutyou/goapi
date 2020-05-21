package main

import (
	"github.com/hakutyou/goapi/account"
	"github.com/hakutyou/goapi/demo"
	_ "github.com/hakutyou/goapi/docs"
	"github.com/hakutyou/goapi/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func LoadConfigure() error {
	v = viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".config.yaml")
	v.SetConfigType("yaml")

	return v.ReadInConfig()
}

func Route(swagger bool) {
	// Swagger 文档
	if swagger {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
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
