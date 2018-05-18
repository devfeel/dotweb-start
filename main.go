package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/devfeel/dotweb-start/const"
	"github.com/devfeel/dotweb-start/core/exception"
	"github.com/devfeel/dotweb-start/protected"
	"github.com/devfeel/dotweb-start/util/file"
	"github.com/devfeel/dotweb-start/config"
	"github.com/devfeel/dotweb-start/server"
	"github.com/devfeel/dotweb-start/global"
	"github.com/devfeel/dotweb-start/task"
)

var (
	configFile  string
	configPath  string
	RunEnv      string
)


const (
	RunEnv_Flag       = "RunEnv"
	RunEnv_Develop    = "develop"
	RunEnv_Test       = "test"
	RunEnv_Production = "production"
)


func main() {
	defer func() {
		if err := recover(); err != nil {
			ex := exception.CatchError(_const.Global_ProjectName+":main", err)
			global.InnerLogger.Error(fmt.Errorf("%v", err), ex.GetDefaultLogString())
			os.Stdout.Write([]byte(ex.GetDefaultLogString()))
		}
	}()

	parseFlag()

	//全局初始化
	err := global.Init(configPath)
	if err!= nil {
		panic(err)
	}

	//加载全局xml配置文件
	config.InitConfig(configFile)

	//服务初始化工作
	err = protected.Init()
	if err != nil{
		global.InnerLogger.Error(err, "protected.InitConfig失败 " + err.Error())
		fmt.Println("protected.InitConfig失败 " + err.Error())
		return
	}

	//启动Task Service
	task.StartTaskService(configPath)

	//监听系统信号
	//go listenSignal()

	//启动监听服务
	err = server.StartServer(configPath)
	if err != nil {
		global.InnerLogger.Error(err, "HttpServer.StartServer失败 " + err.Error())
		fmt.Println("HttpServer.StartServer失败 " + err.Error())
	}

}

func parseFlag(){
	RunEnv = os.Getenv(RunEnv_Flag)
	if RunEnv == "" {
		RunEnv = RunEnv_Develop
	}

	configPath = _file.GetCurrentDirectory() + "/conf/"+ RunEnv
	//load app config
	flag.StringVar(&configFile, "config", "", "配置文件路径")
	if configFile == "" {
		configFile = configPath + "/app.conf"
	}

}

func listenSignal() {
	c := make(chan os.Signal, 1)
	//syscall.SIGSTOP
	signal.Notify(c, syscall.SIGHUP)
	for {
		s := <-c
		global.InnerLogger.Info("signal::ListenSignal [" + s.String() + "]")
		switch s {
		case syscall.SIGHUP: //配置重载
			global.InnerLogger.Info("signal::ListenSignal reload config begin...")
			//重新加载配置文件
			config.InitConfig(configFile)
			global.InnerLogger.Info("signal::ListenSignal reload config end")
		default:
			return
		}
	}
}
