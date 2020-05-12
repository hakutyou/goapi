package main

import (
	"fmt"
	"net/http"

	"github.com/hakutyou/goapi/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name string `gorm:"index;unique;not null;size:255"`
	Age  uint   `binding:"required"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var db *gorm.DB
var err error

func main() {
	// gorm 迁移
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect success")
		db.SingularTable(true)
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	// gin
	r := gin.Default()
	// 中间件
	r.Use(middleware.LoggerMiddleware)

	// 路由
	r.GET("/go/ping/", doRequest)
	r.POST("/go/people/", createPeople)
	r.Run(":8080")
}

func doRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func createPeople(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{1, "参数格式错误"})
		return
	}

	if dbc := db.Create(&user); dbc.Error != nil {
		driverErr := dbc.Error.Error()
		c.JSON(http.StatusBadRequest, Response{1, driverErr})
		return
	}
	c.JSON(200, user)
}
