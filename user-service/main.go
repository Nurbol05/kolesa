package main

import (
	"fmt"
	"github.com/Nurbol05/kolesa/user-service/logging"
	"github.com/Nurbol05/kolesa/user-service/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// API маршруты
	auth := r.Group("api/v1/user")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
		auth.GET("/", userHandler.GetAll)
		auth.PUT("/users/update", userHandler.UpdateUser)
		auth.DELETE("/users/delete", userHandler.DeleteUser)
	}

	// Серверді бастау
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Әдепкі порт
	}
	r.Run(fmt.Sprintf(":%s", port))
}
