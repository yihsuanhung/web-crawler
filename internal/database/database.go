package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	fmt.Println("Connecting db...")

	var dbName = "crawler"

	dsn := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName)

	fmt.Println("dsn", dsn)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
