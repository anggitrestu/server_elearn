package service_test

import (
	"server_elearn/models/courses"
	"server_elearn/models/orders"
	"server_elearn/models/users"
	"server_elearn/service"
	"server_elearn/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var MIDTRANS_CLIENT_KEY string = "dummy_midtrans_client_key"
var MIDTRANS_SERVER_KEY string = "dummy_midtrans_server_key"

var orderRepository mocks.MockOrderRepository
var orderService service.ServiceOrder
var servicePayment mocks.MockPaymentService
var orderDomain orders.Order

func setupOrder() {
	servicePayment := service.NewServicePayment(MIDTRANS_CLIENT_KEY, MIDTRANS_SERVER_KEY)
	orderService = service.NewServiceOrder(&orderRepository, *servicePayment)
	orderDomain = orders.Order{
		ID:       1,
		Status:   "pending",
		CourseID: 1,
		UserID:   2,
		SnapURL:  "https://app.sandbox.midtrans.com/snap/v2/vtweb/2b161101-2309-468f-9ea6-2659b9358a89",
	}
}

func TestCreateOrder__Success(t *testing.T){
	setupOrder()
	orderRepository.On("Save", mock.Anything).Return(orderDomain, nil)
	order, err := orderService.CreateOrder(users.User{
		ID: 1,
		Name: "anggit",
		Email: "anggit@gmail.com",
	}, courses.Course{
		ID: 1,
		Price: 250000,
	})
	assert.Nil(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, order.ID, orderDomain.ID)
}

func TestGetOrders__Success(t *testing.T){
	setupOrder()
	var orders []orders.Order
	orders = append(orders, orderDomain)
	orderRepository.On("FindAllByUserID", mock.AnythingOfType("int")).Return(orders, nil).Once()
	allOrder, err := orderService.GetOrders(1)
	assert.Nil(t, err)
	assert.NotNil(t, allOrder)
	assert.Equal(t, allOrder, orders)
}

func TestProcessOrder__Success(t *testing.T){
	setupOrder()
	orderRepository.On("GetByID", mock.AnythingOfType("int")).Return(orderDomain, nil).Once()
	orderRepository.On("UpdateOrder", mock.Anything).Return(orderDomain, nil).Once()
	order, err := orderService.ProcessOrder(orders.TransactionNotificationInput{
		TransactionStatus: "success",
		OrderID: "1",
		PaymentType: "BCA",
		FraudStatus: "accept",
	})
	assert.Nil(t, err)
	assert.NotNil(t, order)
}

func TestUpdateOrder__Success(t *testing.T){
	orderRepository.On("GetByID", mock.AnythingOfType("int")).Return(orderDomain, nil).Once()
	servicePayment.On("GetPaymentURL", mock.AnythingOfType("int"), mock.Anything, mock.Anything).Return("www.payment.url", nil).Once()
	orderRepository.On("UpdateOrder", mock.Anything).Return(orderDomain, nil).Once()
	_, err := orderService.UpdateOrder(1, users.User{
		Name: "anggit",
		Email: "anggit@gmail.com",
	}, courses.Course{
		ID: 1,
		Price: 20000,
	})
	assert.NotNil(t, err)
	// assert.Equal(t, order, orderDomain)
}