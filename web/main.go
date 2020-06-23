package main

import (
	"errors"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/web/account"
	"github.com/hakutyou/goapi/web/demo"
	"github.com/hakutyou/goapi/web/external"
	"github.com/hakutyou/goapi/web/internal"
	"github.com/hakutyou/goapi/web/middleware"
	"github.com/hakutyou/goapi/web/services"
	"github.com/hakutyou/goapi/web/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	// pid 文件
	if err := newPIDFile("web.pid"); err != nil {
		panic(err)
	}

	// 读取配置文件
	if err := LoadConfigure(); err != nil {
		panic(fmt.Sprintf("无法读取配置文件: %v\n", err))
	}

	// JWT 配置
	// utils.SetEnvironment(v.GetString("JWT_SECRET"))

	// rpcx 配置
	if err := initRpcx(); err != nil {
		panic(err)
	}

	// asynq 配置
	// if err := initAsynq(); err != nil {
	// 	panic(err)
	// }

	// 其他服务设置
	if err := initServices(); err != nil {
		panic(err)
	}
	account.SetTencentSms(tencentSms)

	// API 服务配置
	if err := openBaiduOcrService(); err != nil {
		panic(err)
	}

	// gin
	gin.SetMode(v.GetString("RUN_MODE"))
	r = gin.New()

	MiddleWare()                // 中间件
	Route(v.GetBool("SWAGGER")) // 路由
}

// @title GoAPI
// @version 0.0.1
// @description Gin 的一些 demo
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 日志
	openLogger()
	defer closeLogger()

	account.SetLogger(sugar)
	internal.SetLogger(sugar)
	external.SetLogger(sugar)
	services.SetLogger(sugar)
	utils.SetLogger(sugar)
	middleware.SetLogger(sugar)

	// 连接 Redis
	if err := openRedis(); err != nil {
		panic(err)
	}
	defer closeRedis()

	services.SetRedis(conn)
	demo.SetRedis(conn)

	// 连接 Asynq
	// internal.SetAsynq(client)

	// 运行 gin
	// debugServer(":8080", r)
	// gracehttp. 热更新代码
	if err := gracehttp.Serve(
		&http.Server{Addr: ":8080", Handler: r}); err != nil {
		sugar.Info("Server error:  ", err)
	}
}

func openLogger() {
	var cfg zap.Config

	zapConfig, _ := ioutil.ReadFile(".zap.yaml")
	_ = yaml.Unmarshal(zapConfig, &cfg)

	if err := yaml.Unmarshal(zapConfig, &cfg); err != nil {
		panic(err)
	}
	logger, _ := cfg.Build()
	sugar = logger.Sugar()
}

func closeLogger() error {
	return sugar.Sync()
}

func initServices() error {
	return v.UnmarshalKey("TENCENT_SMS", &tencentSms)
}

func newPIDFile(path string) (err error) {
	var (
		pidByte []byte
	)
	if pidByte, err = ioutil.ReadFile(path); err == nil {
		pid := strings.TrimSpace(string(pidByte))
		if _, err := os.Stat(filepath.Join("/proc", pid)); err == nil {
			return errors.New("进程正在运行")
		}
	}
	if err = os.MkdirAll(filepath.Dir(path), os.FileMode(0755)); err != nil {
		return
	}
	if err = ioutil.WriteFile(path, []byte(fmt.Sprintf("%d", os.Getpid())), 0644); err != nil {
		return
	}
	return
}
