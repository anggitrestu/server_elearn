package service

// controller

import (
	"errors"
	"server_elearn/models/users"
	"server_elearn/repository"

	"golang.org/x/crypto/bcrypt"
)

// membuat mapping dari struct input ke struct user
// mewakili binis logic, kata kerja
type ServiceUser interface {
	RegisterUser(input users.RegisterUserInput) (users.User, error)
	Login(input users.LoginInput)(users.User, error)
	IsEmailAvailable(input users.CheckEmailInput)(bool, error)
	SaveAvatar(ID int, fileLocation string)(users.User, error)
	GetUserByID(ID int)(users.User, error)
}

type serviceUser struct {
	repository repository.UserRepository
}

func NewServiceUser(repository repository.UserRepository) *serviceUser {
	return &serviceUser{repository}
}

func(s *serviceUser) RegisterUser(input users.RegisterUserInput)(users.User, error){
	// s.repository.Save(user)
	user := users.User{}
	user.Name = input.Name
	user.Email  = input.Email
	user.Profession = input.Profession
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	
	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *serviceUser) Login(input users.LoginInput) (users.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, nil
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil

}

func (s *serviceUser) IsEmailAvailable(input users.CheckEmailInput)(bool, error){
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	
	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}


func (s *serviceUser) SaveAvatar(ID int, fileLocation string)(users.User, error) {
	user, err := s.repository.FindById(ID)
	
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func(s *serviceUser) GetUserByID(ID int)(users.User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found with that ID")
	}
	 return user, nil

}
