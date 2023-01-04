package crawler

import (
	"fmt"
)

func Worker(jobChan <-chan *Job) {
	for job := range jobChan {
		fmt.Println("處理任務", *job)
		data := Parse(job.Url)

		job.Result = data

	}
}

var JobChan chan *Job

func Init() {
	if JobChan != nil {
		return
	}

	JobChan = Build()

	go Worker(JobChan)
}

func CreateTask(id, url string) *Job {
	return &Job{
		Id:     id,
		Url:    url,
		Status: "ready",
		Result: "unknown",
	}
}

func Enq(job *Job) {

	if JobChan == nil {
		fmt.Println("Job Queue is not ready")
		return
	}

	fmt.Printf("任務 %s 加入隊列\n", job.Id)

	JobChan <- job
}
