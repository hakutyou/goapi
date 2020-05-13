package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
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
	// 打印接收内容
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("method: %v", c.Request.Method)
	log.Printf("body: %v", string(data))
	// 记录返回
	w := responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = &w

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	c.Next()
	// 打印返回内容
	log.Printf("response[%v]: %v", w.Status(), w.body.String())
}
