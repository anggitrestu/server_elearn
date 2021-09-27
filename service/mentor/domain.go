package mentor

import (
	"server_elearn/models/courses"
	"time"

	"gorm.io/gorm"
)

type Mentor struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	Profile    string           `json:"profile"`
	Email      string           `json:"email"`
	Profession string           `json:"profession"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
	DeletedAt  gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
	Courses    []courses.Course `gorm:"foreignKey:MentorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
