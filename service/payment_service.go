package service

import (
	"server_elearn/models/orders"
	"strconv"
	"time"

	"github.com/veritrans/go-midtrans"
)

type servicePayment struct {
}

type ServicePayment interface {
	GetPaymentURL(data orders.CreateOrderInput) (string, error)
}

func NewServicePayment() *servicePayment {
	return &servicePayment{}
}

func generateOrderID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func (s *servicePayment) GetPaymentURL(data orders.CreateOrderInput) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-m2OkDmszlvtFNFT6XDpW2dbA"
	midclient.ClientKey =  "SB-Mid-client-t-V2YcQBWyf-JhN5"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: data.User.Email,
			FName: data.User.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: generateOrderID(),
			GrossAmt: int64(data.Course.Price),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}


	return snapTokenResp.RedirectURL, nil
}
