package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	// method 1 一條龍 ⭐️
	// TODO extract request body form post context
	// TODO call pkg/crawler (body)
	// TODO error handle => response reject
	// TODO response accept

	// method 2 ⭐️⭐️⭐️⭐️⭐️️️
	// 思考，是否在這裡預處理
	// TODO extract request body form post context
	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ERROR",
		})
		return
	}

	url := request.URL

	crawler.Parse(url)

	// TODO enqueue pkg/taskQueue (body)
	// TODO error handle => response reject
	// TODO response accept

	// c.JSON(200, "Hellov1")
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"url":    url,
	})

}
