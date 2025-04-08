package main

import (
	"database/sql"
	"kolesa/auth"
	"kolesa/database"
	"kolesa/pkg/logger"
	"kolesa/routes"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Logger initialized")

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatal("Error loading .env file")
	}

	// Connect to the database
	db, err := database.ConnectPostgres()
	if err != nil {
		logger.Log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize services
	authService := setupAuthService(db)

	// Setup the router with defined routes
	router := routes.SetupRoutes(authService)

	// Start the HTTP server
	logger.Log.Info("Server running on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Log.Fatalf("Server failed: %v", err)
	}
}

// Initialize AuthService (assuming it takes a db connection)
func setupAuthService(db *sql.DB) *auth.AuthService {
	userRepo := auth.NewUserRepository(db)
	authService := auth.NewAuthService(userRepo)
	return authService
}
