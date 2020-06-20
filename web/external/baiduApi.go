package external

import (
	"fmt"
	"net/http"

	"github.com/hakutyou/goapi/web/utils"

	"github.com/gin-gonic/gin"
)

// @Summary	识别身份证
// @Description	识别身份证
// @Tags 外部调用
// @Accept	mpfd
// @Produce	json
// @Param	image			formData	string	true	"图片的 base64 形式"
// @Param	id_card_side	formData	string	true	"front/back 表示 照片面/国徽面"
// @success	200	{object}	utils.ResponseDataResult	"code 为 0 表示成功"
// @success	400	{object}	utils.ResponseResult		"message 返回错误信息"
// @Router	/go/external/cache	[post]
func idCardRecognize(c *gin.Context) {
	var setRequest = struct {
		Image      string `binding:"required" form:"image" json:"image"`
		IdCardSide string `binding:"required" form:"id_card_side" json:"id_card_side"`
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
