package request

import "project/e-comerce/features/orders"


type Order struct{
	Price 	int 	`form:"price" json:"price"`
	CartID 	[]int	`form:"cartid" json:"cartid"`
	Address Address	`form:"address" json:"address"`
	Payment Payment	`form:"payment" json:"payment"`
}


type Address struct{
	City 		string 	`form:"city" json:"city"`
	Province 	string	`form:"province" json:"province"`
	PostalCode 	string	`form:"postalcode" json:"postalcode"`
	Street 		string	`form:"street" json:"street"`
}

type Payment struct{
	PaymentName string	`form:"paymentname" json:"paymentname"`
	NumberCard 	string	`form:"numbercard" json:"numbercard"`
	PaymentCode string	`form:"paymentcode" json:"paymentcode"`
}

func ToCore(reqData Order)orders.Core{
	return orders.Core{
		Price: reqData.Price,
		Address: orders.AddressCore{
			City: reqData.Address.City,
			Province: reqData.Address.Province,
			PostalCode: reqData.Address.PostalCode,
			Street: reqData.Address.Street,
		},
		Payment: orders.PaymentCore{
			PaymentName: reqData.Payment.PaymentName,
			NumberCard: reqData.Payment.NumberCard,
			PaymentCode: reqData.Payment.PaymentCode,
		},
	}
}