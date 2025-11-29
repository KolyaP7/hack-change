package models

type Project struct {
    ID        uint
    Name      string `gorm:"not null"`
    CreatedBy uint   `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

