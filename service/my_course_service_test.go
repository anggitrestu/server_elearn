package service_test

import (
	"server_elearn/models/courses"
	"server_elearn/models/mycourses"
	"server_elearn/service"
	"server_elearn/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mycourseRepository mocks.MockMyCourseRepository
var mycourseService service.ServiceMyCourse
var mycourseDomain mycourses.MyCourse

func setupMycourse(){
	mycourseService = service.NewServiceMyCourse(&mycourseRepository)
	mycourseDomain = mycourses.MyCourse{
		ID: 1,
		CourseID: 1,
		UserID: 2,
		Course: courses.Course{
			ID: 1,
			Name: "React JS",
			Certificate: true,
			Thumbnail: "images/thumbnail1.jpg",
			Type: "free",
			Status: "published",
			Price: 230000,
			Level: "intermediete",
			Description: "kelas bagus",
			MentorID: 1,
		},
	}
}

func TestCreateMyCourse__Success(t *testing.T){
	setupMycourse()
	mycourseRepository.On("Save", mock.Anything).Return(mycourseDomain, nil).Once()
	mycourse, err := mycourseService.CreateMyCourse(mycourses.CreateMyCourseInput{
		CourseID: 1,
	}, 1)
	assert.Nil(t, err)
	assert.NotNil(t, mycourse)
	assert.Equal(t, mycourse, mycourseDomain)
}

func TestGetAllMyCourse__Success(t *testing.T){
	setupMycourse()
	var  mycourses []mycourses.MyCourse
	mycourses = append(mycourses, mycourseDomain)
	mycourseRepository.On("FindAllByUserID", mock.AnythingOfType("string")).Return(mycourses, nil).Once()
	allMyCourses , err := mycourseService.GetAllMyCourse(1)
	assert.Nil(t, err)
	assert.NotNil(t, allMyCourses)
	assert.Equal(t, mycourses, allMyCourses)
}

func TestIsExistMyCourse__Success(t *testing.T){
	setupMycourse()
	mycourseRepository.On("CheckCourse", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mycourseDomain, nil).Once()
	mycourse, err := mycourseService.IsExistMyCourse(mycourses.CreateMyCourseInput{
		CourseID: 1,
	}, 1)
	assert.Nil(t, err)
	assert.NotNil(t, mycourse)
	assert.Equal(t, mycourse, mycourseDomain)
}

func TestCreatePremiumAccess__Success(t *testing.T){
	setupMycourse()
	mycourseRepository.On("Save", mock.Anything).Return(mycourseDomain, nil).Once()
	err := mycourseService.CreatePremiumAccess(1, 1)
	assert.Nil(t, err)
}