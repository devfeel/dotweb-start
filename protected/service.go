package protected

import (
	"github.com/devfeel/dotweb-start/config"
	"errors"
	"github.com/devfeel/dotlog"
	"github.com/devfeel/dotweb-start/const"
)

type ServiceLoader func()

const (
	DefaultRedisID    = "default"
	defaultDatabaseID = "demodb"
)

var (
	DefaultConfig   *ServiceConfig
	serviceLoaderMap map[string]ServiceLoader
	ServiceLogger   dotlog.Logger
)

func init() {
	serviceLoaderMap = make(map[string]ServiceLoader)
}

func InitLogger() {
	ServiceLogger = dotlog.GetLogger(_const.LoggerName_Service)
}

func Init() error {
	InitLogger()
	var err error

	//获取数据库连接字符串
	dbInfo, exists := config.GetDataBaseInfo(defaultDatabaseID)
	if !exists || dbInfo.ServerUrl == ""{
		err = errors.New("no config " + defaultDatabaseID + " database config")
		ServiceLogger.Error(err, "no config " + defaultDatabaseID + " database config")
		return err
	}

	DefaultConfig = &ServiceConfig{
		DefaultDBConn: dbInfo.ServerUrl,
	}

	//初始化Redis配置信息
	redis, exists := config.GetRedisInfo(DefaultRedisID)
	if !exists {
		err = errors.New("no exists " + DefaultRedisID + " logger config")
		ServiceLogger.Error(err, "not exists "+DefaultRedisID+" logger config")
		return err
	}
	DefaultConfig.DefaultRedisConn = redis.ServerUrl

	//执行已注册的配置初始化接口
	for _, loader := range serviceLoaderMap {
		loader()
	}
	return nil
}

// RegisterServiceLoader 注册服务加载接口
func RegisterServiceLoader(serviceName string, service ServiceLoader) {
	_, exists:= serviceLoaderMap[serviceName]
	if exists{
		panic("Service.RegisterServiceLoader already exists service " + serviceName)
	}
	serviceLoaderMap[serviceName] = service
}