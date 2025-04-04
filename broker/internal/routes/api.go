package routes

import (
	"broker/internal/handlers"
	"broker/internal/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Api(memberHandler *handlers.MemberHandler) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.Cors())

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": true,
		})
	})

	router.POST("/members", func(c *gin.Context) {
		memberHandler.CreateMember(c)
	})

	return router
}
