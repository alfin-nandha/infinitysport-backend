package response

import "project/e-comerce/features/carts"

type Cart struct {
	ID        int `json:"id" form:"id"`
	ProductID int `json:"product_id" form:"product_id"`
	UserID    int `json:"user_id" form:"user_id"`
	Price     int `json:"price" form:"price"`
	Qty       int `json:"quantity" form:"quantity"`
}

func FromCore(data carts.Core) Cart {
	return Cart{
		ID:        data.ID,
		ProductID: data.ProductID,
		UserID:    data.UserID,
		Price:     data.Product.Price,
		Qty:       data.Qty,
	}
}

func FromCoreList(data []carts.Core) []Cart {
	result := []Cart{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
