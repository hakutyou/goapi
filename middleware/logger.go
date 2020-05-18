package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware(c *gin.Context) {
	// 获取 request_id
	requestId := c.GetHeader("request_id")
	if requestId == "" {
		requestId = uuid.NewV4().String()
	}

	// 打印接收内容
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}
	sugar.Infow("接收请求",
		"request_id", requestId,
		"path", c.Request.URL,
		"method", c.Request.Method,
		"body", string(data))
	// 记录返回
	w := responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = &w

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	c.Next()
	// 打印返回内容

	sugar.Infow("返回请求",
		"request_id", requestId,
		"status", w.Status(),
		"response", w.body.String(),
	)
}
