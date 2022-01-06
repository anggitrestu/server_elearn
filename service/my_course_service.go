package service

import (
	"server_elearn/models/mycourses"
	"server_elearn/repository"
)

type ServiceMyCourse interface {
	GetAllMyCourse(userID int) ([]mycourses.MyCourse, error)
	CreateMyCourse(input mycourses.CreateMyCourseInput, userID int) (mycourses.MyCourse, error)
	IsExistMyCourse(input mycourses.CreateMyCourseInput, userID int) (mycourses.MyCourse, error)
	CreatePremiumAccess(userID int, courseID int) error
}

type serviceMyCourse struct {
	repositoryMyCourse repository.MyCourseRepository
}

func NewServiceMyCourse(repositoryMyCourse repository.MyCourseRepository) *serviceMyCourse {
	return &serviceMyCourse{repositoryMyCourse}
}

func (s *serviceMyCourse) CreateMyCourse(input mycourses.CreateMyCourseInput, userID int) (mycourses.MyCourse, error) {
	mycourse := mycourses.MyCourse{}
	mycourse.UserID = userID
	mycourse.CourseID = input.CourseID

	newMyCourse, err := s.repositoryMyCourse.Save(mycourse)
	if err != nil {
		return newMyCourse, err
	}

	return newMyCourse, nil

}

func (s *serviceMyCourse) GetAllMyCourse(userID int) ([]mycourses.MyCourse, error) {
	mycourse, err := s.repositoryMyCourse.FindAllByUserID(userID)
	if err != nil {
		return mycourse, err
	}
	return mycourse, nil
}

func (s *serviceMyCourse) IsExistMyCourse(input mycourses.CreateMyCourseInput, userID int) (mycourses.MyCourse, error) {
	mycourse, err := s.repositoryMyCourse.CheckCourse(input.CourseID, userID)
	if err != nil {
		return mycourse, err
	}
	return mycourse, nil
}

func (s *serviceMyCourse) CreatePremiumAccess(userID int, courseID int) error {
	mycourse := mycourses.MyCourse{}
	mycourse.UserID = userID
	mycourse.CourseID = courseID

	_, err := s.repositoryMyCourse.Save(mycourse)
	if err != nil {
		return err
	}

	return nil

}
