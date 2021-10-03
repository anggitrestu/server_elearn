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
	serviceUser service.ServiceUser
	serviceOrder service.ServiceOrder
}

func NewMyCourseHandler(serviceMyCourse service.ServiceMyCourse, serviceCourse service.ServiceCourse, serviceUser service.ServiceUser, serviceOrder service.ServiceOrder)*myCourseHandler {
	return &myCourseHandler{serviceMyCourse, serviceCourse, serviceUser, serviceOrder}
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



	
	if course.Type == "premium" {
		if course.Price == 0 {
					response := helper.APIResponse("Price Can't be 0", http.StatusMethodNotAllowed, "error", nil ) 
					c.JSON(http.StatusBadRequest, response)
					return
		}
			
		order , err := h.serviceOrder.CreateOrder(currentUser, course)
		if err != nil {
			response := helper.APIResponse("Create order failed", http.StatusBadRequest, "error", nil ) 
			c.JSON(http.StatusBadRequest, response)
			return
		}

		newOrder, err := h.serviceOrder.UpdateOrder(order.ID, currentUser, course)
		if err != nil {
			response := helper.APIResponse("Create order failed", http.StatusBadRequest, "error", nil ) 
			c.JSON(http.StatusBadRequest, response)
			return
		}
		reponse := helper.APIResponse("Success create order", http.StatusOK, "success", newOrder)
		c.JSON(http.StatusOK, reponse)
	} else {
		
		newMyCourse, err := h.serviceMyCourse.CreateMyCourse(input, userID)
		if err != nil {

		response := helper.APIResponse("Create my course failed", http.StatusBadRequest, "error", nil ) 
		c.JSON(http.StatusBadRequest, response)
		return
		}

				
		reponse := helper.APIResponse("Success create my course", http.StatusOK, "success", mycourses.FormatMyCourse(newMyCourse))
		c.JSON(http.StatusOK, reponse)

	}


	
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

	if len(myCourse) == 0 {
		response := helper.APIResponse("you don't have a course", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	
	response := helper.APIResponse("Success to get detail my course", http.StatusOK, "success", mycourses.FormatMyAllCourses(myCourse))
	c.JSON(http.StatusOK, response)
}

