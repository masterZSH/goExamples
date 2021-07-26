package services

import (
	"github.com/gin-gonic/gin"
)

func GetArea(c *gin.Context) {
	c.JSON(200, gin.H{
		"errcode": 0,
		"result":  []string{"1", "2"},
	})
}
