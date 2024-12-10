package main

import (
	"database/sql"
	"log"

	"github.com/cyberkillua/go-api-template/internal/config"
	"github.com/cyberkillua/go-api-template/internal/database"
	"github.com/cyberkillua/go-api-template/internal/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}
	// Verify critical environment variables
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	// Database connection
	connection, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize database queries
	db := database.New(connection)

	// Create and start server
	srv := server.New(cfg, db)
	if err := srv.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
