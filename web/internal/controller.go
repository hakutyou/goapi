package internal

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/web/utils"
	"net/http"
)

// @Summary	过滤敏感词
// @Router	/go/internal/sensitive	[post]
func sensitiveFilter(c *gin.Context) {
	var (
		err        error
		setRequest = struct {
			Sensitive string `binding:"required" form:"word" json:"word"`
		}{}
		reply = &struct {
			Filter string
		}{}
	)

	if err = c.ShouldBind(&setRequest); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}
	xclient := Client.DoConnect("DFA")
	defer xclient.Close()
	if err = xclient.Call(context.Background(), "DFAFilter",
		setRequest, reply); err != nil {
		sugar.Errorw("RPCx服务调用失败",
			"error", err.Error())
		utils.Response(c, http.StatusBadRequest, 1, "服务器繁忙")
		return
	}
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", gin.H{
		"word": reply.Filter,
	})
	return
}

// @Summary	导出 Excel
// @Router	/go/internal/excel[post]
func excelDemo(c *gin.Context) {
	type Args struct {
		Title string
	}
	type Reply struct {
		Url string
	}
	print(c.Request.RemoteAddr)
	xclient := Client.DoConnect("Excel")
	defer xclient.Close()
	args := &Args{
		Title: "wx",
	}
	reply := &Reply{}
	call, err := xclient.Go(context.Background(), "GenerateExcel", args, reply, nil)
	if err != nil {
		sugar.Errorw("RPCx服务调用失败",
			"error", err.Error())
		utils.Response(c, http.StatusBadRequest, 1, "服务器繁忙")
		return
	}
	replyCall := <-call.Done
	if replyCall.Error != nil {
		sugar.Errorw("RPCx服务调用失败",
			"error", replyCall.Error.Error())
		utils.Response(c, http.StatusBadRequest, 1, "服务器繁忙")
		return
	}
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", reply)
	return
}

// // @Summary	asynq 服务测试
// // @Description	asynq 服务测试
// // @Router	/go/internal/delay	[post]
// func runAsynq(c *gin.Context) {
// 	t1 := asynq.NewTask(
// 		"send_welcome_email",
// 		map[string]interface{}{"user_id": 42})
// 	// 立即执行
// 	err := aclient.Enqueue(t1)
// 	// 延迟执行, 24 小时
// 	// err = aclient.EnqueueIn(24*time.Hour, t2)
// 	if err != nil {
// 		utils.Response(c, http.StatusBadRequest, 1, "服务器繁忙")
// 		return
// 	}
// 	utils.Response(c, http.StatusOK, 0, "操作成功")
// 	return
// }
