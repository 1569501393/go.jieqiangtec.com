package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
