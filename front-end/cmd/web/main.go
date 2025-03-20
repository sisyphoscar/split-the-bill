package main

import (
	"front-end/internal/configs"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()

	router := gin.Default()

	router.LoadHTMLGlob("resources/templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "bill.html", gin.H{
			"title":          "Split The Bill",
			"brokerEndpoint": configs.Endpoints.Broker,
		})
	})

	log.Println("Server is running: " + configs.App.Domain + ":" + configs.App.Port)

	router.Run(":" + configs.App.Port)
}
