package service

import (
	"server_elearn/models/chapters"
	"server_elearn/repository"
)

type ServiceChapter interface {
	 AddChapter(input chapters.CreateChapterInput)(chapters.Chapter, error)
	 GetChapter(input chapters.GetChapterInput)(chapters.Chapter, error)
	 GetChapters(courseID int)([]chapters.Chapter, error)
	 UpdateChapter(inputID chapters.GetChapterInput, inputData chapters.UpdateChapterInput)(chapters.Chapter, error)
	 DeleteChapter(inputID chapters.GetChapterInput)(bool, error)
}

type serviceChapter struct {
	repositoryChapter repository.ChapterRepository
}

func NewServiceChapter(repositoryChapter repository.ChapterRepository) *serviceChapter {
	return &serviceChapter{repositoryChapter}
}

func (s *serviceChapter) AddChapter(input chapters.CreateChapterInput)(chapters.Chapter, error) {
	chapter := chapters.Chapter{}
	chapter.Name = input.Name
	chapter.CourseID = input.CourseID

	newChapter , err := s.repositoryChapter.Save(chapter)
	if err != nil {
		return newChapter, err
	}

	return newChapter, nil

}

func(s *serviceChapter) GetChapter(input chapters.GetChapterInput)(chapters.Chapter, error) {
	chapter, err := s.repositoryChapter.FindByID(input.ID)
	if err != nil {
		return chapter, err
	}

	return chapter, nil
}

func(s *serviceChapter)	GetChapters(courseID int)([]chapters.Chapter, error) {
	if courseID != 0 {
		chapters, err := s.repositoryChapter.FindByCourseID(courseID)
		if err != nil {
			return chapters, err
		}
		return chapters, nil
	}
	chapters, err := s.repositoryChapter.FindAll()
	if err != nil {
		return chapters, err
	}

	return chapters, nil
}

func(s *serviceChapter) UpdateChapter(inputID chapters.GetChapterInput, inputData chapters.UpdateChapterInput)(chapters.Chapter, error) {
	chapter, err := s.repositoryChapter.FindByID(inputID.ID)
	if err != nil {
		return chapter, err
	}

	chapter.Name = inputData.Name
	chapter.CourseID = inputData.CourseID

	updateChapter, err := s.repositoryChapter.Update(chapter)
	if err != nil {
		return updateChapter, err
	}
	return updateChapter, nil

}


func (s *serviceChapter) DeleteChapter(inputID chapters.GetChapterInput)(bool, error) {
	_, err := s.repositoryChapter.Delete(inputID.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
