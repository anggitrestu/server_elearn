package mocks

import (
	imagecourses "server_elearn/models/image_courses"

	"github.com/stretchr/testify/mock"
)
type MockImageCourseRepository struct {
	mock.Mock
}

func (_m *MockImageCourseRepository) Save(imageCourse imagecourses.ImageCourse) (imagecourses.ImageCourse, error) {
	ret := _m.Called(imageCourse)

	var r0 imagecourses.ImageCourse
	if rf, ok := ret.Get(0).(func(imagecourses.ImageCourse) imagecourses.ImageCourse); ok {
		r0 = rf(imageCourse)
	} else {
		r0 = ret.Get(0).(imagecourses.ImageCourse)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(imagecourses.ImageCourse)error); ok {
		r1 = rf(imageCourse)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockImageCourseRepository) FindByID(ID int) (imagecourses.ImageCourse, error) {
		ret := _m.Called(ID)

	var r0 imagecourses.ImageCourse
	if rf, ok := ret.Get(0).(func(int) imagecourses.ImageCourse); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(imagecourses.ImageCourse)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockImageCourseRepository)Delete(ID int)(imagecourses.ImageCourse, error) {
	ret := _m.Called(ID)

	var r0 imagecourses.ImageCourse
	if rf, ok := ret.Get(0).(func(int) imagecourses.ImageCourse); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(imagecourses.ImageCourse)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

