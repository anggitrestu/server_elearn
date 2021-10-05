package mocks

import (
	"server_elearn/models/mycourses"

	"github.com/stretchr/testify/mock"
)

type MockMyCourseRepository struct {
	mock.Mock
}

func(_m *MockMyCourseRepository) FindAllByUserID(userID int)([]mycourses.MyCourse, error) {
	ret := _m.Called()

	var r0 []mycourses.MyCourse
	if rf, ok := ret.Get(0).(func() []mycourses.MyCourse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]mycourses.MyCourse)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func()error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(_m *MockMyCourseRepository) CheckCourse(courseID int, userID int)(mycourses.MyCourse, error){
	ret := _m.Called(courseID, userID )

	var r0 mycourses.MyCourse
	if rf, ok := ret.Get(0).(func(int, int) mycourses.MyCourse); ok {
		r0 = rf(courseID, userID )
	} else {
		r0 = ret.Get(0).(mycourses.MyCourse)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int, int)error); ok {
		r1 = rf(courseID, userID )
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(_m *MockMyCourseRepository) Save(myCourse mycourses.MyCourse)(mycourses.MyCourse, error){
	ret := _m.Called(myCourse)

	var r0 mycourses.MyCourse
	if rf, ok := ret.Get(0).(func(mycourses.MyCourse) mycourses.MyCourse); ok {
		r0 = rf(myCourse)
	} else {
		r0 = ret.Get(0).(mycourses.MyCourse)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(mycourses.MyCourse)error); ok {
		r1 = rf(myCourse)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}