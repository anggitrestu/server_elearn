package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int 				`gorm:"primaryKey" json:"id"`
	Name           string			`gorm:"size:256" json:"name"`
	Email          string			`gorm:"unique" json:"email"`
	PasswordHash   string			`gorm:"size:256" json:"password_hash"`
	Profession     string			`gorm:"size:256" json:"profession"`
	AvatarFileName string			`gorm:"size:256" json:"avatar_file_name"`
	Role           string			`gorm:"size:256" json:"role"`
	CreatedAt      time.Time		`json:"created_at"`
	UpdatedAt      time.Time		`json:"updated_at"`
	DeletedAt 	   gorm.DeletedAt	`gorm:"index" json:"deleted_at"`
}