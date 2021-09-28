package repository

import (
	"server_elearn/models/reviews"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	Save(review reviews.Review)(reviews.Review, error) 
	Update(review reviews.Review)(reviews.Review, error)
	FindByID(ID int)(reviews.Review, error) 
	Delete(ID int)(bool, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *reviewRepository {
	return &reviewRepository{db}
}

func(r *reviewRepository) Save(review reviews.Review)(reviews.Review, error) {
	err := r.db.Create(&review).Error
	if err != nil {
		return review, err
	}

	return review, nil
}


func(r *reviewRepository)Update(review reviews.Review)(reviews.Review, error){
	err := r.db.Save(&review).Error
	if err != nil {
		return review, err
	}

	return review, nil
}

func(r *reviewRepository) FindByID(ID int)(reviews.Review, error) {
	var review reviews.Review
	err := r.db.Where("id = ? ", ID).Find(&review).Error
	if err != nil {
		return review, err
	}

	return review, nil
}

func(r *reviewRepository) Delete(ID int)(bool, error) {
	review := reviews.Review{}
	err := r.db.Where("id = ?", ID).Delete(&review).Error
	if err != nil {
		return false, err
	}
	return true, nil
}