package repository

import (
	"server_elearn/models/lessons"

	"gorm.io/gorm"
)

type LessonRepository interface {
	Save(lesson lessons.Lesson)(lessons.Lesson, error)
	FindByID(ID int)(lessons.Lesson, error)
	FindAll()([]lessons.Lesson, error)
	FindByChapterID(chapterID int)([]lessons.Lesson, error)
	Update(lesson lessons.Lesson)(lessons.Lesson, error)
	 Delete(ID int)(bool, error)
}

type lessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) *lessonRepository {
	return &lessonRepository{db}
}

func(r *lessonRepository) Save(lesson lessons.Lesson)(lessons.Lesson, error) {
	err := r.db.Create(&lesson).Error
	if err != nil {
		return lesson, err
	}

	return lesson, nil
}


func(r *lessonRepository) FindByID(ID int)(lessons.Lesson, error){
	var lesson lessons.Lesson
	err := r.db.Where("id = ?", ID).Find(&lesson).Error
	if err != nil {
		return lesson, err
	}
	return lesson, nil
}

func(r *lessonRepository) FindAll()([]lessons.Lesson, error){
	var lessons []lessons.Lesson

	err := r.db.Find(&lessons).Error
	if err != nil {
		return lessons, err
	}

	return lessons, nil
}

func(r *lessonRepository)FindByChapterID(chapterID int)([]lessons.Lesson, error){
	var lessons []lessons.Lesson

	err := r.db.Where("chapter_id = ?", chapterID).Find(&lessons).Error
	if err != nil {
		return lessons, err
	}

	return lessons, nil
}

func(r *lessonRepository) Update(lesson lessons.Lesson)(lessons.Lesson, error){
	err := r.db.Save(&lesson).Error
	if err != nil {
		return lesson, err
	}
	return lesson, nil
}

func(r *lessonRepository) Delete(ID int)(bool, error){
	lesson := lessons.Lesson{}
	err := r.db.Where("id = ? ", ID).Delete(&lesson).Error
	if err != nil {
		return false, err
	}

	return true, nil
}