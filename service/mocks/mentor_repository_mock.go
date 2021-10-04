package mocks

import (
	"server_elearn/models/mentors"

	"github.com/stretchr/testify/mock"
)

type MockMentorRepository struct {
	mock.Mock
}

func (_m *MockMentorRepository) Save(mentor mentors.Mentor)(mentors.Mentor, error) {
	ret := _m.Called(mentor)

	var r0 mentors.Mentor
	if rf, ok := ret.Get(0).(func(mentors.Mentor) mentors.Mentor); ok {
		r0 = rf(mentor)
	} else {
		r0 = ret.Get(0).(mentors.Mentor)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(mentors.Mentor)error); ok {
		r1 = rf(mentor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockMentorRepository) FindByID(ID int) (mentors.Mentor, error) {
	ret := _m.Called(ID)

	var r0 mentors.Mentor
	if rf, ok := ret.Get(0).(func(int) mentors.Mentor); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(mentors.Mentor)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(_m *MockMentorRepository) FindAll() ([]mentors.Mentor, error) {
	ret := _m.Called()

	var r0 []mentors.Mentor
	if rf, ok := ret.Get(0).(func() []mentors.Mentor); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]mentors.Mentor)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func()error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}

func (_m *MockMentorRepository)Update(mentor mentors.Mentor)(mentors.Mentor, error){

	ret := _m.Called(mentor)

	var r0 mentors.Mentor
	if rf, ok := ret.Get(0).(func(mentors.Mentor) mentors.Mentor); ok {
		r0 = rf(mentor)
	} else {
		r0 = ret.Get(0).(mentors.Mentor)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(mentors.Mentor)error); ok {
		r1 = rf(mentor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockMentorRepository) Delete(ID int)(mentors.Mentor, error) {
	ret := _m.Called(ID)

	var r0 mentors.Mentor
	if rf, ok := ret.Get(0).(func(int) mentors.Mentor); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(mentors.Mentor)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(_m *MockMentorRepository)CheckMentorByID(mentorID int)(mentors.Mentor, error) {
	ret := _m.Called(mentorID)

	var r0 mentors.Mentor
	if rf, ok := ret.Get(0).(func(int) mentors.Mentor); ok {
		r0 = rf(mentorID)
	} else {
		r0 = ret.Get(0).(mentors.Mentor)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(mentorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
