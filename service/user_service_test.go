package service_test

import (
	"server_elearn/models/users"
	"server_elearn/service/mocks"

	"server_elearn/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


var userRepository = mocks.MockUserRepository{Mock: mock.Mock{}}
var userServiceMock = service.NewServiceUser(&userRepository)
var userDomain = users.User {
		ID:         2,
		Name:       "anggit restu",
		Email:      "anggit@gmail.com",
		Password:   "$2a$04$PX8dnTskiZTuqeELE6bh2O.VqCQipZD54WJP55xmSwLBoOYcdlHM2",
		Profession: "Fullstack Developer",
		AvatarFileName: "images/avataruser1.jpg",
		Role: "student",
}

var userDomainNil = users.User {}


func TestRegisterUser(t *testing.T){
	userRepository.On("Save",mock.Anything).Return(userDomain, nil)
	t.Run("Test Case 1 | Succes Register", func(t *testing.T) {
		user, err := userServiceMock.RegisterUser(users.RegisterUserInput{
			Name: "anggit restu",
			Profession: "fullstack",
			Email: "anggit@gmail.com",
			Password: "password",
		})
		assert.Nil(t, err)
		assert.Equal(t, userDomain, user)
	})
}

func TestLoginSuccess(t *testing.T){
	t.Run("Test Case 1 | Login Success", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.Anything).Return(userDomain, nil).Once()
		user, err := userServiceMock.Login(users.LoginInput{
			Email: "anggit@gmail.com",
			Password: "password",
		})
		assert.Nil(t, err)
		assert.Equal(t, "anggit@gmail.com", user.Email)
	})
}

func TestEmailAvaibility(t *testing.T){	
		userDomainNil = users.User{}
		userRepository.On("FindByEmail",mock.AnythingOfType("string")).Return(userDomainNil, nil)
		user, err := userServiceMock.IsEmailAvailable(users.CheckEmailInput{
			Email: "wkkwkwkwkwkwkkwk@gmail.com",
		})
		assert.Nil(t, err)
		assert.NotNil(t,user)
		assert.Equal(t, true, user)
}


func TestEmailNotAvaibility(t *testing.T){
		userRepository.On("FindByEmail", mock.AnythingOfType("string")).Return(userDomain, nil)
		user, err := userServiceMock.IsEmailAvailable(users.CheckEmailInput{
			Email: "anggit@gmail.com",
		})
		assert.Nil(t, err)
		assert.Equal(t, true, user)
}

func TestSaveAvatar(t *testing.T){
	userRepository.On("FindById", mock.AnythingOfType("int")).Return(userDomain, nil)
	userRepository.On("Update", mock.Anything).Return(userDomain,nil)
	user, err := userServiceMock.SaveAvatar(1, "images/avataruser1.jpg")
	assert.Nil(t, err)
	assert.NotNil(t,user)
	assert.Equal(t, user.AvatarFileName, userDomain.AvatarFileName)

}

func TestGetUserByID(t *testing.T){
	userRepository.On("FindById",2).Return(userDomain, nil)
	user, err := userServiceMock.GetUserByID(2)
	assert.Nil(t, err)
	assert.NotNil(t,user)
	assert.Equal(t, user.Name, userDomain.Name)
}

func TestGetUserByID__NotFound(t *testing.T){
	userRepository.On("FindById",2).Return(userDomainNil, nil)
	user, err := userServiceMock.GetUserByID(2)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.ID, 2)
}