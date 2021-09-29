package service

import (
	"encoding/json"
	"server_elearn/models/orders"
	"server_elearn/repository"
)

type ServiceOrder interface {
	CreateOrder(input orders.CreateOrderInput)(orders.Order, error)
	GetOrders(userID int)([]orders.Order, error)
}

type serviceOrder struct {
	repositoryOrder repository.OrderRepository
	servicePayment servicePayment
}

func NewServiceOrder(repositoryOrder repository.OrderRepository, servicePayment servicePayment) *serviceOrder {
	return &serviceOrder{repositoryOrder, servicePayment}
}

func(s *serviceOrder) CreateOrder(input orders.CreateOrderInput)(orders.Order, error){

	metadata := orders.Metadata {
		Course_id : input.Course.ID,
		Course_price : input.Course.Price,
		Course_name : input.Course.Name,
		Course_thumbnail : input.Course.Thumbnail,
		Course_level : input.Course.Level,
	}
	
	metadataToJson, _ := json.Marshal(metadata)

	order := orders.Order{}
	order.Status = "pending"
	order.CourseID = input.Course.ID
	order.UserID = input.User.ID
	order.Metadata = metadataToJson
	paymentUrl , err := s.servicePayment.GetPaymentURL(input)
	if err != nil {
		return order, err
	}
	order.SnapURL = paymentUrl

	newOrder , err := s.repositoryOrder.Save(order)
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

// func (s *serviceOrder) ProcessPayment(input orders.TransactionNotificationInput) error {
// 	transaction_id, _ := strconv.Atoi(input.OrderID)

// 	transaction, err := s.repository.GetByID(transaction_id)
// 	if err != nil {
// 		return err
// 	}

// 	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
// 		transaction.Status = "paid"
// 	} else if input.TransactionStatus == "settlement" {
// 		transaction.Status = "paid"
// 	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
// 		transaction.Status = "cancelled"
// 	}

// 	updatedTransaction, err := s.repository.Update(transaction)
// 	if err != nil {
// 		return err
// 	}

// 	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignID)
// 	if err != nil {
// 		return err
// 	}

// 	if updatedTransaction.Status == "paid" {
// 		campaign.BackerCount = campaign.BackerCount + 1
// 		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

// 		_, err := s.campaignRepository.Update(campaign)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }


