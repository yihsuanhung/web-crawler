package db

import (
	"gorm.io/gorm"
)

var SqlInstance *gorm.DB

func Init() {
	// single instance
	if SqlInstance != nil {
		return
	}
	cfg := DefaultConfig()
	SqlInstance = cfg.Build()

	// testing

	// dsn := "root:root@tcp(appsDB:3306)/crawler?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:root@tcp(127.0.0.1:3306)/crawler"
	// _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	//  fmt.Println(err)
	// 	fmt.Println("簡易連線：失敗")
	// } else {
	// 	fmt.Println("簡易連線：成功")
	// }

}
