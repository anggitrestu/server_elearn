package chapters

import (
	"server_elearn/models/lessons"
	"time"

	"gorm.io/gorm"
)

type Chapter struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:256" json:"name"`
	CourseID  int            `json:"course_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Lessons   []lessons.Lesson `gorm:"foreignKey:ChapterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}


