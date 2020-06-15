package demo

import (
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/web/utils"
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

func slowLoading(c *gin.Context) {
	i := 0
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		if i >= 3 {
			break
		}

		select {
		case <-t.C:
			// 1s 触发一次
			utils.Response(c, http.StatusOK, 0, "操作成功")
		case <-c.Request.Context().Done():
			// 客户端中断
			print("closed")
			return
		}
		print("OK")
		i += 1
	}
}
