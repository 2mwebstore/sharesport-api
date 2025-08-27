package models

import "time"

type Highlights struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Image     string    `json:"image"`
	Content   string    `json:"content"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
