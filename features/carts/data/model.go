package data

import (
	//"project/e-comerce/features/carts"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ProductID int
	UserID    int
	Price     int `json:"price" form:"price"`
	Qty       int `json:"quantity" form:"quantity"`
	Product   Product
	User      User
}

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
	Cart   []Cart
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Product  []Product
}
