package test

import (
	"github.com/devfeel/dotweb-start/const"
	"github.com/devfeel/dotweb"
)

func Index(ctx dotweb.Context) error {
	ctx.ViewData().Set("version", _const.Global_Version)
	ctx.Response().SetHeader("testh", "testh-v")
	return ctx.WriteString("welcome to ", _const.Global_ProjectName)
}
