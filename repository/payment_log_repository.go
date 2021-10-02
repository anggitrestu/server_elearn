package repository

import "gorm.io/gorm"

type PaymentLogRepository interface {
	Save(data interface{}) (error)
}

type paymentLogRepository struct {
	db *gorm.DB
}

func NewPaymentLogRepository(db *gorm.DB) *paymentLogRepository {
	return &paymentLogRepository{db}
}

func(r *paymentLogRepository)Save(data interface{}) (error){
	err := r.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}