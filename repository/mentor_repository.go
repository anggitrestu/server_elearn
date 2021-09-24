package repository

import (
	"server_elearn/models/mentors"

	"gorm.io/gorm"
)

type MentorRepository interface {
	Save(mentor mentors.Mentor)(mentors.Mentor, error)
	FindByID(ID int) (mentors.Mentor, error)
	FindAll() ([]mentors.Mentor, error)
	Update(mentor mentors.Mentor)(mentors.Mentor, error)
	Delete(ID int)(mentors.Mentor, error)
	CheckMentorByID(mentorID int)(mentors.Mentor, error)
}

type mentorRepository struct {
	db *gorm.DB
}

func NewMentorRepository(db *gorm.DB) *mentorRepository {

	return &mentorRepository{db}

}

func (r *mentorRepository) Save(mentor mentors.Mentor)(mentors.Mentor, error) {
	err := r.db.Create(&mentor).Error
	if err != nil {
		return mentor, err
	}

	return mentor, nil
}

func (r *mentorRepository) FindByID(ID int) (mentors.Mentor, error) {
	var mentor mentors.Mentor
	err := r.db.Where("id = ?", ID).Find(&mentor).Error
	if err != nil {
		return mentor, err
	}

	return mentor, nil
}

func(r *mentorRepository) FindAll() ([]mentors.Mentor, error) {
	var mentors []mentors.Mentor

	err := r.db.Find(&mentors).Error 
	if err != nil {
		return mentors, err
	} 

	return mentors, nil

}

func (r *mentorRepository)Update(mentor mentors.Mentor)(mentors.Mentor, error){

	err := r.db.Save(&mentor).Error
	if err != nil {
		return mentor, err
	}

	return mentor, nil
}

func (r *mentorRepository) Delete(ID int)(mentors.Mentor, error) {
	mentor := mentors.Mentor{}
	err := r.db.Where("id = ?", ID).Delete(&mentor).Error
	if err != nil {
		return mentor, err
	}

	return mentor, nil
}

func(r *mentorRepository)CheckMentorByID(mentorID int)(mentors.Mentor, error) {
	var mentor mentors.Mentor
	err := r.db.Where("id = ?", mentorID).Find(&mentor).Error
	if err != nil {
		return mentor, err
	}

	return mentor, nil
}

