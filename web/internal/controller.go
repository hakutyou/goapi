package internal

import (
	"context"
	"log"
	"net/http"

	"github.com/hakutyou/goapi/web/utils"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	rpcx "github.com/smallnest/rpcx/client"
)

func init() {
	rpcxService = rpcx.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
}

// @Summary	过滤敏感词
// @Router	/go/internal/sensitive	[post]
func sensitiveFilter(c *gin.Context) {
	var setRequest = struct {
		Sensitive string `binding:"required" form:"word" json:"word"`
	}{}
	if err := c.ShouldBind(&setRequest); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}

	// 参数
	type Reply struct {
		Filter string
	}
	xclient := rpcx.NewXClient("DFA", rpcx.Failtry, rpcx.RandomSelect, rpcxService, rpcx.DefaultOption)
	defer xclient.Close()
	reply := &Reply{}
	err := xclient.Call(context.Background(), "DFAFilter", setRequest, reply)
	if err != nil {
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
	xclient := rpcx.NewXClient("Excel", rpcx.Failtry, rpcx.RandomSelect, rpcxService, rpcx.DefaultOption)
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

// @Summary	rpcx 服务测试
// @Description	rpcx 服务测试
// @Router	/go/internal/rpcx	[post]
func rpcxDemo(c *gin.Context) {
	type Args struct {
		A int
		B int
	}
	type Reply struct {
		C int
	}
	xclient := rpcx.NewXClient("Arith", rpcx.Failtry, rpcx.RandomSelect, rpcxService, rpcx.DefaultOption)
	defer xclient.Close()

	args := &Args{
		A: 10,
		B: 20,
	}
	reply := &Reply{}
	// err := xclient.Call(context.Background(), "Mul", args, reply)
	call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	replyCall := <-call.Done

	// if err != nil {
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	utils.Response(c, http.StatusOK, 0, "操作成功")
	return
}

// @Summary	asynq 服务测试
// @Description	asynq 服务测试
// @Router	/go/internal/delay	[post]
func runAsynq(c *gin.Context) {
	t1 := asynq.NewTask(
		"send_welcome_email",
		map[string]interface{}{"user_id": 42})
	// 立即执行
	err := client.Enqueue(t1)
	// 延迟执行, 24 小时
	// err = client.EnqueueIn(24*time.Hour, t2)
	if err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "服务器繁忙")
		return
	}
	utils.Response(c, http.StatusOK, 0, "操作成功")
	return
}
