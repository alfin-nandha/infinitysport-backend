package products

import (
	"time"
)

type Core struct{
	ID				int
	Name 			string
	ProductName 	string
	ProductDetail 	string
	Stock 			int
	Price 			int
	Photo 			string
	PhotoUrl		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	UserID			int
	User			User
}

type User struct{
	ID int
	Name string
	Email string
}

type Business interface{
	GetAllProduct()(data []Core, err error)
	GetProductByID(param int)(data Core, err error)
	InsertProduct(dataReq Core)(err error)
	DeleteProductByID(id int, userId int)(err error)
	UpdateProductByID(dataReq Core, id int, userId int)(err error)
	GetProductByUserID(id_user int)(data []Core, err error)
}

type Data interface{
	SelectData()(data []Core, err error)
	SelectDataByID(param int)(data Core, err error)
	InsertData(dataReq Core)(err error)
	DeleteDataByID(id int, userId int)(err error)
	UpdateDataByID(dataReq map[string]interface{}, id int, userId int)(err error)
	SelectDataByUserID(id_user int)(data []Core, err error)
}