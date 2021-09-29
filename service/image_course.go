package service

import (
	imagecourses "server_elearn/models/image_courses"
	"server_elearn/repository"
)

type ServiceImageCourse interface {
	CreateImageCourse(input imagecourses.CreateImageCourseInput)(imagecourses.ImageCourse, error)
	GetImageCourseByID(input imagecourses.GetImageCourseInput)(imagecourses.ImageCourse, error)
	DeleteImageCourse(input imagecourses.GetImageCourseInput)(imagecourses.ImageCourse, error)
}

type serviceImageCourse struct {
	repositoryImageCourse repository.ImageCourseRepository
}

func NewServiceImageCourse(repositoryImageCourse repository.ImageCourseRepository)*serviceImageCourse {
	return &serviceImageCourse{repositoryImageCourse}
}

func(s *serviceImageCourse) CreateImageCourse(input imagecourses.CreateImageCourseInput)(imagecourses.ImageCourse, error) {
	imageCourse := imagecourses.ImageCourse{}
	imageCourse.Image = input.Image
	imageCourse.CourseID = input.CourseID

	newImageCourse, err := s.repositoryImageCourse.Save(imageCourse)
	if err != nil {
		return newImageCourse, err
	}

	return newImageCourse, nil
}

func(s *serviceImageCourse) GetImageCourseByID(input imagecourses.GetImageCourseInput)(imagecourses.ImageCourse, error) {
	imageCourse, err := s.repositoryImageCourse.FindByID(input.ID)
	if err != nil {
		return imageCourse, err
	}

	return imageCourse, nil
}

func(s *serviceImageCourse) DeleteImageCourse(input imagecourses.GetImageCourseInput)(imagecourses.ImageCourse, error) {
	imagecourses,err := s.repositoryImageCourse.Delete(input.ID)
	if err != nil {
		return imagecourses, err
	}

	return imagecourses , nil
}
