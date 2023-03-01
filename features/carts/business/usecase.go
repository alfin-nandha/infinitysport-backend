package business

import (
	Carts "project/e-comerce/features/carts"
)

type cartUseCase struct {
	cartData Carts.Data
}

func NewCartBusiness(usrData Carts.Data) Carts.Business {
	return &cartUseCase{
		cartData: usrData,
	}
}

func (uc *cartUseCase) GetAllCart(UserId int) (response []Carts.Core, err error) {
	response, err = uc.cartData.SelectData(UserId)
	return response, err
}

func (uc *cartUseCase) AddCart(data Carts.Core) (result int, err error) {
	resultCart, idCart, Qty, err := uc.cartData.CheckProductInCart(data.UserID, data.ProductID)
	if err != nil {
		return 0, err
	}
	newQty := Qty + 1

	if !resultCart {
		result, err = uc.cartData.InsertData(data)

	} else {
		UserId := data.UserID
		result, err = uc.cartData.Update(UserId, idCart, newQty)
	}
	return result, err
}

func (uc *cartUseCase) UpdateCart(data Carts.Core, idCart int) (result int, err error) {
	UserId := data.UserID
	newQty := data.Qty

	result, err = uc.cartData.Update(UserId, idCart, newQty)
	return result, err
}

func (uc *cartUseCase) DestroyCart(UserId, idCart int) (result int, err error) {
	result, err = uc.cartData.Destroy(UserId, idCart)
	return result, err
}
