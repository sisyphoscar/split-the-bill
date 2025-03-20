package main

import (
	"broker/internal/configs"
	"broker/internal/routes"
	"log"
)

func main() {
	configs.LoadConfig()

	router := routes.Api()

	log.Println("Server is running: " + configs.App.Domain + ":" + configs.App.Port)

	router.Run(":" + configs.App.Port)
}
