package handler

import (
	"net/http"
	"server_elearn/helper"
	"server_elearn/models/mycourses"
	"server_elearn/models/users"
	"server_elearn/service"

	"github.com/gin-gonic/gin"
)

type myCourseHandler struct {
	serviceMyCourse service.ServiceMyCourse
	serviceCourse service.ServiceCourse
}

func NewMyCourseHandler(serviceMyCourse service.ServiceMyCourse, serviceCourse service.ServiceCourse)*myCourseHandler {
	return &myCourseHandler{serviceMyCourse, serviceCourse}
}

func(h *myCourseHandler) CreateMyCourse(c *gin.Context){
	var input mycourses.CreateMyCourseInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create review", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	course, err := h.serviceCourse.GetCourseByCourseID(input.CourseID)
	if err != nil || course.ID < 1 {
		errorMessage := "Failed check course"
		if course.ID < 1 {
			errorMessage = "Courses is not available"
		}
		response := helper.APIResponse(errorMessage, http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID

	isExistMyCourse, err := h.serviceMyCourse.IsExistMyCourse(input, userID)
	if err != nil {
		response := helper.APIResponse("Failed check my course", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if isExistMyCourse.ID != 0 {
		response := helper.APIResponse("User already take this course", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newMyCourse, err := h.serviceMyCourse.CreateMyCourse(input, userID)
	if err != nil {

		response := helper.APIResponse("Create my course failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success create my course", http.StatusOK, "success", newMyCourse)
	c.JSON(http.StatusOK, reponse)



	// if course.Type == "premium" {
	// 	if course.Price == 0 {
	// 		response := helper.APIResponse("Price Can't be 0", http.StatusMethodNotAllowed, "error", nil ) 
	// 		c.JSON(http.StatusBadRequest, response)
	// 		return
	// 	}

	// 	order := 
	// }

}

func(h *myCourseHandler) GetAllMyCourse(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID

	myCourse , err := h.serviceMyCourse.GetAllMyCourse(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get my course", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	// if myCourse.ID == 0 {
	// 	response := helper.APIResponse("course Not Found", http.StatusNotFound, "error", nil)
	// 	c.JSON(http.StatusNotFound, response)
	// 	return
	// }
	
	response := helper.APIResponse("Success to get detail my course", http.StatusOK, "success", myCourse)
	c.JSON(http.StatusOK, response)
}