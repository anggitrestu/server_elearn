package service

import (
	"fmt"
	"server_elearn/models/courses"
	"server_elearn/models/users"
	"strconv"
	"time"

	"github.com/veritrans/go-midtrans"
)

type servicePayment struct {
	MIDTRANS_CLIENT_KEY string
	MIDTRANS_SERVER_KEY string
}

type ServicePayment interface {
	GetPaymentURL(orderID int, user users.User, course courses.Course) (string, error)
}

func NewServicePayment(MIDTRANS_CLIENT_KEY string, MIDTRANS_SERVER_KEY string) *servicePayment {
	return &servicePayment{MIDTRANS_CLIENT_KEY, MIDTRANS_SERVER_KEY}
}

func generateOrderID(orderID int) string {
	rand := strconv.FormatInt(time.Now().UnixNano(), 10)
	path := fmt.Sprintf("%d-%s", orderID, rand) 
	return path
}

func (s *servicePayment) GetPaymentURL(orderID int, user users.User, course courses.Course) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = s.MIDTRANS_SERVER_KEY
	midclient.ClientKey =  s.MIDTRANS_CLIENT_KEY
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email:user.Email,
			FName:user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: generateOrderID(orderID),
			GrossAmt: int64(course.Price),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}


	return snapTokenResp.RedirectURL, nil
}
