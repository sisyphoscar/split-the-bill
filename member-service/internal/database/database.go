package database

import (
	"database/sql"
	"log"
	"member-service/internal/configs"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", configs.Database.PostgresDSN)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

func CloseDB(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}
