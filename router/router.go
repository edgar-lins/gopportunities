package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inialize a Gin router
	router := gin.Default()
	// Define a route and handler
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Start the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
