package server

import (
	"github.com/devfeel/dotweb-start/server/handlers/test"
	"github.com/devfeel/dotweb"
)

func InitRoute(server *dotweb.HttpServer) {
	g := server.Group("/test")
	g.GET("/index", test.Index)
	g.GET("/json", test.Json)
	g.GET("/redis", test.Redis)

}
