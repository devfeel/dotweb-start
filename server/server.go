package server

import (
	"github.com/devfeel/dotweb-start/global"
	_ "fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/config"
	"github.com/devfeel/middleware/cors"
	"strconv"
)

func StartServer(configPath string) error {
	//初始化DotServer
	appConfig := config.MustInitConfig(configPath + "/dotweb.conf")
	global.DotApp = dotweb.ClassicWithConf(appConfig)

	global.DotApp.SetDevelopmentMode()
	global.DotApp.UseRequestLog()
	global.DotApp.Use(cors.Middleware(cors.NewConfig().UseDefault()))

	//设置路由
	InitRoute(global.DotApp.HttpServer)

	global.InnerLogger.Debug("dotweb.StartServer => " + strconv.Itoa(appConfig.Server.Port))
	err := global.DotApp.Start()
	if err != nil {
		panic(err)
	}
	return err
}
