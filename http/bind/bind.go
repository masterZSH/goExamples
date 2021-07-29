package main

import (
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	ID   uint64 `uri:"id" binding:"required"`
	Data struct {
		Name string `json:"name" binding:"required"`
		Slug string `json:"slug" binding:"required"`
	}
}

func main() {
	r := gin.Default()
	r.GET("update/:id", func(c *gin.Context) {
		var params UpdateRequest
		c.ShouldBindUri(&params)
		c.ShouldBindJSON(&params.Data)

		c.JSON(200, params)
	})
	r.Run(":8080")
}
