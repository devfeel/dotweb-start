package global

import (
	"github.com/devfeel/dotweb-start/core"
	"github.com/devfeel/dotweb"
)

//全局map
var GlobalContext *core.CMap
var DotApp *dotweb.DotWeb

func init() {
	GlobalContext = core.NewCMap()
}
