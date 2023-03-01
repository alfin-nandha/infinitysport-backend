package response

import Carts "project/e-comerce/features/carts"

type Cart struct {
	ID          int    `json:"id" form:"id"`
	NameProduct string `json:"product_name" form:"product_name"`
	Image       string `json:"image" form:"image"`
	UserID      int    `json:"user_id" form:"user_id"`
	PriceOrder  int    `json:"price" form:"price"`
	Qty         int    `json:"quantity" form:"quantity"`
}

func FromCore(data Carts.Core) Cart {
	return Cart{
		ID:          data.ID,
		UserID:      data.UserID,
		NameProduct: data.Product.Name,
		Image:       data.Product.PhotoUrl,
		PriceOrder:  data.Product.Price,
		Qty:         data.Qty,
	}
}

func FromCoreList(data []Carts.Core) []Cart {
	result := []Cart{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
