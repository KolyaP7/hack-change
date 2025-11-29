package user

import (
    "hack-change/backend/pkg/models"
    "gorm.io/gorm"
    "log"
)

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

func CreateUser(user *models.User) error {
    result = db.Model(user).Creates(map[string]interface{}{
        "user_name": user.UserName,
        "password_hash": user.PasswordHash,
    })
    
    if result.Error != nil {
        return result.Error
    }
    
    return nil
}



