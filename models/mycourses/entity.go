package mycourses

import (
	"time"

	"gorm.io/gorm"
)

type MyCourse struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	CourseID  int            `json:"course_id"`
	UserID   int             `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}