package external

func Routes(r *gin.RouterGroup) {
	r.POST("/proxy", doProxy)
}
