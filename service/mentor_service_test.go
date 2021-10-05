package service_test

import (
	"server_elearn/models/mentors"
	"server_elearn/service"
	"server_elearn/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mentorRepository = mocks.MockMentorRepository{Mock: mock.Mock{}}
var mentorServiceMock = service.NewServiceMentor(&mentorRepository)
var mentorDomain = mentors.Mentor {
	ID: 1,
	Name: "budi atmoko",
	Email: "budi@gmail.com",
	Profile: "images/tutor/tutor1.jpg",
	Profession: "Front End Developer",
}

func TestAddMentor__Success(t *testing.T){
	mentorRepository.On("Save", mock.Anything).Return(mentorDomain, nil)
	newMentor, err := mentorServiceMock.AddMentor(mentors.AddMentorInput{
		Name: "budi atmoko",
		Email: "budi@gmail.com",
		Profile: "images/tutor/tutor1.jpg",
		Profession: "Front End Developer",
	})
	assert.Nil(t, err)
	assert.Equal(t, mentorDomain, newMentor)
}

func TestGetMentorByID__Success(t *testing.T){
	mentorRepository.On("FindByID", mock.AnythingOfType("int")).Return(mentorDomain, nil)
	mentor, err := mentorServiceMock.GetMentorByID(mentors.GetMentorInput{
		ID: 1,
	})
	assert.Nil(t, err)
	assert.Equal(t, mentorDomain, mentor)
}

func TestGetListMentor__Success(t *testing.T){
	var mentors []mentors.Mentor
	mentorRepository.On("FindAll").Return(mentors, nil)
	allMentor, err := mentorServiceMock.GetListMentor()
	assert.Nil(t, err)
	assert.Equal(t, mentors, allMentor)
}

func TestUpdateMentor__Success(t *testing.T){
	mentorRepository.On("FindByID", 1).Return(mentorDomain, nil)
	mentorRepository.On("Update", mock.Anything).Return(mentorDomain, nil)
	mentor, err := mentorServiceMock.UpdateMentor(mentors.GetMentorInput{
		ID: 1,
	}, mentors.AddMentorInput{
		Name: "budi atmoko",
		Email: "budi@gmail.com",
		Profile: "images/tutor/tutor1.jpg",
		Profession: "Front End Developer",
	})
	assert.Nil(t, err)
	assert.Equal(t, mentorDomain, mentor)

}

func TestDeleteMentor__Success(t *testing.T){
	mentorRepository.On("Delete", mock.AnythingOfType("int")).Return(mentorDomain, nil)
	_, err := mentorServiceMock.DeleteMentor(mentors.GetMentorInput{ID: 1})
	assert.Nil(t, err)
}

func TestMentorIsAvaibility__Success(t *testing.T){
	mentorRepository.On("CheckMentorByID", mock.AnythingOfType("int")).Return(mentorDomain, nil)
	mentor, err := mentorServiceMock.MentorIsAvaibility(1)
	assert.Nil(t, err)
	assert.Equal(t, true, mentor)
}