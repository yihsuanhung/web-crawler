package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	URL string `json:"url" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/scraping", func(c *gin.Context) {

		var request Request

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "ERROR",
			})
			return
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		url := request.URL

		c.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"url":    url,
		})
	})

	router.OPTIONS("/scraping", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.AbortWithStatus(http.StatusOK)
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
