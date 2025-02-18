package data

import (
	Carts "project/e-comerce/features/carts"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ProductID int
	UserID    int
	Qty       int `json:"quantity" form:"quantity"`
	Product   Product
	User      User
}

type Product struct {
	gorm.Model
	Name   string
	Detail string
	Photo  string
	URL    string
	Stock  int
	Price  int
	UserID int
	User   User
	Cart   []Cart
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Product  []Product
}

func (data *Cart) toCore() Carts.Core {
	return Carts.Core{
		ID:        int(data.ID),
		ProductID: data.ProductID,
		UserID:    data.UserID,
		Qty:       data.Qty,
		Product: Carts.Product{
			ID:       int(data.Product.ID),
			Name:     data.Product.Name,
			PhotoUrl: data.Product.URL,
			Stock:    data.Product.Stock,
			Price:    data.Product.Price,
		},
	}
}

func ToCoreList(data []Cart) []Carts.Core {
	result := []Carts.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core Carts.Core) Cart {
	return Cart{
		ProductID: core.Product.ID,
		UserID:    core.UserID,
		Qty:       core.Qty,
	}
}

// func toCore(data Cart) carts.Core {
// 	return data.toCore()
// }
