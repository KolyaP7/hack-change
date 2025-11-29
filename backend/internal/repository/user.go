package models

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
    ID        uint
    UserName  string `gorm:"not null;unique"`
    Email     string `gorm:"not null;unique"`
    Password  string `gorm:"not null"`
    Projects  []Project `gorm:"foreignKey:CreatedBy"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
