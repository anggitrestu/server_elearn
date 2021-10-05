package service_test

import (
	"server_elearn/models/reviews"
	"server_elearn/service"
	"server_elearn/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repositoryReview mocks.MockReviewRepository
var serviceReview service.ServiceReview
var reviewDomain reviews.Review

func setupReview(){
	serviceReview = service.NewServiceReview(&repositoryReview)
	reviewDomain = reviews.Review{
		ID: 1,
		CourseID: 1,
		UserID: 1,
		Rating: 4,
		Note: "Kelasnya keren bangett",
	}
}

func TestCreateReview__Succes(t *testing.T){
	setupReview()
	repositoryReview.On("Save", mock.Anything).Return(reviewDomain, nil).Once()
	review, err := serviceReview.CreateReview(reviews.CreateReviewInput{
		CourseID: 2,
		Rating: 3,
		Note: "Kelasnya keren bangett",
	}, 1)
	assert.Nil(t, err)
	assert.NotNil(t, review)
	assert.Equal(t, review, reviewDomain)
}

func TestUpdateReview__Success(t *testing.T){
	setupReview()
	repositoryReview.On("FindByID", mock.AnythingOfType("int")).Return(reviewDomain, nil).Once()
	repositoryReview.On("Update", mock.Anything).Return(reviewDomain, nil).Once()
	review, err := serviceReview.UpdateReview(reviews.GetReviewInput{
		ID: 1,
	}, reviews.CreateReviewInput{
		CourseID: 1,
		Rating: 4,
		Note: "Kelasnya keren bangett",
	})
	assert.Nil(t, err)
	assert.NotNil(t, review)
	assert.Equal(t, review, reviewDomain)
}

func TestDeleteReview__Success(t  *testing.T){
	setupReview()
	repositoryReview.On("Delete", mock.Anything).Return(true, nil).Once()
	success, err := serviceReview.DeleteReview(reviews.GetReviewInput{
		ID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, success)
	assert.Equal(t, success, true)
}

func TestGetReviewByID__Success(t *testing.T){
	setupReview()
	repositoryReview.On("FindByID", mock.AnythingOfType("int")).Return(reviewDomain, nil).Once()
	review, err := serviceReview.GetReviewByID(reviews.GetReviewInput{
		ID: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, review)
	assert.Equal(t, review, reviewDomain)
}