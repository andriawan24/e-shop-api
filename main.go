package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping success",
		})
	})

	router.Run()
}
