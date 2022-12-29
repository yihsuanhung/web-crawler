package main

import (
	"fmt"

	"github.com/yihsuanhung/web-crawler/pkg/handler"
	"github.com/yihsuanhung/web-crawler/pkg/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/crawler"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("db is connected")

	config := server.DefaultConfig()
	instance := config.Build()
	v1 := instance.Group("v1")
	v1.GET("/hello", handler.Hello)
	v1.POST("/crawl", handler.Crawl)
	v1.OPTIONS("/crawl", handler.Preflight)
	if err := instance.Serve(); err != nil {
		panic(err)
	}

	// 建立 gorm.DB 物件

	// db, err := database.ConnectDB()
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.DB().Close()
}
