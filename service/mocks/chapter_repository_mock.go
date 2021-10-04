package mocks

import (
	"server_elearn/models/chapters"

	"github.com/stretchr/testify/mock"
)

type MockChapterRepository struct {
	mock.Mock
}

func (_m *MockChapterRepository) Save(chapter chapters.Chapter) (chapters.Chapter, error) {

	ret := _m.Called(chapter)

	var r0 chapters.Chapter
	if rf, ok := ret.Get(0).(func(chapters.Chapter) chapters.Chapter); ok {
		r0 = rf(chapter)
	} else {
		r0 = ret.Get(0).(chapters.Chapter)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(chapters.Chapter)error); ok {
		r1 = rf(chapter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockChapterRepository) FindByID(ID int) (chapters.Chapter, error) {

	ret := _m.Called(ID)

	var r0 chapters.Chapter
	if rf, ok := ret.Get(0).(func(int) chapters.Chapter); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(chapters.Chapter)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockChapterRepository) FindAll() ([]chapters.Chapter, error) {
	ret := _m.Called()

	var r0 []chapters.Chapter
	if rf, ok := ret.Get(0).(func() []chapters.Chapter); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]chapters.Chapter)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func()error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockChapterRepository) FindByCourseID(courseID int) ([]chapters.Chapter, error) {
	ret := _m.Called(courseID)
	

	var r0 []chapters.Chapter
	if rf, ok := ret.Get(0).(func(int) []chapters.Chapter); ok {
		r0 = rf(courseID)
	} else {
		r0 = ret.Get(0).([]chapters.Chapter)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(courseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockChapterRepository) Update(chapter chapters.Chapter) (chapters.Chapter, error) {
	ret := _m.Called(chapter)

	var r0 chapters.Chapter
	if rf, ok := ret.Get(0).(func(chapters.Chapter) chapters.Chapter); ok {
		r0 = rf(chapter)
	} else {
		r0 = ret.Get(0).(chapters.Chapter)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(chapters.Chapter)error); ok {
		r1 = rf(chapter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockChapterRepository) Delete(ID int) (bool, error) {
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

func (_m *MockChapterRepository) CheckChapterByID(chapterID int) (chapters.Chapter, error) {
	ret := _m.Called(chapterID)

	var r0 chapters.Chapter
	if rf, ok := ret.Get(0).(func(int) chapters.Chapter); ok {
		r0 = rf(chapterID)
	} else {
		r0 = ret.Get(0).(chapters.Chapter)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(chapterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}