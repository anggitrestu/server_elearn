package handler

import (
	"net/http"
	"server_elearn/helper"
	"server_elearn/models/orders"
	"server_elearn/models/users"
	"server_elearn/service"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	serviceOrder service.ServiceOrder
}


func NewOrderHandler(serviceOrder service.ServiceOrder)*orderHandler {
	return &orderHandler{serviceOrder }
}

func(h *orderHandler)CreateOrder(c *gin.Context){

	var inputData orders.CreateOrderInput
	

	err := c.ShouldBind(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to place an order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	newOrder , err := h.serviceOrder.CreateOrder(inputData)
	if err != nil {
		response := helper.APIResponse("Create order failed", http.StatusBadRequest, "error", err.Error() ) 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	reponse := helper.APIResponse("Success create order", http.StatusOK, "success", newOrder)
	c.JSON(http.StatusOK, reponse)

}

func(h *orderHandler) GetOrders(c *gin.Context){
	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID

	orders, err := h.serviceOrder.GetOrders(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get all orders", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get all orders", http.StatusOK, "success", orders)
	c.JSON(http.StatusOK, response)

}


func(h *orderHandler) Webhook(c *gin.Context){
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	
	c.JSON(http.StatusOK, reqBody)
}