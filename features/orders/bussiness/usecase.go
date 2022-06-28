package bussiness

import (
	"fmt"
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
	fmt.Println(cartID)
	dataDetailOrder,errGetDetailOrder := uc.orderData.SelectChart(cartID)
	if errGetDetailOrder != nil{
		log.Print(errGetDetailOrder)
		return errGetDetailOrder
	}
	fmt.Println("usecase" , dataDetailOrder)
	errOrderDetail := uc.orderData.InsertOrderDetail(orderID, dataDetailOrder)
	if errOrderDetail != nil{
		log.Print(errOrderDetail)
		return errOrderDetail
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