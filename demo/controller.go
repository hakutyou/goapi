package demo

import (
	"net/http"

	"github.com/hakutyou/goapi/utils"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func GetCache(c *gin.Context) {
	var getCacheRequest = struct {
		Key  string `binding:"required" json:"key"`
		Once bool   `json:"once,default=true"`
	}{ // 默认值
		Once: true,
	}
	// 获取参数
	if err := c.ShouldBindJSON(&getCacheRequest); err != nil {
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

func SetCache(c *gin.Context) {
	var setCacheRequest = struct {
		Key   string `binding:"required" json:"key"`
		Value string `binding:"required" json:"value"`
	}{}

	if err := c.ShouldBindJSON(&setCacheRequest); err != nil {
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
