package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
