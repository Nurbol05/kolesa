package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kolesa/pkg/logger"
	"os"
)

// ConnectPostgres подключается к PostgreSQL и возвращает экземпляр *gorm.DB
func ConnectPostgres() (*gorm.DB, error) {
	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

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
	return db, nil
}
