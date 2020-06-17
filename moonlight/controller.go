package main

import (
	"github.com/hakutyou/goapi/moonlight/database"
	"github.com/spf13/viper"
)

func LoadConfigure() (err error) {
	var v *viper.Viper

	v = viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".moonlight.yaml")
	v.SetConfigType("yaml")

	if err = v.ReadInConfig(); err != nil {
		return
	}
	if err = v.UnmarshalKey("MOONLIGHT", &cfg); err != nil {
		return
	}
	if err = v.UnmarshalKey("DATABASE", &database.DBCfg); err != nil {
		return
	}
	return
}
