package mocks

import (
	"server_elearn/models/reviews"

	"github.com/stretchr/testify/mock"
)

type MockReviewRepository struct {
	mock.Mock
}


func(_m *MockReviewRepository) Save(review reviews.Review)(reviews.Review, error) {
	ret := _m.Called(review)

	var r0 reviews.Review
	if rf, ok := ret.Get(0).(func(reviews.Review) reviews.Review); ok {
		r0 = rf(review)
	} else {
		r0 = ret.Get(0).(reviews.Review)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(reviews.Review)error); ok {
		r1 = rf(review)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}


func(_m *MockReviewRepository) Update(review reviews.Review)(reviews.Review, error){
	ret := _m.Called(review)

	var r0 reviews.Review
	if rf, ok := ret.Get(0).(func(reviews.Review) reviews.Review); ok {
		r0 = rf(review)
	} else {
		r0 = ret.Get(0).(reviews.Review)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(reviews.Review)error); ok {
		r1 = rf(review)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(_m *MockReviewRepository) FindByID(ID int)(reviews.Review, error) {
	ret := _m.Called(ID)

	var r0 reviews.Review
	if rf, ok := ret.Get(0).(func(int) reviews.Review); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(reviews.Review)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func(_m *MockReviewRepository) Delete(ID int)(bool, error) {
	ret := _m.Called(ID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(bool)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(int)error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}