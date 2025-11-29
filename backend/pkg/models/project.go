package models

import "time"

type Project struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedBy uint      `gorm:"not null" json:"created_by"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type ProjectStatistics struct {
	ProjectID          uint      `gorm:"foreignKey:ID" json:"project_id"` // FIXME: скорее всего неправильно написано у горма
	CreatedAt          time.Time `json:"created_at"`
	NumPositiveReviews int       `json:"num_positive"`
	NumNegativeReviews int       `json:"num_negative"`
	NumNeutralReviews  int       `json:"num_neutral"`
	NumTotalReviews    int       `json:"num_total"`
	AvgRating          float64   `json:"avg_rating"`
}
