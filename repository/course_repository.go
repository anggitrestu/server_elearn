package repository

import (
	"server_elearn/models/courses"

	"gorm.io/gorm"
)

type CourseRepository interface {
	Save(course courses.Course)(courses.Course, error)
	FindByID(ID int)(courses.Course, error)
	FindAll()([]courses.Course, error)
	Update(course courses.Course)(courses.Course, error)
	Delete(ID int)(courses.Course, error)

}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *courseRepository {
	return &courseRepository{db}
}

func(r *courseRepository) Save(course courses.Course)(courses.Course, error){
	
	err := r.db.Create(&course).Error
	if err != nil {
		return course, err
	}

	return course, nil

}

func(r *courseRepository) FindByID(ID int)(courses.Course, error) {
	var course courses.Course
	err := r.db.Where("id = ?", ID).Find(&course).Error
	if err != nil  {
		return course, err
	}
	return course, nil
}


func(r *courseRepository) Update(course courses.Course)(courses.Course, error) {
	err := r.db.Save(&course).Error
	if err != nil {
		return course, err
	}

	return course, nil
}

func(r *courseRepository) FindAll()([]courses.Course, error){
	var courses []courses.Course
	err := r.db.Find(&courses).Error
	if err != nil {
		return courses, err
	}

	return courses, err
}

func (r *courseRepository) Delete(ID int)(courses.Course, error) {
	course := courses.Course{}
	err := r.db.Where("id = ?", ID).Delete(&course).Error
	if err != nil {
		return course, err
	}

	return course, nil
}
