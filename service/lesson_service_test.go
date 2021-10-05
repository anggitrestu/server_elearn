package service_test

import (
	"server_elearn/models/lessons"
	"server_elearn/service"
	"server_elearn/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repositoryLesson mocks.MockLessonRepository
var serviceLesson service.ServiceLesson
var lessonDomain lessons.Lesson

func setupLesson(){
	serviceLesson = service.NewServiceLesson(&repositoryLesson)
	lessonDomain = lessons.Lesson{
		ID: 1,
		Name: "install react js",
		Video: "you.tube/reactjs",
		ChapterID: 1,
	}
}

func TestCreateLesson__Success(t *testing.T){
	setupLesson()
	repositoryLesson.On("Save", mock.Anything).Return(lessonDomain, nil).Once()
	lesson, err := serviceLesson.CreateLesson(lessons.CreateLessonInput{
		Name: "install react js",
		Video: "you.tube/reactjs",
		ChapterID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, lesson)
	assert.Equal(t, lesson, lessonDomain)
}

func TestGetLesson__Success(t *testing.T){
	setupLesson()
	repositoryLesson.On("FindByID", mock.AnythingOfType("int")).Return(lessonDomain, nil).Once()
	lesson, err := serviceLesson.GetLesson(lessons.GetLessonInput{
		ID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, lesson)
	assert.Equal(t, lesson, lessonDomain)
}

func TestGetLessons_Success(t *testing.T){
	setupLesson()
	var lessons []lessons.Lesson
	lessons = append(lessons, lessonDomain)
	repositoryLesson.On("FindByChapterID", mock.AnythingOfType("int")).Return(lessons, nil).Once()
	repositoryLesson.On("FindAll", mock.AnythingOfType("int")).Return(lessons, nil).Once()
	allLesson, err := serviceLesson.GetLessons(1)
	assert.Nil(t, err)
	assert.NotNil(t, allLesson)
	assert.Equal(t, allLesson, lessons)
}

func TestUpdateLesson__Success(t *testing.T){
	setupLesson()
	repositoryLesson.On("FindByID", mock.AnythingOfType("int")).Return(lessonDomain, nil).Once()
	repositoryLesson.On("Update", mock.Anything).Return(lessonDomain, nil).Once()
	lesson, err := serviceLesson.UpdateLesson(lessons.GetLessonInput{
		ID: 1,
	}, lessons.CreateLessonInput{
		Name: "install react js",
		Video: "you.tube/reactjs",
		ChapterID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, lesson)
	assert.Equal(t, lesson, lessonDomain)
}

func TestDeleteLesson__Success(t *testing.T){
	setupLesson()
	repositoryLesson.On("Delete", mock.AnythingOfType("int")).Return(true, nil).Once()
	success, err := serviceLesson.DeleteLesson(lessons.GetLessonInput{
		ID: 1,
	})
	
	assert.Nil(t, err)
	assert.NotNil(t, success)
	assert.Equal(t, success, true)
}