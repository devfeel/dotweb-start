package server

import (
	"github.com/devfeel/dotweb-start/const"
	"github.com/devfeel/dotweb-start/global"
	"fmt"
	"github.com/devfeel/dotlog"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/config"
	"strconv"
)

func StartServer(configPath string) error {
	//初始化DotServer
	global.DotApp = dotweb.New()

	appConfig := config.MustInitConfig(configPath + "/dotweb.conf")
	global.DotApp.SetConfig(appConfig)

	//设置路由
	InitRoute(global.DotApp.HttpServer)

	innerLogger := dotlog.GetLogger(_const.LoggerName_Inner)

	innerLogger.Debug("dotweb.StartServer => " + strconv.Itoa(appConfig.Server.Port))
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(appConfig.Server.Port))
	err := global.DotApp.Start()
	if err != nil {
		panic(err)
	}
	return err
}
