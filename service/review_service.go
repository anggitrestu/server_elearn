package service

import (
	"server_elearn/models/reviews"
	"server_elearn/repository"
)

type ServiceReview interface {
	CreateReview(input reviews.CreateReviewInput, userID int)(reviews.Review, error)
	UpdateReview(inputID reviews.GetReviewInput, inputData reviews.CreateReviewInput)(reviews.Review, error)
	DeleteReview(inputID reviews.GetReviewInput)(bool, error)
	GetReviewByID(input reviews.GetReviewInput)(reviews.Review, error)
}

type serviceReview struct {
	repositoryReview repository.ReviewRepository
}

func NewServiceReview(repositoryReview repository.ReviewRepository)*serviceReview {
	return &serviceReview{repositoryReview}
}

func(s *serviceReview)CreateReview(input reviews.CreateReviewInput, userID int)(reviews.Review, error){
	review := reviews.Review{}
	review.CourseID = input.CourseID
	review.UserID = userID
	review.Rating = input.Rating
	review.Note = input.Note

	newReview, err := s.repositoryReview.Save(review)
	if err != nil {
		return newReview, err
	}

	return newReview, nil
}


func(s *serviceReview)UpdateReview(inputID reviews.GetReviewInput, inputData reviews.CreateReviewInput)(reviews.Review, error){
	review , err := s.repositoryReview.FindByID(inputID.ID) 
	if err != nil {
		return review, err
	}

	review.CourseID = inputData.CourseID
	review.Rating = inputData.Rating
	review.Note = inputData.Note

	updateReview, err := s.repositoryReview.Update(review)
	if err != nil {
		return updateReview, err
	}

	return review, nil
}

func(s *serviceReview) DeleteReview(inputID reviews.GetReviewInput)(bool, error){
	_, err := s.repositoryReview.Delete(inputID.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func(s *serviceReview) GetReviewByID(input reviews.GetReviewInput)(reviews.Review, error){
	review, err := s.repositoryReview.FindByID(input.ID)
	if err != nil {
		return review, err
	}

	return review, nil
}

