package configs

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig load all env config
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file not found")
		return
	}

	LoadAppConfig()
	LoadDatabaseConfig()

	log.Println("Config loaded successfully")
}
