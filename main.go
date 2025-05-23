package main

import (
	"github.com/Nurbol05/kolesa/database"
	"github.com/Nurbol05/kolesa/pkg/logger"
	"github.com/Nurbol05/kolesa/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Logger initialized")

	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatal("Error loading .env file")
	}

	db, err := database.ConnectPostgres()
	if err != nil {
		logger.Log.Fatalf("Failed to connect to the database: %v", err)
	}

	db, err = database.ConnectPostgresTest()
	if err != nil {
		logger.Log.Fatalf("Failed to connect to the database: %v", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)

	logger.Log.Info("Server running on :8080")
	r.Run(":8080")
}
