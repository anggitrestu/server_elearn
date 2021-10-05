package mocks

import (
	"server_elearn/models/courses"
	"server_elearn/models/users"

	"github.com/stretchr/testify/mock"
)

type MockPaymentService struct {
	mock.Mock
}

func (_m *MockPaymentService) GetPaymentURL(orderID int, user users.User, course courses.Course) (string, error) {

	ret := _m.Called(orderID, user, course)


	var r0 string
	if rf, ok := ret.Get(0).(func(int, users.User, courses.Course) string); ok {
		r0 = rf(orderID, user, course )
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, users.User, courses.Course) error); ok {
		r1 =rf(orderID, user, course )
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}