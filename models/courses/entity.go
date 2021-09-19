package courses

import (
	"server_elearn/models/chapters"
	"time"

	"gorm.io/gorm"
)


type Course struct {
	ID          int            `gorm:"primaryKey"`
	Name        string         `gorm:"size:256" json:"name"`
	Certificate bool           `json:"certificate"`
	Thumbnail   string         `gorm:"size:256" json:"thumbnail"`
	Type        string    	   `json:"type" gorm:"type:enum('free', 'premium')"`
	Status      string		   `json:"status" gorm:"type:enum('draft', 'published')"`
	Price       int            `json:"price"`
	Level       string         `json:"level" gorm:"type:enum('all-level', 'beginner', 'intermediete', 'advance')"`
	Description string         `json:"description"`
	MentorID    int            `json:"mentor_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Chapter 	[]chapters.Chapter `gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

