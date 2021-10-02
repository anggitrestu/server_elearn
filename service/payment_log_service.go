package service

import (
	"encoding/json"
	"server_elearn/models/orders"
	paymentlogs "server_elearn/models/payment_logs"
	"server_elearn/repository"
)

type ServicePaymentLog interface {
	CreatePaymentLog(input orders.TransactionNotificationInput,data interface{})(error)
}

type servicePaymentLog struct {
	repositoryPaymentLog repository.PaymentLogRepository
}

func NewServicePaymentLog(repositoryPaymentLog repository.PaymentLogRepository)*servicePaymentLog{
	return &servicePaymentLog{repositoryPaymentLog}
}

func(s *servicePaymentLog)CreatePaymentLog(input orders.TransactionNotificationInput, data interface{}, orderID int)(error){
	rawResponse, _ := json.Marshal(data)
	paymentLog := paymentlogs.PaymentLog{}
	paymentLog.OrderID = orderID
	paymentLog.PaymentType = input.PaymentType
	paymentLog.RawResponse = rawResponse
	paymentLog.Status = input.TransactionStatus

	err := s.repositoryPaymentLog.Save(paymentLog)
	if err != nil {
		 return err
	}
	return nil
}