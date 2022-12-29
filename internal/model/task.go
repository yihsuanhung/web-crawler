package model

type Task struct {
	ID        uint32 `gorm:"primary_key" json:"id"`
	TimeStamp string `json:"timeStamp"`
	Url       string `json:"url"`
	Result    uint32 `json:"result"`
}
