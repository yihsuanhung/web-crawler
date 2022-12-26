package main

import (
	"github.com/yihsuanhung/web-crawler/pkg/handler"
	"github.com/yihsuanhung/web-crawler/pkg/server"
)

func main() {
	config := server.DefaultConfig()
	instance := config.Build()
	v1 := instance.Group("v1")
	v1.GET("/hello", handler.Hello)
	v1.POST("/crawl", handler.Crawl)
	if err := instance.Serve(); err != nil {
		panic(err)
	}
}
