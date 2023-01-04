package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yihsuanhung/web-crawler/internal/mock"
)

type StatusRequest struct {
	ID string `json:"id" binding:"required"`
}

func Status(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var request StatusRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ERROR",
		})
		return
	}

	id := request.ID

	if data, ok := mock.DB[id]; ok {
		fmt.Println(*data)

		// json.Unmarshal()

		d, err := json.Marshal(data.Result)

		if err != nil {
			panic(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   string(d),
		})

	} else {
		fmt.Println("找不到id")
		c.JSON(http.StatusOK, gin.H{
			"status": "Not Found",
		})
	}

}
