package users

import (
	"server_elearn/models/mycourses"
	"server_elearn/models/orders"
	"server_elearn/models/reviews"
	"time"

	"gorm.io/gorm"
)

type DomainUser struct {
	ID             int                  `gorm:"primaryKey" json:"id"`
	Name           string               `gorm:"size:256" json:"name"`
	Email          string               `gorm:"unique" json:"email"`
	Password       string               `gorm:"size:256" json:"password"`
	Profession     string               `gorm:"size:256" json:"profession"`
	AvatarFileName string               `gorm:"size:256" json:"avatar_file_name"`
	Role           string               `gorm:"size:256;type:enum('student', 'admin')" json:"role"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
	DeletedAt      gorm.DeletedAt       `gorm:"index" json:"deleted_at"`
	MyCourses      []mycourses.MyCourse `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Reviews        []reviews.Review     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Orders         []orders.Order       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}