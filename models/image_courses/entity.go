package imagecourses

import (
	"time"

	"gorm.io/gorm"
)

type ImageCourse struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	CourseID  int            `json:"course_id"`
	Image     string         `gorm:"size:256" json:"image"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
