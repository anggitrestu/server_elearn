package service

import (
	"server_elearn/models/mentors"
	"server_elearn/repository"
)

type ServiceMentor interface {
	AddMentor(input mentors.AddMentorInput)(mentors.Mentor, error)
	GetMentorByID(input mentors.GetMentorInput) (mentors.Mentor, error) 
	GetListMentor()([]mentors.Mentor, error)
	UpdateMentor(inputID mentors.GetMentorInput ,inputData mentors.AddMentorInput)(mentors.Mentor, error)
	DeleteMentor(inputID mentors.GetMentorInput)(mentors.Mentor, error)
}

type serviceMentor struct {
	repositoryMentor repository.MentorRepository
}

func NewServiceMentor(repositoryMentor repository.MentorRepository)*serviceMentor {
	return &serviceMentor{repositoryMentor}
}

func (s *serviceMentor) AddMentor(input mentors.AddMentorInput) (mentors.Mentor, error) {
	mentor := mentors.Mentor{}
	mentor.Name = input.Name
	mentor.Profile = input.Profile
	mentor.Email = input.Email
	mentor.Profession = input.Profession

	newMentor, err := s.repositoryMentor.Save(mentor)
	if err!= nil {
		return newMentor, err
	}

	return newMentor, nil
}

func (s *serviceMentor) GetMentorByID(input mentors.GetMentorInput) (mentors.Mentor, error) {
	mentor, err := s.repositoryMentor.FindByID(input.ID)

	if err != nil {
		return mentor, err
	}

	return mentor , nil
}

func (s *serviceMentor) GetListMentor() ([]mentors.Mentor, error) {

	mentors, err := s.repositoryMentor.FindAll()
	
	if err != nil {
		return mentors, err
	}

	return mentors , nil
}

func (s *serviceMentor) UpdateMentor(inputID mentors.GetMentorInput ,inputData mentors.AddMentorInput)(mentors.Mentor, error) {
	mentor, err := s.repositoryMentor.FindByID(inputID.ID)

	if err != nil {
		return mentor, err
	}

	mentor.Name = inputData.Name
	mentor.Profile = inputData.Profile
	mentor.Profession = inputData.Profession
	mentor.Email = inputData.Email

	updateMentor, err := s.repositoryMentor.Update(mentor)
	if err != nil {
		return updateMentor, err
	}

	return updateMentor, nil

}

func (s *serviceMentor) DeleteMentor(inputID mentors.GetMentorInput)(mentors.Mentor, error) {
	mentor, err := s.repositoryMentor.Delete(inputID.ID)
	if err != nil {
		return mentor, err
	}

	return mentor, nil
}