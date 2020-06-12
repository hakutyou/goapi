package cosfs

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
)

var (
	v      *viper.Viper
	CosApi COSApi
)

type COSApi struct {
	Region string `yaml:"Region"`
	Bucket string `yaml:"Bucket"`
	COSID  string `yaml:"COSID"`
	COSKey string `yaml:"COSKey"`
	c      *cos.Client
}

func (cosApi *COSApi) initCOS() (err error) {
	var b *url.URL

	if err = v.UnmarshalKey("COSFS", &cosApi); err != nil {
		return
	}
	if b, err = url.Parse("https://" + cosApi.Bucket + ".cos." +
		cosApi.Region + ".myqcloud.com"); err != nil {
		return
	}
	cosApi.c = cos.NewClient(&cos.BaseURL{
		BucketURL: b,
	}, &http.Client{
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosApi.COSID,
			SecretKey: cosApi.COSKey,
		},
	})
	return
}

func (cosApi *COSApi) WriteFile(filepath string, content io.Reader) (err error) {
	if _, err = cosApi.c.Object.Put(context.Background(),
		filepath, content, nil); err != nil {
		panic(err)
	}
	return
}

func (cosApi *COSApi) ReadFile(filepath string) (err error) {
	var (
		resp *cos.Response
		bs   []byte
	)

	if resp, err = cosApi.c.Object.Get(context.Background(),
		filepath, nil); err != nil {
		return
	}
	bs, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", string(bs))
	return
}
