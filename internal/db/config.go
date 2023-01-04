package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DSN *DSN `json:"dsn" toml:"dsn"` // DSN 地址: mysql://root:secret@tcp(127.0.0.1:3307)/mysql?timeout=20s&readTimeout=20s
}

// ReadConfigFromEnv 返回默認 // TODO 之後可以獨立寫一個config
func DefaultConfig() *Config {
	dsn := &DSN{}
	dsn.Dbname = "crawler"
	dsn.Host = "appsDB"
	dsn.Dialect = "mysql"
	dsn.Username = "root"
	dsn.Password = "root"
	dsn.Port = "3306"
	return &Config{
		DSN: dsn,
	}
}

// Open 開啟連線
func open(config *Config) (*gorm.DB, error) {
	var inner *gorm.DB
	var err error
	inner, err = gorm.Open(mysql.Open(config.DSN.GetMysqlDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return inner, err
}

// Build ...
func (config *Config) Build() *gorm.DB {
	var err error
	db, err := open(config)
	if err != nil {
		fmt.Printf("[DB] open err: %s , addr: %s:%s , config: %v \n", err, config.DSN.Host, config.DSN.Port, config)
		return db
	}
	instance, _ := db.DB()
	if err := instance.Ping(); err != nil {
		fmt.Printf("[DB] ping err: %s , config : %v ", err, config)
	}
	fmt.Printf("[DB] connect! AddrESS: %s:%s, DB name: %s\n", config.DSN.Host, config.DSN.Port, config.DSN.Dbname)
	return db
}
