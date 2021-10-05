package service_test

import (
	imagecourses "server_elearn/models/image_courses"
	"server_elearn/service"
	"server_elearn/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repositoryImageCourse  mocks.MockImageCourseRepository
var serviceImageCourse service.ServiceImageCourse
var imageCourseDomain imagecourses.ImageCourse

func setupImageCourse(){
	serviceImageCourse = service.NewServiceImageCourse(&repositoryImageCourse)
	imageCourseDomain = imagecourses.ImageCourse {
		ID: 1,
		CourseID: 1,
		Image: "imgaes/image1.jpg",
	}
}

func TestCreateImageCourse__Success(t *testing.T){
	setupImageCourse()
	repositoryImageCourse.On("Save", mock.Anything).Return(imageCourseDomain, nil).Once()
	imageCourse, err := serviceImageCourse.CreateImageCourse(imagecourses.CreateImageCourseInput{
		CourseID: 1,
		Image: "imgaes/image1.jpg",
	})
	assert.Nil(t, err)
	assert.NotNil(t, imageCourse)
	assert.Equal(t, imageCourse, imageCourseDomain)
}

func TestGetImageCourseByID__Success(t *testing.T){
	setupImageCourse()
	repositoryImageCourse.On("FindByID", mock.AnythingOfType("int")).Return(imageCourseDomain, nil).Once()
	imageCourse, err := serviceImageCourse.GetImageCourseByID(imagecourses.GetImageCourseInput{
		ID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, imageCourse)
	assert.Equal(t, imageCourse, imageCourseDomain)
}

func TestDeleteImageCourse__Success(t *testing.T){
	setupImageCourse()
	repositoryImageCourse.On("Delete", mock.AnythingOfType("int")).Return(imageCourseDomain, nil).Once()
	imageCourse, err := serviceImageCourse.DeleteImageCourse(imagecourses.GetImageCourseInput{
		ID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, imageCourse)
	assert.Equal(t, imageCourse, imageCourseDomain)
}