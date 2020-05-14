package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/account"
	"github.com/hakutyou/goapi/demo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

func init() {
	// 读取配置文件
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("无法读取 .env 文件")
	}

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
	// 连接数据库
	openDB()
	defer closeDB()

	account.SetDatabase(db)

	// 连接 Redis
	openRedis()
	defer closeRedis()

	demo.SetRedis(conn)

	// 运行 gin
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func openDB() {
	var err error

	// db, err = gorm.Open("sqlite3", "./gorm.db")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbDatabase))
	// db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
	// 	db_host, db_port, db_username, db_password, db_database))

	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("connect success")
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
		log.Panic(err)
	}
}

func closeRedis() {
	conn.Close()
}
