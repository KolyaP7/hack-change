package models

type User struct {
    ID        uint
    UserName  string `gorm:"not null;unique"`
    Email     string `gorm:"not null;unique"`
    Password  string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Project struct {
    ID        uint
    Name      string `gorm:"not null"`
    CreatedBy uint
    // отношения
}

type Review struct {
    ID        uint
    Text      string `gorm:"unique;not null"`
    Rating    int
    JSONData  string
    // остальные поля
}

