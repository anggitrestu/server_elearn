package service_test

import (
	"server_elearn/models/chapters"
	"server_elearn/service"
	"server_elearn/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repositoryChapter mocks.MockChapterRepository
var serviceMockChapter service.ServiceChapter
var chapterDomain chapters.Chapter

func setupChapter() {
	serviceMockChapter = service.NewServiceChapter(&repositoryChapter)
	chapterDomain = chapters.Chapter{
		ID: 1,
		Name: "Pengenalan React Js",
		CourseID: 1,
	}
}

func TestAddChapter__Success(t *testing.T){
	setupChapter()
	repositoryChapter.On("Save", mock.Anything).Return(chapterDomain, nil).Once()
	chapter, err := serviceMockChapter.AddChapter(chapters.CreateChapterInput{
		Name: "Pengenalan React Js",
		CourseID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, chapter)
	assert.Equal(t, chapter, chapterDomain)
}

func TestGetChapter__Success(t *testing.T){
	setupChapter()
	repositoryChapter.On("FindByID", mock.AnythingOfType("int")).Return(chapterDomain, nil).Once()
	chapter, err := serviceMockChapter.GetChapter(chapters.GetChapterInput{
		ID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, chapter)
	assert.Equal(t, chapter, chapterDomain)
}

func TestGetChapters__Success(t *testing.T){
	setupChapter()
	var chapters []chapters.Chapter
	chapters = append(chapters, chapterDomain)
	repositoryChapter.On("FindByCourseID", mock.AnythingOfType("int")).Return(chapters, nil).Once()
	allChapter, err := serviceMockChapter.GetChapters(1)
	assert.Nil(t, err)
	assert.NotNil(t, allChapter)
	assert.Equal(t, allChapter, chapters)
}

func TestUpdateChapter__Success(t *testing.T){
	setupChapter()
	repositoryChapter.On("FindByID", mock.AnythingOfType("int")).Return(chapterDomain, nil).Once()
	repositoryChapter.On("Update", mock.Anything).Return(chapterDomain, nil).Once()

	newChapter, err := serviceMockChapter.UpdateChapter(chapters.GetChapterInput{
		ID: 1,
	}, chapters.UpdateChapterInput{
		Name: "Pengenalan React Js",
		CourseID: 1,
	})

	assert.Nil(t, err)
	assert.NotNil(t, newChapter)
	assert.Equal(t, newChapter, chapterDomain)
}

func TestDeleteChapter__Sucees(t *testing.T){
	setupChapter()
	repositoryChapter.On("Delete", mock.AnythingOfType("int")).Return(true, nil).Once()
	success, err := serviceMockChapter.DeleteChapter(chapters.GetChapterInput{
		ID: 1,
	})
	assert.Nil(t, err)
	assert.Equal(t, success, true)
}

func TestChapterIsAvaibility__Success(t *testing.T){
	setupChapter()
	repositoryChapter.On("CheckChapterByID", mock.AnythingOfType("int")).Return(chapterDomain, nil).Once()
	success, err := serviceMockChapter.ChapterIsAvaibility(1)
	assert.Nil(t, err)
	assert.Equal(t, success, true)
}