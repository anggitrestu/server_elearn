package service

import (
	"encoding/json"
	"server_elearn/models/courses"
	"server_elearn/models/orders"
	"server_elearn/models/users"
	"server_elearn/repository"
	"strconv"
	"strings"
)

type ServiceOrder interface {
	CreateOrder(user users.User, course courses.Course)(orders.Order, error)
	GetOrders(userID int)([]orders.Order, error)
	ProcessOrder(input orders.TransactionNotificationInput) error
	UpdateOrder(orderID int, user users.User, course courses.Course)(orders.Order, error)
}

type serviceOrder struct {
	repositoryOrder repository.OrderRepository
	servicePayment servicePayment
}

func NewServiceOrder(repositoryOrder repository.OrderRepository, servicePayment servicePayment) *serviceOrder {
	return &serviceOrder{repositoryOrder, servicePayment}
}

func(s *serviceOrder) CreateOrder(user users.User, course courses.Course)(orders.Order, error){

	order := orders.Order{}
	order.CourseID = course.ID
	order.UserID = user.ID
	
	newOrder , err := s.repositoryOrder.Save(order)
	if err != nil {
		return newOrder, err
	}

	return newOrder, nil
}

func(s *serviceOrder) UpdateOrder(orderID int, user users.User, course courses.Course)(orders.Order, error){
	order, err := s.repositoryOrder.GetByID(orderID)
	if err != nil {
		return order, err
	}
	metadata := orders.Metadata {
		Course_id : course.ID,
		Course_price : course.Price,
		Course_name : course.Name,
		Course_thumbnail : course.Thumbnail,
		Course_level : course.Level,
	}
	
	metadataToJson, _ := json.Marshal(metadata)

	order.Status = "pending"
	order.CourseID = course.ID
	order.UserID = user.ID
	order.Metadata = metadataToJson
	paymentUrl , err := s.servicePayment.GetPaymentURL(orderID, user, course)
	if err != nil {
		return order, err
	}
	order.SnapURL = paymentUrl

	newOrder , err := s.repositoryOrder.UpdateOrder(order)
	if err != nil {
		return newOrder, err
	}

	return newOrder, nil
}


func(s *serviceOrder) GetOrders(userID int)([]orders.Order, error) {
	orders , err := s.repositoryOrder.FindAllByUserID(userID)
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (s *serviceOrder) ProcessOrder(input orders.TransactionNotificationInput) error {
	realOrderID := strings.Split(input.OrderID, "-")
	orderID,_ := strconv.Atoi(realOrderID[0])

	order, err := s.repositoryOrder.GetByID(orderID)
	if err != nil {
		return err
	}

	if input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		order.Status = "success"
	} else if input.TransactionStatus == "settlement" {
		order.Status = "success"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		order.Status = "cancelled"
	}
 
	err = s.repositoryOrder.Update(order)
	if err != nil {
		return err
	}

	return nil
}
