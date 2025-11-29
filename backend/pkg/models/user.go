package models

import (
	"time"
)

type User struct {
	ID        uint
	UserName  string    `gorm:"not null;unique"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
