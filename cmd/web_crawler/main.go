package main

import (
	"github.com/yihsuanhung/web-crawler/internal/db"
	"github.com/yihsuanhung/web-crawler/internal/mock"
	"github.com/yihsuanhung/web-crawler/pkg/crawler"
	"github.com/yihsuanhung/web-crawler/pkg/handler"
	"github.com/yihsuanhung/web-crawler/pkg/server"
)

func main() {

	// init the db instance
	db.Init()

	// [DEV] init mock db instance
	mock.InitDB()

	// init the crawler job queue
	crawler.Init()

	// set up routes
	config := server.DefaultConfig()
	instance := config.Build()
	v1 := instance.Group("v1")
	v1.GET("/hello", handler.Hello)
	v1.POST("/crawl", handler.Crawl)
	v1.OPTIONS("/crawl", handler.Preflight)
	v1.POST("/status", handler.Status)
	v1.OPTIONS("/status", handler.Preflight)
	if err := instance.Serve(); err != nil {
		panic(err)
	}

}
