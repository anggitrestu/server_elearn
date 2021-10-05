package mocks

import (
	"server_elearn/models/lessons"

	"github.com/stretchr/testify/mock"
)

type MockLessonRepository struct {
	mock.Mock
}

func (_m *MockLessonRepository) Save(lesson lessons.Lesson) (lessons.Lesson, error) {
	ret := _m.Called(lesson)

	var r0 lessons.Lesson
	if rf, ok := ret.Get(0).(func(lessons.Lesson) lessons.Lesson); ok {
		r0 = rf(lesson)
	} else {
		r0 = ret.Get(0).(lessons.Lesson)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(lessons.Lesson)error); ok {
		r1 = rf(lesson)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockLessonRepository) FindByID(ID int) (lessons.Lesson, error) {
	ret := _m.Called(ID)

	var r0 lessons.Lesson
	if rf, ok := ret.Get(0).(func(int) lessons.Lesson); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(lessons.Lesson)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockLessonRepository) FindAll() ([]lessons.Lesson, error) {
	ret := _m.Called()

	var r0 []lessons.Lesson
	if rf, ok := ret.Get(0).(func() []lessons.Lesson); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]lessons.Lesson)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func()error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockLessonRepository) FindByChapterID(chapterID int) ([]lessons.Lesson, error) {
	ret := _m.Called(chapterID)

	var r0 []lessons.Lesson
	if rf, ok := ret.Get(0).(func(int) []lessons.Lesson); ok {
		r0 = rf(chapterID)
	} else {
		r0 = ret.Get(0).([]lessons.Lesson)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(chapterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockLessonRepository) Update(lesson lessons.Lesson) (lessons.Lesson, error) {
	ret := _m.Called(lesson)

	var r0 lessons.Lesson
	if rf, ok := ret.Get(0).(func(lessons.Lesson) lessons.Lesson); ok {
		r0 = rf(lesson)
	} else {
		r0 = ret.Get(0).(lessons.Lesson)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(lessons.Lesson)error); ok {
		r1 = rf(lesson)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockLessonRepository) Delete(ID int) (bool, error) {
	ret := _m.Called(ID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(bool)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}