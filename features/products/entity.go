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
	UserID			int
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type Business interface{
	GetAllProduct()(data []Core, err error)
	GetProductByID(param int)(data Core, err error)
	InsertProduct(dataReq Core)(err error)
	DeleteProductByID(id int)(err error)
	UpdateProductByID(dataReq Core, id int)(err error)
	GetProductByUserID(id_user int)(data []Core, err error)
}

type Data interface{
	SelectData()(data []Core, err error)
	SelectDataByID(param int)(data Core, err error)
	InsertData(dataReq Core)(err error)
	DeleteDataByID(id int)(err error)
	UpdateDataByID(dataReq map[string]interface{}, id int)(err error)
	SelectDataByUserID(id_user int)(data []Core, err error)
}