package main

import (
	"github.com/hakutyou/goapi/core/Mail"
	"github.com/hakutyou/goapi/core/utils/cosfs"
	"github.com/spf13/viper"
)

func LoadConfigure() (err error) {
	var v *viper.Viper

	v = viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".core.yaml")
	v.SetConfigType("yaml")

	if err = v.ReadInConfig(); err != nil {
		return
	}
	if err = v.UnmarshalKey("CORE", &cfg); err != nil {
		return
	}
	if err = v.UnmarshalKey("MAIL", &Mail.MailSetting); err != nil {
		return
	}
	if err = v.UnmarshalKey("COSFS", &cosfs.CosApi); err != nil {
		return
	}
	if err = cosfs.CosApi.InitCOS(); err != nil {
		return
	}
	return
}
