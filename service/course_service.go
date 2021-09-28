package service

import (
	"server_elearn/models/courses"
	"server_elearn/repository"
)

type ServiceCourse interface {
	CreateCourse(input courses.CreateCourseInput) (courses.Course, error)
	GetCourseByID(input courses.GetCourseInput)(courses.Course, error)
	GetCourses()([]courses.Course, error)
	UpdateCourse(inputID courses.GetCourseInput, inputData courses.CreateCourseInput)(courses.Course, error)
	DeleteCourse(inputID courses.GetCourseInput)(courses.Course, error)
	CourseIsAvaibility(coursesID int)(bool, error)
}

type serviceCourse struct {
	repositoryCourse repository.CourseRepository
}

func NewServiceCourse(repositoryCourse repository.CourseRepository) *serviceCourse {
	return &serviceCourse{repositoryCourse}
}

func(s *serviceCourse) CreateCourse(input courses.CreateCourseInput)(courses.Course, error) {
	course := courses.Course{}
	course.Name = input.Name
	course.Certificate = input.Certificate
	course.Type = input.Type
	course.Status = input.Status
	course.Level = input.Level
	course.Description = input.Description 
	course.MentorID = input.MentorID

	newCourse, err := s.repositoryCourse.Save(course)
	if err != nil {
		return newCourse, err
	}

	return newCourse, nil
}

func(s *serviceCourse)GetCourseByID(input courses.GetCourseInput)(courses.Course, error) {
	course, err := s.repositoryCourse.FindByID(input.ID)
	if err != nil {
		return course, err
	}

	return course, nil
}

func(s *serviceCourse)GetCourses()([]courses.Course, error) {
	courses , err := s.repositoryCourse.FindAll();
	if err != nil {
		return courses,err
	}

	return courses, nil
}

func (s *serviceCourse)	UpdateCourse(inputID courses.GetCourseInput, inputData courses.CreateCourseInput)(courses.Course, error) {
	course, err := s.repositoryCourse.FindByID(inputID.ID)

	if err != nil {
		return course, err
	}

	course.Name = inputData.Name
	course.Certificate = inputData.Certificate
	course.Type = inputData.Type
	course.Status = inputData.Status
	course.Level = inputData.Level
	course.Description = inputData.Description 
	course.MentorID = inputData.MentorID

	updateCourse, err := s.repositoryCourse.Update(course)
	if err != nil {
		return updateCourse, err
	}

	return updateCourse, nil
}


func(s *serviceCourse) DeleteCourse(inputID courses.GetCourseInput)(courses.Course, error) {
	course, err := s.repositoryCourse.Delete(inputID.ID)
	if err != nil  {
		return course, err
	}

	return course, nil
}


func(s *serviceCourse) CourseIsAvaibility(coursesID int)(bool, error) {
	course, err := s.repositoryCourse.FindByCourseID(coursesID)
	if err != nil {
		return false, err
	}

	if course.ID == 0 {
		return false, nil
	}

	return true, nil
}


