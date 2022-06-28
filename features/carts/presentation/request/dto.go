package request

import "project/e-comerce/features/carts"

type Cart struct {
	ProductID int `json:"product_id" form:"product_id"`
	Qty       int `json:"quantity" form:"quantity"`
}

func ToCore(cartReq Cart) carts.Core {
	cartCore := carts.Core{
		ProductID: cartReq.ProductID,
		Qty:       cartReq.Qty,
	}
	return cartCore
}
