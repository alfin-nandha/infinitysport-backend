package data

import (
	CartData "project/e-comerce/features/carts/data"
	Orders "project/e-comerce/features/orders"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    int
	Price     int
	Status    string
	AddressID int
	Address   Address
}

type OrderDetail struct {
	ID          int
	OrderID     int
	ProductName string
	Price       int
	Url         string
	Qty         int
}

type Cart struct {
	gorm.Model
	ProductID int
	UserID    int
	Qty       int
	Product   CartData.Product
}

type Address struct {
	ID       int
	Receiver string
	Phone    string
	Address  string
}

type Payment struct {
	ID          int
	PaymentName string
	NumberCard  string
	PaymentCode string
	orderID     int
}

func (data *Order) toCore() Orders.Core {
	return Orders.Core{
		ID:        int(data.ID),
		Price:     data.Price,
		Status:    data.Status,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
		AddressID: data.AddressID,
		Address: Orders.AddressCore{
			ID:       data.Address.ID,
			Receiver: data.Address.Receiver,
			Phone:    data.Address.Phone,
			Address:  data.Address.Address,
		},
	}
}

func ToCoreList(data []Order) []Orders.Core {
	result := []Orders.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *OrderDetail) toOrderDetailCore() Orders.OrderDetail {
	return Orders.OrderDetail{
		ID:          int(data.ID),
		OrderID:     data.OrderID,
		Price:       data.Price,
		Qty:         data.Qty,
		ProductName: data.ProductName,
	}
}

func ToOrderDetailCoreList(data []OrderDetail) []Orders.OrderDetail {
	result := []Orders.OrderDetail{}
	for key := range data {
		result = append(result, data[key].toOrderDetailCore())
	}
	return result
}

func (data *Cart) toOrderDetailFromCart() Orders.OrderDetail {
	return Orders.OrderDetail{
		ID:          int(data.ID),
		Qty:         data.Qty,
		ProductName: data.Product.Name,
		Price:       data.Product.Price,
	}
}

func ToOrderDetailCoreListFromCart(data []Cart) []Orders.OrderDetail {
	result := []Orders.OrderDetail{}
	for key := range data {
		result = append(result, data[key].toOrderDetailFromCart())
	}
	return result
}

func fromCore(core Orders.Core) Order {
	return Order{
		AddressID: core.AddressID,
		Price:     core.Price,
		UserID:    core.UserID,
	}
}

func fromOrderDetailCore(orderDetailCore Orders.OrderDetail) OrderDetail {
	return OrderDetail{
		ProductName: orderDetailCore.ProductName,
		Price:       orderDetailCore.Price,
		OrderID:     orderDetailCore.OrderID,
		Qty:         orderDetailCore.Qty,
	}
}

func fromOrderDetailCoreList(data []Orders.OrderDetail) []OrderDetail {
	result := []OrderDetail{}
	for key := range data {
		result = append(result, fromOrderDetailCore(data[key]))
	}
	return result
}

func fromAddressCore(addressCore Orders.AddressCore) Address {
	return Address{
		Receiver: addressCore.Receiver,
		Phone:    addressCore.Phone,
		Address:  addressCore.Address,
	}
}

func fromPaymentCore(paymentCore Orders.PaymentCore) Payment {
	return Payment{
		NumberCard:  paymentCore.NumberCard,
		PaymentCode: paymentCore.PaymentCode,
		PaymentName: paymentCore.PaymentName,
		orderID:     paymentCore.OrderID,
	}
}
