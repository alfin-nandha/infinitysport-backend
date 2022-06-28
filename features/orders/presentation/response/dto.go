package response

import (
	_order "project/e-comerce/features/orders"
	"time"
)

type Order struct{
	ID int `json:"id" form:"id"`
	Time time.Time `json:"time" form:"time"`
	Payment string `json:"payment" form:"payment"`
	Price int `json:"price" form:"price"`
	UserID int `json:"userid" form:"userid"`
}

type OrderDetail struct{
	ID int `json:"id" form:"id"`
	ProductName string `json:"productname" form:"Productname"`
	Qty int `json:"qty" form:"qty"`
	Price int `json:"price" form:"price"`
}

func OrderFromCore(data _order.Core) Order{
	return Order{
		ID: data.ID,
		Time: data.CreatedAt,
		Price: data.Price,
		UserID: data.UserID,
	}
}
func OrderDetailFromOrderDetailCore(data _order.OrderDetail) OrderDetail{
	return OrderDetail{
		ID: data.ID,
		ProductName: data.ProductName,
		Price: data.Price,
	}
}
