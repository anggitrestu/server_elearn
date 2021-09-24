package service

import (
	"server_elearn/models/lessons"
	"server_elearn/repository"
)

type ServiceLesson interface {
	CreateLesson(input lessons.CreateLessonInput)(lessons.Lesson, error)
	GetLesson(input lessons.GetLessonInput)(lessons.Lesson, error)
	GetLessons(chapterID int)([]lessons.Lesson, error)
	UpdateLesson(inputID lessons.GetLessonInput, inputData lessons.CreateLessonInput)(lessons.Lesson, error)
	DeleteLesson(inputID lessons.GetLessonInput)(bool, error) 
}

type serviceLesson struct {
	repositoryLesson repository.LessonRepository
}

func NewServiceLesson(repositoryLesson repository.LessonRepository) *serviceLesson {
	return &serviceLesson{repositoryLesson}
}

func(s *serviceLesson) CreateLesson(input lessons.CreateLessonInput)(lessons.Lesson, error){
	
	lesson := lessons.Lesson{}
	lesson.Name = input.Name
	lesson.Video = input.Video
	lesson.ChapterID = input.ChapterID

	newLesson, err := s.repositoryLesson.Save(lesson)
	if err != nil {
		return newLesson, err
	}
	return newLesson, nil

}

func(s *serviceLesson) GetLesson(input lessons.GetLessonInput)(lessons.Lesson, error) {
	lesson, err := s.repositoryLesson.FindByID(input.ID)
	if err != nil {
		return lesson, err
	}

	return lesson, nil
}

func(s *serviceLesson) GetLessons(chapterID int)([]lessons.Lesson, error) {
	if chapterID != 0 {
		lessons, err := s.repositoryLesson.FindByChapterID(chapterID)
		if err != nil {
			return lessons, err
		}
		return lessons, nil
	}

	lessons, err := s.repositoryLesson.FindAll()
	if err != nil {
		return lessons, err
	}

	return lessons, nil
}

func(s *serviceLesson) UpdateLesson(inputID lessons.GetLessonInput, inputData lessons.CreateLessonInput)(lessons.Lesson, error) {
	lesson, err := s.repositoryLesson.FindByID(inputID.ID)
	if err != nil {
		return lesson, err
	}

	lesson.Name = inputData.Name
	lesson.ChapterID = inputData.ChapterID

	updateLesson, err := s.repositoryLesson.Update(lesson)
	if err != nil {
		return updateLesson, err
	}

	return updateLesson, nil
}


func(s *serviceLesson) DeleteLesson(inputID lessons.GetLessonInput)(bool, error) {
	_, err := s.repositoryLesson.Delete(inputID.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}