package crawler

type Job struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Status string `json:"status"`
	Result string `json:"result"`
}

func Build() chan *Job {
	ch := make(chan *Job, 1024)
	return ch
}
