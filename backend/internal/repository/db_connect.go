package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=postgres dbname=analyzer password=yourpass port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Ошибка подключения к БД:", err)
    }
    DB = db
    
    // Автомиграция моделей
    DB.AutoMigrate(&User{})
}
