package server

import (
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb-start/server/handlers/index"
	"github.com/devfeel/dotweb-start/server/handlers/test"
)

func InitRoute(server *dotweb.HttpServer) {
	server.GET("/index", index.Index)

	g := server.Group("/test")
	g.GET("/index", test.Index)
	g.GET("/json", test.Json)
	g.GET("/redis", test.Redis)

	server.ServerFile("/public/*filepath", "../server/public")

}
