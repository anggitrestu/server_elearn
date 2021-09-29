package handler

import (
	"fmt"
	"net/http"
	"server_elearn/helper"
	"server_elearn/models/courses"
	"server_elearn/service"

	"github.com/gin-gonic/gin"
)

type courseHandler struct {
	serviceCourses service.ServiceCourse
	serviceMentor service.ServiceMentor
	
}

func NewCourseHandler(serviceCourses service.ServiceCourse, serviceMentor service.ServiceMentor)*courseHandler {
	return &courseHandler{serviceCourses, serviceMentor}
}

func(h *courseHandler) CreateCourse(c *gin.Context) {
	var input courses.CreateCourseInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create course", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	mentorIsAvaibility , err := h.serviceMentor.MentorIsAvaibility(input.MentorID)
	if err != nil {
		response := helper.APIResponse("Failed check mentor", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !mentorIsAvaibility {
		response := helper.APIResponse("Mentor is not available", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCourse, err := h.serviceCourses.CreateCourse(input)
	if err != nil {

		response := helper.APIResponse("Create course failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success create course", http.StatusOK, "success", courses.FormatCourse(newCourse))
	c.JSON(http.StatusOK, reponse)

}

func (h *courseHandler) GetCourse(c *gin.Context) {
	var input courses.GetCourseInput
	err:= c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get detail course", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}



	detailCourse , err := h.serviceCourses.GetCourseByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get course", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	if detailCourse.ID < 1 {
		response := helper.APIResponse("course Not Found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	
	response := helper.APIResponse("Success to get detail course", http.StatusOK, "success", detailCourse)
	c.JSON(http.StatusOK, response)

}

func(h *courseHandler) GetCourses(c *gin.Context){
	
	listCourses, err := h.serviceCourses.GetCourses()
	if err != nil {
		response := helper.APIResponse("Failed to get all courses", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get all courses", http.StatusOK, "success", courses.FormatCourses(listCourses))
	c.JSON(http.StatusOK, response)
}

func(h *courseHandler) UpdateCourse(c *gin.Context){
	var inputID courses.GetCourseInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update courses", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	course, err := h.serviceCourses.GetCourseByID(inputID)
	if err != nil || course.ID < 1 {
		message := "Failed to get course"
		if course.ID < 1 {
			message = "course not found"
		} 

		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var inputData courses.CreateCourseInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update mentor", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateCourse, err := h.serviceCourses.UpdateCourse(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Update course failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success update course", http.StatusOK, "success", courses.FormatCourse(updateCourse))
	c.JSON(http.StatusOK, reponse)

}

func(h *courseHandler) DeleteCourse(c *gin.Context){
	var inputID courses.GetCourseInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed delete course", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	course, err := h.serviceCourses.GetCourseByID(inputID)
	if err != nil || course.ID < 1 {
		message :=  "Failed to get course"
		if course.ID < 1 {
			message = "course Not Found"
		}
		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	_, err = h.serviceCourses.DeleteCourse(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete mentor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	message := fmt.Sprintf("success delete mentor id : %d", inputID.ID)
	response := helper.APIResponse(message, http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)
}

