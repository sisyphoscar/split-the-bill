package main

import (
	"front-end/internal/configs"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	log.Println("Server is running: " + configs.App.Domain + ":" + configs.App.Port)

	router.Run(":" + configs.App.Port)
}
