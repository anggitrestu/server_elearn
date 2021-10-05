package mocks

import (
	"server_elearn/models/orders"

	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (_m *MockOrderRepository) Save(order orders.Order) (orders.Order, error) {
	ret := _m.Called(order)

	var r0 orders.Order
	if rf, ok := ret.Get(0).(func(orders.Order) orders.Order); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Get(0).(orders.Order)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(orders.Order) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockOrderRepository) UpdateOrder(order orders.Order) (orders.Order, error) {
	ret := _m.Called(order)

	var r0 orders.Order
	if rf, ok := ret.Get(0).(func(orders.Order) orders.Order); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Get(0).(orders.Order)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(orders.Order) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockOrderRepository) FindAllByUserID(userID int) ([]orders.Order, error) {
	ret := _m.Called(userID)

	var r0 []orders.Order
	if rf, ok := ret.Get(0).(func(int) []orders.Order); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).([]orders.Order)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockOrderRepository) GetByID(ID int) (orders.Order, error) {
	ret := _m.Called(ID)

	var r0 orders.Order
	if rf, ok := ret.Get(0).(func(int) orders.Order); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(orders.Order)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
