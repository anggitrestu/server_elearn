package handler

import (
	"fmt"
	"net/http"
	"server_elearn/auth"
	"server_elearn/helper"
	"server_elearn/models/users"
	"server_elearn/service"

	"github.com/gin-gonic/gin"
)
type userHandler struct {
	userService service.ServiceUser
	authService auth.Service
}

func NewUserHandler(userService service.ServiceUser, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {	
	//tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service 

	var input users.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors" : errors}
		response := helper.APIResponse("Accout Register failed", http.StatusUnprocessableEntity, "error", errorMessage )
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	
	if err != nil {
		response := helper.APIResponse("Accout account failed", http.StatusBadRequest, "error", nil )
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return	
	}
	formatter := users.FormatUser(newUser, token)
	response := helper.APIResponse("Accout has been register", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukan input (email dan password)
	// input ditangkap handler
	// mappng dari input user ke input struct
	// input struct passting service
	// di service mencari dg bantuan rpository user dengan email
	// mencocokan password

	var input users.LoginInput
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors" : err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	token, err := h.authService.GenerateToken(loggedinUser.ID)

	if err != nil {
		response := helper.APIResponse("Login Failed",http.StatusBadRequest, "errors", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Succesfullt Loggedin", http.StatusOK,"success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvaibility(c *gin.Context){
	var input users.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_available" : isEmailAvailable}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded:": false}
		response := helper.APIResponse("failed to upload avatar imagee", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return 
	}

	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Save upload avatar image", http.StatusBadRequest, "error", data)
	
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar succesfully uploaded", http.StatusOK, "error", data)
	
	c.JSON(http.StatusOK, response)
	
}

func (h *userHandler) FetchUser(c *gin.Context){

	currentUser := c.MustGet("currentUser").(users.User)

	formatter := users.FormatUser(currentUser, "")

	response := helper.APIResponse("Succesfully fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}