package bussiness

import (
	"errors"
	"log"
	"project/e-comerce/features/orders"
)

type orderUseCase struct{
	orderData orders.Data
}

func NewOrderBusiness(usrData orders.Data) orders.Bussiness {
	return &orderUseCase{
		orderData: usrData,
	}
}

func (uc orderUseCase)AddOrder(core orders.Core, cartID []int)(error){
	dataDetailOrder, productId, totalPrice, errGetDetailOrder := uc.orderData.SelectCart(cartID)
	if errGetDetailOrder != nil{
		log.Print(errGetDetailOrder)
		return errGetDetailOrder
	}

	var remain_stock_list []int
	for i,val := range productId{
		stockProduct, errCheckStock := uc.orderData.SelectProduct(val)
		if errCheckStock != nil{
			log.Print(errCheckStock)
			return errCheckStock
		}
		remain_stock := stockProduct - dataDetailOrder[i].Qty
		remain_stock_list = append(remain_stock_list, remain_stock) 
		if remain_stock < 0{
			return errors.New("product stock is not enough")
		}
	}

	addID, errAdd := uc.orderData.InsertAddress(core.Address)
	if errAdd != nil{
		log.Print(errAdd)
		return errAdd
	}
	
	payID, errPay := uc.orderData.InsertPayment(core.Payment)
	if errPay != nil{
		log.Print(errPay)
		return errPay
	}
	
	orderID, errOrder := uc.orderData.InsertOrder(core, payID, addID)
	if errOrder != nil{
		log.Print(errOrder)
		return errOrder
	}

	errOrderDetail := uc.orderData.InsertOrderDetail(orderID, dataDetailOrder)
	if errOrderDetail != nil{
		log.Print(errOrderDetail)
		return errOrderDetail
	}

	errUpdateOrder := uc.orderData.UpdateOrder(orderID, totalPrice)
	if errUpdateOrder != nil{
		log.Print(errUpdateOrder)
		return errUpdateOrder
	}

	for i,val := range productId{
		errUpdateProduct := uc.orderData.UpdateStockProduct(val,remain_stock_list[i])
		if errUpdateProduct != nil{
			log.Print(errUpdateProduct)
			return errUpdateProduct
		}
	}
	return nil
}

func (uc orderUseCase)GetOrder(userId int)([]orders.Core,error){
	respData, err := uc.orderData.SelectOrder(userId)
	return respData,err
}

func (uc orderUseCase)GetOrderDetail(orderId int)([]orders.OrderDetail ,error){
	respData, err := uc.orderData.SelectOrderDetailByOrderID(orderId)
	return respData,err
}