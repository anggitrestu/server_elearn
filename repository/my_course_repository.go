package repository

import (
	"server_elearn/models/mycourses"

	"gorm.io/gorm"
)

type MyCourseRepository interface {
	FindAllByUserID(userID int)([]mycourses.MyCourse, error)
	CheckCourse(courseID int, userID int)(mycourses.MyCourse, error)
	Save(myCourse mycourses.MyCourse)(mycourses.MyCourse, error)
	// UpdateToPremium(mycou)
}

type myCourseRepository struct {
	db *gorm.DB
}

func NewMyCourseRepository(db *gorm.DB) *myCourseRepository {
	return &myCourseRepository{db}
}

func(r *myCourseRepository)	FindAllByUserID(userID int)([]mycourses.MyCourse, error) {
	var myCourse []mycourses.MyCourse
	err := r.db.Where("user_id = ?", userID).Find(&myCourse).Error
	if err != nil {
		return myCourse, err
	}

	return myCourse, nil
}

func(r *myCourseRepository) CheckCourse(courseID int, userID int)(mycourses.MyCourse, error){
	var myCourse mycourses.MyCourse
	err := r.db.Where("course_id = ? ", courseID).Where("user_id = ?", userID).Find(&myCourse).Error
	if err != nil {
		return myCourse, err
	}

	return myCourse, nil
}

func(r *myCourseRepository) Save(myCourse mycourses.MyCourse)(mycourses.MyCourse, error){
	err := r.db.Create(&myCourse).Error
	if err != nil {
		return myCourse, err
	}
	return myCourse, nil
}

// kurang find one by user



