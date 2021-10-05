package service_test

import (
	"server_elearn/models/courses"
	"server_elearn/models/users"
	"server_elearn/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

var paymentService service.ServicePayment


func setupPayment(){
	paymentService = service.NewServicePayment(MIDTRANS_CLIENT_KEY, MIDTRANS_SERVER_KEY)
}

func TestGetPaymentURL__Success(t *testing.T){
	setupPayment()
	snapUrl, err := paymentService.GetPaymentURL(1, users.User{
		Name: "anggit",
		Email: "anggitrestu60@gmail.com",
	}, courses.Course{
		Price: 280000,
	})
	assert.NotNil(t, err)
	assert.Equal(t,snapUrl, "" )
}