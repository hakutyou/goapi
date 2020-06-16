package internal

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/web/utils"
	"net/http"
)

type Message struct {
	To      []string `binding:"required" form:"to" json:"to"`
	Subject string   `binding:"required" form:"subject" json:"subject"`
	Body    string   `binding:"required" form:"body" json:"body"`
}

func sendMail(c *gin.Context) {
	var (
		err            error
		messageRequest Message
		reply          struct{}
	)

	if err = c.ShouldBind(&messageRequest); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}
	xclient := Client.DoConnect("Mail")
	defer xclient.Close()
	if err = xclient.Call(context.Background(), "SendMail",
		messageRequest, &reply); err != nil {
		sugar.Errorw("RPCx服务调用错误",
			"error", err.Error())
		utils.Response(c, http.StatusBadRequest, 1, err.Error())
		return
	}
	utils.Response(c, http.StatusOK, 0, "发送成功")
	return
}
