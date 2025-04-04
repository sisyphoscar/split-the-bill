package database

import (
	"database/sql"
	"log"
	"member-service/internal/configs"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", configs.Database.PostgresDSN)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Connected to database")

	return db
}

func CloseDB(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}
