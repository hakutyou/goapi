package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func LoggerMiddleware(c *gin.Context) {
	// 打印接收内容
	data, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("method: %v", c.Request.Method)
	log.Printf("body: %v", string(data))
	// 记录返回
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w
	c.Next()
	// 打印返回内容
	log.Printf("response: %v", w.body.String())
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
