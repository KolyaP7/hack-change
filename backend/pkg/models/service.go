package user

import (
    "hack-change/backend/pkg/models"
    "gorm.io/gorm"
    "log"
)

func CreateUser(user *models.User) error {
    result := db.Create(user)
    if result.Error != nil {
        log.Printf("Ошибка создания пользователя: %v", result.Error)
        return result.Error
    }
    return nil
}

