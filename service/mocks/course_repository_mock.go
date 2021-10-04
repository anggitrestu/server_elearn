package mocks

import (
	"server_elearn/models/courses"

	"github.com/stretchr/testify/mock"
)

type MockCourseRepository struct {
	mock.Mock
}

func(_m *MockCourseRepository) Save(course courses.Course)(courses.Course, error){
	
	ret := _m.Called(course)

	var r0 courses.Course
	if rf, ok := ret.Get(0).(func(courses.Course) courses.Course); ok {
		r0 = rf(course)
	} else {
		r0 = ret.Get(0).(courses.Course)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(courses.Course)error); ok {
		r1 = rf(course)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}

func(_m *MockCourseRepository) FindByID(ID int)(courses.Course, error) {
	ret := _m.Called(ID)

	var r0 courses.Course
	if rf, ok := ret.Get(0).(func(int) courses.Course); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(courses.Course)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}


func(_m *MockCourseRepository) Update(course courses.Course)(courses.Course, error) {
	ret := _m.Called(course)

	var r0 courses.Course
	if rf, ok := ret.Get(0).(func(courses.Course) courses.Course); ok {
		r0 = rf(course)
	} else {
		r0 = ret.Get(0).(courses.Course)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(courses.Course)error); ok {
		r1 = rf(course)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(_m *MockCourseRepository) FindAll()([]courses.Course, error){
	ret := _m.Called()

	var r0 []courses.Course
	if rf, ok := ret.Get(0).(func() []courses.Course); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]courses.Course)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func()error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockCourseRepository) Delete(ID int)(courses.Course, error) {
	ret := _m.Called(ID)

	var r0 courses.Course
	if rf, ok := ret.Get(0).(func(int) courses.Course); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(courses.Course)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockCourseRepository) FindByCourseID(ID int)(courses.Course, error) {

	ret := _m.Called(ID)

	var r0 courses.Course
	if rf, ok := ret.Get(0).(func(int) courses.Course); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(courses.Course)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}