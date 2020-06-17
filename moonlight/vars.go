package main

var (
	cfg serverConfig
)

type serverConfig struct {
	Port int `yaml:"Port"`
}
