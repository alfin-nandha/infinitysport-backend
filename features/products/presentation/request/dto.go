package request

import Products "project/e-comerce/features/products"

type Product struct {
	Name   string `json:"name" form:"name"`
	Detail string `json:"detail" form:"detail"`
	Stock  int    `json:"stock" form:"stock"`
	Price  int    `json:"price" form:"price"`
}

func ToCore(productReq Product) Products.Core {
	productCore := Products.Core{
		Name:          productReq.Name,
		ProductDetail: productReq.Detail,
		Stock:         productReq.Stock,
		Price:         productReq.Price,
	}
	return productCore
}
