package main

import (
	"log"
	"member-service/internal/configs"
	"member-service/internal/grpc"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()

	go grpc.ListenAndServe()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	log.Println("Server is running: " + configs.App.Domain + ":" + configs.App.Port)

	router.Run(":" + configs.App.Port)
}
