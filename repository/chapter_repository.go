package repository

import (
	"server_elearn/models/chapters"

	"gorm.io/gorm"
)

type ChapterRepository interface {
	Save(chapter chapters.Chapter)(chapters.Chapter, error)
	FindByCourseID(courseID int)([]chapters.Chapter, error)
	FindByID(ID int)(chapters.Chapter, error)
	FindAll()([]chapters.Chapter,error)
	Update(chapter chapters.Chapter)(chapters.Chapter, error)
	Delete(ID int)(bool, error)
}

type chapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) *chapterRepository {
	return &chapterRepository{db}
}

func (r *chapterRepository) Save(chapter chapters.Chapter)(chapters.Chapter, error) {

	err := r.db.Create(&chapter).Error
	if err != nil {
		return chapter, err
	}

	return chapter, nil

}

func(r *chapterRepository)FindByID(ID int)(chapters.Chapter, error) {
	
	var chapter chapters.Chapter
	err := r.db.Where("id = ?", ID).Find(&chapter).Error
	if err != nil {
		return chapter, err
	}

	return chapter, nil
}

func(r *chapterRepository)FindAll()([]chapters.Chapter,error){
	var chapters []chapters.Chapter
	err := r.db.Find(&chapters).Error
	if err != nil {
		return chapters, err
	}
	return chapters, nil
}

func(r *chapterRepository)FindByCourseID(courseID int)([]chapters.Chapter, error){
	var chapters []chapters.Chapter

	err := r.db.Where("course_id = ? ", courseID).Preload("Lessons").Find(&chapters).Error
	if err != nil {
		return chapters, err
	}
	return chapters, nil
}

func(r *chapterRepository)Update(chapter chapters.Chapter)(chapters.Chapter, error) {
	err := r.db.Save(&chapter).Error
	if err != nil {
		return chapter, err
	}

	return chapter, nil
}

func(r *chapterRepository) Delete(ID int)(bool, error) {
	chapter := chapters.Chapter{}
	err := r.db.Where("id = ?", ID).Delete(&chapter).Error
	if err != nil {
		return false, err
	}

	return true, nil
}