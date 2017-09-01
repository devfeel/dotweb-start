package test

import (
	"github.com/devfeel/dotlog"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb-start/config"
	"github.com/devfeel/dotweb-start/const"
	"github.com/devfeel/dotweb-start/global"
	"github.com/devfeel/dotweb-start/util/redis"
)

func Index(ctx dotweb.Context) error {
	dotlog.GetLogger("HandlerLogger").Debug("test Index")
	global.GlobalContext.Set("1", 1)
	ctx.WriteString("Index Version - ", _const.Global_Version)
	return nil
}

func Json(ctx dotweb.Context) error {
	type info struct {
		Name string
		Age  int
	}
	i := &info{
		Name: "test name",
		Age:  1,
	}
	ctx.WriteJson(i)
	return nil
}

func Redis(ctx dotweb.Context) error {
	redisConf, exists := config.GetRedisInfo("demoredis")
	if !exists {
		ctx.WriteString("no exists redis")
		return nil
	}

	redisClient := redisutil.GetRedisClient(redisConf.ServerIP, redisConf.DB, redisConf.Password)
	redisClient.Set("apitest", 1)
	ctx.WriteString(redisClient.Get("apitest"))
	return nil
}
