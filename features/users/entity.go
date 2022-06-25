package users

import (
	"time"
)

type Core struct{
	ID 			int
	Name 		string
	Email 		string
	Password 	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type Business interface{
	GetAllData(param string)(data []Core, err error)
	GetDataById(param int)(data Core, err error)
	InsertData(dataReq Core)(err error)
	DeleteData(id int)(err error)
	UpdateData(dataReq Core, id int)(err error)
}

type Data interface{
	SelectData(param string)(data []Core, err error)
	SelectDataById(param int)(data Core, err error)
	InsertData(dataReq Core)(err error)
	DeleteData(id int)(err error)
	UpdateData(dataReq map[string]interface{}, id int)(err error)
}