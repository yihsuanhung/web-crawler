package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/yihsuanhung/web-crawler/internal/mock"
	"github.com/yihsuanhung/web-crawler/pkg/crawler"
)

type Request struct {
	URL string `json:"url" binding:"required"`
}

// producer

func Crawl(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ERROR",
		})
		return
	}

	// [Preprocess] create id
	taskId, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 10)
	if err != nil {
		panic(err.Error())
	}

	// [Preprocess] extract url
	url := request.URL

	// [Preprocess] prepare job
	job := crawler.CreateTask(taskId, url)

	// [Preprocess] add job to db
	mock.DB[taskId] = job
	fmt.Println("新增任務", *job)

	// TODO error handle => response reject

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"url":    url,
		"taskId": taskId,
	})

	// [Preprocess] enqueue
	crawler.Enq(job)
}
