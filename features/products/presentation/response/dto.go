package response

import "project/e-comerce/features/products"

type Product struct{
	ID int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Detail string `json:"detail" form:"detail"`
	Photo string `json:"photo" form:"photo"`
	URL string `json:"url" form:"url"`
	Stock int `json:"stock" form:"stock"`
	Price int `json:"price" form:"price"`
	UserID int `json:"userid" form:"userid"`
}

func FromCore(data products.Core) Product{
	return Product{
		ID: data.ID,
		Name: data.Name,
		Detail: data.ProductDetail,
		Photo: data.Photo,
		URL: data.PhotoUrl,
		Stock: data.Stock,
		Price: data.Price,
		UserID: data.UserID,
	}
}

func FromCoreList(data []products.Core) []Product{
	result := []Product{}
	for key := range data{
		result = append(result, FromCore(data[key]))
	}
	return result
}