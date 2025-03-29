package main

import (
	"log"
	"member-service/internal/configs"
	"member-service/internal/database"
	"member-service/internal/grpc"
	"member-service/internal/repositories"
	"member-service/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()

	db := database.NewDB()
	defer database.CloseDB(db)

	repo := repositories.NewMemberRepository(db)
	service := services.NewMemberService(repo)

	go grpc.ListenAndServe(service)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	log.Println("Server is running: " + configs.App.Domain + ":" + configs.App.Port)

	router.Run(":" + configs.App.Port)
}
