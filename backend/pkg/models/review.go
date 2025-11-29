package models

import (
    "gorm.io/gorm"
    "time"
)

// Review структура для хранения отзывов
type Review struct {
    ID         uint
    Text       string `gorm:"unique;not null"`
    Source     string
    JSONData   string
    ProjectName string
    CreatedAt  time.Time
    UpdatedAt  time.Time
    RatingSource int `gorm:"foreignKey"`
    TeacherRate int
    ModelRate  int
}
