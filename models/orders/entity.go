package orders

import (
	paymentlogs "server_elearn/models/payment_logs"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Order struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Status 	  string 		 `gorm:"size:256" json:"status"`
	CourseID  int            `json:"course_id"`
	UserID    int            `json:"user_id"`
	SnapURL	  string 		 `gorm:"size:256" json:"snap_url"`	
	Metadata  datatypes.JSON  `json:"metadata"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	PaymentLogs paymentlogs.PaymentLog `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Metadata struct {
	Course_id int	`json:"course_id"`
	Course_price int `json:"course_price"`
	Course_name string`json:"course_name"`
	Course_thumbnail string `json:"course_thumbnail"`
	Course_level string `json:"level"`
}