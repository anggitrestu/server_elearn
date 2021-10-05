package service_test

import (
	"server_elearn/models/courses"
	"server_elearn/service"
	"server_elearn/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var courseRepository mocks.MockCourseRepository

var courseServiceMock service.ServiceCourse
var courseDomain courses.Course

func setup() {
	courseServiceMock = service.NewServiceCourse(&courseRepository)
	courseDomain = courses.Course {
		ID: 1,
		Name: "Course React JS",
		Certificate: true,
		Thumbnail: "images/thumbnail/course1.jpg",
		Type: "premium",
		Status: "published",
		Price: 280000,
		Level: "intermediete",
		Description: "mempelajari react js mulai dari fundamental",
		MentorID: 1,
	}
}

func TestCreateCourse__Success(t *testing.T){
	setup()
	courseRepository.On("Save", mock.Anything).Return(courseDomain, nil).Once()
	course, err := courseServiceMock.CreateCourse(courses.CreateCourseInput{
		Name: "Course React JS",
		Certificate: false,
		Type:"premium",
		Status: "published",
		Level: "intermediete",
		Description: "mempelajari react js mulai dari fundamental",
		MentorID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, course)
	assert.Equal(t, course, courseDomain)
}


func TestGetCourseByID__Success(t *testing.T){
	setup()
	courseRepository.On("FindByID", mock.AnythingOfType("int")).Return(courseDomain, nil).Once()
	course, err := courseServiceMock.GetCourseByID(courses.GetCourseInput{
		ID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, course)
	assert.Equal(t, course.ID, 1)
}

func TestGetCourses__Success(t *testing.T){
	setup()
	var courses []courses.Course
	courses = append(courses, courseDomain)
	courses = append(courses, courseDomain)
	courseRepository.On("FindAll").Return(courses, nil).Once()
	allCourse, err := courseServiceMock.GetCourses()
	assert.Nil(t, err)
	assert.Equal(t, courses, allCourse)
}

func TestUpdateCourse_Success(t *testing.T){
	setup()
	courseRepository.On("FindByID", mock.AnythingOfType("int")).Return(courseDomain, nil).Once()
	courseRepository.On("Update", mock.Anything).Return(courseDomain, nil).Once()
	course, err := courseServiceMock.UpdateCourse(courses.GetCourseInput{
		ID: 1,
	},courses.CreateCourseInput{
		Name: "Course React JS",
		Certificate: false,
		Type:"premium",
		Status: "published",
		Level: "intermediete",
		Description: "mempelajari react js mulai dari fundamental",
		MentorID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, course)
	assert.Equal(t, course, courseDomain)
}

func TestDeleteCourse__Succeess(t *testing.T){
	setup()
	courseRepository.On("Delete", mock.AnythingOfType("int")).Return(courseDomain, nil).Once()
	_, err := courseServiceMock.DeleteCourse(courses.GetCourseInput{
		ID: 1,
	})
	assert.Nil(t, err)
}

func TestCourseIsAvaibility__Success(t *testing.T){
	setup()
	courseRepository.On("FindByCourseID", mock.AnythingOfType("int")).Return(courseDomain, nil).Once()
	isFound, err := courseServiceMock.CourseIsAvaibility(1)
	assert.Nil(t, err)
	assert.Equal(t, true, isFound)
}

func TestGetCourseByCourseID__Success(t *testing.T){
	setup()
	courseRepository.On("FindByCourseID", mock.AnythingOfType("int")).Return(courseDomain, nil).Once()
	course, err := courseServiceMock.GetCourseByCourseID(1)
	assert.Nil(t, err)
	assert.Equal(t, course, courseDomain)
}

