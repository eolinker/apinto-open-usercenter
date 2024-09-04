//go:build debugPlugin

package main

import (
	"github.com/eolinker/apinto-dashboard/config"
	apinto_module "github.com/eolinker/apinto-dashboard/module"
	_ "github.com/eolinker/apinto-dashboard/plugin/go-plugin/plugin"
	"github.com/eolinker/apinto-open-usercenter/controller"
	"github.com/eolinker/eosc/common/bean"
	"github.com/eolinker/eosc/log"
	"github.com/gin-gonic/gin"
	"net"
)

func Run() {
	InitGrpc(GetConsoleUrl())
	
	gin.SetMode(gin.ReleaseMode)
	config.InitDb()
	config.InitRedis()
	plugin := controller.NewPlugin()
	
	//初始化超管账号 和清除超管缓存
	err := bean.Check()
	if err != nil {
		log.Fatal(err)
	}
	
	engine := gin.Default()
	for _, a := range plugin.Middleware() {
		if handle, has := a.RequestHandler(); has {
			engine.Use(func(ctx *gin.Context) {
				r := &apinto_module.MiddlewareRequest{
					FulPath: ctx.Request.URL.Path,
					Url:     ctx.Request.URL.String(),
					Method:  ctx.Request.Method,
					Header:  ctx.Request.Header,
					Keys:    make(map[string]any),
				}
				
				handle(ctx, r, new(apinto_module.MiddlewareResponse))
				
			})
		}
	}
	for _, a := range plugin.Apis() {
		engine.Handle(a.Method, a.Path, a.HandlerFunc)
	}
	l, err := net.Listen("tcp", ":9900")
	if err != nil {
		log.Fatal(err)
	}
	
	err = engine.RunListener(l)
	if err != nil {
		log.Fatal(err)
		return
	}
}
