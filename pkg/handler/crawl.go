package handler

import "github.com/gin-gonic/gin"

// producer

func Crawl(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	c.JSON(200, "Hellov1")

	// method 1 一條龍 ⭐️
	// TODO extract request body form post context
	// TODO call pkg/crawler (body)
	// TODO error handle => response reject
	// TODO response accept

	// method 2 ⭐️⭐️⭐️⭐️⭐️️️
	// 思考，是否在這裡預處理
	// TODO extract request body form post context
	// TODO enqueue pkg/taskQueue (body)
	// TODO error handle => response reject
	// TODO response accept

}
