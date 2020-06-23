package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/web/account"
	"github.com/hakutyou/goapi/web/demo"
	_ "github.com/hakutyou/goapi/web/docs"
	"github.com/hakutyou/goapi/web/external"
	"github.com/hakutyou/goapi/web/internal"
	"github.com/hakutyou/goapi/web/middleware"
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
	external.Routes(r.Group("/go/external"))
	internal.Routes(r.Group("/go/internal"))
}

func MiddleWare() {
	r.Use(middleware.LoggerMiddleware)
	r.Use(gin.Recovery())
}
