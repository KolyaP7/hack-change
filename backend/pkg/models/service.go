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

func GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    result := db.Where("email = ?", email).First(&user)
    
    if result.Error != nil {
        if gorm.IsRecordNotFoundError(result.Error) {
            return nil, errors.New("Пользователь не найден")
        }
        return nil, result.Error
    }
    
    return user.PasswordHash, user.ID, nil
}


