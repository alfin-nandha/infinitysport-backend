package data

import (
	//"project/e-comerce/features/carts"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ProductID int
	UserID    int
	Qty       int 
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
