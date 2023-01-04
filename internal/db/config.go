package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type Rdb struct {
// 	Name            string `envField:"rdb:name"`
// 	Debug           bool   `envField:"rdb:debug"`
// 	MaxIdle         int    `envField:"rdb:max_idle"`
// 	MaxOpen         int    `envField:"rdb:max_open"`
// 	ConnMaxLifetime string `envField:"rdb:conn_max_lifetime"`
// 	OnDialError     string `envField:"rdb:on_dial_error"`
// 	SlowThreshold   string `envField:"rdb:slow_threshold"`
// 	DialTimeout     string `envField:"rdb:dial_timeout"`
// 	DetailSQL       bool   `envField:"rdb:detail_sql"`
// }

type Config struct {
	Name            string
	DSN             *DSN          `json:"dsn" toml:"dsn"`                         // DSN 地址: mysql://root:secret@tcp(127.0.0.1:3307)/mysql?timeout=20s&readTimeout=20s
	Debug           bool          `json:"debug" toml:"debug"`                     // Debug
	MaxIdle         int           `json:"maxIdle" toml:"maxIdle"`                 // 最大空閑連線數量
	MaxOpen         int           `json:"maxOpen" toml:"maxOpen"`                 // 最大活動連線數量
	ConnMaxLifetime time.Duration `json:"connMaxLifetime" toml:"connMaxLifetime"` // 連線的最大存活時間
	OnDialError     string        `json:"level" toml:"level"`                     // 創建連線時的錯誤級別，=panic時，如果創建失敗，立刻 panic
	SlowThreshold   time.Duration `json:"slowThreshold" toml:"slowThreshold"`     // 慢日志
	DialTimeout     time.Duration `json:"dialTimeout" toml:"dialTimeout"`         // 播接超時
	DetailSQL       bool          `json:"detailSql" toml:"detailSql"`
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
		// Debug:           false,
		// MaxIdle:         10,
		// MaxOpen:         10,
		// ConnMaxLifetime: 100000,
		// OnDialError:     "panic",
		// SlowThreshold:   utils.Duration(config.RdbConfig.SlowThreshold),
		// DialTimeout:     utils.Duration(config.RdbConfig.DialTimeout),
	}
}

// Open 開啟連線
func open(config *Config) (*gorm.DB, error) {
	var inner *gorm.DB
	var err error
	inner, err = gorm.Open(mysql.Open(config.DSN.GetMysqlDSN()), &gorm.Config{})
	// inner, err = gorm.Open(mysql.Open(config.DSN.GetMysqlDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 連線配置
	// instance, _ := inner.DB()
	// instance.SetMaxIdleConns(config.MaxIdle)
	// instance.SetMaxOpenConns(config.MaxOpen)
	// if config.ConnMaxLifetime != 0 {
	// 	instance.SetConnMaxLifetime(config.ConnMaxLifetime)
	// }
	return inner, err
}

// Build ...
func (config *Config) Build() *gorm.DB {
	var err error
	db, err := open(config)
	if err != nil {
		if config.OnDialError == "panic" {
			fmt.Printf("[DB] open err: %s , addr: %s:%s , config: %v \n", err, config.DSN.Host, config.DSN.Port, config)
		} else {
			fmt.Printf("[DB] open err: %s , addr: %s:%s , config: %v \n", err, config.DSN.Host, config.DSN.Port, config)
			return db
		}
	}
	instance, _ := db.DB()
	if err := instance.Ping(); err != nil {
		fmt.Printf("[DB] ping err: %s , config : %v ", err, config)
	}
	fmt.Printf("[DB] connect! AddrESS: %s:%s, DB name: %s\n", config.DSN.Host, config.DSN.Port, config.DSN.Dbname)
	return db
}
