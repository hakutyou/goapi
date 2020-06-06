package demo

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hakutyou/goapi/web/utils"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	rpcx "github.com/smallnest/rpcx/client"
)

// @Summary	获取 Redis 缓存
// @Description	获取 Redis 缓存
// @Tags Demo
// @Accept	mpfd
// @Produce	json
// @Param	key		query	string	true	"键"
// @Param	once	query	bool	false	"是否删除"
// @success	200	{object}	utils.ResponseDataResult	"code 为 0 表示成功"
// @success	400	{object}	utils.ResponseResult		"message 返回错误信息"
// @Router	/go/demo/cache	[get]
func getCache(c *gin.Context) {
	var getCacheRequest = struct {
		Key  string `binding:"required" form:"key" json:"key"`
		Once bool   `form:"once" json:"once"`
	}{ // 默认值
		Once: true,
	}
	// 获取参数
	if err := c.ShouldBind(&getCacheRequest); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}
	// 获得 Redis 的值
	value, err := redis.String(conn.Do("GET", getCacheRequest.Key))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, 2, "键不存在")
		return
	}
	// 检测 Once 参数
	if getCacheRequest.Once {
		_, err = conn.Do("DEL", getCacheRequest.Key)
		if err != nil {
			utils.Response(c, http.StatusBadRequest, 99, "服务器繁忙")
			return
		}
	}
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", gin.H{
		"value": value,
	})
	return
}

// @Summary	设置 Redis 缓存
// @Description	设置 Redis 缓存
// @Tags Demo
// @Accept	mpfd
// @Produce	json
// @Param	key		formData	string	true	"键"
// @Param	value	formData	string	true	"值"
// @success	200	{object}	utils.ResponseDataResult	"code 为 0 表示成功"
// @success	400	{object}	utils.ResponseResult		"message 返回错误信息"
// @Router	/go/demo/cache	[post]
func setCache(c *gin.Context) {
	var setCacheRequest = struct {
		Key   string `binding:"required" form:"key" json:"key"`
		Value string `binding:"required" form:"value" json:"value"`
	}{}

	if err := c.ShouldBind(&setCacheRequest); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}

	if _, err := conn.Do("SET", setCacheRequest.Key, setCacheRequest.Value); err != nil {
		utils.Response(c, http.StatusBadRequest, 99, "服务器繁忙")
		return
	}

	utils.Response(c, http.StatusOK, 0, "操作成功")
	return
}

// @Summary	识别身份证
// @Description	识别身份证
// @Tags Demo
// @Accept	mpfd
// @Produce	json
// @Param	image			formData	string	true	"图片的 base64 形式"
// @Param	id_card_side	formData	string	true	"front/back 表示 照片面/国徽面"
// @success	200	{object}	utils.ResponseDataResult	"code 为 0 表示成功"
// @success	400	{object}	utils.ResponseResult		"message 返回错误信息"
// @Router	/go/demo/cache	[post]
func idCardRecognize(c *gin.Context) {
	var setRequest = struct {
		Image      string `form:"image" json:"image"`
		IdCardSide string `form:"id_card_side" json:"id_card_side"`
	}{}

	if err := c.ShouldBind(&setRequest); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}

	retJson, err := baiduOcr.IdCardRecognition(c.MustGet("request_id").(string), setRequest.Image, setRequest.IdCardSide)
	if err != nil {
		utils.Response(c, http.StatusBadRequest, 99, fmt.Sprintf("服务器繁忙 - (%s)", err.Error()))
		return
	}

	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", retJson)
	return
}

func rpcxDemo(c *gin.Context) {
	type Args struct {
		A int
		B int
	}
	type Reply struct {
		C int
	}

	d := rpcx.NewPeer2PeerDiscovery("tcp@localhost:8972", "")

	xclient := rpcx.NewXClient("Arith", rpcx.Failtry, rpcx.RandomSelect, d, rpcx.DefaultOption)
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
