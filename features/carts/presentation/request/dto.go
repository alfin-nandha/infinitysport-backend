package request

import Carts "project/e-comerce/features/carts"

type Cart struct {
	ProductID int `json:"product_id" form:"product_id"`
	Qty       int `json:"quantity" form:"quantity"`
}

func ToCore(cartReq Cart) Carts.Core {
	cartCore := Carts.Core{
		ProductID: cartReq.ProductID,
		Qty:       cartReq.Qty,
	}
	return cartCore
}

func ToUpdateCore(cartReq Cart) Carts.Core {
	cartCore := Carts.Core{
		Qty: cartReq.Qty,
	}
	return cartCore
}
