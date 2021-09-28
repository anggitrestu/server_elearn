package handler

import (
	"fmt"
	"net/http"
	"server_elearn/helper"
	"server_elearn/models/reviews"
	"server_elearn/models/users"
	"server_elearn/service"

	"github.com/gin-gonic/gin"
)

type reviewHandler struct {
	serviceReview service.ServiceReview
	serviceCourse service.ServiceCourse
}

func NewReviewHandler(serviceReview service.ServiceReview, serviceCourse service.ServiceCourse)*reviewHandler {
	return &reviewHandler{serviceReview, serviceCourse}
}


func(h *reviewHandler)CreateReview(c *gin.Context){
	var input reviews.CreateReviewInput
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create review", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	courseIsAvaibility, err := h.serviceCourse.CourseIsAvaibility(input.CourseID)
	if err != nil {
		response := helper.APIResponse("Failed check course", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !courseIsAvaibility {
		response := helper.APIResponse("Courses is not available", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID

	newReview, err := h.serviceReview.CreateReview(input, userID)
	if err != nil {
		response := helper.APIResponse("Create review failed", http.StatusBadRequest, "error", err.Error() ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}



	reponse := helper.APIResponse("Success create chapter", http.StatusOK, "success", reviews.FormatReview(newReview))
	c.JSON(http.StatusOK, reponse)
}

func(h *reviewHandler)UpdateReview(c *gin.Context){
	var inputID reviews.GetReviewInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update courses", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	review, err := h.serviceReview.GetReviewByID(inputID)
	if err != nil || review.ID < 1 {
		message := "Failed to get course"
		if review.ID < 1 {
			message = "review not found"
		} 

		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var inputData reviews.CreateReviewInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create review", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateReview, err := h.serviceReview.UpdateReview(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Update course failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success update course", http.StatusOK, "success", reviews.FormatReview(updateReview))
	c.JSON(http.StatusOK, reponse)

}

func(h *reviewHandler) DeleteReview(c *gin.Context){
	var inputID reviews.GetReviewInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed delete review", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	review, err := h.serviceReview.GetReviewByID(inputID)
	if err != nil || review.ID < 1 {
		message :=  "Failed to get review"
		if review.ID < 1 {
			message = "review Not Found"
		}
		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	_, err = h.serviceReview.DeleteReview(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete mentor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	message := fmt.Sprintf("success delete review id : %d", inputID.ID)
	response := helper.APIResponse(message, http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)
}