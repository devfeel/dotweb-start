package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/devfeel/dotlog"
	"github.com/devfeel/dotweb-start/config"
	"github.com/devfeel/dotweb-start/const"
	"github.com/devfeel/dotweb-start/core/exception"
	"github.com/devfeel/dotweb-start/server"
	"github.com/devfeel/dotweb-start/task"
	"github.com/devfeel/dotweb-start/util/file"
)

var (
	innerLogger dotlog.Logger
	configFile  string
	currBinPath string
)

func init() {
	//start log service
	currBinPath = file.GetCurrentDirectory()
	err := dotlog.StartLogService(currBinPath + "/dotlog.conf")
	if err != nil {
		os.Stdout.Write([]byte("log service start error => " + err.Error()))
	}
	innerLogger = dotlog.GetLogger(_const.LoggerName_Inner)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			ex := exception.CatchError(_const.Global_ProjectName+":main", err)
			innerLogger.Error(fmt.Errorf("%v", err), ex.GetDefaultLogString())
			os.Stdout.Write([]byte(ex.GetDefaultLogString()))
		}
	}()

	//load app config
	flag.StringVar(&configFile, "config", "", "配置文件路径")
	if configFile == "" {
		configFile = currBinPath + "/app.conf"
	}

	//加载xml配置文件
	config.InitConfig(configFile)

	//监听系统信号
	//go listenSignal()

	//启动Task服务
	task.StartTaskService(currBinPath)

	err := server.StartServer(currBinPath)
	if err != nil {
		innerLogger.Warn("HttpServer.StartServer失败 " + err.Error())
		fmt.Println("HttpServer.StartServer失败 " + err.Error())
	}

}

func listenSignal() {
	c := make(chan os.Signal, 1)
	//syscall.SIGSTOP
	signal.Notify(c, syscall.SIGHUP)
	for {
		s := <-c
		innerLogger.Info("signal::ListenSignal [" + s.String() + "]")
		switch s {
		case syscall.SIGHUP: //配置重载
			innerLogger.Info("signal::ListenSignal reload config begin...")
			//重新加载配置文件
			config.InitConfig(configFile)
			innerLogger.Info("signal::ListenSignal reload config end")
		default:
			return
		}
	}
}
