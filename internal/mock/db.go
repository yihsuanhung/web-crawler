package mock

import (
	"github.com/yihsuanhung/web-crawler/pkg/crawler"
)

var DB map[string]*crawler.Job

func InitDB() map[string]*crawler.Job {

	DB = make(map[string]*crawler.Job)

	return DB
}
