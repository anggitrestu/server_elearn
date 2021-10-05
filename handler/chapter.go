package handler

import (
	"fmt"
	"net/http"
	"server_elearn/helper"
	"server_elearn/models/chapters"
	"server_elearn/service"
	"strconv"

	"github.com/gin-gonic/gin"
)


type chapterHandler struct {
	serviceChapter service.ServiceChapter
	serviceCourse service.ServiceCourse
}

func NewChapterHandler(serviceChapter service.ServiceChapter, serviceCourse service.ServiceCourse)*chapterHandler {
	return &chapterHandler{serviceChapter, serviceCourse}
}
	

func(h *chapterHandler) CreateChapter(c *gin.Context) {

	var input chapters.CreateChapterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Crate Chapter Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	courseIsAvaibility, err := h.serviceCourse.CourseIsAvaibility(input.CourseID)
	if err != nil {
		response := helper.APIResponse("Failed check chapter", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !courseIsAvaibility {
		response := helper.APIResponse("Courses is not available", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newChapter, err := h.serviceChapter.AddChapter(input)
	if err != nil {
		response := helper.APIResponse("Create chapter failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success create chapter", http.StatusOK, "success", chapters.FormatChapter(newChapter))
	c.JSON(http.StatusOK, reponse)

}

func(h *chapterHandler)GetChapter(c *gin.Context) {
	var input chapters.GetChapterInput

	err:= c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Get chapter failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	chapter , err := h.serviceChapter.GetChapter(input)
	if err != nil {
		response := helper.APIResponse("get chapter failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success get chapter", http.StatusOK, "success", chapters.FormatChapter(chapter))
	c.JSON(http.StatusOK, reponse)

}

func(h *chapterHandler)GetChapters(c *gin.Context) {

	course_id, _ := strconv.Atoi(c.Query("course_id"))

	allChapter, err := h.serviceChapter.GetChapters(course_id)
	if err != nil {
		response := helper.APIResponse("get all chapter failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success get all chapter", http.StatusOK, "success", chapters.FormatChapters(allChapter))
	c.JSON(http.StatusOK, reponse)
}

func(h *chapterHandler)UpdateChapter(c *gin.Context){
	var inputID chapters.GetChapterInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update courses", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	chapter , err := h.serviceChapter.GetChapter(inputID)
	if err != nil || chapter.ID < 1 {
		message := "Failed get chapter"
		if chapter.ID < 1 {
			message = "chapter not found"
		} 

		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var inputData chapters.UpdateChapterInput
	c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update chapter", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateChapter, err := h.serviceChapter.UpdateChapter(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Update chapter failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success update chapter", http.StatusOK, "success", chapters.FormatChapter(updateChapter))
	c.JSON(http.StatusOK, reponse)

}

func(h *chapterHandler)DeleteChapter(c *gin.Context) {
	var inputID chapters.GetChapterInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed delete chapter", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	chapter, err := h.serviceChapter.GetChapter(inputID)
	if err != nil || chapter.ID < 1 {
		message :=  "Failed to get chapter"
		if chapter.ID < 1 {
			message = "chapter Not Found"
		}
		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	_, err = h.serviceChapter.DeleteChapter(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete mentor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	message := fmt.Sprintf("success delete chapter id : %d", inputID.ID)
	response := helper.APIResponse(message, http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)

}
