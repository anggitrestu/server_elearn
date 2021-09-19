package handler

import (
	"fmt"
	"net/http"
	"server_elearn/helper"
	"server_elearn/models/mentors"
	"server_elearn/service"

	"github.com/gin-gonic/gin"
)

type mentorHandler struct {
	serviceMentor service.ServiceMentor
}

func NewMentorHandler(serviceMentor service.ServiceMentor)*mentorHandler {
	return &mentorHandler{serviceMentor}
}

func (h *mentorHandler) AddMentor(c *gin.Context) {
	var input mentors.AddMentorInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to add mentor", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMentor, err := h.serviceMentor.AddMentor(input)
	if err != nil {
		errors := err.Error()
		response := helper.APIResponse(errors, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	
	reponse := helper.APIResponse("Success to add mentor", http.StatusOK, "success", mentors.FormatMentor(newMentor))
	c.JSON(http.StatusOK, reponse)

}

func (h *mentorHandler) GetMentor(c *gin.Context) {
	var input mentors.GetMentorInput
	err:= c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get detail mentor", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}



	mentorDetail , err := h.serviceMentor.GetMentorByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get mentor", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	if mentorDetail.ID < 1 {
		response := helper.APIResponse("Mentor Not Found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success to get detail mentor", http.StatusOK, "success", mentors.FormatMentor(mentorDetail))
	c.JSON(http.StatusOK, response)

}


func (h *mentorHandler) GetListMentor(c *gin.Context){
	
	listMentors , err := h.serviceMentor.GetListMentor()
	if err != nil {
		response := helper.APIResponse("Failed to get all mentor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get all mentor", http.StatusOK, "success", mentors.FormatMentors(listMentors))
	c.JSON(http.StatusOK, response)

}

func (h *mentorHandler) UpdateMentor(c *gin.Context) {
	var inputID mentors.GetMentorInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update mentor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mentor , err := h.serviceMentor.GetMentorByID(inputID)
	if err != nil || mentor.ID < 1 {
		message := "Failed to get mentor"
		if mentor.ID < 1 {
			message = "mentor not found"
		} 

		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var inputData mentors.AddMentorInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update mentor", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	

	updateMentor , err := h.serviceMentor.UpdateMentor(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update mentor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update mentor", http.StatusOK, "success", mentors.FormatMentor(updateMentor))
	c.JSON(http.StatusOK, response)
	
	
}

func (h *mentorHandler) DeleteMentor(c *gin.Context) {
	var inputID mentors.GetMentorInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete mentor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mentor , err := h.serviceMentor.GetMentorByID(inputID)
	if err != nil || mentor.ID < 1 {
		message :=  "Failed to get mentor"
		if mentor.ID < 1 {
			message = "Mentor Not Found"
		}
		response := helper.APIResponse(message, http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	_, err = h.serviceMentor.DeleteMentor(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update mentor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	message := fmt.Sprintf("success delete mentor id : %d", inputID.ID)
	response := helper.APIResponse(message, http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)
	

}