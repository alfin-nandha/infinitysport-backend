package orders

import (
	"time"
)

type Core struct{
	ID int
	PaymentID int
	UserID int
	AddressID int
	Price int
	Address AddressCore
	Payment PaymentCore
	CreatedAt time.Time
}

type OrderDetail struct{
	ID int
	OrderID int
	ProductName string
	Price int
	Qty int
}

type AddressCore struct{
	ID int
	City string
	Province string
	PostalCode string
	Street string
}

type PaymentCore struct{
	ID int
	PaymentName string
	NumberCard string
	PaymentCode string
}

type Bussiness interface{
	AddOrder(core Core,orderdetailid []int)(error)
	GetOrder(orderId int)(orderData []Core ,err error)
	GetOrderDetail(orderId int)([]OrderDetail, error)
}

type Data interface {
	InsertOrder(core Core ,payID int ,addID int)(orderID int, err error)
	InsertAddress(AddressCore)(addID int,err error)
	InsertPayment(PaymentCore)(payID int,err error)
	InsertOrderDetail(int,[]OrderDetail)(error)
	SelectOrder(int)([]Core ,error)
	SelectChart([]int)([]OrderDetail,error)
	SelectOrderDetailByOrderID(orderID int)([]OrderDetail, error)
}