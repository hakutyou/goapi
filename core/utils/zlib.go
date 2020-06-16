package utils

import (
	"bytes"
	"compress/zlib"
	"io"
)

// zlib 压缩
func Compress(src []byte) (outData []byte, err error) {
	var in bytes.Buffer

	w := zlib.NewWriter(&in)
	_, err = w.Write(src)
	if err != nil {
		return
	}
	err = w.Close()
	if err != nil {
		return
	}
	outData = in.Bytes()
	return
}

// zlib 解压
func UnCompress(compressSrc []byte) (outData []byte, err error) {
	var (
		r   io.ReadCloser
		out bytes.Buffer
	)

	b := bytes.NewReader(compressSrc)
	r, err = zlib.NewReader(b)
	if err != nil {
		return
	}
	_, err = io.Copy(&out, r)
	if err != nil {
		return
	}
	outData = out.Bytes()
	return
}
