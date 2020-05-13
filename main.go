package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/account"
	"github.com/hakutyou/goapi/demo"
	"github.com/hakutyou/goapi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

var (
	r  *gin.Engine
	db *gorm.DB
)

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
	// gin.SetMode(gin.ReleaseMode)
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
	// 连接数据库
	openDB()
	defer closeDB()

	_ = r.Run(":8080")
}
