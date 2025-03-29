package configs

import (
	"os"
)

type DatabaseConfig struct {
	PostgresDSN string
}

var Database DatabaseConfig

// LoadDatabaseConfig load database config
func LoadDatabaseConfig() {
	Database = DatabaseConfig{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
	}
}
