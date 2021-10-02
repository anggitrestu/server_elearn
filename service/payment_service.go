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
}

type ServicePayment interface {
	GetPaymentURL(orderID int, user users.User, course courses.Course) (string, error)
}

func NewServicePayment() *servicePayment {
	return &servicePayment{}
}

func generateOrderID(orderID int) string {
	rand := strconv.FormatInt(time.Now().UnixNano(), 10)
	path := fmt.Sprintf("%d-%s", orderID, rand) 
	return path
}

func (s *servicePayment) GetPaymentURL(orderID int, user users.User, course courses.Course) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-m2OkDmszlvtFNFT6XDpW2dbA"
	midclient.ClientKey =  "SB-Mid-client-t-V2YcQBWyf-JhN5"
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
