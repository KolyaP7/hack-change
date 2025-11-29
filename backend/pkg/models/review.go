package models

type Review struct {
	ID           uint
	Text         string `gorm:"unique;not null"`
	Source       string
	JSONData     string
	RatingSource int `gorm: foreign-key` // FIXME: исправить, написал наобум
	TeacherRate  int
	ModelRate    int
}
