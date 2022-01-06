package repository

import (
	"server_elearn/models/users"

	"gorm.io/gorm"
)

// R besar berarti bersifat public
// interface nanti akan mengacunya ke userRepository
type UserRepository interface {
	Save(user users.User) (users.User, error)
	FindByEmail(email string) (users.User, error)
	FindById(ID int) (users.User, error)
	Update(user users.User) (users.User, error)
}

// r kecil berarti bersifat private
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user users.User) (users.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *userRepository) FindByEmail(email string) (users.User, error) {
	var user users.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindById(ID int) (users.User, error) {
	var user users.User

	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *userRepository) Update(user users.User) (users.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *userRepository) FindAll() ([]users.User, error) {
	var users []users.User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil

}
