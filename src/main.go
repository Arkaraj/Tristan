package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)

func api(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is a simple Portfolio REST API",
	})
}

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	server.GET("/api", api)

	port := ":8080"

	server.Run(port) // Runs on port http://localhost:8080 (by default)
}
