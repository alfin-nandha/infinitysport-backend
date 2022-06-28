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

func (data *Cart) toCore() carts.Core {
  return carts.Core {
    ID: int(data.ID)
    ProductID int 
    UserID int 
    Price int 
    Qty int 
    Product: carts.Product {
      ID: int(data.Product.ID),
      Name: data.Product.ID,
      PhotoUrl: data.Product.PhotoUrl,
      Stock: data.Product.Stock,
      Price: data.Product.Price,
    },
    User: carts.User {
      ID: int(data.User.ID),
      Name: data.User.Name,
      Email: data.User.Email,
    },
  }
}

func ToCoreList(data []Cart)[]carts.Core {
	result :=[]carts.Core{}
	for key := range data{
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core carts.Core) Cart {
  return Cart {
    ID: core.ID,
    ProductID: core.Product.ID,
    UserID: core.User.ID,
    Price: core.Product.Price,
    Qty: core.Product.Qty,
    Product: Product {
      Name: core.Product.Name,
      PhotoUrl: core.Product.PhotoUrl,
      Stock: core.Product.Stock,
      Price: core.Product.Price,
    },
    User: User {
      Name: core.User.Name,
      Email: core.User.Email,
    },
  }
}

func toCore(data Cart)carts.Core{
	return data.toCore()
}