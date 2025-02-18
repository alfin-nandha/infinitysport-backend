package request

import Orders "project/e-comerce/features/orders"

type Order struct {
	//Price 	int 	`form:"price" json:"price"`
	CartID  []int   `form:"cartid" json:"cartid"`
	Address Address `form:"address" json:"address"`
}

type Address struct {
	Receiver string `form:"receiver" json:"receiver"`
	Phone    string `form:"phone" json:"phone"`
	Address  string `form:"address" json:"address"`
}

type Payment struct {
	PaymentName string `form:"paymentname" json:"paymentname"`
	NumberCard  string `form:"numbercard" json:"numbercard"`
	PaymentCode string `form:"paymentcode" json:"paymentcode"`
}

func ToCore(reqData Order) Orders.Core {
	return Orders.Core{
		Address: Orders.AddressCore{
			Receiver: reqData.Address.Receiver,
			Phone:    reqData.Address.Phone,
			Address:  reqData.Address.Address,
		},
	}
}

func ToPaymentCore(reqData Payment) Orders.PaymentCore {
	return Orders.PaymentCore{
		PaymentName: reqData.PaymentName,
		NumberCard:  reqData.NumberCard,
		PaymentCode: reqData.PaymentCode,
	}
}
