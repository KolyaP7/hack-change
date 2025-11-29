package models

import (
    "gorm.io/gorm"
    "time"
)

// Review структура для хранения отзывов
type Review struct {
    ID          uint
    Text        string `gorm:"unique;not null"`
    Source      string
    TeacherRate int    // Оценка преподавателя
    ModelRate   int    // Оценка модели
    ProjectID   uint   // Изменили на uint для соответствия GORM
    CreatedAt   time.Time
}
