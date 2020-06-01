package external

import (
	"fmt"
	"github.com/hakutyou/goapi/web/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func doProxy(c *gin.Context) {
	var (
		statusCode int
		retText    string
		err        error
	)

	var requestInfo = struct {
		Url    string `binding:"required" form:"url" json:"url"`
		Method string `form:"method" json:"method"`
		Json   string `form:"json" json:"json"`
	}{ // 默认值
		Method: "post",
		Json:   "",
	}
	// 获取参数
	if err := c.ShouldBind(&requestInfo); err != nil {
		fmt.Printf("%v\n", err)
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}

	statusCode, retText, err = utils.ServiceProxy(c.MustGet("request_id").(string),
		requestInfo.Method, requestInfo.Url, requestInfo.Json)
	if err != nil {
		return
	}
	c.String(statusCode, retText)
}
