package mocks

import (
	"server_elearn/models/users"

	"github.com/stretchr/testify/mock"
)


type MockUserRepository struct {
	mock.Mock
}

func(_m *MockUserRepository)Save(user users.User)(users.User, error){
	ret := _m.Called(user)

	var r0 users.User
	if rf, ok := ret.Get(0).(func(users.User) users.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(users.User)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(users.User)error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(_m *MockUserRepository)FindByEmail(email string) (users.User, error){
	
	ret := _m.Called(email)

	var r0 users.User
	if rf, ok := ret.Get(0).(func(string) users.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(users.User)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(string)error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(_m *MockUserRepository)FindById(ID int)(users.User, error){

	ret := _m.Called(ID)

	var r0 users.User
	if rf, ok := ret.Get(0).(func(int) users.User); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(users.User)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}


func(_m *MockUserRepository)Update(user users.User)(users.User, error){
	ret := _m.Called(user)

	var r0 users.User
	if rf, ok := ret.Get(0).(func(users.User) users.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(users.User)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(users.User)error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}