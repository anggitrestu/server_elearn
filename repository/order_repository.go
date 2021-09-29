package repository

import (
	"server_elearn/models/orders"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Save(order orders.Order)(orders.Order , error)
	FindAllByUserID(userID int)([]orders.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func(r *orderRepository) Save(order orders.Order)(orders.Order , error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func(r *orderRepository) FindAllByUserID(userID int)([]orders.Order, error) {
	var orders []orders.Order
	err := r.db.Where("user_id = ? ", userID).Find(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}