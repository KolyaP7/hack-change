package db

import (
	"fmt"
	. "hack-change-backend/pkg/getenv"
	"hack-change-backend/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", GetValue("DB_HOST", "localhost"), GetValue("DB_USER", "postgres"), GetValue("DB_NAME", "hakaton"), GetValue("DB_PASSWORD", "changeme"), GetValue("DB_PORT", ":5432"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
		return err
	}
	DB = db

	// Автомиграция моделей
	err = DB.AutoMigrate(&models.User{})
	return err

}
