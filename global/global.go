package global

import (
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotlog"
	"github.com/devfeel/dotweb-start/const"
	"errors"
	"github.com/devfeel/dotweb-start/core"
	"github.com/devfeel/dottask"
)

//全局map
var GlobalItemMap *core.CMap
var DotApp *dotweb.DotWeb
var DotTask *task.TaskService
var InnerLogger dotlog.Logger

func Init(configPath string) error{
	GlobalItemMap = core.NewCMap()
	err := dotlog.StartLogService(configPath + "/dotlog.conf")
	if err != nil {
		return errors.New("log service start error => " + err.Error())
	}
	InnerLogger = dotlog.GetLogger(_const.LoggerName_Inner)
	return nil
}