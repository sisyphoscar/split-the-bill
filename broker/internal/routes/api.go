package routes

import (
	"broker/internal/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.Cors())

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": true,
		})
	})

	// TODO: gRPC to member service
	router.POST("/members", func(c *gin.Context) {
		name := c.PostForm("username")
		email := c.PostForm("email")

		log.Printf("name: %s, 電子郵件: %s", name, email)
		c.JSON(http.StatusCreated, gin.H{
			"message": "新增使用者成功, 成員: " + name + ", 電子郵件: " + email,
		})
	})

	return router
}
