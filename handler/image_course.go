package handler

import (
	"fmt"
	"net/http"
	"server_elearn/helper"
	imagecourses "server_elearn/models/image_courses"
	"server_elearn/service"

	"github.com/gin-gonic/gin"
)

type imageCourseHandler struct {
	serviceImageCourse service.ServiceImageCourse
	serviceCourse service.ServiceCourse
}

func NewImageCourseHandler(serviceImageCourse service.ServiceImageCourse, serviceCourse service.ServiceCourse) *imageCourseHandler {
	return &imageCourseHandler{serviceImageCourse, serviceCourse}
}

func(h *imageCourseHandler) CreateImageCourse(c *gin.Context) {
	var input imagecourses.CreateImageCourseInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create image course", http.StatusUnprocessableEntity, "error", errorMessage)
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
		response := helper.APIResponse("Course is not available", http.StatusNotFound, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newImageCourse, err := h.serviceImageCourse.CreateImageCourse(input)
	if err != nil {

		response := helper.APIResponse("Create image course failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success create image course", http.StatusOK, "success", imagecourses.FormatImageCourse(newImageCourse))
	c.JSON(http.StatusOK, reponse)

}

func( h*imageCourseHandler) DeleteImageCourse(c *gin.Context) {
	var inputID imagecourses.GetImageCourseInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed delete image course", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	imageCourse, err := h.serviceImageCourse.GetImageCourseByID(inputID)
	if err != nil || imageCourse.ID < 1 {
		message :=  "Failed to get course"
		if imageCourse.ID < 1 {
			message = "course Not Found"
		}
		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	_, err = h.serviceImageCourse.DeleteImageCourse(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete image course", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	message := fmt.Sprintf("success delete image course id : %d", inputID.ID)
	response := helper.APIResponse(message, http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)

}