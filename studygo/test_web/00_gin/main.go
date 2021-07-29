package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")

	})

	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":8089")
	if err != nil {
		return
	}
}
