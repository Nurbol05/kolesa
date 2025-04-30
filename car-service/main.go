package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kolesa/car-service/car"
	"kolesa/car-service/logging"
	"log"
	"os"
)

func init() {
	// .env файлын жүктеу
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// .env файлынан дерекқор параметрлерін алу
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),     // host
		os.Getenv("DB_USER"),     // user
		os.Getenv("DB_PASSWORD"), // password
		os.Getenv("DB_NAME"),     // dbname
		os.Getenv("DB_PORT"))     // port

	// PostgreSQL дерекқорына қосылу
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Журнал жүргізу middleware қосу
	r := gin.New()
	r.Use(logging.RequestLogger())

	// Репозиторий, қызмет және хендлер объектілерін құру
	carRepo := car.NewCarRepository(db)
	carService := car.NewCarService(carRepo)
	carHandler := car.NewCarHandler(carService)

	// API маршруты
	carGroup := r.Group("/api/v1/car")
	{
		carGroup.POST("/", carHandler.Create)
		carGroup.GET("/", carHandler.GetAll)
		carGroup.GET("/:id", carHandler.GetByID)
		carGroup.PUT("/:id", carHandler.Update)
		carGroup.DELETE("/:id", carHandler.Delete)
	}

	// Серверді бастау
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082" // Әдепкі порт
	}
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
