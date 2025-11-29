package models

type User struct {
	ID           uint
	UserName     string `gorm:"not null;unique"`
	Email        string `gorm:"not null;unique"`
	PasswordHash string `gorm:"not null"`
}
