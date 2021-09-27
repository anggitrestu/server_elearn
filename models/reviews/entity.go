package reviews

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	CourseID  int            `json:"course_id"`
	UserID    int            `json:"user_id"`
	Rating 		int 		  `json:"rating"`
	Note 		int 		  `json:"note"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}