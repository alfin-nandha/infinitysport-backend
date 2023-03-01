package data

import (
	Products "project/e-comerce/features/products"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	Detail string `json:"detail" form:"detail"`
	Photo  string `json:"photo" form:"photo"`
	URL    string `json:"url" form:"url"`
	Stock  int    `json:"stock" form:"stock"`
	Price  int    `json:"price" form:"price"`
	UserID int
	User   User
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Product  []Product
}

//DTO

func (data *Product) toCore() Products.Core {
	return Products.Core{
		ID:            int(data.ID),
		Name:          data.Name,
		ProductDetail: data.Detail,
		Photo:         data.Photo,
		PhotoUrl:      data.URL,
		Stock:         data.Stock,
		Price:         data.Price,
		UserID:        data.UserID,
		User: Products.User{
			ID:    int(data.User.ID),
			Name:  data.User.Name,
			Email: data.User.Email,
		},
	}
}

func ToCoreList(data []Product) []Products.Core {
	result := []Products.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core Products.Core) Product {
	return Product{
		Name:   core.Name,
		Detail: core.ProductDetail,
		Photo:  core.Photo,
		URL:    core.PhotoUrl,
		Stock:  core.Stock,
		Price:  core.Price,
		UserID: core.UserID,
		User: User{
			Name:  core.User.Name,
			Email: core.User.Email,
		},
	}
}

func toCore(data Product) Products.Core {
	return data.toCore()
}
