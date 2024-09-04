//go:build !debugPlugin

package main

import (
	"github.com/eolinker/apinto-dashboard/config"
	"github.com/eolinker/apinto-dashboard/plugin/go-plugin/plugin"
	"github.com/eolinker/apinto-open-usercenter/controller"
	"github.com/eolinker/eosc/common/bean"
	"github.com/eolinker/eosc/log"
	"github.com/gin-gonic/gin"
)

func Run() {
	InitGrpc(GetConsoleUrl())

	gin.SetMode(gin.ReleaseMode)
	config.InitDb()
	config.InitRedis()
	ps := plugin.NewPlugin(controller.NewPlugin())

	//初始化超管账号 和清除超管缓存
	err := bean.Check()
	if err != nil {
		log.Fatal(err)
	}

	ps.Server()
}
