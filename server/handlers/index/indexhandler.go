package index

import (
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb-start/const"
)

func Index(ctx dotweb.Context) error {
	ctx.ViewData().Set("version", _const.Global_Version)
	err := ctx.View("index/index.html")
	return err
}
