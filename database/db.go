package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kolesa/pkg/logger"
	"os"
)

var DB *gorm.DB

func ConnectPostgres() (*gorm.DB, error) {
	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),     // host
		os.Getenv("DB_USER"),     // user
		os.Getenv("DB_PASSWORD"), // password
		os.Getenv("DB_NAME"),     // dbname
		os.Getenv("DB_PORT")) // port

	// Подключаемся к базе данных через GORM и pgx
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// Проверка соединения с БД
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	logger.Log.Info("Connected to PostgreSQL")
	DB = db
	return db, nil
}

// Функция для получения переменной окружения с дефолтным значением
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
