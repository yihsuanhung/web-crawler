package db

import "fmt"

// DSN ...
type DSN struct {
	Dialect   string `default:"mysql"`
	Username  string
	Password  string
	Net       string // Network type
	Host      string
	Port      string
	Dbname    string            // Database name
	Params    map[string]string // Connection parameters
	Migration bool              `default:"false"`
}

func (d *DSN) GetMysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.Dbname)
}
