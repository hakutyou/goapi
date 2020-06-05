package middleware

import (
	"bytes"
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
	c.Set("request_id", requestId)

	// 打印接收内容
	var (
		requestData string
	)
	data, err := c.GetRawData()
	if err != nil {
		sugar.Error(err.Error())
	}

	if len(string(data)) > 300 {
		// w.body.Truncate(300)
		requestData = "<data>"
	} else {
		requestData = string(data)
	}
	sugar.Infow("接收请求",
		"request_id", requestId,
		"type", c.GetHeader("Content-Type"),
		"path", c.Request.URL,
		"method", c.Request.Method,
		"body", requestData)
	// 记录返回
	w := responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = &w

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	c.Next()
	// 打印返回内容

	var (
		response string
	)
	if len(w.body.String()) > 300 {
		// w.body.Truncate(300)
		response = "<data>"
	} else {
		response = w.body.String()
	}
	sugar.Infow("返回请求",
		"request_id", requestId,
		"status", w.Status(),
		"response", response,
	)
}
