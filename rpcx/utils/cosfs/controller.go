package cosfs

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	// 读取配置文件
	if err := LoadConfigure(); err != nil {
		panic(fmt.Sprintf("无法读取配置文件: %v\n", err))
	}
	if err := CosApi.initCOS(); err != nil {
		panic(err)
	}
}

func LoadConfigure() error {
	v = viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".rpcx.yaml")
	v.SetConfigType("yaml")

	return v.ReadInConfig()
}
