package main

import (
	"even"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/even", func(c *gin.Context) {
		fmt.Print(c)
		id := c.Query("id")
		fmt.Println(id)
		intS, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		name := even.Returnc(intS)
		c.String(http.StatusOK, "Hello %b", name)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
