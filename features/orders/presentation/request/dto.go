package request

import "project/e-comerce/features/orders"

type Order struct {
	//Price 	int 	`form:"price" json:"price"`
	CartID  []int   `form:"cartid" json:"cartid"`
	Address Address `form:"address" json:"address"`
}

type Address struct {
	Receiver string `form:"city" json:"city"`
	Phone    string `form:"province" json:"province"`
	Address  string `form:"postalcode" json:"postalcode"`
}

type Payment struct {
	PaymentName string `form:"paymentname" json:"paymentname"`
	NumberCard  string `form:"numbercard" json:"numbercard"`
	PaymentCode string `form:"paymentcode" json:"paymentcode"`
}

func ToCore(reqData Order) orders.Core {
	return orders.Core{
		Address: orders.AddressCore{
			Receiver: reqData.Address.Receiver,
			Phone:    reqData.Address.Phone,
			Address:  reqData.Address.Address,
		},
	}
}

func ToPaymentCore(reqData Payment) orders.PaymentCore {
	return orders.PaymentCore{
		PaymentName: reqData.PaymentName,
		NumberCard:  reqData.NumberCard,
		PaymentCode: reqData.PaymentCode,
	}
}
