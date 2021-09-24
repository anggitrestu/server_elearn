package handler

import (
	"fmt"
	"net/http"
	"server_elearn/helper"
	"server_elearn/models/lessons"
	"server_elearn/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type lessonHandler struct {
	serviceLesson service.ServiceLesson
}

func NewLessonHandler(serviceLesson service.ServiceLesson)*lessonHandler {
	return &lessonHandler{serviceLesson}
}

func(h *lessonHandler) CreateLesson(c *gin.Context) {
	var input lessons.CreateLessonInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create chapter", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newLesson, err := h.serviceLesson.CreateLesson(input)
	if err != nil {
		errors := err.Error()
		response := helper.APIResponse(errors, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success create chapter", http.StatusOK, "success", lessons.FormatLesson(newLesson))
	c.JSON(http.StatusOK, reponse)

}

func(h *lessonHandler) GetLesson(c *gin.Context){
	var input lessons.GetLessonInput
	err:= c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get detail mentor", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	getLesson,err := h.serviceLesson.GetLesson(input)
	if err != nil {
		response := helper.APIResponse("Failed get lesson", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	if getLesson.ID < 1 {
		response := helper.APIResponse("lesson Not Found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success to get detail mentor", http.StatusOK, "success", lessons.FormatLesson(getLesson))
	c.JSON(http.StatusOK, response)

}

func(h *lessonHandler)GetLessons(c *gin.Context){
	chapter_id, _ := strconv.Atoi(c.Query("chapter_id"))
	
	allLesson, err := h.serviceLesson.GetLessons(chapter_id)
	if err != nil {
		response := helper.APIResponse("get all lesson failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success get all lesson", http.StatusOK, "success", lessons.FormatLessons(allLesson))
	c.JSON(http.StatusOK, reponse)
}

func(h *lessonHandler)UpdateLesson(c *gin.Context){
	var inputID lessons.GetLessonInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update courses", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getLesson, err := h.serviceLesson.GetLesson(inputID)
		if err != nil || getLesson.ID < 1 {
		message := "Failed get lesson"
		if getLesson.ID < 1 {
			message = "lesson not found"
		} 

		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var inputData lessons.CreateLessonInput
	c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update lesson", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateLesson, err := h.serviceLesson.UpdateLesson(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Update lesson failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success update lesson", http.StatusOK, "success", lessons.FormatLesson(updateLesson))
	c.JSON(http.StatusOK, reponse)

}


func(h *lessonHandler) DeleteLesson(c *gin.Context){
	var inputID lessons.GetLessonInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed delete chapter", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	lesson, err := h.serviceLesson.GetLesson(inputID)
	if err != nil || lesson.ID < 1 {
		message :=  "Failed to get lesson"
		if lesson.ID < 1 {
			message = "lesson Not Found"
		}
		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	_, err = h.serviceLesson.DeleteLesson(inputID)
	if err != nil {
		response := helper.APIResponse("Failed delete lesson", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	message := fmt.Sprintf("success delete lesson id : %d", inputID.ID)
	response := helper.APIResponse(message, http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)

}
