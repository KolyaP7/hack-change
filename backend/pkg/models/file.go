package models

type File struct {
	Text   string `gorm:"not null" json:"text"`
	Rating int    `gorm:"not null" json:"rating"`
	Src    string `gorm:"not null" json:"src"`
}
