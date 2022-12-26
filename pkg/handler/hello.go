package handler

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(200, "Hello v1")
}
