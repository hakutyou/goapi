package main

import (
	"github.com/hakutyou/goapi/account/Account"
	"github.com/hakutyou/goapi/account/database"
	"github.com/spf13/viper"
)

func LoadConfigure() (err error) {
	var v *viper.Viper

	v = viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".account.yaml")
	v.SetConfigType("yaml")

	if err = v.ReadInConfig(); err != nil {
		return
	}
	if err = v.UnmarshalKey("ACCOUNT", &cfg); err != nil {
		return
	}
	if err = v.UnmarshalKey("DATABASE", &database.DBCfg); err != nil {
		return
	}
	if err = v.UnmarshalKey("JWT_SECRET", &Account.JwtCfg); err != nil {
		return
	}
	return
}
