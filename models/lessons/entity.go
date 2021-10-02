package lessons

import (
	"time"

	"gorm.io/gorm"
)

type Lesson struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:256" json:"name"`
	Video     string         `gorm:"size:256" json:"video"`
	ChapterID int            `json:"chapter_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

