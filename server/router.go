package server

import (
	"github.com/devfeel/dotweb"

	"github.com/devfeel/dotweb-start/server/handlers/test"
	"github.com/devfeel/dotweb-start/server/handlers/demo"
)

func InitRoute(server *dotweb.HttpServer) {
	g := server.Group("/test")
	g.GET("/index", test.Index)

	g = server.Group("/demo")
	g.GET("/queryinfo", demo.QueryDemoInfo)
	g.GET("/info/:demoid", demo.QueryDemoInfo)
	g.GET("/querylist", demo.QueryDemoList)
	g.GET("/add", demo.AddDemo)
}
