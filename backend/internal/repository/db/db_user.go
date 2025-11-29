package db

import (
	"errors"
	"hack-change-backend/pkg/models"

	"gorm.io/gorm"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}

func CreateUser(user *models.User) error {
	result := DB.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
