package mentors

// yang di create ke database
import (
	"time"

	"gorm.io/gorm"
)

type Mentor struct {
	ID             	int 			`gorm:"primaryKey" json:"id"`
	Name           	string			`gorm:"size:256" json:"name"`
	Profile        	string			`gorm:"size:256" json:"profile"`
	Email          	string			`gorm:"unique" json:"email"`
	Profession 		string			`gorm:"size:256" json:"profession"`
	CreatedAt      	time.Time		`json:"created_at"`
	UpdatedAt      	time.Time		`json:"updated_at"`
	DeletedAt 	   	gorm.DeletedAt	`gorm:"index" json:"deleted_at"`
}
