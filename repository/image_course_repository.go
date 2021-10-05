package repository

import (
	imagecourses "server_elearn/models/image_courses"

	"gorm.io/gorm"
)

type ImageCourseRepository interface {
	FindByID(ID int)(imagecourses.ImageCourse, error)
	Delete(ID int)(imagecourses.ImageCourse, error)
	Save(imageCourse imagecourses.ImageCourse)(imagecourses.ImageCourse, error)
}

type imageCourseRepository struct {
	db *gorm.DB
}

func NewImageCourseRepository(db *gorm.DB) *imageCourseRepository {
	return &imageCourseRepository{db}
}

func(r *imageCourseRepository) Save(imageCourse imagecourses.ImageCourse)(imagecourses.ImageCourse, error) {
	err := r.db.Create(&imageCourse).Error
	if err != nil {
		return imageCourse, err
	}

	return imageCourse, nil
}

func(r *imageCourseRepository) FindByID(ID int)(imagecourses.ImageCourse, error) {
	var imageCourse imagecourses.ImageCourse
	err := r.db.Where("id = ? ", ID).Find(&imageCourse).Error
	if err != nil {
		return imageCourse, err
	}

	return imageCourse, nil
}

func(r *imageCourseRepository) Delete(ID int)(imagecourses.ImageCourse, error) {
	imageCourse := imagecourses.ImageCourse{}
	err := r.db.Where("id = ? ", ID).Delete(&imageCourse).Error
	if err != nil {
		return imageCourse, err
	} 

	return imageCourse, nil
}