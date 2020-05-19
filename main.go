package main

import (
	"context"
	"encoding/json"
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
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func init() {
	// 读取配置文件
	err := godotenv.Load(".env")
	if err != nil {
		panic("无法读取 .env 文件")
	}

	// JWT 配置
	utils.SetEnvironment(os.Getenv("JWT_SECRET"))

	// 数据库配置
	openDB()
	defer closeDB()

	// gin
	gin.SetMode(os.Getenv("RUN_MODE"))
	r = gin.New()

	MiddleWare() // 中间件
	Migrations() // 数据库迁移
	Route()      // 路由
}

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

func openDB() {
	var (
		err     error
		command string
	)

	database := os.Getenv("DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	if database == "mysql" {
		command = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUsername, dbPassword, dbHost, dbPort, dbDatabase)
	} else if database == "postgres" {
		command = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
			dbHost, dbPort, dbUsername, dbPassword, dbDatabase)
	} else {
		database = "sqlite3"
		command = "./gorm.db"
	}

	db, err = gorm.Open(database, command)
	if err != nil {
		panic(err)
	} else {
		db.SingularTable(true)
	}
}

func closeDB() {
	_ = db.Close()
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
	var (
		err error
		cfg zap.Config
	)

	zapConfig, _ := ioutil.ReadFile("zap.config")
	if err = json.Unmarshal(zapConfig, &cfg); err != nil {
		panic(err)
	}
	logger, _ := cfg.Build() // zap.NewProduction()
	sugar = logger.Sugar()
}

func closeLogger() {
	_ = sugar.Sync()
}
