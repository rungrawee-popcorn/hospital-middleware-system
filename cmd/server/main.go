package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rungrawee-popcorn/hospital-middleware-system/configs"
)

func main() {

	// Load environment variables
	configs.LoadConfig()


	// Connect database
	configs.ConnectDatabase()


	// Create database tables
	configs.AutoMigrate()


	router := gin.Default()


	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hospital Middleware System API is running",
		})

	})


	router.Run(":8080")
}