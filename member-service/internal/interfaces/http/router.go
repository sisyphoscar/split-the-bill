package http

import (
	"log"
	"member-service/internal/configs"

	"github.com/gin-gonic/gin"
)

func Listen() {
	router := gin.Default()

	router.GET("/", healthCheck)

	log.Println("HTTP Server is running: " + configs.App.Domain + ":" + configs.App.Port)

	router.Run(":" + configs.App.Port)
}
