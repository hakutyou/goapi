package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hakutyou/goapi/account"
	"github.com/hakutyou/goapi/demo"
	"github.com/hakutyou/goapi/middleware"
	"github.com/hakutyou/goapi/utils"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

func init() {
	// 读取配置文件
	if err := LoadConfigure(); err != nil {
		panic(fmt.Sprintf("无法读取配置文件: %v\n", err))
	}

	// JWT 配置
	utils.SetEnvironment(v.GetString("JWT_SECRET"))

	// 数据库配置
	openDB()
	defer closeDB()

	// gin
	gin.SetMode(v.GetString("RUN_MODE"))
	r = gin.New()

	MiddleWare()                // 中间件
	Migrations()                // 数据库迁移
	Route(v.GetBool("SWAGGER")) // 路由
}

// @title GoAPI
// @version 0.0.1
// @description Gin 的一些 demo
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 日志
	openLogger()
	defer closeLogger()

	middleware.SetLogger(sugar)

	// 连接数据库
	openDB()
	defer closeDB()

	account.SetDatabase(db)

	// 连接 Redis
	openRedis()
	defer closeRedis()

	demo.SetRedis(conn)

	// 运行 gin
	// TODO: 需要一个热更新代码的方式, gracehttp
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	sugar.Info("Server started")
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("listen: %s\n", err))
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	sugar.Info("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		sugar.Info("Server Shutdown: ", err)
	}
	sugar.Info("Server exiting")
}

func openRedis() {
	var err error

	conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
}

func closeRedis() {
	conn.Close()
}

func openLogger() {
	var cfg zap.Config

	zapConfig, _ := ioutil.ReadFile(".zap.yaml")
	_ = yaml.Unmarshal(zapConfig, &cfg)

	if err := yaml.Unmarshal(zapConfig, &cfg); err != nil {
		panic(err)
	}
	logger, _ := cfg.Build()
	sugar = logger.Sugar()
}

func closeLogger() {
	_ = sugar.Sync()
}
