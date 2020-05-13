package account

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createPeople(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("%v", err.Error())
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
